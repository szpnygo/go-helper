package neo

import (
	"strconv"
	"time"
)

// StringInSlice ...
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

//20060102150405
func TimeStrToTime(format string, t string) (time.Time, error) {
	loc, _ := time.LoadLocation("PRC")
	return time.ParseInLocation("20060102150405", "20200114222001", loc)
}

func GetUnixTime() string {
	return strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
}
