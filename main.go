package wifi

import (
	"runtime"

	"github.com/play175/wifiNotifier"
	wifiname "github.com/yelinaung/wifi-name"
)

func GetSSID() string {
	var wifi string
	if runtime.GOOS == "windows" {
		wifi = getSSIDForWindows()
	} else {
		wifi = getSSIDForMacos()
	}
	return wifi
}

//for windows
func getSSIDForWindows() string {
	return wifiNotifier.GetCurrentSSID()
}

//for macos
func getSSIDForMacos() string {
	return wifiname.WifiName()
}
