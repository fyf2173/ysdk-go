package util

import "time"

const TimeFormat = "2006-01-02 15:04:05"

func GetRunningTime() string {
	return time.Now().Format(TimeFormat)
}

func StringToTime(timestr string) (time.Time, error) {
	return time.Parse(TimeFormat[:len(timestr)], timestr)
}
