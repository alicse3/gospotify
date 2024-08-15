package apis

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/alicse3/gospotify/consts"
	"github.com/alicse3/gospotify/models"
	"github.com/alicse3/gospotify/utils"
)

// ArtistService interface defines the methods for interacting with the Spotify Artist's API.
type ArtistService interface {
	GetArtist(input models.GetArtistRequest) (*models.Artist, error)
	GetArtists(input models.GetArtistsRequest) (*models.Artists, error)
	GetArtistAlbums(input models.GetArtistAlbumsRequest) (*models.ArtistAlbums, error)
	GetArtistTopTracks(input models.GetArtistTopTracksRequest) (*models.ArtistTopTracks, error)
	GetRelatedArtists(input models.GetRelatedArtistsRequest) (*models.Artists, error)
}

// DefaultArtistService is a struct that implements ArtistService interface.
type DefaultArtistService struct {
	client *utils.HttpClient
}

// NewDefaultArtistService initializes the DefaultArtistService with given dependencies.
func NewDefaultArtistService(client *utils.HttpClient) *DefaultArtistService {
	return &DefaultArtistService{client}
}

// GetArtist implements the ArtistService's interface GetArtist method.
func (service *DefaultArtistService) GetArtist(input models.GetArtistRequest) (*models.Artist, error) {
	// Validate the input
	if input.Id == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointArtist, input.Id)

	// Make an API call
	res, err := service.client.Get(context.Background(), endpoint, nil)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetArtist, Err: err}
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

	// Unmarshal the response data into Artist struct
	var artist models.Artist
	if err := json.Unmarshal(data, &artist); err != nil {
		return nil, err
	}

	// Return the Artist
	return &artist, nil
}

// GetArtists implements the ArtistService's interface GetArtists method.
func (service *DefaultArtistService) GetArtists(input models.GetArtistsRequest) (*models.Artists, error) {
	// Validate the input
	if input.Ids == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdsRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"ids": input.Ids}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointArtists, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetArtists, Err: err}
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

	// Unmarshal the response data into Artists struct
	var artists models.Artists
	if err := json.Unmarshal(data, &artists); err != nil {
		return nil, err
	}

	// Return the Artists
	return &artists, nil
}

// GetArtistAlbums implements the ArtistService's interface GetArtistAlbums method.
func (service *DefaultArtistService) GetArtistAlbums(input models.GetArtistAlbumsRequest) (*models.ArtistAlbums, error) {
	// Validate the input
	if input.Id == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointArtistAlbums, input.Id)

	// Add inputs to the query parameters
	params := map[string]string{"id": input.Id, "include_groups": input.IncludeGroups, "market": input.Market, "limit": strconv.Itoa(input.Limit), "offset": strconv.Itoa(input.Offset)}

	// Make an API call
	res, err := service.client.Get(context.Background(), endpoint, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetArtistAlbums, Err: err}
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

	// Unmarshal the response data into ArtistAlbums struct
	var artistAlbums models.ArtistAlbums
	if err := json.Unmarshal(data, &artistAlbums); err != nil {
		return nil, err
	}

	// Return the Artists
	return &artistAlbums, nil
}

// GetArtistTopTracks implements the ArtistService's interface GetArtistTopTracks method.
func (service *DefaultArtistService) GetArtistTopTracks(input models.GetArtistTopTracksRequest) (*models.ArtistTopTracks, error) {
	// Validate the input
	if input.Id == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointArtistTopTracks, input.Id)

	// Add inputs to the query parameters
	params := map[string]string{"id": input.Id, "market": input.Market}

	// Make an API call
	res, err := service.client.Get(context.Background(), endpoint, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetArtistTopTracks, Err: err}
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

	// Unmarshal the response data into ArtistTopTracks struct
	var artistTopTracks models.ArtistTopTracks
	if err := json.Unmarshal(data, &artistTopTracks); err != nil {
		return nil, err
	}

	// Return the ArtistTopTracks
	return &artistTopTracks, nil
}

// GetRelatedArtists implements the ArtistService's interface GetRelatedArtists method.
func (service *DefaultArtistService) GetRelatedArtists(input models.GetRelatedArtistsRequest) (*models.Artists, error) {
	// Validate the input
	if input.Id == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointRelatedArtists, input.Id)

	// Add inputs to the query parameters
	params := map[string]string{"id": input.Id}

	// Make an API call
	res, err := service.client.Get(context.Background(), endpoint, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetRelatedArtists, Err: err}
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

	// Unmarshal the response data into Artists struct
	var relatedArtists models.Artists
	if err := json.Unmarshal(data, &relatedArtists); err != nil {
		return nil, err
	}

	// Return the Artists
	return &relatedArtists, nil
}
