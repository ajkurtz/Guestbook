package service

import (
	"encoding/json"
	"fmt"
	"guestbook/config"
	"guestbook/domain"
	"net/http"
)

func APIGet(writer http.ResponseWriter, request *http.Request) {

	signatures := []domain.Signature{}

	for _, s := range getStrings(config.DataFilename) {
		signatures = append(signatures, domain.Signature{Signature: s})
	}

	guestbook := domain.Guestbook{
		Signatures: signatures,
	}

	output, err := json.Marshal(&guestbook)
	checkError(err)

	writer.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(writer, string(output))

}
