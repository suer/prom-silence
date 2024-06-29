package main

import (
	"net/http"
	"net/url"
)

func DeleteSilence(endpoint string, silenceId string) ([]byte, error) {
	url, err := url.JoinPath(endpoint, "/alertmanager/api/v2/silence", silenceId)
	if err != nil {
		return nil, err
	}

	return RequestWithSigv4(http.MethodDelete, url, []byte{})
}
