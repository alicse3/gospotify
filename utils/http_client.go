package utils

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/alicse3/gospotify/consts"
	"github.com/alicse3/gospotify/models"
)

const (
	// default http client timeout
	defaultHttpClientTimeout = 10 * time.Second
)

// HttpClient is a struct that wraps the standard http.Client
// and provides a convenient way to make HTTP requests with a base URL.
type HttpClient struct {
	client      *http.Client      // Client for making HTTP requests
	baseUrl     string            // Base url for api requests
	authToken   *models.AuthToken // Auth token for authenticating requests
	credentials *Credentials      // For refreshing the tokens
	mu          sync.Mutex
}

// NewHttpClient returns a new HttpClient instance with a default timeout of 10 seconds.
func NewHttpClient(baseUrl string) *HttpClient {
	return &HttpClient{
		client:  &http.Client{Timeout: defaultHttpClientTimeout},
		baseUrl: baseUrl,
	}
}

// NewHttpClientWithToken creates an httpClient instance with the given dependencies.
func NewHttpClientWithToken(baseUrl string, authToken *models.AuthToken, credentials *Credentials) *HttpClient {
	return &HttpClient{
		client:      &http.Client{Timeout: defaultHttpClientTimeout},
		baseUrl:     baseUrl,
		authToken:   authToken,
		credentials: credentials,
	}
}

// refreshToken refreshes the access token using the refresh token.
func (hc *HttpClient) refreshToken() error {
	// To make sure the dependencies are initialized before refreshing the tokens
	if hc.credentials == nil {
		return &Error{Type: AppErrorType, AppError: &AppError{Status: http.StatusInternalServerError, Message: consts.MsgCredentialsNotInitialized}}
	}

	// Generating base64 endoded(client id and client secret) string for authorization.
	// For details, visit: https://developer.spotify.com/documentation/web-api/tutorials/refreshing-tokens
	base64Encoded := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", hc.credentials.ClientId, hc.credentials.ClientSecret)))

	// Set the required headers
	headers := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": "Basic " + base64Encoded, // As per the Spotify document, this is only required for the Authorization Code
	}

	// Set the form values for the token refresh request
	formValues := map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": hc.authToken.RefreshToken,
	}

	// Make a POST request to the token endpoint
	res, err := hc.Post(context.Background(), consts.EndpointToken, headers, nil, formValues, nil)
	if err != nil {
		return &AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToRefreshTokens, Err: err}
	}

	// Handle Spotify API error
	if res.StatusCode != http.StatusOK {
		return ParseSpotifyError(res, AuthErrorType)
	}

	// Read the response body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return &AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToReadResponseBody, Err: err}
	}
	defer res.Body.Close()

	// Unmarshal the response data into AuthToken struct
	var authToken models.AuthToken
	if err := json.Unmarshal(data, &authToken); err != nil {
		return &AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToUnmarshalResponseData, Err: err}
	}

	// Update the client's authToken except the refreshToken
	// Not updating the refresh token, because the Spotify's refresh tokens do not expire by default unless they are explicitly revoked by the user or by Spotify.
	hc.authToken.AccessToken = authToken.AccessToken
	hc.authToken.TokenType = authToken.TokenType
	hc.authToken.ExpiresIn = authToken.ExpiresIn
	hc.authToken.Scope = authToken.Scope
	hc.authToken.SetExpiryTime()

	return nil
}

// checkAndRefreshTokens checks for the AuthToken expiry and then triggers refresh tokens call if needed.
func (hc *HttpClient) checkAndRefreshTokens() error {
	hc.mu.Lock()
	defer hc.mu.Unlock()

	// Check if the token has expired
	if hc.authToken.IsExpired() {
		// Refresh the token
		if err := hc.refreshToken(); err != nil {
			return err
		}
	}

	return nil
}

// do sends an HTTP request and automatically handles token expiration.
func (hc *HttpClient) do(req *http.Request) (*http.Response, error) {
	// If auth token is set, check and refresh the token if needed
	if hc.authToken != nil {
		if err := hc.checkAndRefreshTokens(); err != nil {
			return nil, err
		}

		// Add the access token to the Authorization header
		req.Header.Set("Authorization", "Bearer "+hc.authToken.AccessToken)
	}

	// Send the request
	return hc.client.Do(req)
}

// Post makes an HTTP POST request to the specified endpoint with optional headers, query params, form values, and request body.
// It returns the HTTP response and any error that occurred during the request.
func (hc *HttpClient) Post(ctx context.Context, endpoint string, headers, queryParams, formValues map[string]string, body any) (*http.Response, error) {
	// Construct full url
	fullUrl := hc.baseUrl + endpoint

	// Parse the URL and handle any errors
	u, err := url.ParseRequestURI(fullUrl)
	if err != nil {
		return nil, &Error{Type: AppErrorType, AppError: &AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToParseUrl, Err: err}}
	}

	// Marshal the request body (if provided) to JSON
	var jsonData []byte
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, &Error{Type: AppErrorType, AppError: &AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToMarshalRequestData, Err: err}}
		}
		jsonData = data
	}

	// Encode the form values (if provided) and append them to the URL
	if formValues != nil {
		values := url.Values{}
		for key, val := range formValues {
			values.Add(key, val)
		}
		u.RawQuery = values.Encode()
	}

	// Create a new HTTP request with the provided context, method, URL, and request body
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, &Error{Type: AppErrorType, AppError: &AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToCreatePostRequest, Err: err}}
	}

	// Set the request headers (if provided)
	for key, val := range headers {
		req.Header.Set(key, val)
	}

	// Set the query params in the request
	if queryParams != nil {
		query := req.URL.Query()
		for key, val := range queryParams {
			query.Add(key, val)
		}
		req.URL.RawQuery = query.Encode()
	}

	// Send the HTTP request and return the response and any error that occurred
	res, err := hc.do(req)
	if err != nil {
		return nil, &Error{Type: AppErrorType, AppError: &AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToSendRequest, Err: err}}
	}

	// Return the Response
	return res, nil
}

