package tool

import (
	"strconv"
	"time"
)

func GetTime(timestamp time.Time) (string, string) {
	var dayStr string
	var yourStr string
	var monthStr string
	var hourStr string
	var minStr string
	var secStr string

	your := timestamp.Year()
	yourStr = strconv.Itoa(your)

	day := timestamp.Day()
	if day < 10 {
		dayStr = "0" + strconv.Itoa(day)
	} else {
		dayStr = strconv.Itoa(day)
	}

	month := int(timestamp.Month())
	if month < 10 {
		monthStr = "0" + strconv.Itoa(month)
	} else {
		monthStr = strconv.Itoa(month)
	}

	hour := timestamp.Hour()
	if hour < 10 {
		hourStr = "0" + strconv.Itoa(hour)
	} else {
		hourStr = strconv.Itoa(hour)
	}

	min := timestamp.Minute()
	if min < 10 {
		minStr = "0" + strconv.Itoa(min)
	} else {
		minStr = strconv.Itoa(min)
	}

	sec := timestamp.Second()
	if sec < 10 {
		secStr = "0" + strconv.Itoa(sec)
	} else {
		secStr = strconv.Itoa(sec)
	}

	return yourStr+"."+monthStr+"."+dayStr+"_"+hourStr+":"+minStr+":"+secStr,
		yourStr+monthStr+dayStr+"_"+hourStr+minStr+secStr
}
