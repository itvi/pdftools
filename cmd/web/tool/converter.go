package tool

import (
	"io/ioutil"
	"log"
	"os/exec"
	"pdftools/cmd/web/util"
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

func imageToPDF(files []string, combine bool) (out []string, err error) {
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

	// combine images to single pdf file (convert *.png output.pdf)
	if combine {
		// merge pdf files
		out, err := mergePDF(pdfFiles)
		if err != nil {
			log.Println("merger error:", err)
			return nil, err
		}
		pdfFiles = []string{out}
	}
	return pdfFiles, err
}

// pdfToImage can convert to many type of image(jpg,png...)
func pdfToImage(files []string, format string, combine bool) (out []string, err error) {
	var plainCmd string
	var imgFiles []string
	fileVars := strings.Join(files, " ")

	// convert multiple pdfs to a single image file
	if combine {
		imgFile := util.RandString(10) + "." + format
		plainCmd = "convert " + fileVars + " -append " + imgFile
		log.Println(plainCmd)
		sliceCmd := strings.Fields(plainCmd)
		cmd := exec.Command(sliceCmd[0], sliceCmd[1:]...)
		if err := cmd.Run(); err != nil {
			log.Println("cmd run error:", err)
			return nil, err
		}
		imgFiles = []string{imgFile}
	}

	if !combine {
		plainCmd = "mogrify -format " + format + " -- " + fileVars
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

		for _, f := range dir {
			nameSlice := strings.Split(f.Name(), ".") // [file-1 pdf]
			name := nameSlice[0]
			ext := nameSlice[1]
			for _, pf := range files {
				pdfName := strings.Split(pf, "./upload/")[1] // file.pdf
				nameSlice := strings.Split(pdfName, ".")
				fileNameWithoutExt := nameSlice[0] // file

				if strings.Contains(name, fileNameWithoutExt) && ext == format {
					imgFiles = append(imgFiles, "./upload/"+f.Name())
				}
			}
		}
	}

	return imgFiles, err
}
