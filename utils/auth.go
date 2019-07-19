package utils

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"context"

	"golang.org/x/oauth2"
)

// References:
// https://gist.github.com/billmccord/4247b0c4d2a6b5a4d09f
// https://gist.github.com/jfcote87/89eca3032cd5f9705ba3
// https://github.com/golang/oauth2/issues/84

// TokenNotifyFunc is a function that accepts an oauth2 Token upon refresh, and
// returns an error if it should not be used.
type TokenNotifyFunc func(*oauth2.Token) error

// NotifyRefreshTokenSource is essentially `oauth2.ResuseTokenSource` with `TokenNotifyFunc` added.
type NotifyRefreshTokenSource struct {
	new oauth2.TokenSource
	mu  sync.Mutex // guards t
	t   *oauth2.Token
	f   TokenNotifyFunc // called when token refreshed so new refresh token can be persisted
}

func StoreNewToken(t *oauth2.Token) error {
	// persist token
	return nil // or error
}

// Token returns the current token if it's still valid, else will
// refresh the current token (using r.Context for HTTP client
// information) and return the new one.
func (s *NotifyRefreshTokenSource) Token() (*oauth2.Token, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.t.Valid() {
		fmt.Println("returning existing token")
		return s.t, nil
	}
	t, err := s.new.Token()
	if err != nil {
		return nil, err
	}
	s.t = t
	return t, s.f(t)
}

func EncryptTokenCache(refreshToken string) {
	// encrypt with AES, store in memory with realmID
}

// OAuth2Service is a service to perform oauth
type OAuth2Service struct {
	Config *oauth2.Config
	Token  *oauth2.Token
}

// NewOAuth2Service makes a new OAuth2Service
func NewOAuth2Service(clientID string, clientSecret string, scopes []string, tokenURL string, authURL string) *OAuth2Service {
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       scopes,
		Endpoint: oauth2.Endpoint{
			TokenURL: tokenURL,
			AuthURL:  authURL,
		},
	}
	return &OAuth2Service{
		Config: conf,
	}
}

// GetAuthURL returns the authorization URL for the user to visit
func (s *OAuth2Service) GetAuthURL(state string, opts ...oauth2.AuthCodeOption) string {
	return s.Config.AuthCodeURL(state, opts...)
}

// ExchangeAuthCodeForToken exchanges the auth code for a token
func (s *OAuth2Service) ExchangeAuthCodeForToken(code string) error {
	ctx := context.Background()
	httpClient := &http.Client{Timeout: 2 * time.Second}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)
	token, err := s.Config.Exchange(ctx, code)
	if err != nil {
		return err
	}
	s.Token = token
	return nil
}

// ProvideToken provides a new token
func (s *OAuth2Service) ProvideToken(accessToken string, refreshToken string, expiry time.Time, tokenType string) {
	s.Token = &oauth2.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expiry:       expiry,
		TokenType:    tokenType,
	}
}

// GetToken returns the token
func (s *OAuth2Service) GetToken() (*oauth2.Token, error) {
	// tokenSource := s.Config.TokenSource(oauth2.NoContext, new(oauth2.Token))
	// return tokenSource.Token()
	return s.Token, nil
}

// VerifyState returns whether the state is valid or not
func (s *OAuth2Service) VerifyState(state string) (bool, error) {
	return true, nil
}

// ParseResponse parses an http Response
func (s *OAuth2Service) ParseResponse(state string) (bool, error) {
	return true, nil
}

// GetClient returns the HTTP client
func (s *OAuth2Service) GetClient() *http.Client {
	ctx := context.Background()
	httpClient := &http.Client{Timeout: 2 * time.Second}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)
	nrts := &NotifyRefreshTokenSource{
		new: s.Config.TokenSource(ctx, s.Token),
		t:   s.Token,
		f: func(token *oauth2.Token) error {
			s.Token = token
			return nil
		},
	}
	return oauth2.NewClient(ctx, nrts)
}
