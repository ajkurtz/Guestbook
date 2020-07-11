package service

import (
	"fmt"
	"guestbook/config"
	"os"
)

func insert(signature string) {

	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile(config.DataFilename, options, os.FileMode(0600))
	checkError(err)

	_, err = fmt.Fprintln(file, signature)
	checkError(err)

	err = file.Close()
	checkError(err)

}
