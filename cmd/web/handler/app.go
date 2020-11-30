package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"pdftools/cmd/web/helper"
)

// PageData pass different string to page
type PageData struct {
	Title   string
	Header  string
	BtnText string
}

// Home page
func Home(w http.ResponseWriter, r *http.Request) {
	Render(w, r, "./ui/html/home.html", "Hello Home Page")

}

// ImageToPDF is the main page of convert image to pdf
func ImageToPDF(w http.ResponseWriter, r *http.Request) {
	Render(w, r, "./ui/html/img2pdf.html", PageData{"Image to PDF", "图片转PDF", "转换"})
}

// PDFToImage with jpg,png etc...
func PDFToImage(w http.ResponseWriter, r *http.Request) {
	Render(w, r, "./ui/html/pdf2img.html", PageData{"PDF to Image", "PDF转图片", "转换"})
}

// Upload upload file(s) to server. Return a file name from server.
func Upload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		log.Println("ParseMultipartForm error:", err)
		return
	}
	files := r.MultipartForm.File["filepond"]
	action := r.PostFormValue("action") // merge|img2pdf...
	clockwise := r.PostFormValue("cw")
	counterclockwise := r.PostFormValue("ccw")
	var direction string
	if clockwise == "true" {
		direction = "cw"
	}
	if counterclockwise == "true" {
		direction = "ccw"
	}
	degrees := r.PostFormValue("degree")
	format := r.PostFormValue("format")
	log.Println("action:", action, "degrees:", degrees)
	out := helper.UploadFiles(files, action, direction, degrees, format)
	log.Println("The file from server is :", out)

	// upload + zip + download

	j, err := json.Marshal(out)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// MergePDF combine PDFs in the order you want
func MergePDF(w http.ResponseWriter, r *http.Request) {
	Render(w, r, "./ui/html/mergepdf.html", PageData{"Merge PDF", "合并PDF", "合并"})
}

func SplitPDF(w http.ResponseWriter, r *http.Request) {
	Render(w, r, "./ui/html/splitpdf.html", PageData{"Split PDF", "拆分PDF", "拆分"})
}

func RotatePDF(w http.ResponseWriter, r *http.Request) {
	Render(w, r, "./ui/html/rotatepdf.html", PageData{"Rotate PDF", "旋转PDF", "旋转"})
}
