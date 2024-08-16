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

// PlaylistService interface defines the methods for interacting with the Spotify Playlist's API.
type PlaylistService interface {
	// Get a playlist owned by a Spotify user.
	GetPlaylist(input models.GetPlaylistRequest) (*models.Playlist, error)

	// Change a playlist's name and public/private state. (The user must, of course, own the playlist.)
	// Authorization scopes: playlist-modify-public, playlist-modify-private
	ChangePlaylistDetails(input models.ChangePlaylistDetailsRequest) error

	// Get full details of the items of a playlist owned by a Spotify user.
	// Authorization scopes: playlist-read-private
	GetPlaylistItems(input models.GetPlaylistItemsRequest) (*models.PlaylistItems, error)

	// Either reorder or replace items in a playlist depending on the request's parameters.
	// To reorder items, include range_start, insert_before, range_length and snapshot_id in the request's body.
	// To replace items, include uris as either a query parameter or in the request's body.
	// Replacing items in a playlist will overwrite its existing items.
	// This operation can be used for replacing or clearing items in a playlist.

	// Note: Replace and reorder are mutually exclusive operations which share the same endpoint, but have different parameters.
	// These operations can't be applied together in a single request.

	// Authorization scopes: playlist-modify-public, playlist-modify-private
	UpdatePlaylistItems(input models.UpdatePlaylistItemsRequest) (*models.UpdatePlaylistItems, error)

	// Add one or more items to a user's playlist.
	// Authorization scopes: playlist-modify-public, playlist-modify-private
	AddPlaylistItems(input models.AddPlaylistItemsRequest) (*models.AddPlaylistItems, error)

	// Remove one or more items from a user's playlist.
	// Authorization scopes: playlist-modify-public, playlist-modify-private
	RemovePlaylistItems(input models.RemovePlaylistItemsRequest) (*models.RemovePlaylistItems, error)

	// Get a list of the playlists owned or followed by the current Spotify user.
	// Authorization scopes: playlist-read-private
	GetCurrentUserPlaylists(input models.GetCurrentUsersPlaylistsRequest) (*models.Playlists, error)

	// Get a list of the playlists owned or followed by a Spotify user.
	// Authorization scopes: playlist-read-private, playlist-read-collaborative
	GetUserPlaylists(input models.GetUsersPlaylistsRequest) (*models.Playlists, error)

	// Create a playlist for a Spotify user. (The playlist will be empty until you add tracks.) Each user is generally limited to a maximum of 11000 playlists.
	// Authorization scopes: playlist-modify-public, playlist-modify-private
	CreatePlaylist(input models.CreatePlaylistRequest) (*models.Playlist, error)

	// Get a list of Spotify featured playlists (shown, for example, on a Spotify player's 'Browse' tab).
	GetFeaturedPlaylists(input models.GetFeaturedPlaylistsRequest) (*models.FeaturedPlaylists, error)

	// Get a list of Spotify playlists tagged with a particular category.
	GetCategoryPlaylists(input models.GetCategoryPlaylistsRequest) (*models.CategoryPlaylists, error)

	// Get the current image associated with a specific playlist.
	GetPlaylistCoverImage(input models.GetPlaylistCoverImageRequest) (*models.PlaylistCoverImage, error)

	// Replace the image used to represent a specific playlist.
	// Authorization scopes: ugc-image-upload, playlist-modify-public, playlist-modify-private
	AddCustomPlaylistCoverImage(input models.GetCustomPlaylistCoverImageRequest) error
}

// DefaultPlaylistService is a struct that implements PlaylistService interface.
type DefaultPlaylistService struct {
	client *utils.HttpClient
}

// NewDefaultPlaylistService initializes the DefaultPlaylistService with given dependencies.
func NewDefaultPlaylistService(client *utils.HttpClient) *DefaultPlaylistService {
	return &DefaultPlaylistService{client}
}

