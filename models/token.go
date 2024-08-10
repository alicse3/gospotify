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
	ExpiryTime   time.Time
}
