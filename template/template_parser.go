package template

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func GetTemplate(path, name string) (temp, absPath, fileName string) {
	timeNow := time.Now()
	timestamp := strconv.Itoa(int(timeNow.Unix()))
	timeString, timeLightString := getTime(timeNow)
	structName := strings.ToUpper(name[:1])+name[1:]+"_"+timeLightString
	fileName = timeLightString+"_"+name+".go"
	absPath, packageName := getAbsPathAndPackage(path)

	temp = strings.Replace(templateSource, "${packageName}", packageName, -1)
	temp = strings.Replace(temp, "${structName}", structName, -1)
	temp = strings.Replace(temp, "${timestamp}", timestamp, -1)
	temp = strings.Replace(temp, "${timeString}", timeString, -1)
	temp = strings.Replace(temp, "${timeString}", timeString, -1)

	return temp, absPath, fileName
}

func getAbsPathAndPackage(path string) (absPath, packageName string) {
	var err error

	if path[len(path)-1:] == "/" {
		path = path[:len(path)-1]
	}

	if path[:1] != "." {
		absPath = path
		path = ""
	} else {
		path = path[1:]
		absPath, err = filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			panic(err)
		}
	}

	absPath += path
	i := strings.LastIndex(absPath, "/")
	packageName = absPath[i+1:]

	return
}

func getTime(timestamp time.Time) (string, string) {
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
