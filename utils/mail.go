package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendResetEmail(to, token string) error {
	from := os.Getenv("SMTP_EMAIL")
	pass := os.Getenv("SMTP_PASSWORD")

	msg := "Subject: Reset Your Password\n\n" +
		"Click the link to reset your password:\n" +
		fmt.Sprintf("http://yourfrontend.com/reset?token=%s", token)

	return smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))
}
