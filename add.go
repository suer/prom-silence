package main

import (
	"io"
	"net/http"
	"net/url"
)

func AddSilence(endpoint string, reader io.Reader) ([]byte, error) {
	url, err := url.JoinPath(endpoint, "/alertmanager/api/v2/silences")
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return RequestWithSigv4(http.MethodPost, url, body)
}
