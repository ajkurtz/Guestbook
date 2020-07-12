package main

import (
	"guestbook/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/guestbook", viewHandler).Methods("GET")
	r.HandleFunc("/guestbook/new", newHandler).Methods("GET")
	r.HandleFunc("/guestbook/create", createHandler).Methods("POST")
	r.HandleFunc("/guestbook/api/v1/signatures", apiGetHandler).Methods("GET")
	r.HandleFunc("/guestbook/api/v1/signatures", apiPostHandler).Methods("POST")
	http.Handle("/", r)

	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	service.View(writer)
}

func newHandler(writer http.ResponseWriter, request *http.Request) {
	service.New(writer)
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	service.Create(writer, request)
}

func apiGetHandler(writer http.ResponseWriter, request *http.Request) {
	service.APIGet(writer, request)
}

func apiPostHandler(writer http.ResponseWriter, request *http.Request) {
	service.APIPost(writer, request)
}
