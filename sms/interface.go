package sms

type SmsAdapter interface {
	SendMessage(to string, text string) bool
}
