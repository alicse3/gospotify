package models

// Markets represents the markets information retrieved from the Spotify API.
type Markets struct {
	Markets []string `json:"markets"`
}
