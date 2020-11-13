package helper

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

// FileUpload upload single file
func FileUpload(r *http.Request, inputName string) (string, error) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile(inputName) // retrieve the file from  form data
	if err != nil {
		return "", err
	}
	defer file.Close()

	f, err := os.OpenFile("./upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()

	io.Copy(f, file)

	return handler.Filename, nil
}

// Upload can upload many files
func Upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20) // 32MB is the default used by FormFile
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	files := r.MultipartForm.File["filepond"]

	for i := range files { // loop through the files one by one
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			log.Println("Open file error:", err)
			return
		}

		out, err := os.Create("./upload/" + files[i].Filename)
		defer out.Close()
		if err != nil {
			log.Println("Unable to create the file for writing. Check your write access privilege")
			return
		}

		_, err = io.Copy(out, file) // file not files[i] !

		if err != nil {
			log.Println("io copy file error:", err)
			return
		}

		log.Printf("File %s uploaded successfully!", files[i].Filename)

		// convert image to pdf
		imgFile := files[i].Filename
		fileName := strings.Split(imgFile, ".")[0]
		pdfFile := fileName + ".pdf"
		from := "./upload/" + imgFile
		to := "./upload/" + pdfFile

		if err = convert(from, to); err != nil {
			log.Println("convert error:", err)
			return
		}
		log.Printf("File %s convert to %s successfully!", imgFile, pdfFile)
	}
}

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
