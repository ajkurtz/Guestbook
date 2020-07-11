package service

import (
	"encoding/json"
	"guestbook/domain"
	"io/ioutil"
	"net/http"
)

func APIPost(writer http.ResponseWriter, request *http.Request) {

	signature := domain.Signature{}

	body, err := ioutil.ReadAll(request.Body)
	checkError(err)

	json.Unmarshal(body, &signature)

	insert(signature.Signature)

	writer.WriteHeader(http.StatusCreated)

}
