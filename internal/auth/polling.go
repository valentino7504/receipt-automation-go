package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	AuthorizationPending  = "authorization_pending"
	AuthorizationDeclined = "authorization_declined"
	BadVerificationCode   = "bad_verification_code"
	ExpiredToken          = "expired_token"
)

func postRequest(uri, params string) (*UserToken, error) {
	res, err := http.Post(
		uri,
		"application/x-www-form-urlencoded",
		strings.NewReader(params),
	)
	if err != nil {
		return nil, err
	}
	var token UserToken
	defer func() {
		if resErr := res.Body.Close(); resErr != nil {
			fmt.Println("error closing response body")
		}
	}()
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusBadRequest {
		return nil, fmt.Errorf("HTTP %d", res.StatusCode)
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(resBody, &token); err != nil {
		return nil, err
	}
	return &token, nil
}

func pollEndpoint(ctx context.Context, uri string, interval int, params string) (*UserToken, error) {
	duration := time.Duration(interval) * time.Second
	ticker := time.NewTicker(duration)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			token, err := postRequest(uri, params)
			if err != nil {
				fmt.Println(err.Error())
				return nil, fmt.Errorf("error in POST request for /token: %s", err)
			}
			if token.Error == "" {
				return token, nil
			}
			if token.Error == AuthorizationPending {
				continue
			}
			fmt.Println("an error 1")
			return nil, fmt.Errorf("auth error: %s - %s", token.Error, token.ErrorDescription)
		}
	}
}
