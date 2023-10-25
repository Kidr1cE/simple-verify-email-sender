package sender

import (
	"bytes"
	"fmt"
	"net/smtp"
	"net/textproto"
	"os"
	"text/template"
	"time"

	"github.com/jordan-wright/email"
)

type V struct {
	Verify string
}
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
	// init Auth
	auth := config.Auth

	// init Email
	sender.Email = auth.Username

	// init Pool
	pool, err := email.NewPool(
		config.Address,
		4,
		smtp.PlainAuth("", auth.Username, auth.Password, auth.Host),
	)
	sender.Pool = pool

	// init ccontent info
	sender.Template = config.Template

	// init HTMLTemplate config
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

	sender.HTMLTemplate, err = template.New("HTML").Parse(string(html))
	if err != nil {
		return err
	}
	return nil
}

func (sender *VerifyEmailSender) SendTo(address, verify string) error {
	// get html
	v := &V{
		Verify: verify,
	}
	var buf bytes.Buffer
	sender.HTMLTemplate.Execute(&buf, v)

	// structure email
	var email = &email.Email{
		To:      []string{address},
		From:    fmt.Sprintf("%s <%s>", sender.Template.Name, sender.Email),
		Subject: sender.Template.Subject,
		HTML:    []byte(buf.String()),
		Headers: textproto.MIMEHeader{},
	}

	// send email
	sender.Pool.Send(email, 10*time.Second)
	return nil
}
