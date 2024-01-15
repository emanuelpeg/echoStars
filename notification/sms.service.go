package notification

import (
	"echoStars/util"
	"encoding/json"
	"fmt"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type SMSServiceConfig struct {
	accountSid      string
	authToken       string
	fromPhoneNumber string
}

type SMSService struct {
	config SMSServiceConfig
}

func newSMSService(config SMSServiceConfig) *SMSService {
	return &SMSService{config: config}
}

func loadSMSServiceConfig() (SMSServiceConfig, error) {
	config := SMSServiceConfig{
		accountSid:      util.ApplicationConfig.GetString("notification.account.sid"),
		authToken:       util.ApplicationConfig.GetString("notification.auth.token"),
		fromPhoneNumber: util.ApplicationConfig.GetString("notification.from.phone.number"),
	}
	return config, nil
}

func (s *SMSService) SendNotification(subject, body, recipient string) error {

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: s.config.accountSid,
		Password: s.config.authToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetFrom(s.config.fromPhoneNumber)
	params.SetTo(recipient)
	params.SetBody(body)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
		return err
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}

	return nil
}
