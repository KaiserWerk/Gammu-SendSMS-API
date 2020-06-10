package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func getHeaderIfSet(r *http.Request, key string) (string, error) {
	header := r.Header.Get(key)
	if header == "" {
		return "", errors.New("header is not set or empty")
	}
	return header, nil
}

func sendSMS(m string, r string) {
	cmd := exec.Command("timeout", "10", "gammu", "-c", "/opt/.gammurc", "sendsms", "TEXT", r, "-text", m)
	o, err := cmd.Output()
	if err != nil {
		fmt.Println("could not send sms; aborting: " + err.Error())
		return
	}
	c := string(o)
	if strings.Contains(c, "Sending SMS") {
		fmt.Println("SMS sent!")
	} else {
		fmt.Println("error sending sms")
	}
}

func setupDB() error {
	dbFile := "./sendsms-api.db"
	// if db file does not exist, create it and import schema
	if _, err := os.Stat(dbFile); err != nil || os.IsNotExist(err) {
		file, err := os.Create(dbFile)
		if err != nil {
			return err
		}
		defer file.Close()

		db, err := sql.Open("sqlite3", dbFile)
		if err != nil {
			return err
		}
		defer db.Close()

		_, err = db.Exec(schemaString)
		if err != nil {
			return err
		}
	}

	return nil
}
