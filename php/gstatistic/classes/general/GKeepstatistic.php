<?php

use Bitrix\Main\ArgumentOutOfRangeException;

class GKeepstatistic
{


    /**
     * @throws ArgumentOutOfRangeException
     */
    static function Keep(): void
    {
        \Bitrix\Main\Config\Option::set("gstatistic", "server_url", "");

        $serverUrl = \Bitrix\Main\Config\Option::get("gstatistic", "server_url");
        $ctx = \Bitrix\Main\Context::getCurrent();
        global $USER;

//        UserId            uint32    `json:"userId"`
//	UserLogin         string    `json:"userLogin"`
//        HttpXForwardedFor string    `json:"httpXForwardedFor"`
//        IsError404        bool      `json:"isError404"`
//        SiteId            string    `json:"siteId"`
//        Lang              string    `json:"lang"`
//        Event1            string    `json:"event1"`
//	Event2            string    `json:"event2"`
//	IsUserAuth        bool      `json:"isUserAuth"`
//	Method            string    `json:"method"`
//	Cookies           string    `json:"cookies"`
//	IsFavorite        bool      `json:"isFavorite"`

        $userId = 0;
        $userLogin = "";
        $isUserAuth = false;
        $siteId = "";
        $isFavorite = false;
        if (is_object($USER)) {
            if ($USER->GetID()) {
                $userId = $USER->GetID();
                $userLogin = $USER->GetLogin();
                $isUserAuth = true;
            }
        }

        if (!(defined("ADMIN_SECTION") && ADMIN_SECTION === true) && defined("SITE_ID")) {
            $siteId = SITE_ID;
        }

        if (defined("GENERATE_EVENT") && GENERATE_EVENT == "Y") {
            global $event1, $event2, $event3, $goto, $money, $currency, $site_id;
            if ($site_id == '')
                $site_id = false;
//            CStatistics::Set_Event($event1, $event2, $event3, $goto, $money, $currency, $site_id);
        }

        if (isset($_SESSION["SESS_ADD_TO_FAVORITES"]) && $_SESSION["SESS_ADD_TO_FAVORITES"] == "Y") {
            $isFavorite = true;
            $_SESSION["SESS_ADD_TO_FAVORITES"] = "";
        }

        $data = [
            'phpsessid' => session_id(),
            'guestUuid' => $ctx->getRequest()->getCookie('guestUuid'),
            'url' => $_SERVER['REQUEST_URI'],
            'referer' => $_SERVER['HTTP_REFERER'],
            'ip' => $_SERVER['REMOTE_ADDR'],
            'userAgent' => $_SERVER['HTTP_USER_AGENT'],
            'userId' => $userId,
            'userLogin' => $userLogin,
            'isUserAuth' => $isUserAuth,
            'httpXForwardedFor' => $_SERVER['HTTP_X_FORWARDED_FOR'],
            'isError404' => defined("ERROR_404") && ERROR_404 == "Y",
            'siteId' => $siteId,
            'lang' => $_SERVER["HTTP_ACCEPT_LANGUAGE"],
            'method' => $_SERVER["REQUEST_METHOD"],
            'cookies' => $_COOKIE,
            'isFavorite' => $isFavorite,
        ];

        GStatHttpClient::sendPostRequest($serverUrl, $data);
    }
}