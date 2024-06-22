package utils

import "time"

func GenerateDate(date time.Time, day int) time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.Date(date.Year(), date.Month(), day, 0, 0, 0, 0, loc)
}
