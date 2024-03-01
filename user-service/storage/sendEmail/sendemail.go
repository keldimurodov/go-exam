package sendemail

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"math/rand"
)

func GenerateCode(length int) string {
	const charset = "0123456789"
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}

	return string(code)
}


func SendEmail(email string) string {
	// sender data
	from := "keldimurodovsardor@gmail.com"
	password := "ghgybsvnstezkztx"

	// Receiver email address
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, _ := template.ParseFiles("template.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Your verification code \n%s\n\n", mimeHeaders)))

	genCode := GenerateCode(6)

	t.Execute(&body, struct {
		Passwd string
	}{

		Passwd: genCode,
	})

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return "eror"
	}
	fmt.Println("Email sended to: ", email)
	return genCode
}
