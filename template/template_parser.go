package template

import (
	"strconv"
	"strings"
	"time"

	"github.com/darkfoxs96/gomigrate/tool"
)

func GetTemplate(path, name string) (temp, absPath, fileName string) {
	timeNow := time.Now()
	timestamp := strconv.Itoa(int(timeNow.Unix()))
	_, timeLightString := tool.GetTime(timeNow)
	structName := strings.ToUpper(name[:1])+name[1:]+"_"+timeLightString
	fileName = timeLightString+"_"+name+".go"
	absPath, packageName := tool.GetAbsPathAndPackage(path)

	temp = strings.Replace(templateSource, "${packageName}", packageName, -1)
	temp = strings.Replace(temp, "${structName}", structName, -1)
	temp = strings.Replace(temp, "${timestamp}", timestamp, -1)
	temp = strings.Replace(temp, "${timeString}", timeLightString, -1)

	return temp, absPath, fileName
}
