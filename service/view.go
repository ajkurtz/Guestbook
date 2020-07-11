package service

import (
	"guestbook/config"
	"guestbook/domain"
	"net/http"
	"text/template"
)

func View(writer http.ResponseWriter) {
	signatures := getStrings(config.DataFilename)

	guestbook := domain.ViewTemplate{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}

	html, err := template.ParseFiles(config.ViewTemplateFilename)
	checkError(err)

	err = html.Execute(writer, guestbook)
	checkError(err)

}
