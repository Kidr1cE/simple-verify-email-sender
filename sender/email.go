package sender

import (
	"email_verify/sender/utils"
	"net/smtp"
	"net/textproto"
	"os"
	"text/template"
	"time"

	"github.com/jordan-wright/email"
)

type VerifyEmailSender struct {
	Email        string
	Template     *Template
	HTMLTemplate *template.Template
	Pool         *email.Pool
}

func (sender *VerifyEmailSender) SetTemplate(path string) error {

	return nil
}

func (sender *VerifyEmailSender) ReadConfig(config *VerifyEmailConfig) error {
	auth := config.Auth
	sender.Email = auth.Username
	pool, err := email.NewPool(
		config.Address,
		4,
		smtp.PlainAuth("", auth.Username, auth.Password, auth.Host),
	)
	sender.Pool = pool
	sender.Template = config.Template

	var path string
	if config.Verify.Type == "token" {
		path = "./sender/template/default_token.html"
	} else {
		path = "./sender/template/default_code.html"
	}
	html, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	sender.HTMLTemplate, err = utils.GetTemplate(string(html))

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
		HTML:    utils.GetContent(sender.HTMLTemplate, []string{verify}),
		Headers: textproto.MIMEHeader{},
	}

	// send email
	sender.Pool.Send(email, 10*time.Second)
	return nil
}
