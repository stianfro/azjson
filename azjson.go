// Package azjson provides functions for doing http requests to APIs that use Azure for authentication
package azjson

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// Get returns the json response from an api with authentication
func Get(url string, token azcore.AccessToken) (json string, err error) {
	httpClient := &http.Client{}
	emptyResponse := "{}"

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, url, http.NoBody)
	if err != nil {
		return emptyResponse, err
	}

	req.Header.Add("Authorization", "Bearer "+token.Token)

	res, err := httpClient.Do(req)
	if err != nil {
		return emptyResponse, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return emptyResponse, err
	}

	return string(body), err
}

// Post sends a JSON payload to an api with authentication
func Post(url string, payload []byte, token azcore.AccessToken, id string) error {
	httpClient := &http.Client{}
	bodyReader := bytes.NewReader(payload)

	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Bearer "+token.Token)
	req.Header.Add("Content-Type", "application/json")

	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode == http.StatusForbidden {
		return err
	}

	return nil
}
