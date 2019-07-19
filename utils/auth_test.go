package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

var (
	testClientID     = "clientID"
	testClientSecret = "clientSecret"
	testAccessToken  = "accessToken"
	testRefreshToken = "refreshToken"
	testTokenURL     = "http://www.google.com/token"
	testAuthURL      = "http://www.google.com/auth"
	testState        = "state"
	testScopes       = []string{"scope1", "scope2"}
	testCode         = "code"
)

func TestOAuthServiceURL(t *testing.T) {
	expectedURL := "http://www.google.com/auth?client_id=clientID&duration=permanent&response_type=code&scope=scope1+scope2&state=state"
	service := NewOAuth2Service(testClientID, testClientSecret, testScopes, testTokenURL, testAuthURL)
	authURL := service.GetAuthURL(testState, oauth2.SetAuthURLParam("duration", "permanent"))
	assert.Equal(t, authURL, expectedURL)
}

// TODO: Need to build out
func TestOAuthServiceExchange(t *testing.T) {
	service := NewOAuth2Service(testClientID, testClientSecret, testScopes, testTokenURL, testAuthURL)
	service.ExchangeAuthCodeForToken(testCode)
}

// TODO: Need to build out
func TestOAuthVerifyState(t *testing.T) {
	service := NewOAuth2Service(testClientID, testClientSecret, testScopes, testTokenURL, testAuthURL)
	service.ExchangeAuthCodeForToken(testCode)
}
