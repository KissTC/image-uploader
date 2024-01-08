package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-mail/mail/v2"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("SMTP_HOST")
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		panic(err)
	}
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	from := "julio.code@gmail.com"
	to := "jc@kissi.dev"
	subject := "this is a test"
	plainText := "this is the body of the email"
	html := `<h1>hello bro!</h1>
	<p>this is a test</p>`

	msg := mail.NewMessage()
	msg.SetHeader("To", to)
	msg.SetHeader("From", from)
	msg.SetHeader("subject", subject)
	msg.SetBody("text/plain", plainText)
	msg.AddAlternative("text/html", html)

	// para conectarse con smtp
	dialer := mail.NewDialer(host, port, username, password)
	err = dialer.DialAndSend(msg)
	if err != nil {
		panic(err)
	}

	fmt.Println("message sent")
}
