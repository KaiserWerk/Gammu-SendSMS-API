package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

func main() {


	// setup DB


	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			input, _, err := reader.ReadLine()
			if err != nil {
				fmt.Printf("could not process input %v\n", input)
			}
			check := string(input)
			if check == "help" {

			} else if check == "list tokens" {

			} else if check == "generate token" {

			} else if strings.Contains(check, "remove token") {
				parts := strings.Split(check, " ")
				if len(parts) != 3 {
					fmt.Println("Usage: remove token <id>")
				} else {
					// sql remove by id
				}
			}
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleHomePage)
	mux.HandleFunc("/send_sms", handleSendSMS)

	srv := &http.Server{
		ReadTimeout: 2 * time.Second,
		WriteTimeout: 4 * time.Second,
		IdleTimeout: 20 * time.Second,
		Handler: mux,
		Addr: ":5050",
	}

	done := make(chan bool)
	notify := make(chan os.Signal)

	signal.Notify(notify, os.Interrupt)

	go func() {
		<-notify
		ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		err := srv.Shutdown(ctx)
		if err != nil {
			panic("Could not gracefully shut down server: " + err.Error())
		}
		close(done)
	}()

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server could not be started")
	}
	<-done
	fmt.Println("Server shutdown complete")
}