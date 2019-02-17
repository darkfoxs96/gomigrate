package main

import (
	"fmt"
	"os"

	"github.com/darkfoxs96/gomigrate/template"
)

func main() {
	args := os.Args

	for _, arg := range args {
		if arg == "-help" {
			help()
			return
		}
	}

	if len(args) < 3 {
		fmt.Println("Error: need 2 args. Path to package and name record.")
		fmt.Println("gomigrate -help")
	}

	temp, path, fileName := template.GetTemplate(args[1], args[2])

	file, err := os.Create(path+"/"+fileName)
	if err != nil{
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	_, _ = file.WriteString(temp)
}

func help() {
	fmt.Println("Need 2 args. Path to package and name record.")
	fmt.Println("Example:")
	fmt.Println("gomigrate ./ init_db")
	fmt.Println("Generated: ./20190223_201930_init_db.go")
}
