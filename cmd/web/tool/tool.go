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

func splitPDF(file, pageNums string) (out []string, err error) {
	// file is like: ./upload/rMqpsgoOCbBook.pdf

	var pages []string
	var plainCmd string
	randFileName := util.RandString(10)

	if pageNums == "" {
		plainCmd := "qpdf " + file + " ./upload/" + randFileName + ".pdf --split-pages"
		ps, err := splitPages(plainCmd, randFileName)
		if err != nil {
			log.Println("Split pages error:", err)
			return nil, err
		}
		pages = ps
	} else {
		plainCmd = "qpdf " + "--empty " + "--pages " + file + " " + pageNums + " -- " + "./upload/" + randFileName + ".pdf"
		// split page range will return a single pdf file.
		sliceCmd := strings.Fields(plainCmd)
		cmd := exec.Command(sliceCmd[0], sliceCmd[1:]...)
		if err := cmd.Run(); err != nil {
			log.Println("cmd run error:", err)
			return nil, err
		}

		// split again
		newName := util.RandString(10)
		plainCmd = "qpdf " + "./upload/" + randFileName + ".pdf" + " ./upload/" + newName + ".pdf --split-pages"

		ps, err := splitPages(plainCmd, newName)
		if err != nil {
			log.Println("Split pages error:", err)
			return nil, err
		}
		pages = ps
	}

	return pages, err
}
func splitPages(plainCmd, randFileName string) (out []string, err error) {
	sliceCmd := strings.Fields(plainCmd)
	cmd := exec.Command(sliceCmd[0], sliceCmd[1:]...)
	if err := cmd.Run(); err != nil {
		log.Println("cmd run error:", err)
		return nil, err
	}

	dir, err := ioutil.ReadDir("./upload")
	if err != nil {
		log.Println("read dir error:", err)
		return nil, err
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
