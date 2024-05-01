package notification

import (
	"encoding/json"
	"fmt"
	"os"

	"ecommerce.com/config"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type NotificationClient interface {
	SendSMS(phone string, message string)
}

type notificationClient struct {
	config config.AppConfig
}

func (c notificationClient) SendSMS(phone string, message string) {
	accountSid := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo("+3GFDSGDS27363") //MY PHONE NUMBER
	params.SetFrom("+15017250604") //FROM TWILLIO
	params.SetBody("Hello from Go!")

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
}

func NewNotificationClient(config config.AppConfig) NotificationClient {

	return &notificationClient{
		config: config,
	}
}
