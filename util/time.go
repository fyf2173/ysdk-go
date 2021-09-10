package util

import (
	"fmt"
	"strconv"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

func GetRunningTime() string {
	return time.Now().Format(TimeFormat)
}

func StringToTime(timestr string) (time.Time, error) {
	return time.Parse(TimeFormat[:len(timestr)], timestr)
}

// ParseTimestr 解析时间
func ParseTimeStr(timeStr string) time.Time {
	t, _ := time.ParseInLocation(TimeFormat, timeStr, time.Local)
	return t
}

// GetDateGap 计算开始时间到结束时间相隔n年，n月，n周，n天
func GetDateGap(start, end int64) (totalYear, totalMonth, totalWeek, totalDay int) {
	timeStart := time.Unix(start, 0)
	timeEnd := time.Unix(end, 0)

	if timeEnd.Before(timeStart) {
		return 0, 0, 0, 0
	}

	sy, sw := timeStart.ISOWeek()
	ey, ew := timeEnd.ISOWeek()
	sd, ed := timeStart.YearDay(), timeEnd.YearDay()
	sm, em := timeStart.Format("1"), timeEnd.Format("01")

	y := ey - sy
	w := ew - sw
	d := ed - sd
	tsm, _ := strconv.Atoi(sm)
	tem, _ := strconv.Atoi(em)
	m := tem - tsm

	m += 12 * y
	for i := sy; i < ey; i++ {
		for j := 31; j > 24; j++ {
			t, _ := time.Parse("2006-01-02", fmt.Sprintf("%+v-12-%+v", i, j))
			if _, tw := t.ISOWeek(); tw != 1 {
				w += tw
				break
			}
		}
		if i%4 == 0 && (i%100 != 0 || i%400 == 0) {
			d += 366
		} else {
			d += 365
		}
	}

	return y, m, w, d
}
