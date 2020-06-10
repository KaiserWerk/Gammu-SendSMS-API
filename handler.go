package main

import (
	"errors"
	"fmt"
	"net/http"
)

func getHeaderIfSet(r *http.Request, key string) (string, error) {
	header := r.Header.Get(key)
	if header == "" {
		return "", errors.New("header is not set or empty")
	}
	return header, nil
}

// For teh lulz
func handleHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sorry Mario, the pricess is in another castle...")
}

func handleSendSMS(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		return
	}


}
