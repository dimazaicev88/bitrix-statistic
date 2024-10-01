package main

import (
	"net/url"
	"os"
	"strings"
)

//type User struct {
//	UserGroups []int
//}
//
//var USER User
//
//func GetOptionString(category, option string) string {
//	// Имитация функции получения опций, заменить на фактическую реализацию
//	if category == "statistic" && option == "SKIP_STATISTIC_WHAT" {
//		return "ranges" // Пример опции
//	} else if category == "statistic" && option == "SKIP_STATISTIC_GROUPS" {
//		return "1,2,3" // Пример групп
//	} else if category == "statistic" && option == "SKIP_STATISTIC_IP_RANGES" {
//		return "192.168.1.1-192.168.1.255\n10.0.0.1-10.0.0.255" // Пример IP диапазонов
//	}
//	return ""
//}
//
//func CheckSkip() bool {
//	GO := true
//	skipMode := GetOptionString("statistic", "SKIP_STATISTIC_WHAT")
//
//	switch skipMode {
//	case "none":
//		break
//	case "both", "groups":
//		arUserGroups := USER.UserGroups
//		arSkipGroups := strings.Split(GetOptionString("statistic", "SKIP_STATISTIC_GROUPS"), ",")
//		for _, value := range arSkipGroups {
//			groupID, _ := strconv.Atoi(value)
//			for _, userGroup := range arUserGroups {
//				if userGroup == groupID {
//					GO = false
//					break
//				}
//			}
//			if !GO {
//				break
//			}
//		}
//		if skipMode == "groups" {
//			break
//		}
//		fallthrough
//	case "ranges":
//		if skipMode == "both" && !GO {
//			break
//		}
//		GO = true
//		ip := net.ParseIP(GetIPAddress())
//		arSkipIPRanges := strings.Split(GetOptionString("statistic", "SKIP_STATISTIC_IP_RANGES"), "\n")
//
//		for _, value := range arSkipIPRanges {
//			ipRange := strings.Split(value, "-")
//			if len(ipRange) == 2 {
//				startIP := net.ParseIP(strings.TrimSpace(ipRange[0]))
//				endIP := net.ParseIP(strings.TrimSpace(ipRange[1]))
//
//				if ipInRange(ip, startIP, endIP) {
//					GO = false
//					break
//				}
//			}
//		}
//		break
//	}
//	return GO
//}
//
//func ipInRange(ip, startIP, endIP net.IP) bool {
//	for i := 0; i < len(ip); i++ {
//		if ip[i] < startIP[i] || ip[i] > endIP[i] {
//			return false
//		}
//	}
//	return true
//}
//
//func GetIPAddress() string {
//	// Здесь можно использовать os.Getenv("REMOTE_ADDR") для получения IP-адреса пользователя
//	return os.Getenv("REMOTE_ADDR")
//}

// Mock function to simulate getting options from a config
func GetOptionString(category, option string) string {
	if category == "statistic" && option == "IMPORTANT_PAGE_PARAMS" {
		return "param1,param2" // Пример важных параметров
	} else if category == "statistic" && option == "DIRECTORY_INDEX" {
		return "index.html,index.php" // Пример индексов директорий
	}
	return ""
}

// GetCurrentDir simulates getting the current directory
func GetCurrentDir() string {
	return "/current/directory" // Пример текущей директории
}

// Rel2Abs simulates converting a relative path to an absolute path
func Rel2Abs(curdir, page string) string {
	return strings.TrimSuffix(curdir, "/") + "/" + page
}

// GetPage processes the page URL and parameters
func GetPage(page string, withImpParams bool, curdir string) string {
	if page == "" {
		page = os.Getenv("REQUEST_URI")
		curdir = ""
	} else {
		page = strings.ReplaceAll(page, "\\", "/")
		if !strings.HasPrefix(page, "/") && !strings.Contains(page, "://") {
			if curdir == "" {
				curdir = GetCurrentDir()
			}
			page = Rel2Abs(curdir, page)
		}
	}

	// Split URL into path and query parts
	parsedURL, err := url.Parse(page)
	if err != nil {
		return ""
	}
	sPath := parsedURL.Path

	if sPath != "" {
		if stat, err := os.Stat(os.Getenv("DOCUMENT_ROOT") + sPath); err == nil && stat.IsDir() {
			sPath += "/"
		}
	}

	if withImpParams {
		impParams := strings.Split(GetOptionString("statistic", "IMPORTANT_PAGE_PARAMS"), ",")
		queryParams := parsedURL.Query()

		for i, key := range impParams {
			value := queryParams.Get(key)
			if value != "" {
				if i > 0 {
					sPath += "&"
				} else {
					sPath += "?"
				}
				sPath += url.QueryEscape(key) + "=" + url.QueryEscape(value)
			}
		}
	}

	// Remove directory index if present
	directoryIndexes := strings.Split(GetOptionString("statistic", "DIRECTORY_INDEX"), ",")
	for _, index := range directoryIndexes {
		index = "/" + strings.TrimSpace(index)
		if strings.HasSuffix(sPath, index) {
			sPath = sPath[:len(sPath)-len(index)+1]
			break
		}
	}

	return sPath + "?" + parsedURL.RawQuery
}

func main() {

}
