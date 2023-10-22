package verify

import (
	"email_verify/verify/utils"
	"log"
	"net/smtp"
	"net/textproto"
	"os"

	"github.com/jordan-wright/email"
)

type VerifyEmail struct {
	To      string
	Auth    smtp.Auth
	Content []byte
	Config  *VerifyEmailConfig
}

// NewVerifyEmail: init Sender,smtp Auth
func NewVerifyEmail(config *EmailConfig) *VerifyEmail {
	var email VerifyEmail
	email.Auth = smtp.PlainAuth("", "alco89963@163.com", "VBARJVXGVBJYBLJR", "smtp.163.com")
	return nil
}

// Set HTML
func (ve *VerifyEmail) SetContent(path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	ve.Content = file
	return nil
}

func (ve *VerifyEmail) Send(to []string) {
	e := &email.Email{
		To:      to,
		From:    utils.FromFormat(ve.Config.Content.Name, ve.Config.Auth.Username),
		Subject: ve.Config.Content.Subject,
		HTML:    ve.Content,
		Headers: textproto.MIMEHeader{},
	}

	err := e.Send("smtp.163.com:25", ve.Auth)
	if err != nil {
		log.Println(err)
	}
}

// VBARJVXGVBJYBLJR
