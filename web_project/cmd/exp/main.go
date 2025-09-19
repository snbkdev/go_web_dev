package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"web_project/models"

	"github.com/joho/godotenv"
)

const (
	host     = "***"
	port     = 1234
	username = "***"
	password = "***"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("SMTP_HOST")
	portStr := os.Getenv("SMTP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(err)
	}
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	es := models.NewEmailService(models.SMTPConfig{
		Host: host,
		Port: port,
		Username: username,
		Password: password,
	})
	err = es.ForgotPassword("***", "***")
	fmt.Println("Message sent")
}
