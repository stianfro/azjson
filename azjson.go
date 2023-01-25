package azjson

import (
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
