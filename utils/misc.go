package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// StateGenerator interface defines the methods for generating a state string for the Spotify authentication flow.
// This is used to prevent CSRF attacks by ensuring that the application that initiated the authentication flow
// is the same as the application that received the callback.
type StateGenerator interface {
	GetRandomState(length int) (string, error)
}

// DefaultStateGenerator is a struct which implements StateGenerator interface.
type DefaultStateGenerator struct{}

// GetState generates a random string of 16 characters
// to be used as the state parameter in the Spotify authorization flow.
func (dsg *DefaultStateGenerator) GetRandomState(length int) (string, error) {
	// Create a byte slice with given length
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// Convert the byte slice to a base64-encoded string
	// and return the first 16 characters.
	return base64.URLEncoding.EncodeToString(bytes), nil
}
