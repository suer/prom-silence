package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
)

func DeleteSilence(endpoint string, silenceId string) ([]byte, error) {
	url, err := url.JoinPath(endpoint, "/alertmanager/api/v2/silence", silenceId)
	if err != nil {
		return nil, err
	}

	hash := sha256.Sum256([]byte{})
	hashString := hex.EncodeToString(hash[:])

	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	credentials, err := cfg.Credentials.Retrieve(ctx)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	signer := v4.NewSigner()
	signer.SignHTTP(ctx, credentials, req, hashString, "aps", cfg.Region, time.Now().UTC())
	httpClient := new(http.Client)
	response, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("failed to delete silence: Status code %d: response: %s", response.StatusCode, string(responseBody))
	}

	return responseBody, nil
}
