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

// AudiobookService interface defines the methods for interacting with the Spotify Audiobook's API.
type AudiobookService interface {
	GetAudiobook(input models.GetAudiobookRequest) (*models.Audiobook, error)
	GetAudiobooks(input models.GetAudiobooksRequest) (*models.Audiobooks, error)
	GetAudiobookChapters(input models.GetAudiobookChaptersRequest) (*models.AudiobookChapters, error)
	GetSavedAudiobooks(input models.GetSavedAudiobooksRequest) (*models.SavedAudiobooks, error)
	SaveAudiobooks(input models.SaveAudiobooksRequest) error
	DeleteAudiobooks(input models.RemoveAudiobooksRequest) error
	CheckSavedAudiobooks(input models.CheckSavedAudiobooksRequest) (*models.CheckSavedAudiobooks, error)
}

// DefaultAudiobookService is a struct that implements AudiobookService interface.
type DefaultAudiobookService struct {
	client *utils.HttpClient
}

// NewDefaultAudiobookService initializes the DefaultAudiobookService with given dependencies.
func NewDefaultAudiobookService(client *utils.HttpClient) *DefaultAudiobookService {
	return &DefaultAudiobookService{client}
}

// GetAudiobook implements the AudiobookService's interface GetAudiobook method.
func (service *DefaultAudiobookService) GetAudiobook(input models.GetAudiobookRequest) (*models.Audiobook, error) {
	// Validate the input
	if input.Id == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointAudiobook, input.Id)

	// Make an API call
	res, err := service.client.Get(context.Background(), endpoint, nil)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetAudiobook, Err: err}
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

	// Unmarshal the response data into an Audiobook struct
	var audiobook models.Audiobook
	if err := json.Unmarshal(data, &audiobook); err != nil {
		return nil, err
	}

	// Return the Audiobook
	return &audiobook, nil
}

// GetAudiobooks implements the AudiobookService's interface GetAudiobooks method.
func (service *DefaultAudiobookService) GetAudiobooks(input models.GetAudiobooksRequest) (*models.Audiobooks, error) {
	// Validate the input
	if input.Ids == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdsRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"ids": input.Ids, "market": input.Market}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointAudiobooks, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetAudiobooks, Err: err}
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

	// Unmarshal the response data into an Audiobooks struct
	var audiobooks models.Audiobooks
	if err := json.Unmarshal(data, &audiobooks); err != nil {
		return nil, err
	}

	// Return the Audiobooks
	return &audiobooks, nil
}

// GetAudiobookChapters implements the AudiobookService's interface GetAudiobookChapters method.
func (service *DefaultAudiobookService) GetAudiobookChapters(input models.GetAudiobookChaptersRequest) (*models.AudiobookChapters, error) {
	// Validate the input
	if input.Id == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointAudiobookChapters, input.Id)

	// Add inputs to the query parameters
	params := map[string]string{"id": input.Id, "market": input.Market, "limit": strconv.Itoa(input.Limit), "offset": strconv.Itoa(input.Offset)}

	// Make an API call
	res, err := service.client.Get(context.Background(), endpoint, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetAudiobookChapters, Err: err}
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

	// Unmarshal the response data into an AudiobookChapters struct
	var audiobookChapters models.AudiobookChapters
	if err := json.Unmarshal(data, &audiobookChapters); err != nil {
		return nil, err
	}

	// Return the AudiobookChapters
	return &audiobookChapters, nil
}

// GetSavedAudiobooks implements the AudiobookService's interface GetSavedAudiobooks method.
func (service *DefaultAudiobookService) GetSavedAudiobooks(input models.GetSavedAudiobooksRequest) (*models.SavedAudiobooks, error) {
	// Add inputs to the query parameters
	params := map[string]string{"limit": strconv.Itoa(input.Limit), "offset": strconv.Itoa(input.Offset)}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointMyAudiobooks, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetSavedAudiobooks, Err: err}
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

	// Unmarshal the response data into an SavedAudiobooks struct
	var savedAudiobooks models.SavedAudiobooks
	if err := json.Unmarshal(data, &savedAudiobooks); err != nil {
		return nil, err
	}

	// Return the SavedAudiobooks
	return &savedAudiobooks, nil
}

// SaveAudiobooks implements the AudiobookService's interface SaveAudiobooks method.
func (service *DefaultAudiobookService) SaveAudiobooks(input models.SaveAudiobooksRequest) error {
	// Validate the input
	if input.Ids == "" {
		return &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdsRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"ids": input.Ids}

	// Make an API call
	res, err := service.client.Put(context.Background(), consts.EndpointMyAudiobooks, nil, params, nil)
	if err != nil {
		return &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToSaveAudiobooks, Err: err}
	}

	// Handle Spotify API error
	if res.StatusCode != http.StatusOK {
		return utils.ParseSpotifyError(res, utils.RegErrorType)
	}

	// Return the empty response
	return nil
}

// DeleteAudiobooks implements the AudiobookService's interface DeleteAudiobooks method.
func (service *DefaultAudiobookService) DeleteAudiobooks(input models.RemoveAudiobooksRequest) error {
	// Validate the input
	if input.Ids == "" {
		return &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdsRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"ids": input.Ids}

	// Make an API call
	res, err := service.client.Delete(context.Background(), consts.EndpointMyAudiobooks, nil, params, nil)
	if err != nil {
		return &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToSaveAudiobooks, Err: err}
	}

	// Handle Spotify API error
	if res.StatusCode != http.StatusOK {
		return utils.ParseSpotifyError(res, utils.RegErrorType)
	}

	// Return the empty response
	return nil
}

// CheckSavedAudiobooks implements the AudiobookService's interface CheckSavedAudiobooks method.
func (service *DefaultAudiobookService) CheckSavedAudiobooks(input models.CheckSavedAudiobooksRequest) (*models.CheckSavedAudiobooks, error) {
	// Validate the input
	if input.Ids == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdsRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"ids": input.Ids}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointMySavedAudiobooks, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToCheckSavedAudiobooks, Err: err}
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

	// Unmarshal the response data into an CheckSavedAudiobooks struct
	var checkSavedAudiobooks models.CheckSavedAudiobooks
	if err := json.Unmarshal(data, &checkSavedAudiobooks); err != nil {
		return nil, err
	}

	// Return the CheckSavedAudiobooks
	return &checkSavedAudiobooks, nil
}
