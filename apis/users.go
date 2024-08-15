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

// UserService interface defines the methods for interacting with the Spotify User's API.
type UserService interface {
	GetCurrentUserProfile() (*models.User, error)
	GetUserTopItems(input models.GetUsersTopItemsRequest) (*models.UserTopItems, error)
	GetUsersProfile(input models.GetUsersProfileRequest) (*models.UserProfile, error)
	FollowPlaylist(input models.FollowPlaylistRequest) error
	UnfollowPlaylist(input models.UnfollowPlaylistRequest) error
	GetFollowedArtists(input models.GetFollowedArtistsRequest) (*models.FollowedArtists, error)
	FollowArtistsOrUsers(input models.FollowArtistsOrUsersRequest) error
	UnfollowArtistsOrUsers(input models.UnfollowArtistsOrUsersRequest) error
	CheckUserFollowsArtistsOrUsers(input models.UserFollowsArtistsOrUsersRequest) (*models.CheckUserFollowsArtistsOrUsers, error)
	CheckCurrentUserFollowsPlaylist(input models.CurrentUserFollowsPlaylistRequest) (*models.CheckCurrentUserFollowsPlaylist, error)
}

// DefaultUserService is a struct that implements UserService interface.
type DefaultUserService struct {
	client *utils.HttpClient
}

// NewDefultUserService initializes the DefaultUserService with given dependencies.
func NewDefultUserService(client *utils.HttpClient) *DefaultUserService {
	return &DefaultUserService{client}
}

// GetCurrentUserProfile implements the UserService's interface GetCurrentUserProfile method.
func (service *DefaultUserService) GetCurrentUserProfile() (*models.User, error) {
	// Make a Get call
	res, err := service.client.Get(context.Background(), consts.EndpointMe, nil)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetCurrentUserProfile, Err: err}
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

	// Unmarshal the response data into User struct
	var user models.User
	if err := json.Unmarshal(data, &user); err != nil {
		return nil, err
	}

	// Return the User
	return &user, nil
}

// GetUserTopItems implements the UserService's interface GetUserTopItems method.
func (service *DefaultUserService) GetUserTopItems(input models.GetUsersTopItemsRequest) (*models.UserTopItems, error) {
	// Validate the input
	if input.Type == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgTypeRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointUserTopItems, input.Type)

	// Add inputs to the query parameters
	params := map[string]string{"type": input.Type, "time_range": input.TimeRange, "limit": strconv.Itoa(input.Limit), "offset": strconv.Itoa(input.Offset)}

	// Make an API call
	res, err := service.client.Get(context.Background(), endpoint, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetUserTopItems, Err: err}
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

	// Unmarshal the response data into UserTopItems struct
	var userTopItems models.UserTopItems
	if err := json.Unmarshal(data, &userTopItems); err != nil {
		return nil, err
	}

	// Return the UserTopItems
	return &userTopItems, nil
}

// GetUsersProfile implements the UserService's interface GetUsersProfile method.
func (service *DefaultUserService) GetUsersProfile(input models.GetUsersProfileRequest) (*models.UserProfile, error) {
	// Validate the input
	if input.UserId == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgUserIdRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointUserProfile, input.UserId)

	// Add inputs to the query parameters
	params := map[string]string{"user_id": input.UserId}

	// Make an API call
	res, err := service.client.Get(context.Background(), endpoint, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetUserProfile, Err: err}
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

	// Unmarshal the response data into UserProfile struct
	var userProfile models.UserProfile
	if err := json.Unmarshal(data, &userProfile); err != nil {
		return nil, err
	}

	// Return the UserProfile
	return &userProfile, nil
}

// FollowPlaylist implements the UserService's interface FollowPlaylist method.
func (service *DefaultUserService) FollowPlaylist(input models.FollowPlaylistRequest) error {
	// Validate the input
	if input.PlaylistId == "" {
		return &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgPlaylistIdRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointFollowers, input.PlaylistId)

	// Add inputs to the query parameters
	params := map[string]string{"playlist_id": input.PlaylistId}

	// Make an API call
	res, err := service.client.Put(context.Background(), endpoint, nil, params, input.Body)
	if err != nil {
		return &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToFollowPlaylist, Err: err}
	}

	// Handle Spotify API error
	if res.StatusCode != http.StatusOK {
		return utils.ParseSpotifyError(res, utils.RegErrorType)
	}

	// Return the empty response
	return nil
}

// UnfollowPlaylist implements the UserService's interface UnfollowPlaylist method.
func (service *DefaultUserService) UnfollowPlaylist(input models.UnfollowPlaylistRequest) error {
	// Validate the input
	if input.PlaylistId == "" {
		return &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgPlaylistIdRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointFollowers, input.PlaylistId)

	// Add inputs to the query parameters
	params := map[string]string{"playlist_id": input.PlaylistId}

	// Make an API call
	res, err := service.client.Delete(context.Background(), endpoint, nil, params, nil)
	if err != nil {
		return &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToUnfollowPlaylist, Err: err}
	}

	// Handle Spotify API error
	if res.StatusCode != http.StatusOK {
		return utils.ParseSpotifyError(res, utils.RegErrorType)
	}

	// Return the empty response
	return nil
}

