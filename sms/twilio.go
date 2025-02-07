package sms

import (
	"fmt"
	"log"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type twilioClient struct {
	tc *twilio.RestClient
}

func NewTwilioClient(accountSid string, authToken string) SmsAdapter {
	tc := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
	return &twilioClient{tc: tc}
}

// func (rc *twilioClient) CheckBalance(accountSid string) {
// 	params := &openapi.FetchBalanceParams{}
// 	params.SetPathAccountSid(accountSid)
// 	balance, err := rc.tc.Api.FetchBalance(params)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("Twilio account balance: %s %s", *balance.Balance, *balance.Currency)
// }

func (rc *twilioClient) SendMessage(to string, otp string) bool {
	params := &openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom("+84976447558")

	msg := fmt.Sprintf("Your OTP is %s", otp)
	params.SetBody(msg)

	_, err := rc.tc.Api.CreateMessage(params)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	log.Println("SMS sent successfully!")

	return true
}
