package auth

import (
	"context"
	"fmt"
	"time"
)

func Login() (*UserToken, error) {
	// Request device code from Microsoft device code endpoint
	deviceCode, err := requestDeviceCode()
	if err != nil {
		return nil, err
	}

	// Create context and get the user to authenticate
	expiryTime := time.Now().Add(time.Duration(deviceCode.ExpiresIn) * time.Second)
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(deviceCode.ExpiresIn)*time.Second,
	)
	defer cancel()
	fmt.Printf("Device code: %s\n", deviceCode.UserCode)
	fmt.Printf("Expires by %v\n", expiryTime.Format("3:04pm"))
	fmt.Printf("URL: %s\n", deviceCode.VerificationURI)

	// Gets Access Token
	userToken, err := getAccessToken(ctx, deviceCode)
	if err != nil {
		return nil, err
	}
	return userToken, nil
}
