package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "ApiKey 1234567")
	expected := "1234567"
	actual, _ := GetAPIKey(header)
	if actual != expected {
		t.Errorf("%s does not equal %s", expected, actual)
	}

}

func TestGetApiKeyNoAuth(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "Bearer 1234567")
	_, err := GetAPIKey(header)
	if err == nil {
		t.Errorf("Did not return error")
	}

}
