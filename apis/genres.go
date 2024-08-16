package apis

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/alicse3/gospotify/consts"
	"github.com/alicse3/gospotify/models"
	"github.com/alicse3/gospotify/utils"
)

// GenreService interface defines the methods for interacting with the Spotify Genre's API.
type GenreService interface {
	// Retrieve a list of available genres seed parameter values for recommendations.
	GetAvailableGenresSeeds() (*models.Genres, error)
}

// DefaultGenreService is a struct that implements GenreService interface.
type DefaultGenreService struct {
	client *utils.HttpClient
}

// NewDefaultGenreService initializes the DefaultGenreService with given dependencies.
func NewDefaultGenreService(client *utils.HttpClient) *DefaultGenreService {
	return &DefaultGenreService{client}
}

// GetAvailableGenresSeeds implements the DefaultGenreService's interface GetAvailableGenresSeeds method.
func (service *DefaultGenreService) GetAvailableGenresSeeds() (*models.Genres, error) {
	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointGetAvailableGenreSeeds, nil)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetAvailableGenreSeeds, Err: err}
	}

	// Handle Spotify API error
	if res.StatusCode != http.StatusOK {
		return nil, utils.ParseSpotifyError(res, utils.RegErrorType)
	}

	// Read the response body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Unmarshal the response data into Genres struct
	var genres models.Genres
	if err := json.Unmarshal(data, &genres); err != nil {
		return nil, err
	}

	// Return the Genres
	return &genres, nil
}
