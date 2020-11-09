package main

import (
	"log"
	"net/http"
	"pdftools/cmd/web/handler"
	"pdftools/cmd/web/helper"
)

func main() {

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static"))))
	mux.Handle("/upload/", http.StripPrefix("/upload/", http.FileServer(http.Dir("./upload"))))
	mux.HandleFunc("/", handler.ImageToPDF)
	mux.HandleFunc("/img2pdf", handler.ImageToPDF)
	// mux.HandleFunc("/upload", handler.Convert)
	mux.HandleFunc("/upload", handler.Upload)
	mux.HandleFunc("/mergepdf", handler.MergePDF) //

	localIP := helper.GetLocalIP()
	log.Println("Starting server on " + localIP + ":1234")
	err := http.ListenAndServe(":1234", mux)
	log.Fatal(err)
}
