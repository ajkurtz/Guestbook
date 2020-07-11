package service

import (
	"guestbook/config"
	"net/http"
	"text/template"
)

func New(writer http.ResponseWriter) {

	html, err := template.ParseFiles(config.NewTemplateFilename)
	checkError(err)

	err = html.Execute(writer, nil)
	checkError(err)

}
