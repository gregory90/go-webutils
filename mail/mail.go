package mail

import (
	mailgun "github.com/mailgun/mailgun-go"

	. "bitbucket.org/pqstudio/go-webutils/logger"
)

var (
	gun mailgun.Mailgun
)

func Init(mailgunDomain string, mailgunPrivateKey string, mailgunPublicKey string) {
	gun = mailgun.NewMailgun(mailgunDomain, mailgunPrivateKey, mailgunPublicKey)
}

func Send(sender string, subject string, message string, recipient string) error {
	m := mailgun.NewMessage(sender, subject, message, recipient)
	m.SetHtml(message)
	_, _, err := gun.Send(m)
	Log.Debug("%+v", err)

	return err
}
