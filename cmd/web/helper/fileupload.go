package helper

import (
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
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
func Upload(files []*multipart.FileHeader) (result string) {
	// ready for zip file
	var convertedFiles []string

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

		// convert image to pdf
		pdfFile := img2pdf(files[i].Filename)

		// put files to zip files slice
		convertedFiles = append(convertedFiles, pdfFile)
	}

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
