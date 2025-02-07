package bootstrap

import (
	"github.com/sportgo-app/sportgo-go/sms"
)

func NewTwilioClient(env *Env) sms.TwilioClient {
	accountSid := env.TwilioAccountSID
	authToken := env.TwilioAuthToken

	client := sms.NewTwilioClient(accountSid, authToken)

	client.CheckBalance(accountSid)

	return client
}

func NewUnimtxClient(env *Env) sms.UnimtxClient {
	return sms.NewUnimtxClient(env.UnimtxAccessKeyID, env.UnimtxAccessKeySecret)
}
