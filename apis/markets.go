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

// MarketService interface defines the methods for interacting with the Spotify Market's API.
type MarketService interface {
	// Get the list of markets where Spotify is available.
	GetAvailableMarkets() (*models.Markets, error)
}

// DefaultMarketService is a struct that implements MarketService interface.
type DefaultMarketService struct {
	client *utils.HttpClient
}

// NewDefaultMarketService initializes the DefaultMarketService with given dependencies.
func NewDefaultMarketService(client *utils.HttpClient) *DefaultMarketService {
	return &DefaultMarketService{client}
}

// GetAvailableMarkets implements the DefaultMarketService's interface GetAvailableMarkets method.
func (service *DefaultMarketService) GetAvailableMarkets() (*models.Markets, error) {
	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointGetAvailableMarkets, nil)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetAvailableMarkets, Err: err}
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

	// Unmarshal the response data into Markets struct
	var markets models.Markets
	if err := json.Unmarshal(data, &markets); err != nil {
		return nil, err
	}

	// Return the Markets
	return &markets, nil
}
