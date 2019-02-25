package tool

import (
	"os"
	"path/filepath"
	"strings"
)

func GetAbsPathAndPackage(path string) (absPath, packageName string) {
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
