package helper

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// FileUpload ...
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

// Upload files
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
			fmt.Fprintln(w, err)
			return
		}

		out, err := os.Create("./upload/" + files[i].Filename)
		defer out.Close()
		if err != nil {
			fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
			return
		}

		_, err = io.Copy(out, file) // file not files[i] !

		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		fmt.Fprintf(w, "Files uploaded successfully : ")
		fmt.Fprintf(w, files[i].Filename+"\n")

	}
}
