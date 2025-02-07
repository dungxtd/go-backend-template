package sms

import (
	"fmt"

	"github.com/unimtx/uni-go-sdk"
)

type unimtxClient struct {
	cl *uni.UniClient
}

func NewUnimtxClient(unimtxAccessKeyID string, unimtxAccessKeySecret string) SmsAdapter {

	client := uni.NewClient()
	client.AccessKeyId = unimtxAccessKeyID
	client.AccessKeySecret = unimtxAccessKeySecret

	return &unimtxClient{cl: client}
}

func (rc *unimtxClient) SendMessage(to string, msg string) bool {
	res, err := rc.cl.Messages.Send(&uni.MessageSendParams{
		To:   to, // in E.164 format
		Text: msg,
	})
	if err != nil {
		fmt.Println(err)
		return false
	}

	fmt.Println(res.Valid)
	return true
}

func (rc *unimtxClient) VerifyOTP(to string, code string) {
	res, err := rc.cl.Otp.Verify(&uni.OtpVerifyParams{
		To:   to,
		Code: code, // the code user provided
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func (rc *unimtxClient) SendOTP(to string) {
	res, err := rc.cl.Otp.Send(&uni.OtpSendParams{
		To: to,
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.Status, res.Data)
	}
}
