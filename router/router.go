package router

import (
	"log"
	"net/http"
	"testpicx/control"
)

func Test() {
	log.Println("testmodel")
}

func Run() {
	http.HandleFunc("/", control.IndexView)
	http.HandleFunc("/upload", control.UploadView)
	http.ListenAndServe(":8080", nil)
}
