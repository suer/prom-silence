package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
)

func RequestWithSigv4(httpMethod string, url string, body []byte) ([]byte, error) {
	reqBody := strings.NewReader(string(body))
	hash := sha256.Sum256(body)
	hashString := hex.EncodeToString(hash[:])

	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	credentials, err := cfg.Credentials.Retrieve(ctx)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(httpMethod, url, reqBody)
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
		return nil, fmt.Errorf("failed to add silence: Status code %d: response: %s", response.StatusCode, string(responseBody))
	}

	return responseBody, nil

}
