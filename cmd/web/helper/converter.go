package helper

import (
	"log"
	"os/exec"
	"strings"
)

// convert from type to type
// Remark: icon file not .ico
func convert(from, to string) error {
	// Windows use "cmd /c magick convert from to"
	// app := "cmd"
	// arg0 := "/c"
	app := "convert"
	arg1 := from
	arg2 := to
	err := exec.Command(app, arg1, arg2).Run()
	if err != nil {
		log.Println(err)
	}
	return err
}

// from a.jpg to a.pdf
func img2pdf(imgFile string) string {
	fileName := strings.Split(imgFile, ".")[0]

	img := "./upload/" + imgFile
	pdf := "./upload/" + fileName + ".pdf"

	if err := convert(img, pdf); err != nil {
		log.Println("convert image to pdf error:", err)
		return ""
	}
	log.Println(imgFile, "convert successfully!")
	return pdf
}

func imageToPDF(files []string) []string {
	log.Println("files is:", files)
	//	mogrify -format pdf -- a.jpg c.png
	app := "mogrify"
	arg1 := "-format pdf --"
	arg2 := strings.Join(files, " ")
	log.Println("arg2:", arg2)
	err := exec.Command(app, arg1, arg2).Run()
	if err != nil {
		log.Println(err)
	}

	// a.jpg b.jpg => a.pdf b.pdf
	var pdfFiles []string
	for _, file := range files {
		pdfFile := "." + strings.Split(file, ".")[1] + ".pdf"
		pdfFiles = append(pdfFiles, pdfFile)
	}
	return pdfFiles
}