// GetPlaylist implements the DefaultPlaylistService's interface GetPlaylist method.
func (service *DefaultPlaylistService) GetPlaylist(input models.GetPlaylistRequest) (*models.Playlist, error) {
	// Validate the input
	if input.PlaylistId == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgPlaylistIdRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointPlaylists, input.PlaylistId)

	// Add inputs to the query parameters
	params := map[string]string{"playlist_id": input.PlaylistId, "market": input.Market, "fields": input.Fields}
	if input.AdditionalTypes != "" {
		params["additional_types"] = input.AdditionalTypes
	}

	// Make an API call
	res, err := service.client.Get(context.Background(), endpoint, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetPlaylist, Err: err}
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

	// Unmarshal the response data into Playlist struct
	var playlist models.Playlist
	if err := json.Unmarshal(data, &playlist); err != nil {
		return nil, err
	}

	// Return the Playlist
	return &playlist, nil
}

// ChangePlaylistDetails implements the DefaultPlaylistService's interface ChangePlaylistDetails method.
func (service *DefaultPlaylistService) ChangePlaylistDetails(input models.ChangePlaylistDetailsRequest) error {
	// Validate the input
	if input.PlaylistId == "" {
		return &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgPlaylistIdRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointPlaylists, input.PlaylistId)

	// Add inputs to the query parameters
	params := map[string]string{"playlist_id": input.PlaylistId}

	// Make an API call
	res, err := service.client.Put(context.Background(), endpoint, nil, params, input.Body)
	if err != nil {
		return &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToChangePlaylistDetails, Err: err}
	}

	// Handle Spotify API error
	if res.StatusCode != http.StatusOK {
		return utils.ParseSpotifyError(res, utils.RegErrorType)
	}

	// Return the empty response
	return nil
}

// GetPlaylistItems implements the DefaultPlaylistService's interface GetPlaylistItems method.
func (service *DefaultPlaylistService) GetPlaylistItems(input models.GetPlaylistItemsRequest) (*models.PlaylistItems, error) {
	// Validate the input
	if input.PlaylistId == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgPlaylistIdRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointPlaylistItems, input.PlaylistId)

	// Add inputs to the query parameters
	params := map[string]string{"playlist_id": input.PlaylistId, "market": input.Market, "fields": input.Fields, "limit": strconv.Itoa(input.Limit), "offset": strconv.Itoa(input.Offset)}
	if input.AdditionalTypes != "" {
		params["additional_types"] = input.AdditionalTypes
	}

	// Make an API call
	res, err := service.client.Get(context.Background(), endpoint, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetPlaylistItems, Err: err}
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

	// Unmarshal the response data into PlaylistItems struct
	var playlistItems models.PlaylistItems
	if err := json.Unmarshal(data, &playlistItems); err != nil {
		return nil, err
	}

	// Return the PlaylistItems
	return &playlistItems, nil
}

// UpdatePlaylistItems implements the DefaultPlaylistService's interface UpdatePlaylistItems method.
func (service *DefaultPlaylistService) UpdatePlaylistItems(input models.UpdatePlaylistItemsRequest) (*models.UpdatePlaylistItems, error) {
	// Validate the input
	if input.PlaylistId == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgPlaylistIdRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointPlaylistItems, input.PlaylistId)

	// Add inputs to the query parameters
	params := map[string]string{"playlist_id": input.PlaylistId, "uris": input.Uris}

	// Make an API call
	res, err := service.client.Put(context.Background(), endpoint, nil, params, input.Body)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToUpdatePlaylistItems, Err: err}
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

	// Unmarshal the response data into UpdatePlaylistItems struct
	var updatePlaylistItems models.UpdatePlaylistItems
	if err := json.Unmarshal(data, &updatePlaylistItems); err != nil {
		return nil, err
	}

	// Return the UpdatePlaylistItems
	return &updatePlaylistItems, nil
}