// GetFollowedArtists implements the UserService's interface GetFollowedArtists method.
func (service *DefaultUserService) GetFollowedArtists(input models.GetFollowedArtistsRequest) (*models.FollowedArtists, error) {
	// Validate the input
	if input.Type == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgTypeRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"type": input.Type, "after": input.After, "limit": strconv.Itoa(input.Limit)}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointFollowing, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetFollowedArtists, Err: err}
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

	// Unmarshal the response data into FollowedArtists struct
	var followedArtists models.FollowedArtists
	if err := json.Unmarshal(data, &followedArtists); err != nil {
		return nil, err
	}

	// Return the FollowedArtists
	return &followedArtists, nil
}

// FollowArtistsOrUsers implements the UserService's interface FollowArtistsOrUsers method.
func (service *DefaultUserService) FollowArtistsOrUsers(input models.FollowArtistsOrUsersRequest) error {
	// Validate the input
	if input.Type == "" {
		return &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgTypeRequired}
	}
	if input.Ids == "" {
		return &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdsRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"type": input.Type, "ids": input.Ids}

	// Make an API call
	res, err := service.client.Put(context.Background(), consts.EndpointFollowing, nil, params, input.Body)
	if err != nil {
		return &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToFollowArtistsOrUsers, Err: err}
	}

	// Handle Spotify API error
	if res.StatusCode != http.StatusOK {
		return utils.ParseSpotifyError(res, utils.RegErrorType)
	}

	// Return the empty response
	return nil
}

// UnfollowArtistsOrUsers implements the UserService's interface UnfollowArtistsOrUsers method.
func (service *DefaultUserService) UnfollowArtistsOrUsers(input models.UnfollowArtistsOrUsersRequest) error {
	// Validate the input
	if input.Type == "" {
		return &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgTypeRequired}
	}
	if input.Ids == "" {
		return &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdsRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"type": input.Type, "ids": input.Ids}

	// Make an API call
	res, err := service.client.Delete(context.Background(), consts.EndpointFollowing, nil, params, input.Body)
	if err != nil {
		return &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToUnfollowArtistsOrUsers, Err: err}
	}

	// Handle Spotify API error
	if res.StatusCode != http.StatusOK {
		return utils.ParseSpotifyError(res, utils.RegErrorType)
	}

	// Return the empty response
	return nil
}

// CheckUserFollowsArtistsOrUsers implements the UserService's interface CheckUserFollowsArtistsOrUsers method.
func (service *DefaultUserService) CheckUserFollowsArtistsOrUsers(input models.UserFollowsArtistsOrUsersRequest) (*models.CheckUserFollowsArtistsOrUsers, error) {
	// Validate the input
	if input.Type == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgTypeRequired}
	}
	if input.Ids == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdsRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"type": input.Type, "ids": input.Ids}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointUserFollowsArtistsOrUsers, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToCheckIfUserFollowsArtistsOrUsers, Err: err}
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

	// Unmarshal the response data into CheckUserFollowsArtistsOrUsers struct
	var checkUserFollowsArtistsOrUsers models.CheckUserFollowsArtistsOrUsers
	if err := json.Unmarshal(data, &checkUserFollowsArtistsOrUsers); err != nil {
		return nil, err
	}

	// Return the CheckUserFollowsArtistsOrUsers
	return &checkUserFollowsArtistsOrUsers, nil
}

// CheckCurrentUserFollowsPlaylist implements the UserService's interface CheckCurrentUserFollowsPlaylist method.
func (service *DefaultUserService) CheckCurrentUserFollowsPlaylist(input models.CurrentUserFollowsPlaylistRequest) (*models.CheckCurrentUserFollowsPlaylist, error) {
	// Validate the input
	if input.PlaylistId == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgPlaylistIdRequired}
	}
	if input.Ids == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdsRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointCurrentUserFollowsPlaylist, input.PlaylistId)

	// Add inputs to the query parameters
	params := map[string]string{"playlist_id": input.PlaylistId, "ids": input.Ids}

	// Make an API call
	res, err := service.client.Get(context.Background(), endpoint, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToCheckIfCurrentUserFollowsPlaylist, Err: err}
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

	// Unmarshal the response data into CheckCurrentUserFollowsPlaylist struct
	var checkCurrentUserFollowsPlaylist models.CheckCurrentUserFollowsPlaylist
	if err := json.Unmarshal(data, &checkCurrentUserFollowsPlaylist); err != nil {
		return nil, err
	}

	// Return the CheckCurrentUserFollowsPlaylist
	return &checkCurrentUserFollowsPlaylist, nil
}
