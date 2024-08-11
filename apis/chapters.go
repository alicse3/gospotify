package apis

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/alicse3/gospotify/consts"
	"github.com/alicse3/gospotify/models"
	"github.com/alicse3/gospotify/utils"
)

// ChapterService interface defines the methods for interacting with the Spotify Chapter's API.
type ChapterService interface {
	GetChapter(input models.GetChapterRequest) (*models.Chapter, error)
	GetChapters(input models.GetChaptersRequest) (*models.Chapters, error)
}

// DefaultChapterService is a struct that implements ChapterService interface.
type DefaultChapterService struct {
	client *utils.HttpClient
}

// NewDefaultChapterService initializes the DefaultChapterService with given dependencies.
func NewDefaultChapterService(client *utils.HttpClient) *DefaultChapterService {
	return &DefaultChapterService{client}
}

// GetChapter implements the ChapterService's interface GetChapter method.
func (service *DefaultChapterService) GetChapter(input models.GetChapterRequest) (*models.Chapter, error) {
	// Validate the input
	if input.Id == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointGetChapter, input.Id)

	// Make an API call
	res, err := service.client.Get(context.Background(), endpoint, nil)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetChapter, Err: err}
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

	// Unmarshal the response data into an Chapter struct
	var chapter models.Chapter
	if err := json.Unmarshal(data, &chapter); err != nil {
		return nil, err
	}

	// Return the Chapter
	return &chapter, nil
}

// GetChapters implements the ChapterService's interface GetChapters method.
func (service *DefaultChapterService) GetChapters(input models.GetChaptersRequest) (*models.Chapters, error) {
	// Validate the input
	if input.Ids == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdsRequired}
	}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointGetChapters, nil)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetChapters, Err: err}
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

	// Unmarshal the response data into an Chapters struct
	var chapters models.Chapters
	if err := json.Unmarshal(data, &chapters); err != nil {
		return nil, err
	}

	// Return the Chapters
	return &chapters, nil
}
