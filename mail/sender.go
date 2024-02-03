package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
	"github.com/raaafly/powerup-client-service-golang/utils"
)

type EmailTemplate struct {
	Username         string
	ResetPasswordURL string
	Token			 string
}

const (
	smtpAuthAddress   = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:587"
)

type GmailSender struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

type EmailSender interface {
	SendEmail(
		to []string,
		cc []string,
		bcc []string,
	) error
}

func NewGmailSender(name string, fromEmailAddress string, fromEmailPassword string) EmailSender {
	return &GmailSender{
		name:              name,
		fromEmailAddress:  fromEmailAddress,
		fromEmailPassword: fromEmailPassword,
	}
}

func (sender *GmailSender) SendEmail(
	to []string,
	cc []string,
	bcc []string,
) error {
	userData := EmailTemplate{
		Username:        "Nama Pengguna",
		Token: utils.GenerateRandomCode(8),		
	}

	htmlTemplate := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	  <meta charset="UTF-8">
	  <meta name="viewport" content="width=device-width, initial-scale=1.0">
	  <title>Reset Password</title>
	  <style>
	    /* ... (gaya CSS Anda di sini) ... */
	  </style>
	</head>
	<body>
	  <div class="container">
	    <h2>Reset Password</h2>
	    <p>Hello {{.Username}},</p>
	    <p>Kami menerima permintaan untuk mereset password akun Anda. Untuk melanjutkan proses reset password, silakan klik tautan di bawah ini:</p>

	    <p><a href="{{.ResetPasswordURL}}" class="button">Reset Password</a></p>

	    <p>Jika Anda tidak meminta reset password, Anda dapat mengabaikan email ini.</p>

	    <div class="footer">
	      <p>Terima kasih,<br>Tim Support</p>
	    </div>
	  </div>
	</body>
	</html>
	`

	tmpl, err := template.New("emailTemplate").Parse(htmlTemplate)
	if err != nil {
		log.Fatal(err)
	}

	var bodyMessage bytes.Buffer
	userData.ResetPasswordURL = "https://example.com/reset-password?token=" + userData.Token

	err = tmpl.Execute(&bodyMessage, userData)
	if err != nil {
		log.Fatal(err)
	}

	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.name, sender.fromEmailAddress)
	e.Subject = "[PowerUP Reset Password]"
	e.HTML = bodyMessage.Bytes()
	e.To = to
	e.Cc = cc
	e.Bcc = bcc

	smtpAuth := smtp.PlainAuth("", sender.fromEmailAddress, sender.fromEmailPassword, smtpAuthAddress)
	return e.Send(smtpServerAddress, smtpAuth)
}