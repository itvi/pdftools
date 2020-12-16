package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"pdftools/cmd/web/handler"
	"pdftools/cmd/web/util"
)

func main() {

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static"))))
	mux.Handle("/upload/", http.StripPrefix("/upload/", http.FileServer(http.Dir("./upload"))))
	mux.Handle("/download/", http.StripPrefix("/download/", http.FileServer(http.Dir("./download"))))

	mux.HandleFunc("/", handler.ImageToPDF)
	mux.HandleFunc("/img2pdf", handler.ImageToPDF)
	mux.HandleFunc("/pdf2img", handler.PDFToImage)
	mux.HandleFunc("/upload", handler.Upload)
	mux.HandleFunc("/mergepdf", handler.MergePDF) //
	mux.HandleFunc("/splitpdf", handler.SplitPDF)
	mux.HandleFunc("/rotatepdf", handler.RotatePDF)
	mux.HandleFunc("/dl", dl)

	localIP := util.GetLocalIP()

	server := &http.Server{
		Addr:    ":12345",
		Handler: mux,
	}
	log.Println("Starting server on " + localIP + server.Addr)
	err := server.ListenAndServe()
	log.Fatal(err)
}

func dl(w http.ResponseWriter, r *http.Request) {
	x("./download")
	x("./upload")
}

func x(dirName string) {
	dir, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Println(err)
	}
	for _, f := range dir {
		err = os.RemoveAll(path.Join([]string{dirName, f.Name()}...))
		if err != nil {
			log.Println(err)
		}
	}
}
