package models

// Genres represents the genres information retrieved from the Spotify API.
type Genres struct {
	Genres []string `json:"genres"`
}
