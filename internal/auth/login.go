package auth

import (
	"context"
	"fmt"
)

func Login(ctx context.Context) error {
	deviceCode, err := requestDeviceCode()
	if err != nil {
		return err
	}
	fmt.Printf("Device code: %s\n", deviceCode.UserCode)
	fmt.Printf("URL: %s\n", deviceCode.VerificationURI)
	userToken, err := getAccessToken(ctx, deviceCode)
	if err != nil {
		return nil
	}
	fmt.Println(*userToken)
	return nil
}
