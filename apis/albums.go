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

// AlbumService interface defines the methods for interacting with the Spotify Album's API.
type AlbumService interface {
	// Get Spotify catalog information for a single album.
	GetAlbum(input models.GetAlbumRequest) (*models.Album, error)

	// Get Spotify catalog information for multiple albums identified by their Spotify IDs.
	GetAlbums(input models.GetAlbumsRequest) (*models.Albums, error)

	// Get Spotify catalog information about an album’s tracks. Optional parameters can be used to limit the number of tracks returned.
	GetAlbumTracks(input models.GetAlbumTracksRequest) (*models.AlbumTracks, error)

	// Get a list of the albums saved in the current Spotify user's 'Your Music' library.
	// Authorization scopes: user-library-read
	GetSavedAlbums(input models.GetSavedAlbumsRequest) (*models.SavedAlbums, error)

	// Save one or more albums to the current user's 'Your Music' library.
	// Authorization scopes: user-library-modify
	SaveAlbums(input models.SaveAlbumsRequest) error

	// Remove one or more albums from the current user's 'Your Music' library.
	// Authorization scopes: user-library-modify
	RemoveAlbums(input models.RemoveAlbumsRequest) error

	// Check if one or more albums is already saved in the current Spotify user's 'Your Music' library.
	// Authorization scopes: user-library-read
	CheckSavedAlbums(input models.CheckSavedAlbumsRequest) (*models.CheckSavedAlbums, error)

	// Get a list of new album releases featured in Spotify (shown, for example, on a Spotify player’s “Browse” tab).
	GetNewReleases(input models.GetNewReleasesRequest) (*models.NewlyReleasedAlbums, error)
}

// DefaultAlbumService is a struct that implements AlbumService interface.
type DefaultAlbumService struct {
	client *utils.HttpClient
}

// NewDefaultAlbumService initializes the DefaultAlbumService with given dependencies.
func NewDefaultAlbumService(client *utils.HttpClient) *DefaultAlbumService {
	return &DefaultAlbumService{client}
}

// GetAlbum implements the AlbumService's interface GetAlbum method.
func (service *DefaultAlbumService) GetAlbum(input models.GetAlbumRequest) (*models.Album, error) {
	// Validate the input
	if input.Id == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointAlbum, input.Id)

	// Add inputs to the query parameters
	params := map[string]string{"market": input.Market}

	// Make an API call
	res, err := service.client.Get(context.Background(), endpoint, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetAlbum, Err: err}
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

	// Unmarshal the response data into Album struct
	var album models.Album
	if err := json.Unmarshal(data, &album); err != nil {
		return nil, err
	}

	// Return the Album
	return &album, nil
}

// GetAlbums implements the AlbumService's interface GetAlbums method.
func (service *DefaultAlbumService) GetAlbums(input models.GetAlbumsRequest) (*models.Albums, error) {
	// Validate the input
	if input.Ids == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdsRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"ids": input.Ids, "market": input.Market}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointAlbums, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetAlbums, Err: err}
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

	// Unmarshal the response data into Albums struct
	var albums models.Albums
	if err := json.Unmarshal(data, &albums); err != nil {
		return nil, err
	}

	// Return the Albums
	return &albums, nil
}

// GetAlbumTracks implements the AlbumService's interface GetAlbumTracks method.
func (service *DefaultAlbumService) GetAlbumTracks(input models.GetAlbumTracksRequest) (*models.AlbumTracks, error) {
	// Validate the input
	if input.Id == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointAlbumTracks, input.Id)

	// Add inputs to the query parameters
	params := map[string]string{"market": input.Market, "limit": strconv.Itoa(input.Limit), "offset": strconv.Itoa(input.Offset)}

	// Make an API call
	res, err := service.client.Get(context.Background(), endpoint, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetTracks, Err: err}
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

	// Unmarshal the response data into Tracks struct
	var tracks models.AlbumTracks
	if err := json.Unmarshal(data, &tracks); err != nil {
		return nil, err
	}

	// Return the Tracks
	return &tracks, nil
}

// GetSavedAlbums implements the AlbumService's interface GetSavedAlbums method.
func (service *DefaultAlbumService) GetSavedAlbums(input models.GetSavedAlbumsRequest) (*models.SavedAlbums, error) {
	// Add inputs to the query parameters
	params := map[string]string{"limit": strconv.Itoa(input.Limit), "offset": strconv.Itoa(input.Offset), "market": input.Market}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointMyAlbums, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetTracks, Err: err}
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

	// Unmarshal the response data into SavedAlbums struct
	var savedAlbums models.SavedAlbums
	if err := json.Unmarshal(data, &savedAlbums); err != nil {
		return nil, err
	}

	// Return the SavedAlbums
	return &savedAlbums, nil
}

// SaveAlbums implements the AlbumService's interface SaveAlbums method.
func (service *DefaultAlbumService) SaveAlbums(input models.SaveAlbumsRequest) error {
	// Validate the input
	if input.Ids == "" {
		return &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdsRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"ids": input.Ids}

	// Make an API call
	res, err := service.client.Put(context.Background(), consts.EndpointMyAlbums, nil, params, input.Body)
	if err != nil {
		return &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToSaveAlbums, Err: err}
	}

	// Handle Spotify API error
	if res.StatusCode != http.StatusOK {
		return utils.ParseSpotifyError(res, utils.RegErrorType)
	}

	// Return the empty response
	return nil
}

// RemoveAlbums implements the AlbumService's interface RemoveAlbums method.
func (service *DefaultAlbumService) RemoveAlbums(input models.RemoveAlbumsRequest) error {
	// Validate the input
	if input.Ids == "" {
		return &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdsRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"ids": input.Ids}

	// Make an API call
	res, err := service.client.Delete(context.Background(), consts.EndpointMyAlbums, nil, params, input.Body)
	if err != nil {
		return &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToRemoveAlbums, Err: err}
	}

	// Handle Spotify API error
	if res.StatusCode != http.StatusOK {
		return utils.ParseSpotifyError(res, utils.RegErrorType)
	}

	// Return the empty response
	return nil
}

// CheckSavedAlbums implements the AlbumService's interface CheckSavedAlbums method.
func (service *DefaultAlbumService) CheckSavedAlbums(input models.CheckSavedAlbumsRequest) (*models.CheckSavedAlbums, error) {
	// Validate the input
	if input.Ids == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdsRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"ids": input.Ids}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointCheckMyAlbums, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToCheckSavedAlbums, Err: err}
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

	// Unmarshal the response data into CheckSavedAlbums struct
	var checkSavedAlbums models.CheckSavedAlbums
	if err := json.Unmarshal(data, &checkSavedAlbums); err != nil {
		return nil, err
	}

	// Return the CheckSavedAlbums
	return &checkSavedAlbums, nil
}

// GetNewReleases implements the AlbumService's interface GetNewReleases method.
func (service *DefaultAlbumService) GetNewReleases(input models.GetNewReleasesRequest) (*models.NewlyReleasedAlbums, error) {
	// Add inputs to the query parameters
	params := map[string]string{"limit": strconv.Itoa(input.Limit), "offset": strconv.Itoa(input.Offset)}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointNewReleases, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetNewReleases, Err: err}
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

	// Unmarshal the response data into NewlyReleasedAlbums struct
	var newlyReleasedAlbums models.NewlyReleasedAlbums
	if err := json.Unmarshal(data, &newlyReleasedAlbums); err != nil {
		return nil, err
	}

	// Return the NewlyReleasedAlbums
	return &newlyReleasedAlbums, nil
}
