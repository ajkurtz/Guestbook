package service

import (
	"net/http"
	"strings"
)

func Create(writer http.ResponseWriter, request *http.Request) {
	signature := strings.TrimSpace(request.FormValue("signature"))

	if signature != "" {
		insert(signature)
		http.Redirect(writer, request, "/guestbook", http.StatusFound)
	} else {
		http.Redirect(writer, request, "/guestbook/new", http.StatusFound)
	}

}
