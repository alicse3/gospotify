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

// SearchService interface defines the methods for interacting with the Spotify Search's API.
type SearchService interface {
	// Get Spotify catalog information about albums, artists, playlists, tracks, shows, episodes or audiobooks that match a keyword string.
	// Audiobooks are only available within the US, UK, Canada, Ireland, New Zealand and Australia markets.
	Search(input models.SearchRequest) (*models.SearchResponse, error)
}

// DefaultSearchService is a struct that implements SearchService interface.
type DefaultSearchService struct {
	client *utils.HttpClient
}

// NewDefaultSearchService initializes the DefaultSearchService with given dependencies.
func NewDefaultSearchService(client *utils.HttpClient) *DefaultSearchService {
	return &DefaultSearchService{client}
}

// Search implements the SearchService's interface Search method.
func (service *DefaultSearchService) Search(input models.SearchRequest) (*models.SearchResponse, error) {
	// Validate the input
	if input.Q == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgSearchQueryRequired}
	}
	if len(input.Type) == 0 {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgSearchTypeRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"q": input.Q, "type": input.Type, "market": input.Market, "limit": strconv.Itoa(input.Limit), "offset": strconv.Itoa(input.Offset), "include_external": input.IncludeExternal}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointSearch, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetSearchResults, Err: err}
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

	// Unmarshal the response data into SearchResponse struct
	var searchResponse models.SearchResponse
	if err := json.Unmarshal(data, &searchResponse); err != nil {
		return nil, err
	}

	// Return the SearchResponse
	return &searchResponse, nil
}
