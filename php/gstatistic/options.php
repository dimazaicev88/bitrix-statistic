<?php
IncludeModuleLangFile($_SERVER["DOCUMENT_ROOT"] . BX_ROOT . "/modules/main/options.php");
IncludeModuleLangFile(__FILE__);

$moduleId = "gstatistic";
$STAT_RIGHT = $APPLICATION->GetGroupRight($moduleId);
$strError = "";
$ctx = Bitrix\Main\Context::getCurrent();
if ($STAT_RIGHT >= "R") {

    $arTabs = [
        ["DIV" => "edit1",
            "TAB" => GetMessage("MAIN_TAB_SET"),
            "ICON" => "statistic_settings",
            "TITLE" => GetMessage("MAIN_TAB_TITLE_SET")
        ]
    ];
    $tabControl = new CAdminTabControl("tabControl", $arTabs);

    if ($ctx->getRequest()->isPost() == "POST" && $STAT_RIGHT == "W" && $RestoreDefaults <> '' && check_bitrix_sessid()) {
        COption::RemoveOption($moduleId);
        $z = CGroup::GetList("id", "asc", array("ACTIVE" => "Y", "ADMIN" => "N"));
        while ($zr = $z->Fetch())
            $APPLICATION->DelGroupRight($moduleId, array($zr["ID"]));
        LocalRedirect($APPLICATION->GetCurPage() . "?mid=" . urlencode($mid) . "&lang=" . urlencode(LANGUAGE_ID) . "&back_url_settings=" . urlencode($_REQUEST["back_url_settings"]) . "&" . $tabControl->ActiveTabParam());
    }

    $arOPTIONS = [
        "TAB1" => [
            "ONLINE_INTERVAL" => ["ONLINE_INTERVAL", GetMessage("STAT_OPT_ONLINE_INTERVAL"), ["text", 5]],
            "BASE_CURRENCY" => "",
        ],
    ];

    if ($ctx->getRequest()->isPost() == "POST" && $Update . $Apply <> '' && $STAT_RIGHT >= "W" && check_bitrix_sessid()) {
//if (CheckFDate($next_exec, GetMessage("STAT_OPT_WRONG_NEXT_EXEC"))) {
//    foreach ($arOPTIONS as $arOp) {
//        foreach ($arOp as $arOption) {
//            if (is_array($arOption)) {
//                $name = $arOption[0];
//                $val = $_REQUEST[$name];
//                $type = $arOption[2][0];
//                if ($type == "checkbox" && $val != "Y")
//                    $val = "N";
//                COption::SetOptionString($moduleId, $name, $val);
//                if (${$name . "_clear"} == "Y") {
//                    $func = $arOption[3];
//                    eval($func);
//                }
//            }
//        }
//    }

//    COption::SetOptionString($moduleId, "IP_LOOKUP_CLASS", $IP_LOOKUP_CLASS);
//    if ($recount_base_currency == "Y")
//        CStatistics::RecountBaseCurrency($BASE_CURRENCY);


        $Update = $Update . $Apply;
        ob_start();
        require_once($_SERVER["DOCUMENT_ROOT"] . "/bitrix/modules/main/admin/group_rights.php");
        ob_end_clean();

        if ($strError == "") {
            if ($Update <> '' && $_REQUEST["back_url_settings"] <> '')
                LocalRedirect($_REQUEST["back_url_settings"]);
            else
                LocalRedirect($APPLICATION->GetCurPage() . "?mid=" . urlencode($mid) . "&lang=" . urlencode(LANGUAGE_ID) . "&back_url_settings=" . urlencode($_REQUEST["back_url_settings"]) . "&" . $tabControl->ActiveTabParam());
        }
    }

//    if ($cleanup <> '' && $REQUEST_METHOD == "POST" && $STAT_RIGHT >= "W" && check_bitrix_sessid()) {
//        if (CheckFDate($cleanup_date, GetMessage("STAT_OPT_WRONG_CLEANUP_DATE"))) {
//            set_time_limit(0);
//            ignore_user_abort(true);
//            if (CStatistics::CleanUp($cleanup_date, $arErrors)) {
//                $_SESSION["STAT_strNote"] .= GetMessage("STAT_OPT_CLEAN_UP_OK") . "<br>";
//            } else {
//                $strError .= GetMessage("STAT_OPT_CLEAN_UP_ERRORS") . "<br><pre>" . mydump($arErrors) . "</pre><br>";
//            }
//        }
//        if ($strError == "") {
//            LocalRedirect($APPLICATION->GetCurPage() . "?mid=" . urlencode($mid) . "&lang=" . urlencode(LANGUAGE_ID) . "&back_url_settings=" . urlencode($_REQUEST["back_url_settings"]) . "&" . $tabControl2->ActiveTabParam());
//        }
//    }
//
//    if ($runsql <> '' && $REQUEST_METHOD == "POST" && $STAT_RIGHT >= "W" && check_bitrix_sessid()) {
//        set_time_limit(0);
//        ignore_user_abort(true);
//        $bDone = true;
//        if (count($ar = CStatistics::GetDDL()) > 0) {
//            foreach ($ar as $arDDL) {
//                if (!CStatistics::ExecuteDDL($arDDL["ID"])) {
//                    $strError .= $arDDL["SQL_TEXT"] . ":(" . $statDB->db_Error . ")<br>";
//                    $bDone = false;
//                }
//            }
//        }
//        if ($bDone) {
//            $_SESSION["STAT_strNote"] .= GetMessage("STAT_OPT_INDEXED") . "<br>";
//            COption::RemoveOption("statistic", "sql_to_run");
//        }
//        if ($strError == "") {
//            LocalRedirect($APPLICATION->GetCurPage() . "?mid=" . urlencode($mid) . "&lang=" . urlencode(LANGUAGE_ID) . "&back_url_settings=" . urlencode($_REQUEST["back_url_settings"]) . "&" . $tabControl2->ActiveTabParam());
//        }
//    }
//
//    if ($optimize <> '' && $REQUEST_METHOD == "POST" && $STAT_RIGHT >= "W" && check_bitrix_sessid()) {
//        set_time_limit(0);
//        ignore_user_abort(true);
//        $fname = $_SERVER["DOCUMENT_ROOT"] . "/bitrix/modules/statistic/install/db/" . mb_strtolower($statDB->type) . "/optimize.sql";
//        if (file_exists($fname)) {
//            $arErrors = $statDB->RunSQLBatch($fname);
//            if (!$arErrors)
//                $_SESSION["STAT_strNote"] .= GetMessage("STAT_OPT_OPTIMIZED") . "<br>";
//            else
//                $strError .= GetMessage("STAT_OPT_OPTIMIZE_ERRORS") . "<br>" . mydump($arErrors) . "<br>";
//        }
//        if ($strError == "") {
//            LocalRedirect($APPLICATION->GetCurPage() . "?mid=" . urlencode($mid) . "&lang=" . urlencode(LANGUAGE_ID) . "&back_url_settings=" . urlencode($_REQUEST["back_url_settings"]) . "&" . $tabControl2->ActiveTabParam());
//        }
//    }

    if ($strError <> '')
        CAdminMessage::ShowMessage($strError);

    if ($_SESSION["STAT_strNote"] <> '') {
        CAdminMessage::ShowNote($_SESSION["STAT_strNote"]);
        unset($_SESSION["STAT_strNote"]);
    }

    $tabControl->Begin();
    ?>
    <form name="form_settings" method="POST"
          action="<?php echo $APPLICATION->GetCurPage() ?>?mid=<?= htmlspecialcharsbx($mid) ?>&amp;lang=<?= LANGUAGE_ID ?>">
        <?php $tabControl->BeginNextTab(); ?>
         <tr>
            <td nowrap><?php echo GetMessage("STAT_OPT_DEFENCE_DELAY") ?></td>
            <td><input size="3" type="text" name="DEFENCE_DELAY" id="DEFENCE_DELAY"
                       value="<?= htmlspecialcharsbx($DEFENCE_DELAY) ?>">&nbsp;<?php echo GetMessage("STAT_OPT_DEFENCE_DELAY_MEASURE_SEC") ?>
            </td>
        </tr>

        <?php $tabControl->EndTab(); ?>

        <?php if ($_REQUEST["back_url_settings"] <> ''): ?>
            <input type="hidden" name="back_url_settings"
                   value="<?= htmlspecialcharsbx($_REQUEST["back_url_settings"]) ?>">
        <?php endif ?>
    </form>
    <?php $tabControl->Buttons() ?>

    <input type="submit" name="Update" <?php if ($MOD_RIGHT < 'W') echo 'disabled'; ?>
           value="<?= GetMessage("MAIN_SAVE") ?>" title="<?= GetMessage("MAIN_OPT_SAVE_TITLE") ?>" class="adm-btn-save">
    <input type="submit" name="Apply" <?php if ($MOD_RIGHT < 'W') echo 'disabled'; ?>
           value="<?= GetMessage("MAIN_OPT_APPLY") ?>" title="<?= GetMessage("MAIN_OPT_APPLY_TITLE") ?>">
    <?php if ($_REQUEST["back_url_settings"] <> ''): ?>
        <input type="button" name="Cancel" value="<?= GetMessage("MAIN_OPT_CANCEL") ?>"
               title="<?= GetMessage("MAIN_OPT_CANCEL_TITLE") ?>"
               onclick="window.location='<?php echo htmlspecialcharsbx(CUtil::addslashes($_REQUEST["back_url_settings"])) ?>'">
        <input type="hidden" name="back_url_settings" value="<?= htmlspecialcharsbx($_REQUEST["back_url_settings"]) ?>">
    <?php endif ?>
    <input type="submit" name="RestoreDefaults" <?php if ($MOD_RIGHT < 'W') echo 'disabled'; ?>
           title="<?php echo GetMessage("MAIN_HINT_RESTORE_DEFAULTS") ?>"
           OnClick="return confirm('<?php echo AddSlashes(GetMessage("MAIN_HINT_RESTORE_DEFAULTS_WARNING")) ?>')"
           value="<?php echo GetMessage("MAIN_RESTORE_DEFAULTS") ?>">

    <?php $tabControl->End() ?>
    <?php
}