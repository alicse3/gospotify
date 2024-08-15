package apis

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/alicse3/gospotify/consts"
	"github.com/alicse3/gospotify/models"
	"github.com/alicse3/gospotify/utils"
)

// PlayerService interface defines the methods for interacting with the Spotify Player's API.
type PlayerService interface {
	GetPlaybackState(models.GetPlaybackStateRequest) (*models.PlaybackState, error)
	TransferPlayback(models.TransferPlaybackRequest) error
	GetAvailableDevices() (*models.AvailableDevices, error)
	GetCurrentlyPlayingTrack(models.GetCurrentlyPlayingTrackRequest) (*models.PlaybackState, error)
	StartOrResumePlayback(models.StartOrResumePlaybackRequest) error
	PausePlayback(models.PausePlaybackRequest) error
	SkipToNext(models.SkipToNextRequest) error
	SkipToPrevious(models.SkipToPreviousRequest) error
	SeekToPosition(models.SeekToPositionRequest) error
	SetRepeatMode(models.SetRepeatModeRequest) error
	SetPlaybackVolume(models.SetPlaybackVolumeRequest) error
	TogglePlaybackShuffle(models.TogglePlaybackShuffleRequest) error
	GetRecentlyPlayedTracks(models.GetRecentlyPlayedTracksRequest) (*models.RecentlyPlayedTracks, error)
	GetUsersQueue() (*models.UsersQueue, error)
	AddItemToPlaybackQueue(models.AddItemToPlaybackQueueRequest) error
}

// DefaultPlayerService is a struct that implements PlayerService interface.
type DefaultPlayerService struct {
	client *utils.HttpClient
}

// NewDefaultPlayerService initializes the DefaultPlayerService with given dependencies.
func NewDefaultPlayerService(client *utils.HttpClient) *DefaultPlayerService {
	return &DefaultPlayerService{client}
}

// GetPlaybackState implements the DefaultPlayerService's interface GetPlaybackState method.
func (service *DefaultPlayerService) GetPlaybackState(input models.GetPlaybackStateRequest) (*models.PlaybackState, error) {
	// Add inputs to the query parameters
	params := map[string]string{"market": input.Market, "additional_types": input.AdditionalTypes}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointPlaybackState, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetPlaybackState, Err: err}
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

	// Unmarshal the response data into PlaybackState struct
	var playbackState models.PlaybackState
	if err := json.Unmarshal(data, &playbackState); err != nil {
		return nil, err
	}

	// Return the PlaybackState
	return &playbackState, nil
}

// TransferPlayback implements the DefaultPlayerService's interface TransferPlayback method.
func (service *DefaultPlayerService) TransferPlayback(input models.TransferPlaybackRequest) error {
	// Validate the input
	if len(input.Body.DeviceIds) == 0 {
		return &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgDeviceIdsRequired}
	}

	// Add inputs to the body
	body := &models.TransferPlaybackRequestBody{DeviceIds: input.Body.DeviceIds, Play: input.Body.Play}

	// Make an API call
	res, err := service.client.Put(context.Background(), consts.EndpointPlaybackState, nil, nil, body)
	if err != nil {
		return &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToTransferPlayback, Err: err}
	}

	// Handle Spotify API error
	if res.StatusCode != http.StatusOK {
		return utils.ParseSpotifyError(res, utils.RegErrorType)
	}

	// Return the empty response
	return nil
}

// GetAvailableDevices implements the DefaultPlayerService's interface GetAvailableDevices method.
func (service *DefaultPlayerService) GetAvailableDevices() (*models.AvailableDevices, error) {
	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointAvailableDevices, nil)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetAvailableDevices, Err: err}
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

	// Unmarshal the response data into AvailableDevices struct
	var availableDevices models.AvailableDevices
	if err := json.Unmarshal(data, &availableDevices); err != nil {
		return nil, err
	}

	// Return the AvailableDevices
	return &availableDevices, nil
}

// GetCurrentlyPlayingTrack implements the DefaultPlayerService's interface GetCurrentlyPlayingTrack method.
func (service *DefaultPlayerService) GetCurrentlyPlayingTrack(input models.GetCurrentlyPlayingTrackRequest) (*models.PlaybackState, error) {
	// Add inputs to the query parameters
	params := map[string]string{"market": input.Market, "additional_types": input.AdditionalTypes}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointCurrentlyPlayingTrack, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetCurrentlyPlayingTrack, Err: err}
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

	// Unmarshal the response data into PlaybackState struct
	var currentlyPlayingTrack models.PlaybackState
	if err := json.Unmarshal(data, &currentlyPlayingTrack); err != nil {
		return nil, err
	}

	// Return the PlaybackState
	return &currentlyPlayingTrack, nil
}

