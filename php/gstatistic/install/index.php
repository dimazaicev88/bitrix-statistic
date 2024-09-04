<?php

IncludeModuleLangFile(__FILE__);

if (class_exists("statistic")) return;

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
        $arModuleVersion = array();

        include(__DIR__ . '/version.php');

        if (is_array($arModuleVersion) && array_key_exists("VERSION", $arModuleVersion)) {
            $this->MODULE_VERSION = $arModuleVersion["VERSION"];
            $this->MODULE_VERSION_DATE = $arModuleVersion["VERSION_DATE"];
        } else {
            $this->MODULE_VERSION = $STATISTIC_VERSION;
            $this->MODULE_VERSION_DATE = $STATISTIC_VERSION_DATE;
        }

        $this->MODULE_NAME = GetMessage("STAT_MODULE_NAME");
        $this->MODULE_DESCRIPTION = GetMessage("STAT_MODULE_DESCRIPTION");
    }

    function InstallDB($arParams = array())
    {

        RegisterModule("gstatistic");

        RegisterModuleDependences("main", "OnPageStart", "gstatistic", "GStopList", "Check", "100");
        RegisterModuleDependences("main", "OnBeforeProlog", "gstatistic", "GStatistics", "Keep", "100");
        RegisterModuleDependences("main", "OnAfterEpilog", "gstatistic", "GStatistics", "Keep", "100");
        RegisterModuleDependences("main", "OnEpilog", "gstatistic", "GStatistics", "Set404", "100");
        RegisterModuleDependences("main", "OnBeforeProlog", "gstatistic", "GStatistics", "StartBuffer", "1000");
        RegisterModuleDependences("main", "OnEndBufferContent", "gstatistic", "GStatistics", "EndBuffer", "900");
        RegisterModuleDependences("main", "OnEventLogGetAuditTypes", "gstatistic", "GStatistics", "GetAuditTypes", 10);

        return true;
    }

    function UnInstallDB($arParams = array())
    {
        UnRegisterModuleDependences("main", "OnPageStart", "gstatistic", "GStopList", "Check");
        UnRegisterModuleDependences("main", "OnBeforeProlog", "gstatistic", "GStatistics", "Keep");
        UnRegisterModuleDependences("main", "OnAfterEpilog", "gstatistic", "GStatistics", "Keep");
        UnRegisterModuleDependences("main", "OnEpilog", "gstatistic", "GStatistics", "Set404");
        UnRegisterModuleDependences("main", "OnEventLogGetAuditTypes", "gstatistic", "GStatistics", "GetAuditTypes");
        UnRegisterModuleDependences("main", "OnBeforeProlog", "gstatistic", "GStatistics", "StartBuffer");
        UnRegisterModuleDependences("main", "OnEndBufferContent", "gstatistic", "GStatistics", "EndBuffer");

        UnRegisterModule("statistic");

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
        global $APPLICATION, $step;
        $STAT_RIGHT = $APPLICATION->GetGroupRight("gstatistic");
        $step = intval($step);

        if ($STAT_RIGHT < "W")
            return;

        if (!CBXFeatures::IsFeatureEditable("Analytics")) {
            $this->errors = array(GetMessage("MAIN_FEATURE_ERROR_EDITABLE"));
            $GLOBALS["errors"] = $this->errors;
            $APPLICATION->ThrowException(implode("<br>", $this->errors));
            $APPLICATION->IncludeAdminFile(GetMessage("STAT_INSTALL_TITLE"), $_SERVER["DOCUMENT_ROOT"] . "/bitrix/modules/gstatistic/install/step2.php");
        } elseif ($step < 2) {
            $APPLICATION->IncludeAdminFile(GetMessage("STAT_INSTALL_TITLE"), $_SERVER["DOCUMENT_ROOT"] . "/bitrix/modules/gstatistic/install/step1.php");
        } elseif ($step == 2) {
            $db_install_ok = $this->InstallDB(array(
                "allow_initial" => $_REQUEST["allow_initial"],
                "START_HITS" => $_REQUEST["START_HITS"],
                "START_HOSTS" => $_REQUEST["START_HOSTS"],
                "START_GUESTS" => $_REQUEST["START_GUESTS"],
                "CREATE_I2C_INDEX" => $_REQUEST["CREATE_I2C_INDEX"],
                "DATABASE" => $_REQUEST["DATABASE"],
            ));
            if ($db_install_ok) {
                CBXFeatures::SetFeatureEnabled("Analytics", true);
            }
            $GLOBALS["errors"] = $this->errors;
            $APPLICATION->IncludeAdminFile(GetMessage("STAT_INSTALL_TITLE"), $_SERVER["DOCUMENT_ROOT"] . "/bitrix/modules/gstatistic/install/step2.php");
        }
    }

    function DoUninstall()
    {
        global $APPLICATION, $step;
        $STAT_RIGHT = $APPLICATION->GetGroupRight("statistic");
        if ($STAT_RIGHT >= "W") {
            $step = intval($step);
            if ($step < 2) {
                $APPLICATION->IncludeAdminFile(GetMessage("STAT_UNINSTALL_TITLE"), $_SERVER["DOCUMENT_ROOT"] . "/bitrix/modules/gstatistic/install/unstep1.php");
            } elseif ($step == 2) {
                $this->UnInstallDB(array(
                    "savedata" => $_REQUEST["savedata"],
                ));
                //message types and templates
                if ($_REQUEST["save_templates"] != "Y") {
                    $this->UnInstallEvents();
                }
                $this->UnInstallFiles();
                CBXFeatures::SetFeatureEnabled("Analytics", false);
                $GLOBALS["errors"] = $this->errors;
                $APPLICATION->IncludeAdminFile(GetMessage("STAT_UNINSTALL_TITLE"), $_SERVER["DOCUMENT_ROOT"] . "/bitrix/modules/gstatistic/install/unstep2.php");
            }
        }
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