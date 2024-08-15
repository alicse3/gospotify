package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/alicse3/gospotify/consts"
)

// AuthenticationError represents the Spotify authentication error object.
type AuthenticationError struct {
	Err         string `json:"error"`
	Description string `json:"error_description"`
}

// Error returns the AuthenticationError message.
func (ae *AuthenticationError) Error() string {
	return fmt.Sprintf("Authentication Error: %s - %s", ae.Err, ae.Description)
}

// RegularError represents the Spotify regular error object.
type RegularError struct {
	Err struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"error"`
}

// Error returns the RegularError message.
func (re *RegularError) Error() string {
	return fmt.Sprintf("Regular Error: %d - %s", re.Err.Status, re.Err.Message)
}

// AppError represents a custom error for the SDK.
type AppError struct {
	Status  int
	Message string
	Err     error
}

// Error returns the AppError message.
func (ae *AppError) Error() string {
	if ae.Err != nil {
		return fmt.Sprintf("App Error: %d - %s - %v", ae.Status, ae.Message, ae.Err)
	}
	return fmt.Sprintf("App Error: %d - %s", ae.Status, ae.Message)
}

// ErrorType defines the different types of errors.
type ErrorType int

const (
	AuthErrorType ErrorType = iota
	RegErrorType
	AppErrorType
)

// Error represents a unified error type.
type Error struct {
	Type      ErrorType
	AuthError *AuthenticationError
	RegError  *RegularError
	AppError  *AppError
}

// Error returns the appropriate error message based on the ErrorType.
func (e *Error) Error() string {
	switch e.Type {
	case AuthErrorType:
		return e.AuthError.Error()
	case RegErrorType:
		return e.RegError.Error()
	case AppErrorType:
		return e.AppError.Error()
	default:
		return consts.MsgUknownErrorType
	}
}

// ParseSpotifyError parses the Spotify API error response into a unified Error type.
func ParseSpotifyError(res *http.Response, errorType ErrorType) error {
	// Read response body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return &Error{Type: AppErrorType, AppError: &AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToReadResponseBody, Err: err}}
	}
	defer res.Body.Close()

	// Print the response in case of any Spotify API errors
	fmt.Printf("Spotify API response: %+v\n", string(data))

	// Handle Spotify errors
	if errorType == AuthErrorType {
		var authError AuthenticationError
		if err := json.Unmarshal(data, &authError); err != nil {
			return &Error{Type: AppErrorType, AppError: &AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToUnmarshalResponseData, Err: err}}
		}
		return &Error{Type: AuthErrorType, AuthError: &authError}
	} else if errorType == RegErrorType {
		var regError RegularError
		if err := json.Unmarshal(data, &regError); err != nil {
			return &Error{Type: AppErrorType, AppError: &AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToUnmarshalResponseData, Err: err}}
		}
		return &Error{Type: RegErrorType, RegError: &regError}
	} else {
		return errors.ErrUnsupported
	}
}
