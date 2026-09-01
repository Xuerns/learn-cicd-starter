package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(test *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "ApiKey cdsojncojsdnojcndsojcodsjkcsdkldjkvpkwkcznvojsncjfnnw0-31u29y8943932=2584935894389543-54893589347584389")

	_, err := GetAPIKey(header)

	if err != nil {
		test.Errorf("Mayday maydat auo uo")
	}
}
