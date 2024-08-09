package gospotify

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/alicse3/gospotify/consts"
	"github.com/alicse3/gospotify/models"
	"github.com/alicse3/gospotify/utils"
)

// CredentialsExchanger interface defines the methods for token retrieval.
type CredentialsExchanger interface {
	GetAuthorizationUrl(scopes []string, state string) (string, error)
	ExchangeCodeForTokens(httpClient *utils.HttpClient, code string) (*models.AuthToken, error)
}

// Credentials struct holds the client ID, client secret and redirect url.
type Credentials struct {
	ClientId     string
	ClientSecret string
	RedirectUrl  string
}

// GetAuthorizationUrl generates the URL for initiating the authorization flow.
func (c *Credentials) GetAuthorizationUrl(scopes []string, state string) (string, error) {
	authUrl := consts.BaseUrlAccounts + consts.EndpointAuthorize

	// Create a new URL object with the base URL
	u, err := url.Parse(authUrl)
	if err != nil {
		return "", err
	}

	// Set the query parameters for the authorization URL
	q := u.Query()
	q.Set("client_id", c.ClientId)
	q.Set("client_secret", c.ClientSecret)
	q.Set("redirect_uri", c.RedirectUrl)
	q.Set("response_type", "code")
	q.Set("scope", strings.Join(scopes, " "))
	q.Set("state", state)
	u.RawQuery = q.Encode()

	// Return the constructed authorization URL as a string
	return u.String(), nil
}

// ExchangeCodeForTokens method fetches an access token from the Accounts API.
func (c *Credentials) ExchangeCodeForTokens(httpClient *utils.HttpClient, code string) (*models.AuthToken, error) {
	// Set the required headers
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	// Set the form values for the request
	formValues := map[string]string{
		"grant_type":    "authorization_code",
		"client_id":     c.ClientId,
		"client_secret": c.ClientSecret,
		"redirect_uri":  c.RedirectUrl,
		"code":          code,
	}

	// Make a POST request to the token endpoint
	res, err := httpClient.Post(context.Background(), consts.EndpointToken, headers, formValues, nil)
	if err != nil {
		return nil, err
	}

	// Handle Spotify API error
	if res.StatusCode != http.StatusOK {
		return nil, utils.ParseSpotifyError(res, utils.AuthErrorType)
	}

	// Read the response body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Unmarshal the response data into an AuthToken struct
	var authToken models.AuthToken
	if err := json.Unmarshal(data, &authToken); err != nil {
		return nil, err
	}
	authToken.ExpiryTime = time.Now().Add(time.Duration(authToken.ExpiresIn) * time.Second)

	// Return the AuthToken
	return &authToken, nil
}
