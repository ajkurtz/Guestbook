package main

import (
	"guestbook/service"
	"log"
	"net/http"
)

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	service.View(writer)
}

func newHandler(writer http.ResponseWriter, request *http.Request) {
	service.New(writer)
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	service.Create(writer, request)
}

func apiHandler(writer http.ResponseWriter, request *http.Request) {

	switch request.Method {
	case "GET":
		service.APIGet(writer, request)
	case "POST":
		service.APIPost(writer, request)
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	http.HandleFunc("/guestbook/api/v1/signatures", apiHandler)

	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