// StartOrResumePlayback implements the DefaultPlayerService's interface StartOrResumePlayback method.
func (service *DefaultPlayerService) StartOrResumePlayback(input models.StartOrResumePlaybackRequest) error {
	// Add inputs to the query parameters
	params := map[string]string{"device_id": input.DeviceId}

	// Add inputs to the body
	body := &models.StartOrResumePlaybackRequestBody{ContextUri: input.Body.ContextUri, Uris: input.Body.Uris, Offset: input.Body.Offset, PositionMs: input.Body.PositionMs}

	// Make an API call
	res, err := service.client.Put(context.Background(), consts.EndpointStartOrResumePlayback, nil, params, body)
	if err != nil {
		return &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToStartOrResumePlayback, Err: err}
	}

	// Handle Spotify API error
	if res.StatusCode != http.StatusOK {
		return utils.ParseSpotifyError(res, utils.RegErrorType)
	}

	// Return the empty response
	return nil
}

// PausePlayback implements the DefaultPlayerService's interface PausePlayback method.
func (service *DefaultPlayerService) PausePlayback(input models.PausePlaybackRequest) error {
	// Add inputs to the query parameters
	params := map[string]string{"device_id": input.DeviceId}

	// Make an API call
	res, err := service.client.Put(context.Background(), consts.EndpointPausePlayback, nil, params, nil)
	if err != nil {
		return &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToPausePlayback, Err: err}
	}

	// Handle Spotify API error
	if res.StatusCode != http.StatusOK {
		return utils.ParseSpotifyError(res, utils.RegErrorType)
	}

	// Return the empty response
	return nil
}

// SkipToNext implements the DefaultPlayerService's interface SkipToNext method.
func (service *DefaultPlayerService) SkipToNext(input models.SkipToNextRequest) error {
	// Add inputs to the query parameters
	params := map[string]string{"device_id": input.DeviceId}

	// Make an API call
	res, err := service.client.Post(context.Background(), consts.EndpointSkipToNext, nil, params, nil, nil)
	if err != nil {
		return &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToSkipToNext, Err: err}
	}

	// Handle Spotify API error
	if res.StatusCode != http.StatusOK {
		return utils.ParseSpotifyError(res, utils.RegErrorType)
	}

	// Return the empty response
	return nil
}

// SkipToPrevious implements the DefaultPlayerService's interface SkipToPrevious method.
func (service *DefaultPlayerService) SkipToPrevious(input models.SkipToPreviousRequest) error {
	// Add inputs to the query parameters
	params := map[string]string{"device_id": input.DeviceId}

	// Make an API call
	res, err := service.client.Post(context.Background(), consts.EndpointSkipToPrevious, nil, params, nil, nil)
	if err != nil {
		return &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToSkipToPrevious, Err: err}
	}

	// Handle Spotify API error
	if res.StatusCode != http.StatusOK {
		return utils.ParseSpotifyError(res, utils.RegErrorType)
	}

	// Return the empty response
	return nil
}

// SeekToPosition implements the DefaultPlayerService's interface SeekToPosition method.
func (service *DefaultPlayerService) SeekToPosition(input models.SeekToPositionRequest) error {
	// Validate the input
	if input.PositionMs < 0 {
		return &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgMustBePositiveNumber}
	}

	// Add inputs to the query parameters
	params := map[string]string{"position_ms": strconv.Itoa(input.PositionMs), "device_id": input.DeviceId}

	// Make an API call
	res, err := service.client.Put(context.Background(), consts.EndpointSeekToPosition, nil, params, nil)
	if err != nil {
		return &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToSeekToPosition, Err: err}
	}

	// Handle Spotify API error
	if res.StatusCode != http.StatusOK {
		return utils.ParseSpotifyError(res, utils.RegErrorType)
	}

	// Return the empty response
	return nil
}

