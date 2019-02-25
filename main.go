package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/darkfoxs96/gomigrate/maintemplate"
	"github.com/darkfoxs96/gomigrate/template"
)

func main() {
	args := os.Args

	for _, arg := range args {
		if arg == "-help" {
			help()
			return
		} else if arg == "-build" {
			if len(args) < 6 {
				fmt.Println("Need min 6 args. Path to package, Systems params, -up or -down or 'data', Connect params.")
				fmt.Println("gomigrate -help")
			}

			connect := ""
			for i, p := range args[7:] {
				if i == 0 {
					connect += p
				} else {
					connect += " " + p
				}
			}

			buildBin(args[2], args[3], args[4], args[5], args[6], connect)
			return
		}
	}

	if len(args) < 3 {
		fmt.Println("Error: need 2 args. Path to package and name record.")
		fmt.Println("gomigrate -help")
	}

	buildTmp(args[1], args[2])
}

func help() {
	fmt.Println("build migrate point")
	fmt.Println("Need 2 args. Path to package and name record.")
	fmt.Println("Example:")
	fmt.Println("gomigrate ./ init_db")
	fmt.Println("Generated: ./20190223_201930_init_db.go")

	fmt.Println("////////////////////")

	fmt.Println("build binary migration")
	fmt.Println("Need min 6 args. Path to package, Systems params, -up or -down or 'data', Connect params.")
	fmt.Println("Example:")
	fmt.Println("gomigrate -build ./ GOOS=darwin GOARCH=amd64 -up postgres user=don password=sdef12 dbname=don host=0.0.0.0 port=5432 sslmode=disable")
	fmt.Println("Generated: ./20190223_201930_up")
	fmt.Println("or")
	fmt.Println("gomigrate -build ./ GOOS=darwin GOARCH=amd64 20190216_133221 mysql don:sdef12@/don")
	fmt.Println("Generated: ./20190223_201930_to_20190223_133221")
}

func buildTmp(p, n string) {
	temp, path, fileName := template.GetTemplate(p, n)

	createFile(path+"/"+fileName, []byte(temp))
}

func buildBin(p, goos, goarch, to, db, connect string)  {
	temp, path, fileName, packageName := maintemplate.GetTemplate(p, to, db, connect)
	fileName += ""

	err := os.MkdirAll(path+"/"+fileName, os.ModePerm)
	if err != nil{
		fmt.Println("Unable to create file:")
		panic(err)
	}

	buildPackage(path, fileName, packageName)
	createFile(path+"/"+fileName+"/main.go", []byte(temp))

	command("sh", "-c", "cd "+path+"/"+fileName+" && "+"go get ./..."+" && "+goos+" "+goarch+" go build")

	fileNewName := fileName
	if goos == "GOOS=windows" {
		fileNewName += ".exe"
	}

	removeOtherFiles(path+"/"+fileName, fileNewName)
}

func command(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	var out bytes.Buffer

	cmd.Stdin = strings.NewReader("some input")
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Printf(out.String())
		panic(err)
	}

	fmt.Printf(out.String())
}

func buildPackage(p, fileName, packageName string)  {
	files, err := ioutil.ReadDir(p)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		name := file.Name()

		if !strings.Contains(name, ".go") {
			continue
		}

		d, err := ioutil.ReadFile(p+"/"+name)
		if err != nil {
			panic(err)
		}

		data := string(bytes.Replace(d, []byte("package "+packageName), []byte("package main"), 1))
		createFile(p+"/"+fileName+"/"+name, []byte(data))
	}
}

func removeOtherFiles(p, fileName string) {
	files, err := ioutil.ReadDir(p)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		name := file.Name()

		if name == fileName {
			continue
		}

		err = os.Remove(p+"/"+name)
		if err != nil {
			panic(err)
		}
	}
}

func createFile(path string, data []byte)  {
	file, err := os.Create(path)
	if err != nil{
		fmt.Println("Unable to create file:")
		panic(err)
	}

	_, err = file.Write(data)
	_ = file.Close()
	if err != nil {
		fmt.Println("Unable to write file:")
		panic(err)
	}
}
