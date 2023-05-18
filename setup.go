package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mailgun/mailgun-go/v4"
	"os"
	"time"
)

var domain string = "mg.adomate.ai"

var mg *mailgun.MailgunImpl
var sender string

type RabbitMQConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Queue    string
}

var RMQConfig RabbitMQConfig

func Setup() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	privateKey := os.Getenv("API_KEY")

	RMQConfig = RabbitMQConfig{
		Host:     os.Getenv("RABBIT_HOST"),
		Port:     os.Getenv("RABBIT_PORT"),
		User:     os.Getenv("RABBIT_USER"),
		Password: os.Getenv("RABBIT_PASS"),
		Queue:    os.Getenv("RABBIT_MAIL_QUEUE"),
	}

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
