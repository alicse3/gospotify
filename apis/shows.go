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

// ShowService interface defines the methods for interacting with the Spotify Show's API.
type ShowService interface {
	// Get Spotify catalog information for a single show identified by its unique Spotify ID.
	// Authorization scopes: user-read-playback-position
	GetShow(input models.GetShowRequest) (*models.Show, error)

	// Get Spotify catalog information for several shows based on their Spotify IDs.
	GetShows(input models.GetShowsRequest) (*models.Shows, error)

	// Get Spotify catalog information about an show’s episodes.
	// Optional parameters can be used to limit the number of episodes returned.
	// Authorization scopes: user-read-playback-position
	GetShowEpisodes(input models.GetShowEpisodesRequest) (*models.ShowEpisodes, error)

	// Get a list of shows saved in the current Spotify user's library.
	// Optional parameters can be used to limit the number of shows returned.
	// Authorization scopes: user-library-read
	GetSavedShows(input models.GetSavedShowsRequest) (*models.SavedShows, error)

	// Save one or more shows to current Spotify user's library.
	// Authorization scopes: user-library-modify
	SaveShows(input models.SaveShowsRequest) error

	// Delete one or more shows from current Spotify user's library.
	// Authorization scopes: user-library-modify
	RemoveSavedShows(input models.RemoveShowsRequest) error

	// Check if one or more shows is already saved in the current Spotify user's library.
	// Authorization scopes: user-library-read
	CheckSavedShows(input models.CheckSavedShowsRequest) (*models.CheckSavedShows, error)
}

// DefaultShowService is a struct that implements ShowService interface.
type DefaultShowService struct {
	client *utils.HttpClient
}

// NewDefaultShowService initializes the DefaultShowService with given dependencies.
func NewDefaultShowService(client *utils.HttpClient) *DefaultShowService {
	return &DefaultShowService{client}
}

// GetShow implements the ShowService's interface GetShow method.
func (service *DefaultShowService) GetShow(input models.GetShowRequest) (*models.Show, error) {
	// Validate the input
	if input.Id == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointShow, input.Id)

	// Add inputs to the query parameters
	params := map[string]string{"id": input.Id, "market": input.Market}

	// Make an API call
	res, err := service.client.Get(context.Background(), endpoint, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetShow, Err: err}
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

	// Unmarshal the response data into Show struct
	var show models.Show
	if err := json.Unmarshal(data, &show); err != nil {
		return nil, err
	}

	// Return the Show
	return &show, nil
}

// GetShows implements the ShowService's interface GetShows method.
func (service *DefaultShowService) GetShows(input models.GetShowsRequest) (*models.Shows, error) {
	// Validate the input
	if input.Ids == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdsRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"ids": input.Ids, "market": input.Market}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointShows, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetShows, Err: err}
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

	// Unmarshal the response data into Shows struct
	var shows models.Shows
	if err := json.Unmarshal(data, &shows); err != nil {
		return nil, err
	}

	// Return the Shows
	return &shows, nil
}

// GetShowEpisodes implements the ShowService's interface GetShowEpisodes method.
func (service *DefaultShowService) GetShowEpisodes(input models.GetShowEpisodesRequest) (*models.ShowEpisodes, error) {
	// Validate the input
	if input.Id == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointShowEpisodes, input.Id)

	// Add inputs to the query parameters
	params := map[string]string{"id": input.Id, "market": input.Market, "limit": strconv.Itoa(input.Limit), "offset": strconv.Itoa(input.Offset)}

	// Make an API call
	res, err := service.client.Get(context.Background(), endpoint, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetShowEpisodes, Err: err}
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

	// Unmarshal the response data into ShowEpisodes struct
	var showEpisodes models.ShowEpisodes
	if err := json.Unmarshal(data, &showEpisodes); err != nil {
		return nil, err
	}

	// Return the ShowEpisodes
	return &showEpisodes, nil
}

// GetSavedShows implements the ShowService's interface GetSavedShows method.
func (service *DefaultShowService) GetSavedShows(input models.GetSavedShowsRequest) (*models.SavedShows, error) {
	// Add inputs to the query parameters
	params := map[string]string{"limit": strconv.Itoa(input.Limit), "offset": strconv.Itoa(input.Offset)}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointSaveShows, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetSavedShows, Err: err}
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

	// Unmarshal the response data into SavedShows struct
	var savedShows models.SavedShows
	if err := json.Unmarshal(data, &savedShows); err != nil {
		return nil, err
	}

	// Return the SavedShows
	return &savedShows, nil
}

// SaveShows implements the ShowService's interface SaveShows method.
func (service *DefaultShowService) SaveShows(input models.SaveShowsRequest) error {
	// Validate the input
	if input.Ids == "" {
		return &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdsRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"ids": input.Ids}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointSaveShows, params)
	if err != nil {
		return &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToSaveShows, Err: err}
	}

	// Handle Spotify API error
	if res.StatusCode != http.StatusOK {
		return utils.ParseSpotifyError(res, utils.RegErrorType)
	}

	// Return the empty response
	return nil
}

// RemoveSavedShows implements the ShowService's interface RemoveSavedShows method.
func (service *DefaultShowService) RemoveSavedShows(input models.RemoveShowsRequest) error {
	// Validate the input
	if input.Ids == "" {
		return &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdsRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"ids": input.Ids, "market": input.Market}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointSaveShows, params)
	if err != nil {
		return &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToRemoveSavedShows, Err: err}
	}

	// Handle Spotify API error
	if res.StatusCode != http.StatusOK {
		return utils.ParseSpotifyError(res, utils.RegErrorType)
	}

	// Return the empty response
	return nil
}

// CheckSavedShows implements the ShowService's interface CheckSavedShows method.
func (service *DefaultShowService) CheckSavedShows(input models.CheckSavedShowsRequest) (*models.CheckSavedShows, error) {
	// Validate the input
	if input.Ids == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdsRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"ids": input.Ids}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointCheckSavedShows, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToCheckSavedShows, Err: err}
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

	// Unmarshal the response data into CheckSavedShows struct
	var checkSavedShows models.CheckSavedShows
	if err := json.Unmarshal(data, &checkSavedShows); err != nil {
		return nil, err
	}

	// Return the CheckSavedShows
	return &checkSavedShows, nil
}
