package sms

import (
	"fmt"
	"github.com/twilio/twilio-go"
	"log"

	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type TwilioClient interface {
	SendOTP(to string, otp string) error
	CheckBalance(accountSid string)
}

type twilioClient struct {
	tc *twilio.RestClient
}

func NewTwilioClient(accountSid string, authToken string) TwilioClient {
	tc := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
	return &twilioClient{tc: tc}
}

func (rc *twilioClient) CheckBalance(accountSid string) {
	params := &openapi.FetchBalanceParams{}
	params.SetPathAccountSid(accountSid)
	balance, err := rc.tc.Api.FetchBalance(params)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Twilio account balance: %s %s", *balance.Balance, *balance.Currency)
}

func (rc *twilioClient) SendOTP(to string, otp string) error {
	params := &openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom("+84976447558")

	msg := fmt.Sprintf("Your OTP is %s", otp)
	params.SetBody(msg)

	_, err := rc.tc.Api.CreateMessage(params)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	log.Println("SMS sent successfully!")

	return nil
}
