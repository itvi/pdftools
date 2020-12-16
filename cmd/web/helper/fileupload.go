package helper

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"pdftools/cmd/web/util"
)

// UploadFile upload single file
func UploadFile(r *http.Request, inputName string) (string, error) {
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

// UploadFiles upload multiple files
func UploadFiles(files []*multipart.FileHeader, action, direction, degree, format string, combine, pdf2oneimg bool) (result string) {
	var myFiles []string

	// loop through the files one by one
	for i := range files {
		file, err := files[i].Open()
		fileName := files[i].Filename
		defer file.Close()
		if err != nil {
			log.Println("Open file error:", err)
			result = err.Error()
			return
		}

		//out, err := os.Create("./upload/" + files[i].Filename)
		uploadedFile := "./upload/" + util.RandString(10) + fileName
		out, err := os.Create(uploadedFile)
		defer out.Close()
		if err != nil {
			log.Println("Unable to create the file for writing. Check your write access privilege")
			result = err.Error()
			return
		}

		_, err = io.Copy(out, file) // file not files[i] !

		if err != nil {
			log.Println("io copy file error:", err)
			result = err.Error()
			return
		}

		log.Printf("File %s uploaded successfully!", fileName)

		// // image to pdf
		// if action == "img2pdf" {
		// 	// // convert image to pdf
		// 	// pdfFile := img2pdf(files[i].Filename)

		// 	// // put files to zip files slice
		// 	// myFiles = append(myFiles, pdfFile)

		// 	// collect all image files
		// 	imgFiles = append(imgFiles, imgFile)
		// }

		// if action == "merge" {
		// 	// files[i].Filename is a.jpg
		// 	allFiles = append(allFiles, "./upload/"+files[i].Filename)
		// }

		myFiles = append(myFiles, uploadedFile)
	}

	switch action {
	case "img2pdf":
		out, err := imageToPDF(myFiles, combine)
		if err != nil {
			log.Println("image to pdf error:", err)
			return
		}
		myFiles = out

	case "pdf2img":
		out, err := pdfToImage(myFiles, format, pdf2oneimg)
		if err != nil {
			log.Println("pdf to image error:", err)
			return
		}
		myFiles = out

	case "merge":
		out, err := mergePDF(myFiles)
		// log.Println("out is:", out)
		if err != nil {
			log.Println("merger error:", err)
			return
		}
		// combined pdf file is a single file
		myFiles = []string{out}

	case "split":
		out, err := splitPDF(myFiles[0])
		// log.Println("out is:", out)
		if err != nil {
			log.Println("split error:", err)
			return
		}
		myFiles = out

	case "rotate":
		out, err := rotatePDF(myFiles[0], direction, degree)
		// log.Println("out is:", out)
		if err != nil {
			log.Println("rotate error:", err)
			return
		}
		myFiles = []string{out}
	}

	// Zip files for download
	randString := util.RandString(10)
	zipFile := "./download/" + randString + ".zip"
	if err := util.ZipFiles(zipFile, myFiles); err != nil {
		log.Println("zip files error:", err)
		result = err.Error()
		return
	}
	//	log.Println(zipFile)

	return randString + ".zip"
}
