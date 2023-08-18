package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mailgun/mailgun-go/v4"
	"os"
	"time"
)

var domain = "mg.adomate.ai"

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
	sender = "Adomate Mailman <mailman@mg.adomate.com>"
}

func SendEmail(body Email) (string, error) {
	message := mg.NewMessage(sender, body.Subject, "", body.To)
	message.SetTemplate(body.Template)

	variables := make(map[string]interface{})
	err := json.Unmarshal([]byte(body.Variables), &variables)
	if err != nil {
		return "", err
	}
	for key, value := range variables {
		err := message.AddVariable(key, value)
		if err != nil {
			return "", err
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, id, err := mg.Send(ctx, message)
	return id, err
}
