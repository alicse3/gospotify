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

// EpisodeService interface defines the methods for interacting with the Spotify Episode's API.
type EpisodeService interface {
	GetEpisode(input models.GetEpisodeRequest) (*models.Episode, error)
	GetEpisodes(input models.GetEpisodesRequest) (*models.Episodes, error)
	GetSavedEpisodes(input models.GetSavedEpisodesRequest) (*models.SavedEpisodes, error)
	SaveEpisodes(input models.SaveEpisodesRequest) error
	RemoveEpisodes(input models.RemoveEpisodesRequest) error
	CheckSavedEpisodes(input models.CheckSavedEpisodesRequest) (*models.CheckSavedEpisodes, error)
}

// DefaultEpisodeService is a struct that implements EpisodeService interface.
type DefaultEpisodeService struct {
	client *utils.HttpClient
}

// NewDefaultEpisodeService initializes the DefaultEpisodeService with given dependencies.
func NewDefaultEpisodeService(client *utils.HttpClient) *DefaultEpisodeService {
	return &DefaultEpisodeService{client}
}

// GetEpisode implements the EpisodeService's interface GetEpisode method.
func (service *DefaultEpisodeService) GetEpisode(input models.GetEpisodeRequest) (*models.Episode, error) {
	// Validate the input
	if input.Id == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointEpisode, input.Id)

	// Add inputs to the query parameters
	params := map[string]string{"id": input.Id, "market": input.Market}

	// Make an API call
	res, err := service.client.Get(context.Background(), endpoint, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetEpisode, Err: err}
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

	// Unmarshal the response data into an Episode struct
	var episode models.Episode
	if err := json.Unmarshal(data, &episode); err != nil {
		return nil, err
	}

	// Return the Episode
	return &episode, nil
}

// GetEpisodes implements the EpisodeService's interface GetEpisodes method.
func (service *DefaultEpisodeService) GetEpisodes(input models.GetEpisodesRequest) (*models.Episodes, error) {
	// Validate the input
	if input.Ids == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdsRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"id": input.Ids, "market": input.Market}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointEpisodes, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetEpisodes, Err: err}
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

	// Unmarshal the response data into an Episodes struct
	var episodes models.Episodes
	if err := json.Unmarshal(data, &episodes); err != nil {
		return nil, err
	}

	// Return the Episodes
	return &episodes, nil
}

// GetSavedEpisodes implements the EpisodeService's interface GetSavedEpisodes method.
func (service *DefaultEpisodeService) GetSavedEpisodes(input models.GetSavedEpisodesRequest) (*models.SavedEpisodes, error) {
	// Add inputs to the query parameters
	params := map[string]string{"market": input.Market, "limit": strconv.Itoa(input.Limit), "offset": strconv.Itoa(input.Offset)}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointEpisodes, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetSavedEpisodes, Err: err}
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

	// Unmarshal the response data into an SavedEpisodes struct
	var savedEpisodes models.SavedEpisodes
	if err := json.Unmarshal(data, &savedEpisodes); err != nil {
		return nil, err
	}

	// Return the SavedEpisodes
	return &savedEpisodes, nil
}

// SaveEpisodes implements the EpisodeService's interface SaveEpisodes method.
func (service *DefaultEpisodeService) SaveEpisodes(input models.SaveEpisodesRequest) error {
	// Validate the input
	if input.Ids == "" {
		return &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdsRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"ids": input.Ids}

	// Make an API call
	res, err := service.client.Put(context.Background(), consts.EndpointMyEpisodes, nil, params, nil)
	if err != nil {
		return &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToSaveEpisodes, Err: err}
	}

	// Handle Spotify API error
	if res.StatusCode != http.StatusOK {
		return utils.ParseSpotifyError(res, utils.RegErrorType)
	}

	// Return the empty response
	return nil
}

// RemoveEpisodes implements the EpisodeService's interface RemoveEpisodes method.
func (service *DefaultEpisodeService) RemoveEpisodes(input models.RemoveEpisodesRequest) error {
	// Validate the input
	if input.Ids == "" {
		return &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdsRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"ids": input.Ids}

	// Make an API call
	res, err := service.client.Delete(context.Background(), consts.EndpointMyEpisodes, nil, params, nil)
	if err != nil {
		return &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToRemoveEpisodes, Err: err}
	}

	// Handle Spotify API error
	if res.StatusCode != http.StatusOK {
		return utils.ParseSpotifyError(res, utils.RegErrorType)
	}

	// Return the empty response
	return nil
}

// CheckSavedEpisodes implements the EpisodeService's interface CheckSavedEpisodes method.
func (service *DefaultEpisodeService) CheckSavedEpisodes(input models.CheckSavedEpisodesRequest) (*models.CheckSavedEpisodes, error) {
	// Validate the input
	if input.Ids == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdsRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"id": input.Ids}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointCheckMyEpisodes, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToCheckSavedEpisodes, Err: err}
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

	// Unmarshal the response data into an CheckSavedEpisodes struct
	var checkSavedEpisodes models.CheckSavedEpisodes
	if err := json.Unmarshal(data, &checkSavedEpisodes); err != nil {
		return nil, err
	}

	// Return the CheckSavedEpisodes
	return &checkSavedEpisodes, nil
}
