package sms

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type SmsAdapterConfig struct {
	apiURL   string
	apiToken string
}

type SmsRequest struct {
	To      string `json:"to"`
	Content string `json:"content"`
	SMSType string `json:"sms_type"`
	// Sender  string   `json:"sender"`
}

type SmsResponse struct {
	Status string `json:"status"`
}

func NewSmsSpeedAdapter(apiToken string) SmsAdapter {
	return &SmsAdapterConfig{
		apiURL:   "http://api.speedsms.vn/index.php/sms/send",
		apiToken: apiToken,
	}
}

func (s *SmsAdapterConfig) SendMessage(to string, text string) bool {
	if len(to) == 0 {
		return false
	}

	requestData := SmsRequest{
		To:      to,
		Content: text,
		SMSType: "2",
		// Sender:  brandname,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return false
	}

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("POST", s.apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return false
	}
	req.SetBasicAuth(s.apiToken, ":x")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	var response SmsResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return false
	}

	return response.Status == "success"
}
