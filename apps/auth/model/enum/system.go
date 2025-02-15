package enum

import "strings"

type ClientType string
type DeviceType string

const (
	ClientTypeWeb     ClientType = "WEB"
	ClientTypeApp     ClientType = "APP"
	ClientTypeBot     ClientType = "BOT"
	DeviceTypeDesktop DeviceType = "DESKTOP"
	DeviceTypeMobile  DeviceType = "MOBILE"
)

var appKeywords = []string{"MyApp", "MobileApp", "AndroidApp", "iOSApp"}

func IsApp(ua string) bool {
	for _, v := range appKeywords {
		if strings.Contains(strings.ToLower(ua), strings.ToLower(v)) {
			return true
		}
	}
	return false
}

func IsBot(ua string) bool {
	ua = strings.ToLower(ua)
	return strings.Contains(ua, "bot") || strings.Contains(ua, "spider")
}