// SetRepeatMode implements the DefaultPlayerService's interface SetRepeatMode method.
func (service *DefaultPlayerService) SetRepeatMode(input models.SetRepeatModeRequest) error {
	// Validate the input
	if input.State == "" {
		return &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgStateRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"state": input.State, "device_id": input.DeviceId}

	// Make an API call
	res, err := service.client.Put(context.Background(), consts.EndpointRepeatMode, nil, params, nil)
	if err != nil {
		return &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToSetRepeatMode, Err: err}
	}

	// Handle Spotify API error
	if res.StatusCode != http.StatusOK {
		return utils.ParseSpotifyError(res, utils.RegErrorType)
	}

	// Return the empty response
	return nil
}

// SetPlaybackVolume implements the DefaultPlayerService's interface SetPlaybackVolume method.
func (service *DefaultPlayerService) SetPlaybackVolume(input models.SetPlaybackVolumeRequest) error {
	// Validate the input
	if input.VolumePercent < 0 || input.VolumePercent > 100 {
		return &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgVolumePercentMustBeInclusive}
	}

	// Add inputs to the query parameters
	params := map[string]string{"volume_percent": strconv.Itoa(input.VolumePercent), "device_id": input.DeviceId}

	// Make an API call
	res, err := service.client.Put(context.Background(), consts.EndpointPlaybackVolume, nil, params, nil)
	if err != nil {
		return &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToSetPlaybackVolume, Err: err}
	}

	// Handle Spotify API error
	if res.StatusCode != http.StatusOK {
		return utils.ParseSpotifyError(res, utils.RegErrorType)
	}

	// Return the empty response
	return nil
}

// TogglePlaybackShuffle implements the DefaultPlayerService's interface TogglePlaybackShuffle method.
func (service *DefaultPlayerService) TogglePlaybackShuffle(input models.TogglePlaybackShuffleRequest) error {
	// Add inputs to the query parameters
	params := map[string]string{"state": strconv.FormatBool(input.State), "device_id": input.DeviceId}

	// Make an API call
	res, err := service.client.Put(context.Background(), consts.EndpointTogglePlaybackShuffle, nil, params, nil)
	if err != nil {
		return &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToTogglePlaybackShuffle, Err: err}
	}

	// Handle Spotify API error
	if res.StatusCode != http.StatusOK {
		return utils.ParseSpotifyError(res, utils.RegErrorType)
	}

	// Return the empty response
	return nil
}

// GetRecentlyPlayedTracks implements the DefaultPlayerService's interface GetRecentlyPlayedTracks method.
func (service *DefaultPlayerService) GetRecentlyPlayedTracks(input models.GetRecentlyPlayedTracksRequest) (*models.RecentlyPlayedTracks, error) {
	// Add inputs to the query parameters
	params := map[string]string{"limit": strconv.Itoa(input.Limit), "after": strconv.Itoa(input.After), "before": strconv.Itoa(input.Before)}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointRecentlyPlayedTracks, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetRecentlyPlayedTracks, Err: err}
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

	// Unmarshal the response data into RecentlyPlayedTracks struct
	var recentlyPlayedTracks models.RecentlyPlayedTracks
	if err := json.Unmarshal(data, &recentlyPlayedTracks); err != nil {
		return nil, err
	}

	// Return the RecentlyPlayedTracks
	return &recentlyPlayedTracks, nil
}

// GetUsersQueue implements the DefaultPlayerService's interface GetUsersQueue method.
func (service *DefaultPlayerService) GetUsersQueue() (*models.UsersQueue, error) {
	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointUsersQueue, nil)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetUsersQueue, Err: err}
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

	// Unmarshal the response data into UsersQueue struct
	var usersQueue models.UsersQueue
	if err := json.Unmarshal(data, &usersQueue); err != nil {
		return nil, err
	}

	// Return the UsersQueue
	return &usersQueue, nil
}

// AddItemToPlaybackQueue implements the DefaultPlayerService's interface AddItemToPlaybackQueue method.
func (service *DefaultPlayerService) AddItemToPlaybackQueue(input models.AddItemToPlaybackQueueRequest) error {
	// Validate the input
	if input.Uri == "" {
		return &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgUriRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"uri": input.Uri, "device_id": input.DeviceId}

	// Make an API call
	res, err := service.client.Post(context.Background(), consts.EndpointPlaybackVolume, nil, params, nil, nil)
	if err != nil {
		return &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToAddItemToPlaybackQueue, Err: err}
	}

	// Handle Spotify API error
	if res.StatusCode != http.StatusOK {
		return utils.ParseSpotifyError(res, utils.RegErrorType)
	}

	// Return the empty response
	return nil
}
