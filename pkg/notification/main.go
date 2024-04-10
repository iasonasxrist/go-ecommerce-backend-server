package notification

import "ecommerce.com/config"

type NotificationClient interface {
	SendSMS(phone string, message string)
}

type notificationClient struct {
	config config.AppConfig
}

func (c notificationClient) SendSMS(phone string, message string) {

}

func NewNotificationClient(config config.AppConfig) NotificationClient {

	return &notificationClient{
		config: config,
	}
}
