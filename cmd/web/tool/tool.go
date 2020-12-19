package tool

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os/exec"
	"pdftools/cmd/web/util"
	"strings"
	"time"
)

func mergePDF(files []string) (out string, err error) {
	// log.Println("files is:", files) // [./upload/1.pdf ./upload/5.pdf]
	fileVars := strings.Join(files, " ")
	// log.Println("file vars:", fileVars) // ./upload/1.pdf ./upload/5.pdf

	// ready for out combined file
	rand.Seed(int64(time.Now().UnixNano()))
	out = "./upload/" + util.RandString(10) + ".pdf"

	// cpdf -merge ./upload/1.pdf ./upload/5.pdf -o out12.pdf
	//plainCmd := "cpdf -merge " + fileVars + " -o " + out

	// qpdf --empty --pages 1.pdf 3.pdf -- 13.pdf
	plainCmd := "qpdf --empty --pages " + fileVars + " -- " + out

	// log.Println("plaincmd:", plainCmd)

	sliceA := strings.Fields(plainCmd)
	cmd := exec.Command(sliceA[0], sliceA[1:]...)

	err = cmd.Run()
	if err != nil {
		log.Println("cmd run error:", err)
		return "", err
	}

	return out, err
}

func splitPDF(file string) (out []string, err error) {
	// log.Println("get file(with .pdf) :", file) // ./upload/rMqpsgoOCbC#编程风格.pdf
	randFileName := util.RandString(10)
	plainCmd := "qpdf " + file + " ./upload/" + randFileName + ".pdf --split-pages"
	// log.Println(plainCmd)
	sliceCmd := strings.Fields(plainCmd)
	cmd := exec.Command(sliceCmd[0], sliceCmd[1:]...)
	if err := cmd.Run(); err != nil {
		log.Println("cmd run error:", err)
		return nil, err
	}

	dir, err := ioutil.ReadDir("./upload")
	if err != nil {
		log.Println("read dir error:", err)
		return
	}
	// x-01.pdf x-02.pdf ...
	var pages []string
	for _, f := range dir {
		name := strings.Split(f.Name(), "-")[0]
		if strings.Contains(randFileName, name) {
			pages = append(pages, "./upload/"+f.Name())
		}
	}

	return pages, err
}

func rotatePDF(file, direction, degree string) (out string, err error) {
	// qpdf in.pdf out.pdf --rotate=+90
	outFile := "./upload/" + util.RandString(10) + ".pdf"
	var operator string
	if direction == "cw" {
		operator = "+"
	}
	if direction == "ccw" {
		operator = "-"
	}
	plainCmd := "qpdf " + file + " " + outFile + " --rotate=" + operator + degree
	// log.Println("plain:", plainCmd)
	sliceCmd := strings.Fields(plainCmd)
	cmd := exec.Command(sliceCmd[0], sliceCmd[1:]...)
	if err := cmd.Run(); err != nil {
		log.Println("cmd run error:", err)
		return "", err
	}
	return outFile, nil
}
