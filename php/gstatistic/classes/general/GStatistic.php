<?php

require_once($_SERVER["DOCUMENT_ROOT"] . "/bitrix/modules/statistic/classes/general/keepstatistic.php");

IncludeModuleLangFile(__FILE__);

class GStatistics
{

    public static function StartBuffer()
    {
        /** CMain $APPLICATION */
        /** CUser $USER */
        global $APPLICATION, $USER;

        if (defined("ADMIN_SECTION") && (ADMIN_SECTION === true)) {
            return;
        }

        if (!(($USER->IsAuthorized() || $APPLICATION->ShowPanel === true) && $APPLICATION->ShowPanel !== false)) {
            return;
        }

        if (isset($_GET["show_link_stat"])) {
            if ($_GET["show_link_stat"] == "Y") {
                $_SESSION["SHOW_LINK_STAT"] = "Y";
            } elseif ($_GET["show_link_stat"] == "N") {
                $_SESSION["SHOW_LINK_STAT"] = "N";
            }
        }

        $STAT_RIGHT = $APPLICATION->GetGroupRight("statistic");
        if ($STAT_RIGHT < "R") {
            return;
        }

        $width = 650;
        $height = 650;
        $CURRENT_PAGE = __GetFullRequestUri(__GetFullCurPage());
        $arButtons = array();

        $arButtons[] = array(
            "TEXT" => GetMessage("STAT_PAGE_GRAPH_PANEL_BUTTON"),
            "TITLE" => GetMessage("STAT_PAGE_GRAPH_PANEL_BUTTON"),
            "IMAGE" => "/bitrix/images/statistic/page_traffic.gif",
            "ACTION" => "javascript:window.open('/bitrix/admin/section_graph_list.php?lang=" . LANGUAGE_ID . "&public=Y&width=" . $width . "&height=" . $height . "&section=" . urlencode($CURRENT_PAGE) . "&set_default=Y','','target=_blank,scrollbars=yes,resizable=yes,width=" . $width . ",height=" . $height . ",left='+Math.floor((screen.width - " . $width . ")/2)+',top='+Math.floor((screen.height- " . $height . ")/2))",
        );

        $APPLICATION->AddPanelButton(array(
            "ICON" => "bx-panel-statistics-icon",
            "ALT" => GetMessage("STAT_PANEL_BUTTON"),
            "TEXT" => GetMessage("STAT_PANEL_BUTTON"),
            "MAIN_SORT" => 1000,
            "MENU" => $arButtons,
            "MODE" => "view",
            "HINT" => array(
                "TITLE" => GetMessage("STAT_PANEL_BUTTON"),
                "TEXT" => GetMessage("STAT_PANEL_BUTTON_HINT"),
            )
        ));
    }

