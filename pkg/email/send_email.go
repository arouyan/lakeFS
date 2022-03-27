package email

import (
	"github.com/treeverse/lakefs/pkg/email/params"
	gomail "gopkg.in/gomail.v2"
)

func NewEmail(e params.Email, err error) (*params.Email, error) {
	email := &params.Email{
		SmtpHost:           e.SmtpHost,
		Port:               e.Port,
		Username:           e.Username,
		Password:           e.Password,
		Sender:             e.Sender,
		Receivers:          e.Receivers,
		Subject:            e.Subject,
		Body:               e.Body,
		AttachmentFilePath: e.AttachmentFilePath,
	}
	return email, err
}

func SendEmail(e params.Email) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", e.Sender)
	msg.SetHeader("To", e.Receivers...)
	msg.SetHeader("Subject", e.Subject)
	msg.SetBody("text/html", e.Body)
	if len(e.AttachmentFilePath) > 0 {
		for _, f := range e.AttachmentFilePath {
			msg.Attach(f)
		}
	}
	d := gomail.NewDialer(e.SmtpHost, e.Port, e.Username, e.Password)
	return d.DialAndSend(msg)
}

func InsertAdditionalParamsToEmail(e params.Email, receivers []string, subject string, body string, attachmentFilePath ...string) *params.Email {
	e.Receivers = receivers
	e.Subject = subject
	e.Body = body
	e.AttachmentFilePath = attachmentFilePath
	return &e
}
