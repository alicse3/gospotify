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

// UserService interface defines the methods for interacting with the Spotify User's API.
type UserService interface {
	GetCurrentUserProfile() (*models.User, error)
}

// DefaultUserService is a struct that implements UserService interface.
type DefaultUserService struct {
	client *utils.HttpClient
}

// NewDefultUserService initializes the DefaultUserService with given dependencies.
func NewDefultUserService(client *utils.HttpClient) *DefaultUserService {
	return &DefaultUserService{client}
}

// GetCurrentUserProfile implements one of the UserService interface methods.
func (dus *DefaultUserService) GetCurrentUserProfile() (*models.User, error) {
	// Make a Get call
	res, err := dus.client.Get(context.Background(), consts.EndpointMe, nil)
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