// AddPlaylistItems implements the DefaultPlaylistService's interface AddPlaylistItems method.
func (service *DefaultPlaylistService) AddPlaylistItems(input models.AddPlaylistItemsRequest) (*models.AddPlaylistItems, error) {
	// Validate the input
	if input.PlaylistId == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgPlaylistIdRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointPlaylistItems, input.PlaylistId)

	// Add inputs to the query parameters
	params := map[string]string{"playlist_id": input.PlaylistId, "position": strconv.Itoa(input.Position), "uris": input.Uris}

	// Make an API call
	res, err := service.client.Post(context.Background(), endpoint, nil, params, nil, input.Body)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToAddPlaylistItems, Err: err}
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

	// Unmarshal the response data into AddPlaylistItems struct
	var addPlaylistItems models.AddPlaylistItems
	if err := json.Unmarshal(data, &addPlaylistItems); err != nil {
		return nil, err
	}

	// Return the AddPlaylistItems
	return &addPlaylistItems, nil
}

// RemovePlaylistItems implements the DefaultPlaylistService's interface RemovePlaylistItems method.
func (service *DefaultPlaylistService) RemovePlaylistItems(input models.RemovePlaylistItemsRequest) (*models.RemovePlaylistItems, error) {
	// Validate the input
	if input.PlaylistId == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgPlaylistIdRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointPlaylistItems, input.PlaylistId)

	// Add inputs to the query parameters
	params := map[string]string{"playlist_id": input.PlaylistId}

	// Make an API call
	res, err := service.client.Delete(context.Background(), endpoint, nil, params, input.Body)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToRemovePlaylistItems, Err: err}
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

	// Unmarshal the response data into RemovePlaylistItems struct
	var removePlaylistItems models.RemovePlaylistItems
	if err := json.Unmarshal(data, &removePlaylistItems); err != nil {
		return nil, err
	}

	// Return the RemovePlaylistItems
	return &removePlaylistItems, nil
}

// GetCurrentUserPlaylists implements the DefaultPlaylistService's interface GetCurrentUserPlaylists method.
func (service *DefaultPlaylistService) GetCurrentUserPlaylists(input models.GetCurrentUsersPlaylistsRequest) (*models.Playlists, error) {
	// Add inputs to the query parameters
	params := map[string]string{"limit": strconv.Itoa(input.Limit), "offset": strconv.Itoa(input.Offset)}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointCurrentUsersPlaylists, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetCurrentUsersPlaylists, Err: err}
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

	// Unmarshal the response data into Playlists struct
	var playlists models.Playlists
	if err := json.Unmarshal(data, &playlists); err != nil {
		return nil, err
	}

	// Return the Playlists
	return &playlists, nil
}

// GetUserPlaylists implements the DefaultPlaylistService's interface GetUserPlaylists method.
func (service *DefaultPlaylistService) GetUserPlaylists(input models.GetUsersPlaylistsRequest) (*models.Playlists, error) {
	// Validate the input
	if input.UserId == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgUserIdRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointUsersPlaylists, input.UserId)

	// Add inputs to the query parameters
	params := map[string]string{"user_id": input.UserId, "limit": strconv.Itoa(input.Limit), "offset": strconv.Itoa(input.Offset)}

	// Make an API call
	res, err := service.client.Get(context.Background(), endpoint, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetUsersItems, Err: err}
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

	// Unmarshal the response data into Playlists struct
	var userPlaylists models.Playlists
	if err := json.Unmarshal(data, &userPlaylists); err != nil {
		return nil, err
	}

	// Return the Playlists
	return &userPlaylists, nil
}

// CreatePlaylist implements the DefaultPlaylistService's interface CreatePlaylist method.
func (service *DefaultPlaylistService) CreatePlaylist(input models.CreatePlaylistRequest) (*models.Playlist, error) {
	// Validate the input
	if input.UserId == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgUserIdRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointUsersPlaylists, input.UserId)

	// Add inputs to the query parameters
	params := map[string]string{"user_id": input.UserId}

	// Make an API call
	res, err := service.client.Post(context.Background(), endpoint, nil, params, nil, input.Body)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToCreatePlaylist, Err: err}
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

	// Unmarshal the response data into Playlist struct
	var playlist models.Playlist
	if err := json.Unmarshal(data, &playlist); err != nil {
		return nil, err
	}

	// Return the Playlist
	return &playlist, nil
}

