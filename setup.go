package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/mailgun/mailgun-go/v4"
	"log"
	"os"
	"time"
)

var domain string = "mg.adomate.ai"

var mg *mailgun.MailgunImpl
var sender string

func Setup() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	privateKey := os.Getenv("API_KEY")

	mg = mailgun.NewMailgun(domain, privateKey)
	sender = "Adomate Robot <no-reply@mg.adomate.com>"
}

func SendEmail(to string, subject string, body string) (string, error) {
	message := mg.NewMessage(sender, subject, "", to)
	message.SetHtml(body)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, id, err := mg.Send(ctx, message)
	return id, err
}
