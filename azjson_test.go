package azjson

import (
	"encoding/json"
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
		_, err := w.Write([]byte(`{"key": "value"}`))
		if err != nil {
			t.Error(err)
		}
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

func TestPost(t *testing.T) {
	url := "http://localhost:8080"
	payload := []byte(`{"example": "payload"}`)
	token := azcore.AccessToken{Token: "test-token"}
	id := "123"

	// Start a test server
	server := &http.Server{Addr: ":8080"}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Error("Expected POST method, got:", r.Method)
		}

		// Verify that the request headers match the expected values
		if r.Header.Get("Authorization") != "Bearer "+token.Token {
			t.Error("Expected authorization header to be 'Bearer "+token.Token+"', got:", r.Header.Get("Authorization"))
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Error("Expected content-type header to be 'application/json', got:", r.Header.Get("Content-Type"))
		}

		// Verify that the request body matches the expected payload
		var reqPayload map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&reqPayload)
		if err != nil {
			t.Error("Error decoding request body:", err)
		}
		if reqPayload["example"] != "payload" {
			t.Error("Expected payload to be `{\"example\": \"payload\"}`, got:", reqPayload)
		}
	})

	// Check the error return value of Listen and Serve
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			t.Error("Error starting server:", err)
		}
		defer server.Close()
	}()

	// Test success case
	err := Post(url, payload, token, id)
	if err != nil {
		t.Error("Expected nil error, got:", err)
	}
}
