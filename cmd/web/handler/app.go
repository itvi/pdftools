package handler

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"pdftools/cmd/web/helper"
	"strings"
)

// PageData pass different string to page
type PageData struct {
	Title  string
	Header string
}

// Home page
func Home(w http.ResponseWriter, r *http.Request) {
	helper.Render(w, r, "./ui/html/home.html", "Hello Home Page")

}

// ImageToPDF is the main page of convert image to pdf
func ImageToPDF(w http.ResponseWriter, r *http.Request) {
	helper.Render(w, r, "./ui/html/img2pdf.html", PageData{"Image to PDF", "图片转PDF"})
}

// Upload upload file(s) to server
func Upload(w http.ResponseWriter, r *http.Request) {
	helper.Upload(w, r)
}

// Convert upload image file and convert to pdf file then download
func Convert(w http.ResponseWriter, r *http.Request) {
	fname, err := helper.FileUpload(r, "filepond")
	if err != nil {
		log.Println(err)
		return
	}

	from := "./upload/" + fname
	name := strings.Split(fname, ".")[0]
	to := "./upload/" + name + ".pdf"
	if err = convert(from, to); err != nil {
		log.Println("convert error:", err)
		return
	}
}

// Convert handler
// func Convert(w http.ResponseWriter, r *http.Request) {
// 	//convert("a.jpg", "a.pdf")
// }

// Remark: icon file not .ico
func convert(from, to string) error {
	// Windows use "cmd /c"
	// app := "cmd"
	// arg0 := "/c"
	arg1 := "magick convert"
	arg2 := from
	arg3 := to
	err := exec.Command(arg1, arg2, arg3).Run()
	if err != nil {
		log.Println(err)
	}
	return err
}

// MergePDF combine PDFs in the order you want
func MergePDF(w http.ResponseWriter, r *http.Request) {
	helper.Render(w, r, "./ui/html/mergepdf.html", PageData{"Merge PDF", "合并PDF"})
}

func mergePDF(files []string) error {
	fileNames := fmt.Sprint(files)

	app := "pdftk"
	arg0 := strings.Trim(fileNames, "[]")
	arg1 := "cat output"
	arg2 := "out.pdf1"

	err := exec.Command(app, arg0, arg1, arg2).Run()
	if err != nil {
		log.Println(err)
	}
	return err
}
