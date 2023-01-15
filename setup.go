package main

import (
	"context"
	"github.com/mailgun/mailgun-go/v4"
	"os"
	"time"
)

var domain string = "adomate.com"

var private_key string = os.Getenv("EMAIL_PRIVATE_KEY")

var mg *mailgun.MailgunImpl
var sender string

func Setup() {
	mg = mailgun.NewMailgun(domain, private_key)
	sender = "no-reply@adomate.com"
}

func SendEmail(to string, subject string, body string) (string, string, error) {
	message := mg.NewMessage(sender, subject, body, to)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, id, err := mg.Send(ctx, message)
	if err != nil {
		return "", "", err
	}

	return resp, id, nil
}
