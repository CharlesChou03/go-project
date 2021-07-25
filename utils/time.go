package utils

import "time"

// response format: YYYYMMDDhhmmss
func GetLocationTime(name string, format string) string {
	loc, err := time.LoadLocation(name)
	if err != nil {
		loc, _ = time.LoadLocation("UTC")
	}
	scanTime := time.Now().In(loc).Format(format)
	return scanTime
}

func GetCurrentMillisecondTimestamp() int64 {
	return time.Now().Unix() * 1000
}
