package azjson

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	// Create a test server to mock the API response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"key": "value"}`))
	}))
	defer ts.Close()

	// Create an access token for authentication
	token := azcore.AccessToken{Token: "test-token"}

	// Call the Get function with the test server URL and access token
	json, err := Get(ts.URL, token)

	// Assert that the json response and error are as expected
	assert.Equal(t, `{"key": "value"}`, json)
	assert.NoError(t, err)
}
