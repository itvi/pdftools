package tool

import (
	"log"
	"os/exec"
	"strings"
)

func MergePDF(files []string) error {
	fileNames := strings.Join(files, " ")
	log.Println("merge pdf:", fileNames)
	app := "pdftk"
	//arg0 := strings.Trim(fileNames, "[]")
	arg0 := "./upload/a.pdf ./upload/c.pdf" // TODO:??????
	arg1 := "cat output"
	arg2 := "out.pdf"

	err := exec.Command(app, arg0, arg1, arg2).Run()
	if err != nil {
		log.Println(err)
	}
	return err
}