// GetFeaturedPlaylists implements the DefaultPlaylistService's interface GetFeaturedPlaylists method.
func (service *DefaultPlaylistService) GetFeaturedPlaylists(input models.GetFeaturedPlaylistsRequest) (*models.FeaturedPlaylists, error) {
	// Add inputs to the query parameters
	params := map[string]string{"locale": input.Locale, "limit": strconv.Itoa(input.Limit), "offset": strconv.Itoa(input.Offset)}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointFeaturedPlaylists, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetFeaturedPlaylists, Err: err}
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

	// Unmarshal the response data into FeaturedPlaylists struct
	var featuredPlaylists models.FeaturedPlaylists
	if err := json.Unmarshal(data, &featuredPlaylists); err != nil {
		return nil, err
	}

	// Return the FeaturedPlaylists
	return &featuredPlaylists, nil
}

// GetCategoryPlaylists implements the DefaultPlaylistService's interface GetCategoryPlaylists method.
func (service *DefaultPlaylistService) GetCategoryPlaylists(input models.GetCategoryPlaylistsRequest) (*models.CategoryPlaylists, error) {
	// Validate the input
	if input.CategoryId == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgCategoryIdRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointCategoryPlaylists, input.CategoryId)

	// Add inputs to the query parameters
	params := map[string]string{"category_id": input.CategoryId, "limit": strconv.Itoa(input.Limit), "offset": strconv.Itoa(input.Offset)}

	// Make an API call
	res, err := service.client.Get(context.Background(), endpoint, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetCategoryPlaylists, Err: err}
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

	// Unmarshal the response data into CategoryPlaylists struct
	var categoryPlaylists models.CategoryPlaylists
	if err := json.Unmarshal(data, &categoryPlaylists); err != nil {
		return nil, err
	}

	// Return the CategoryPlaylists
	return &categoryPlaylists, nil
}

// GetPlaylistCoverImage implements the DefaultPlaylistService's interface GetPlaylistCoverImage method.
func (service *DefaultPlaylistService) GetPlaylistCoverImage(input models.GetPlaylistCoverImageRequest) (*models.PlaylistCoverImage, error) {
	// Validate the input
	if input.PlaylistId == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgPlaylistIdRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointPlaylistCoverImage, input.PlaylistId)

	// Add inputs to the query parameters
	params := map[string]string{"playlist_id": input.PlaylistId}

	// Make an API call
	res, err := service.client.Get(context.Background(), endpoint, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetPlaylistCoverImage, Err: err}
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

	// Unmarshal the response data into PlaylistCoverImage struct
	var playlistCoverImage models.PlaylistCoverImage
	if err := json.Unmarshal(data, &playlistCoverImage); err != nil {
		return nil, err
	}

	// Return the PlaylistCoverImage
	return &playlistCoverImage, nil
}

// AddCustomPlaylistCoverImage implements the DefaultPlaylistService's interface AddCustomPlaylistCoverImage method.
func (service *DefaultPlaylistService) AddCustomPlaylistCoverImage(input models.GetCustomPlaylistCoverImageRequest) error {
	// Validate the input
	if input.PlaylistId == "" {
		return &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgPlaylistIdRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"playlist_id": input.PlaylistId}

	// Make an API call
	res, err := service.client.Put(context.Background(), consts.EndpointPlaylistCoverImage, nil, params, input.Body)
	if err != nil {
		return &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToAddCustomPlaylistCoverImage, Err: err}
	}

	// Handle Spotify API error
	if res.StatusCode != http.StatusOK {
		return utils.ParseSpotifyError(res, utils.RegErrorType)
	}

	// Return the empty response
	return nil
}
