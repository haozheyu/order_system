package util

import "time"


func GetTimeNow() (format string) {
	now := time.Now()

	format = now.Format("2006-01-02 15:04:05")
	return
}
