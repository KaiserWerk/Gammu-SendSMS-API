package main

import (
	"fmt"
	"net/http"
)

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
