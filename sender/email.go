package sender

import (
	"email_verify/sender/utils"
	"net/smtp"
	"net/textproto"
	"time"

	"github.com/jordan-wright/email"
)

type VerifyEmailSender struct {
	Email    string
	Template *Template
	Pool     *email.Pool
}

func (sender *VerifyEmailSender) ReadConfig(config *VerifyEmailConfig) error {
	auth := config.Auth
	pool, err := email.NewPool(
		config.Address,
		4,
		smtp.PlainAuth("", auth.Username, auth.Password, auth.Host),
	)
	sender.Pool = pool

	sender.Email = auth.Username

	if err != nil {
		return err
	}
	return nil
}

func (sender *VerifyEmailSender) SendTo(address, verify string) error {
	// structure email
	var email = &email.Email{
		To:      []string{address},
		From:    utils.FromFormat(sender.Template.Name, sender.Email),
		Subject: sender.Template.Subject,
		HTML:    []byte(""),
		Headers: textproto.MIMEHeader{},
	}

	// send email
	sender.Pool.Send(email, 10*time.Second)
	return nil
}
