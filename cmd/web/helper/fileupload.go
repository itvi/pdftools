package helper

import (
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"pdftools/cmd/web/tool"
	"time"
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
func Upload(files []*multipart.FileHeader, action string) (result string) {
	// ready for zip file
	var convertedFiles []string

	var uploadFiles []string

	for i := range files { // loop through the files one by one
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			log.Println("Open file error:", err)
			result = err.Error()
			return
		}

		out, err := os.Create("./upload/" + files[i].Filename)
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

		log.Printf("File %s uploaded successfully!", files[i].Filename)

		log.Println(filepath.Abs(files[i].Filename))

		if action == "img2pdf" {
			// convert image to pdf
			pdfFile := img2pdf(files[i].Filename)

			// put files to zip files slice
			convertedFiles = append(convertedFiles, pdfFile)
		}

		if action == "merge" {
			// files[i].Filename is a.jpg
			uploadFiles = append(uploadFiles, "./upload/"+files[i].Filename)
		}
	}

	if action == "merge" {
		if err := tool.MergePDF(uploadFiles); err != nil {
			log.Println(err)
		}
	}

	// switch action {
	// case "img2pdf": // conver image files to pdf
	// 	convertedFiles = imageToPDF(uploadFiles)
	// 	log.Println(convertedFiles)
	// case "merge":
	// 	if err := tool.MergePDF(uploadFiles); err != nil {
	// 		log.Println(err)
	// 	}
	// }

	// Zip files for download
	rand.Seed(int64(time.Now().UnixNano()))
	log.Println(rand.Int())
	log.Println(randString(10))

	randString := randString(10)
	zipFile := "./download/" + randString + ".zip"
	if err := ZipFiles(zipFile, convertedFiles); err != nil {
		log.Println("zip files error:", err)
		result = err.Error()
		return
	}
	log.Println(zipFile)

	return randString + ".zip"
}

func randString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