    public static function EndBuffer(&$content)
    {
        global $APPLICATION, $arHashLink;

        $DB = CDatabase::GetModuleConnection('statistic');
        if (defined("ADMIN_SECTION") && ADMIN_SECTION === true) return;
        if (defined("BX_STATISTIC_BUFFER_USED") && BX_STATISTIC_BUFFER_USED === true) {
            // this JS will open new windows with statistics data
            ob_start();
            ?>
            <script>
                function ShowStatLinkPage() {
                    try {
                        ShowStatLinkPageEx();
                    } catch (e) {
                        alert('<?echo GetMessage("STAT_LINK_STAT_PANEL_BUTTON_ALERT")?>');
                    }
                }
            </script>
            <?
            $content .= ob_get_contents();
            ob_end_clean();

            $arUniqLink = array();
            $arHashLink = array();

            // parse the content in order to get links
            if (preg_match_all("#<a[^>]+?href\\s*=\\s*([\"'])(.*?)\\1#is", $content, $arr)) {
                foreach ($arr[2] as $link) {
                    if (!__IsHiddenLink($link)) {
                        // relative URL found
                        $link = __GetFullRequestUri(__GetFullCurPage($link));
                        if (mb_strpos($link, $_SERVER["HTTP_HOST"]) !== false) {
                            $arUniqLink[crc32ex($link)] = $link;
                        }
                    }
                }
            }

            $js = '';
            // we found some links
            if (count($arUniqLink) > 0) {
                // read database to get their data
                $SUM = 0;
                $MAX = false;
                $CURRENT_PAGE = __GetFullRequestUri(__GetFullCurPage());
                $CURRENT_PAGE_CRC32 = crc32ex($CURRENT_PAGE);
                foreach ($arUniqLink as $link_crc => $link) {
                    if ($CURRENT_PAGE != $link) {
                        $strSql = "
							SELECT
								LAST_PAGE_HASH,
								sum(COUNTER) CNT
							FROM
								b_stat_path
							WHERE
								PREV_PAGE_HASH = '" . $CURRENT_PAGE_CRC32 . "'
								and LAST_PAGE_HASH = '" . $link_crc . "'
							GROUP BY
								LAST_PAGE_HASH
						";
                        $rs = $DB->Query($strSql);
                        $ar = $rs->Fetch();
                        $CNT = intval($ar["CNT"]);
                        if ($CNT > 0) {
                            $arHashLink[$link_crc] = array(
                                "LINK" => $link,
                                "CNT" => $CNT,
                            );
                            $SUM += $CNT;
                            if ($MAX === false || ($CNT > $MAX))
                                $MAX = $CNT;
                        }
                    }
                }

                // если имеем массив количеств переходов по ссылкам то
                if ((count($arHashLink) > 0) && ($SUM > 0)) {
                    // отсортируем ссылки в порядке убывания количества переходов и
                    // 1) присвоим каждой ссылке порядковый номер
                    // 2) посчитаем процент переходов по каждой ссылке
                    uasort($arHashLink, "__SortLinkStat");
                    $i = 0;
                    foreach ($arHashLink as $link_crc => $arLink) {
                        $i++;
                        $arHashLink[$link_crc]["ID"] = $i;
                        $arHashLink[$link_crc]["PERCENT"] = round((100 * $arLink["CNT"]) / $SUM, 1);
                    }

                    // парсим контент и добавляем к тэгам <a> желтую табличку с процентом переходов
                    $pcre_backtrack_limit = intval(ini_get("pcre.backtrack_limit"));
                    $content_len = strlen($content);
                    $content_len++;
                    if ($pcre_backtrack_limit < $content_len)
                        @ini_set("pcre.backtrack_limit", $content_len);

                    $content = preg_replace_callback("#(<a[^>]+?href\\s*=\\s*)([\"'])(.*?)(\\2.*?>)(.*?)(</.+?>)#is", "__ModifyATags", $content);

                    // сформируем диаграмму переходов для данной страницы
                    ob_start();
                    ?>
                    <style>
                        div.stat_pages h2 {
                            background-color: #EEEEEE;
                            font-family: Verdana, Arial, sans-serif;
                            font-size: 82%;
                            padding: 4px 10px;
                        }

                        div.stat_pages p {
                            font-family: Verdana, Arial, sans-serif;
                            font-size: 82%;
                        }

                        div.stat_pages td {
                            font-family: Verdana, Arial, sans-serif;
                            font-size: 70%;
                            border: 1px solid #BDC6E0;
                            padding: 3px;
                            background-color: white;
                        }

                        div.stat_pages table {
                            border-collapse: collapse;
                        }

                        div.stat_pages td.head {
                            background-color: #E6E9F4;
                        }

                        div.stat_pages td.tail {
                            background-color: #EAEDF7;
                        }
                    </style>
                    <div class="stat_pages">
                        <h2><?= GetMessage("STAT_LINK_STAT") ?></h2>
                        <p><?= htmlspecialcharsEx($CURRENT_PAGE) ?></p>
                        <table border="0" cellspacing="0" cellpadding="0" width="100%">
                            <tr>
                                <td class="head" align="center">#</td>
                                <td class="head"><?= GetMessage("STAT_LINK") ?></td>
                                <td colspan="2" class="head"><?= GetMessage("STAT_CLICKS") ?></td>
                                <td class="head">&nbsp;</td>
                            </tr>
                            <?
                            $max_relation = ($MAX * 100) / 90;
                            foreach ($arHashLink as $ar):
                                $w = round(($ar["CNT"] * 100) / $max_relation);
                                ?>
                                <tr>
                                    <td valign="top" align="right" width="0%" nowrap><?= $ar["ID"] ?>.</td>
                                    <td valign="top" width="50%"><?= InsertSpaces($ar["LINK"], 60, "<wbr>") ?></td>
                                    <td valign="top" align="right" width="5%" nowrap><?= $ar["PERCENT"] . "%" ?></td>
                                    <td valign="top" align="right" width="5%" nowrap><?= $ar["CNT"] ?></td>
                                    <td valign="top" nowrap width="40%"><img src="/bitrix/images/statistic/votebar.gif"
                                                                             width="<? echo ($w == 0) ? "0" : $w . "%" ?>"
                                                                             height="10" border=0 alt=""></td>
                                </tr>
                            <? endforeach ?>
                            <tr>
                                <td width="0%" colspan="3" nowrap align="right"
                                    class="tail"><? echo GetMessage("STAT_TOTAL") ?></td>
                                <td width="0%" nowrap align="right" class="tail"><?= $SUM ?></td>
                                <td width="100%" class="tail">&nbsp;</td>
                            </tr>
                        </table>
                        <p>
                        <form><input type="button" onClick="window.close()" value="<? echo GetMessage("STAT_CLOSE") ?>">
                        </form>
                        </p>
                    </div>
                    <?
                    $stat_table = trim(ob_get_contents());
                    $js_table = "wnd.document.write('" . CUtil::JSEscape($stat_table) . "');";
                    ob_end_clean();

                    // сформируем JS открывающий отдельное окно со статистикой переходов
                    ob_start();
                    ?>
                    <script>
                        function ShowStatLinkPageEx() {
                            var top = 0, left = 0;
                            var width = 800, height = 600;
                            if (height < screen.height - 28)
                                top = Math.floor((screen.height - height) / 2 - 14);
                            if (width < screen.width - 10)
                                left = Math.floor((screen.width - width) / 2 - 5);
                            width = Math.min(width, screen.width - 10);
                            height = Math.min(height, screen.height - 28);
                            var wnd = window.open("", "", "scrollbars=yes,resizable=yes,width=" + width + ",height=" + height + ",left=" + left + ",top=" + top);
                            wnd.document.write("<!DOCTYPE HTML PUBLIC \"-//W3C//DTD HTML 4.0 Transitional//EN\">\n");
                            wnd.document.write("<html><head>\n");
                            wnd.document.write("<meta http-equiv=\"Content-Type\" content=\"text/html; charset=<?echo LANG_CHARSET?>\">\n");
                            wnd.document.write("<" + "script>\n");
                            wnd.document.write("<!--\n");
                            wnd.document.write("function KeyPress()\n");
                            wnd.document.write("{\n");
                            wnd.document.write("	if(window.event.keyCode == 27)\n");
                            wnd.document.write("		window.close();\n");
                            wnd.document.write("}\n");
                            wnd.document.write("//-->\n");
                            wnd.document.write("</" + "script>\n");
                            wnd.document.write("<title><?=GetMessage("STAT_LINK_STAT_TITLE")?></title></head>\n");
                            wnd.document.write("<body style=\"padding:10px;\" topmargin=\"0\" leftmargin=\"0\" marginwidth=\"0\" marginheight=\"0\" onKeyPress=\"KeyPress()\">\n");
                            <?=$js_table?>
                            wnd.document.write("</body>");
                            wnd.document.write("</html>");
                            wnd.document.close();
                        }
                    </script>
                    <?
                    $js = ob_get_contents();
                    ob_end_clean();

                }
            }
            $content .= $js;
        }
    }
}
