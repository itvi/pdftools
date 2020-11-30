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
		log.Printf("File %s convert image to pdf error:%s", img, err)
		return ""
	}
	log.Printf("File %s convert successfully!", imgFile)
	return pdf
}

func imageToPDF(files []string) (out []string, err error) {
	// log.Println("files is:", files)
	fileVars := strings.Join(files, " ")

	//	mogrify -format pdf -- a.jpg c.png
	plainCmd := "mogrify -format pdf -- " + fileVars
	// log.Println("plaincmd:", plainCmd)
	sliceCmd := strings.Fields(plainCmd)
	cmd := exec.Command(sliceCmd[0], sliceCmd[1:]...)
	if err := cmd.Run(); err != nil {
		log.Println("cmd run error:", err)
		return nil, err
	}

	/// a.jpg b.jpg => a.pdf b.pdf
	var pdfFiles []string
	for _, file := range files {
		pdfFile := "." + strings.Split(file, ".")[1] + ".pdf"
		pdfFiles = append(pdfFiles, pdfFile)
	}
	return pdfFiles, err
}

// pdfToImage can convert to many type of image(jpg,png...)
func pdfToImage(files []string, format string) (out []string, err error) {
	fileVars := strings.Join(files, " ")
	plainCmd := "mogrify -format " + format + " -- " + fileVars
	sliceCmd := strings.Fields(plainCmd)
	cmd := exec.Command(sliceCmd[0], sliceCmd[1:]...)
	if err := cmd.Run(); err != nil {
		log.Println("cmd run error:", err)
		return nil, err
	}

	var imgFiles []string
	for _, file := range files {
		imgFile := "." + strings.Split(file, ".")[1] + "." + format
		imgFiles = append(imgFiles, imgFile)
	}

	return imgFiles, err
}
