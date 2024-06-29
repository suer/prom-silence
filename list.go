package main

import (
	"net/http"
	"net/url"
)

func ListSilences(endpoint string) ([]byte, error) {
	url, err := url.JoinPath(endpoint, "/alertmanager/api/v2/silences")
	if err != nil {
		return nil, err
	}

	return RequestWithSigv4(http.MethodGet, url, []byte{})
}