// Get makes an HTTP Get request to the specified endpoint with optional query params.
// It returns the HTTP response and any error that occurred during the request.
func (hc *HttpClient) Get(ctx context.Context, endpoint string, queryParams map[string]string) (*http.Response, error) {
	// Construct full url
	fullUrl := hc.baseUrl + endpoint

	// Parse the URL and handle any errors
	u, err := url.ParseRequestURI(fullUrl)
	if err != nil {
		return nil, &Error{Type: AppErrorType, AppError: &AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToParseUrl, Err: err}}
	}

	// Create a new HTTP request with the provided context, method, and request URL
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, &Error{Type: AppErrorType, AppError: &AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToCreateGetRequest, Err: err}}
	}

	// Set the query params in the request
	if queryParams != nil {
		query := req.URL.Query()
		for key, val := range queryParams {
			query.Add(key, val)
		}
		req.URL.RawQuery = query.Encode()
	}

	// Send the HTTP request and return the response and any error that occurred
	res, err := hc.do(req)
	if err != nil {
		return nil, &Error{Type: AppErrorType, AppError: &AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToSendRequest, Err: err}}
	}

	// Return the Response
	return res, nil
}

// Put makes an HTTP PUT request to the specified endpoint with optional headers, query params, and request body.
// It returns the HTTP response and any error that occurred during the request.
func (hc *HttpClient) Put(ctx context.Context, endpoint string, headers, queryParams map[string]string, body any) (*http.Response, error) {
	// Construct full url
	fullUrl := hc.baseUrl + endpoint

	// Parse the URL and handle any errors
	u, err := url.ParseRequestURI(fullUrl)
	if err != nil {
		return nil, &Error{Type: AppErrorType, AppError: &AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToParseUrl, Err: err}}
	}

	// Marshal the request body (if provided) to JSON
	var jsonData []byte
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, &Error{Type: AppErrorType, AppError: &AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToMarshalRequestData, Err: err}}
		}
		jsonData = data
	}

	// Create a new HTTP request with the provided context, method, URL, and request body
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, u.String(), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, &Error{Type: AppErrorType, AppError: &AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToCreatePutRequest, Err: err}}
	}

	// Set the request headers (if provided)
	for key, val := range headers {
		req.Header.Set(key, val)
	}

	// Set the query params in the request
	if queryParams != nil {
		query := req.URL.Query()
		for key, val := range queryParams {
			query.Add(key, val)
		}
		req.URL.RawQuery = query.Encode()
	}

	// Send the HTTP request and return the response and any error that occurred
	res, err := hc.do(req)
	if err != nil {
		return nil, &Error{Type: AppErrorType, AppError: &AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToSendRequest, Err: err}}
	}

	// Return the Response
	return res, nil
}

// Delete makes an HTTP DELETE request to the specified endpoint with optional headers, query params, and request body.
// It returns the HTTP response and any error that occurred during the request.
func (hc *HttpClient) Delete(ctx context.Context, endpoint string, headers, queryParams map[string]string, body any) (*http.Response, error) {
	// Construct full url
	fullUrl := hc.baseUrl + endpoint

	// Parse the URL and handle any errors
	u, err := url.ParseRequestURI(fullUrl)
	if err != nil {
		return nil, &Error{Type: AppErrorType, AppError: &AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToParseUrl, Err: err}}
	}

	// Marshal the request body (if provided) to JSON
	var jsonData []byte
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, &Error{Type: AppErrorType, AppError: &AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToMarshalRequestData, Err: err}}
		}
		jsonData = data
	}

	// Create a new HTTP request with the provided context, method, URL, and request body
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, u.String(), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, &Error{Type: AppErrorType, AppError: &AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToCreateDeleteRequest, Err: err}}
	}

	// Set the request headers (if provided)
	for key, val := range headers {
		req.Header.Set(key, val)
	}

	// Set the query params in the request
	if queryParams != nil {
		query := req.URL.Query()
		for key, val := range queryParams {
			query.Add(key, val)
		}
		req.URL.RawQuery = query.Encode()
	}

	// Send the HTTP request and return the response and any error that occurred
	res, err := hc.do(req)
	if err != nil {
		return nil, &Error{Type: AppErrorType, AppError: &AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToSendRequest, Err: err}}
	}

	// Return the Response
	return res, nil
}
