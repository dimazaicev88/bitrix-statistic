<?php

IncludeModuleLangFile(__FILE__);

if (class_exists("gstatistic")) return;

class gstatistic extends CModule
{
    var $MODULE_ID = "gstatistic";
    var $MODULE_VERSION;
    var $MODULE_VERSION_DATE;
    var $MODULE_NAME;
    var $MODULE_DESCRIPTION;
    var $MODULE_CSS;
    var $MODULE_GROUP_RIGHTS = "Y";

    var $errors;

    public function __construct()
    {
        $arModuleVersion = [];

        include(__DIR__ . '/version.php');
        $this->MODULE_VERSION = $arModuleVersion["VERSION"];
        $this->MODULE_VERSION_DATE = $arModuleVersion["VERSION_DATE"];
        $this->MODULE_NAME = GetMessage("GSTAT_MODULE_NAME");
        $this->MODULE_DESCRIPTION = GetMessage("GSTAT_MODULE_DESCRIPTION");
    }

    function InstallDB($arParams = array())
    {

        RegisterModule("gstatistic");

        RegisterModuleDependences("main", "OnPageStart", "gstatistic", "GStopList", "Check", "100");
//        RegisterModuleDependences("main", "OnBeforeProlog", "gstatistic", "GStatistics", "Keep", "100");
        RegisterModuleDependences("main", "OnAfterEpilog", "gstatistic", "GKeepStatistics", "Keep", "100");
//        RegisterModuleDependences("main", "OnEpilog", "gstatistic", "GStatistics", "Set404", "100");
//        RegisterModuleDependences("main", "OnBeforeProlog", "gstatistic", "GStatistics", "StartBuffer", "1000");
//        RegisterModuleDependences("main", "OnEndBufferContent", "gstatistic", "GStatistics", "EndBuffer", "900");
//        RegisterModuleDependences("main", "OnEventLogGetAuditTypes", "gstatistic", "GStatistics", "GetAuditTypes", 10);

        return true;
    }

    function UnInstallDB($arParams = array())
    {
        UnRegisterModuleDependences("main", "OnPageStart", "gstatistic", "GStopList", "Check");
//        UnRegisterModuleDependences("main", "OnBeforeProlog", "gstatistic", "GStatistics", "Keep");
        UnRegisterModuleDependences("main", "OnAfterEpilog", "gstatistic", "GStatistics", "Keep");
//        UnRegisterModuleDependences("main", "OnEpilog", "gstatistic", "GStatistics", "Set404");
//        UnRegisterModuleDependences("main", "OnEventLogGetAuditTypes", "gstatistic", "GStatistics", "GetAuditTypes");
//        UnRegisterModuleDependences("main", "OnBeforeProlog", "gstatistic", "GStatistics", "StartBuffer");
//        UnRegisterModuleDependences("main", "OnEndBufferContent", "gstatistic", "GStatistics", "EndBuffer");

        UnRegisterModule("gstatistic");

        return true;
    }

    function InstallEvents()
    {
        return true;
    }

    function UnInstallEvents()
    {
        return true;
    }

    function InstallFiles()
    {
        return true;
    }

    function UnInstallFiles()
    {
        return true;
    }

    function DoInstall()
    {
        global $APPLICATION;

        $this->InstallDB();
        $APPLICATION->IncludeAdminFile(GetMessage("GSTAT_INSTALL_TITLE"), $_SERVER["DOCUMENT_ROOT"] . "/bitrix/modules/gstatistic/install/install_page.php");
    }


    function DoUninstall()
    {
        global $APPLICATION;

        $this->UnInstallDB();
        $APPLICATION->IncludeAdminFile(GetMessage("GSTAT_INSTALL_TITLE"), $_SERVER["DOCUMENT_ROOT"] . "/bitrix/modules/gstatistic/install/uninstall_page.php");
    }

    function GetModuleRightList()
    {
        return [
            "reference_id" => array("D", "M", "R", "W"),
            "reference" => array(
                "[D] " . GetMessage("STAT_DENIED"),
                "[M] " . GetMessage("STAT_VIEW_WITHOUT_MONEY"),
                "[R] " . GetMessage("STAT_VIEW"),
                "[W] " . GetMessage("STAT_ADMIN"))
        ];
    }
}
