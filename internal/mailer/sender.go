package mailer

import "github.com/valentino7504/tax-automation-go/internal/auth"

func SendMail(token auth.UserToken, message Message) error {
	payload, err := buildGraphPayload(message)
	if err != nil {
		return err
	}
	err = postMail(token, payload)
	if err != nil {
		return err
	}
	return nil
}
