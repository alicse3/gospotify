package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

const (
	// default http client timeout
	defaultHttpClientTimeout = 10 * time.Second
)

// HttpClient is a struct that wraps the standard http.Client
// and provides a convenient way to make HTTP requests with a base URL.
type HttpClient struct {
	Client  *http.Client
	BaseUrl string // Base for api requests
}

// NewHttpClient returns a new HttpClient instance with a default timeout of 10 seconds.
func NewHttpClient(baseUrl string) *HttpClient {
	return &HttpClient{
		Client:  &http.Client{Timeout: defaultHttpClientTimeout},
		BaseUrl: baseUrl,
	}
}

// TokenTransport is a struct that adds a Bearer token to the Authorization header of each HTTP request.
type TokenTransport struct {
	Token     string
	Transport http.RoundTripper
}

// RoundTrip adds the Authorization header to each request
func (tt *TokenTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+tt.Token)
	return tt.Transport.RoundTrip(req)
}

// NewHttpClientWithToken creates an http.Client with the token injected into each request.
func NewHttpClientWithToken(baseUrl, token string) *HttpClient {
	return &HttpClient{
		Client: &http.Client{
			Timeout: defaultHttpClientTimeout,
			Transport: &TokenTransport{
				Token:     token,
				Transport: http.DefaultTransport,
			},
		},
		BaseUrl: baseUrl,
	}
}

// Post makes an HTTP POST request to the specified endpoint with optional headers, query params, form values, and request body.
// It returns the HTTP response and any error that occurred during the request.
func (hc *HttpClient) Post(ctx context.Context, endpoint string, headers, queryParams, formValues map[string]string, body any) (*http.Response, error) {
	// Construct full url
	fullUrl := hc.BaseUrl + endpoint

	// Parse the URL and handle any errors
	u, err := url.ParseRequestURI(fullUrl)
	if err != nil {
		return nil, err
	}

	// Marshal the request body (if provided) to JSON
	var jsonData []byte
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, err
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
		return nil, err
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
	res, err := hc.Client.Do(req)
	if err != nil {
		return nil, err
	}

	// Return the Response
	return res, nil
}

// Get makes an HTTP Get request to the specified endpoint with optional query params.
// It returns the HTTP response and any error that occurred during the request.
func (hc *HttpClient) Get(ctx context.Context, endpoint string, queryParams map[string]string) (*http.Response, error) {
	// Construct full url
	fullUrl := hc.BaseUrl + endpoint

	// Parse the URL and handle any errors
	u, err := url.ParseRequestURI(fullUrl)
	if err != nil {
		return nil, err
	}

	// Create a new HTTP request with the provided context, method, and request URL
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
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
	res, err := hc.Client.Do(req)
	if err != nil {
		return nil, err
	}

	// Return the Response
	return res, nil
}

// Put makes an HTTP PUT request to the specified endpoint with optional headers, query params, and request body.
// It returns the HTTP response and any error that occurred during the request.
func (hc *HttpClient) Put(ctx context.Context, endpoint string, headers, queryParams map[string]string, body any) (*http.Response, error) {
	// Construct full url
	fullUrl := hc.BaseUrl + endpoint

	// Parse the URL and handle any errors
	u, err := url.ParseRequestURI(fullUrl)
	if err != nil {
		return nil, err
	}

	// Marshal the request body (if provided) to JSON
	var jsonData []byte
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		jsonData = data
	}

	// Create a new HTTP request with the provided context, method, URL, and request body
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, u.String(), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
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
	res, err := hc.Client.Do(req)
	if err != nil {
		return nil, err
	}

	// Return the Response
	return res, nil
}

// Delete makes an HTTP DELETE request to the specified endpoint with optional headers, query params, and request body.
// It returns the HTTP response and any error that occurred during the request.
func (hc *HttpClient) Delete(ctx context.Context, endpoint string, headers, queryParams map[string]string, body any) (*http.Response, error) {
	// Construct full url
	fullUrl := hc.BaseUrl + endpoint

	// Parse the URL and handle any errors
	u, err := url.ParseRequestURI(fullUrl)
	if err != nil {
		return nil, err
	}

	// Marshal the request body (if provided) to JSON
	var jsonData []byte
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		jsonData = data
	}

	// Create a new HTTP request with the provided context, method, URL, and request body
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, u.String(), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
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
	res, err := hc.Client.Do(req)
	if err != nil {
		return nil, err
	}

	// Return the Response
	return res, nil
}
