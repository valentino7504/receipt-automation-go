package mailer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/valentino7504/tax-automation-go/internal/auth"
)

const (
	contentTypeHTML  = "HTML"
	contentTypeJSON  = "application/json"
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
	req, err := http.NewRequest(
		http.MethodPost,
		graphSendMailURL,
		bytes.NewReader(payload),
	)
	if err != nil {
		fmt.Println("Error creating request")
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	req.Header.Set("Content-type", "application/json")
	client := &http.Client{Timeout: 10 * time.Second}
	_, err = client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %s", err)
	}
	return nil
}
