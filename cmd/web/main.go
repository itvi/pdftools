package main

import (
	"log"
	"net/http"
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
	// mux.HandleFunc("/upload", handler.Convert)
	mux.HandleFunc("/upload", handler.Upload)
	mux.HandleFunc("/mergepdf", handler.MergePDF) //
	mux.HandleFunc("/splitpdf", handler.SplitPDF)
	mux.HandleFunc("/rotatepdf", handler.RotatePDF)

	localIP := util.GetLocalIP()

	server := &http.Server{
		Addr:    ":12345",
		Handler: mux,
	}
	log.Println("Starting server on " + localIP + server.Addr)
	err := server.ListenAndServe()
	log.Fatal(err)
}
