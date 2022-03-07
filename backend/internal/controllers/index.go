package controllers

import (
	"log"
	"net/http"
)

func processError(w http.ResponseWriter, code int, message *string) {
	w.WriteHeader(code)

	totalMessage := http.StatusText(code)
	if message != nil {
		totalMessage += ": " + *message
	}

	if _, err := w.Write([]byte(totalMessage)); err != nil {
		log.Printf("[processError][1]: %+v", err)
	}
}
