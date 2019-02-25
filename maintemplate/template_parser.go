package maintemplate

import (
	"strings"
	"time"

	"github.com/darkfoxs96/gomigrate/tool"
)

func GetTemplate(path, to, driver, connect string) (temp, absPath, fileName, packageName string) {
	toStr := ""
	migrateFunc := ""

	if to == "-up" {
		toStr = "up"
		migrateFunc = "MigrateUp(db)"
	} else if to == "-down" {
		migrateFunc = "MigrateDown(db)"
		toStr = "down"
	} else {
		migrateFunc = `MigrateTo("`+to+`", db)`
		toStr = "to_" + to
	}

	timeNow := time.Now()
	_, timeLightString := tool.GetTime(timeNow)
	fileName = timeLightString+"_"+toStr
	absPath, packageName = tool.GetAbsPathAndPackage(path)
	driverImport := getDriverImport(driver)

	temp = strings.Replace(templateSource, "${packageName}", "main", -1)
	temp = strings.Replace(temp, "${driverImport}", driverImport, -1)
	temp = strings.Replace(temp, "${driverName}", driver, -1)
	temp = strings.Replace(temp, "${connectParams}", connect, -1)
	temp = strings.Replace(temp, "${migrateFunc}", migrateFunc, -1)

	return
}

func getDriverImport(driver string) (driverImport string) {
	if driver == "postgres" {
		driverImport = "github.com/lib/pq"
	}

	return
}
