package models

import (
	"time"
)

// AuthToken represents the response from the Spotify API when requesting an access token.
type AuthToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	ExpiryTime   time.Time
}

// SetExpiryTime sets the ExpiryTime field of the AuthToken struct based on the ExpiresIn field.
func (at *AuthToken) SetExpiryTime() {
	at.ExpiryTime = time.Now().Add(time.Duration(at.ExpiresIn) * time.Second)
}

// IsExpired checks if the current time is after the ExpiryTime of the AuthToken struct.
func (at *AuthToken) IsExpired() bool {
	return time.Now().After(at.ExpiryTime)
}
