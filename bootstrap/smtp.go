package bootstrap

import (
	"github.com/sportgo-app/sportgo-go/email"
)

func NewSMTPMailer(env *Env) email.MailClient {
	client, err := email.NewMailer(env.EmailSMTPHost, env.EmailSMTPPort, env.EmailSMTPUser, env.EmailSMTPPassword)
	if err != nil {
		panic(err)
	}
	return client
}
