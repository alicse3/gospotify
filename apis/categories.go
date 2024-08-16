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

// CategoryService interface defines the methods for interacting with the Spotify Category's API.
type CategoryService interface {
	// Get a list of categories used to tag items in Spotify (on, for example, the Spotify player’s “Browse” tab).
	GetBrowseCategories(input models.GetBrowseCategoriesRequest) (*models.Categories, error)

	// Get a single category used to tag items in Spotify (on, for example, the Spotify player’s “Browse” tab).
	GetBrowseCategory(input models.GetBrowseCategoryRequest) (*models.Category, error)
}

// DefaultCategoryService is a struct that implements CategoryService interface.
type DefaultCategoryService struct {
	client *utils.HttpClient
}

// NewDefaultCategoryService initializes the DefaultCategoryService with given dependencies.
func NewDefaultCategoryService(client *utils.HttpClient) *DefaultCategoryService {
	return &DefaultCategoryService{client}
}

// GetBrowseCategories implements the CategoryService's interface GetBrowseCategories method.
func (service *DefaultCategoryService) GetBrowseCategories(input models.GetBrowseCategoriesRequest) (*models.Categories, error) {
	// Add inputs to the query parameters
	params := map[string]string{"locale": input.Locale, "limit": strconv.Itoa(input.Limit), "offset": strconv.Itoa(input.Offset)}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointBrowseCategories, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetBrowseCategories, Err: err}
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

	// Unmarshal the response data into Categories struct
	var categories models.Categories
	if err := json.Unmarshal(data, &categories); err != nil {
		return nil, err
	}

	// Return the Categories
	return &categories, nil
}

// GetBrowseCategory implements the CategoryService's interface GetBrowseCategory method.
func (service *DefaultCategoryService) GetBrowseCategory(input models.GetBrowseCategoryRequest) (*models.Category, error) {
	// Validate the input
	if input.CategoryId == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgCategoryIdRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"category_id": input.CategoryId, "locale": input.Locale}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointBrowseCategory, input.CategoryId)

	// Make an API call
	res, err := service.client.Get(context.Background(), endpoint, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetBrowseCategory, Err: err}
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

	// Unmarshal the response data into Category struct
	var category models.Category
	if err := json.Unmarshal(data, &category); err != nil {
		return nil, err
	}

	// Return the Category
	return &category, nil
}
