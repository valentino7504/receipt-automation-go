package mailer

import (
	"encoding/json"
	"fmt"

	"github.com/valentino7504/tax-automation-go/internal/auth"
)

const (
	contentTypeHTML  = "HTML"
	graphSendMailURL = "https://graph.microsoft.com/v1.0/me/sendMail"
)

func buildGraphPayload(message Message) ([]byte, error) {
	var graphMsg GraphMessage
	graphMsg.Message.Subject = message.Subject
	graphMsg.Message.Body.ContentType = contentTypeHTML
	graphMsg.Message.Body.Content = message.Body
	graphMsg.Message.ToRecipients = append(
		graphMsg.Message.ToRecipients,
		GraphRecipient{
			EmailAddress: EmailAddress{
				Address: message.To,
			},
		},
	)
	payload, err := json.Marshal(graphMsg)
	if err != nil {
		return nil, fmt.Errorf("error marshalling json")
	}
	return payload, nil
}

func postMail(token auth.UserToken, payload []byte) error {
	return nil
}
