package main

import (
	"fmt"

	"github.com/elgizabbasov/ezGmail"
)

var (
	subject string
	gs      ezGmail.GmailService
)

func extract() string {
	// InitSrv() uses client_secret.json to try to get a OAuth 2.0 token
	gs.InitSrv()

	// We compose a search statement with filter functions
	gs.InInbox().MaxResults(1).NewerThanRel("1d").Match("no-reply@uofcnotify.me").HasAttachment(false)

	// GetMessages() tries to execute the search statement and get a list of messages
	for _, ii := range gs.GetMessages() {
		fmt.Println("\nTrying to get the subject...")
		if ii.HasSubject() {
			subject = ii.GetSubject()
		}
	}

	return subject
}
