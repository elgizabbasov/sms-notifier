package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

var (
	accountSid  string
	authToken   string
	fromPhone   string
	toPhone     string
	client      *twilio.RestClient
	subjectLine string
)

func SendMessage(message string) {
	params := openapi.CreateMessageParams{}
	params.SetTo(toPhone)
	params.SetFrom(fromPhone)
	params.SetBody(message)

	resp, err := client.ApiV2010.CreateMessage(&params)
	if err != nil {
		fmt.Printf("Error creating message: %s\n", err.Error())
		return
	}
	fmt.Printf("Message SID: %s\n", *resp.Sid)
}

func init() {
	enverr := godotenv.Load(".env")
	if enverr != nil {
		fmt.Printf("Error loading environment: %s\n", enverr.Error())
		os.Exit(1)
	}

	accountSid = os.Getenv("ACCOUNT_SID")
	authToken = os.Getenv("AUTH_TOKEN")
	fromPhone = os.Getenv("FROM_PHONE")
	toPhone = os.Getenv("TO_PHONE")

	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
}

func main() {
	fmt.Println("Welcome to Notifier!")

	subjectLine = extract()

	if subjectLine != "" {
		message := fmt.Sprintf(os.Getenv("MSG"), subjectLine)
		SendMessage(message)
	} else {
		fmt.Println("No need to send SMS!")
	}
}
