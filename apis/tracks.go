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

// TrackService interface defines the methods for interacting with the Spotify Track's API.
type TrackService interface {
	// Get Spotify catalog information for a single track identified by its unique Spotify ID.
	GetTrack(input models.GetTrackRequest) (*models.Track, error)

	// Get Spotify catalog information for multiple tracks based on their Spotify IDs.
	GetTracks(input models.GetTracksRequest) (*models.Tracks, error)

	// Get a list of the songs saved in the current Spotify user's 'Your Music' library.
	// Authorization scopes: user-library-read
	GetSavedTracks(input models.GetSavedTracksRequest) (*models.SavedTracks, error)

	// Save one or more tracks to the current user's 'Your Music' library.
	// Authorization scopes: user-library-modify
	SaveTracks(input models.SaveTracksRequest) error

	// Remove one or more tracks from the current user's 'Your Music' library.
	// Authorization scopes: user-library-modify
	RemoveSavedTracks(input models.RemoveTracksRequest) error

	// Check if one or more tracks is already saved in the current Spotify user's 'Your Music' library.
	// Authorization scopes: user-library-read
	CheckSavedTracks(input models.CheckSavedTracksRequest) (*models.CheckSavedTracks, error)

	// Get audio features for multiple tracks based on their Spotify IDs.
	CheckSeveralTracksAudioFeatures(input models.GetSeveralTracksAudioFeaturesRequest) (*models.SeveralTracksAudioFeatures, error)

	// Get audio feature information for a single track identified by its unique Spotify ID.
	CheckTracksAudioFeatures(input models.GetTracksAudioFeaturesRequest) (*models.TracksAudioFeatures, error)

	// Get a low-level audio analysis for a track in the Spotify catalog.
	// The audio analysis describes the track’s structure and musical content, including rhythm, pitch, and timbre.
	CheckTracksAudioAnalysis(input models.GetTracksAudioAnalysisRequest) (*models.TracksAudioAnalysis, error)

	// Recommendations are generated based on the available information for a given seed entity and matched against similar artists and tracks.
	// If there is sufficient information about the provided seeds, a list of tracks will be returned together with pool size details.
	// For artists and tracks that are very new or obscure there might not be enough data to generate a list of tracks.
	GetRecommendations(input models.GetRecommendationsRequest) (*models.GetRecommendations, error)
}

// DefaultTrackService is a struct that implements TrackService interface.
type DefaultTrackService struct {
	client *utils.HttpClient
}

// NewDefaultTrackService initializes the DefaultTrackService with given dependencies.
func NewDefaultTrackService(client *utils.HttpClient) *DefaultTrackService {
	return &DefaultTrackService{client}
}

// GetTrack implements the TrackService's interface GetTrack method.
func (service *DefaultTrackService) GetTrack(input models.GetTrackRequest) (*models.Track, error) {
	// Validate the input
	if input.Id == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointTrack, input.Id)

	// Add inputs to the query parameters
	params := map[string]string{"id": input.Id, "market": input.Market}

	// Make an API call
	res, err := service.client.Get(context.Background(), endpoint, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetTrack, Err: err}
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

	// Unmarshal the response data into Track struct
	var track models.Track
	if err := json.Unmarshal(data, &track); err != nil {
		return nil, err
	}

	// Return the Track
	return &track, nil
}

// GetTracks implements the TrackService's interface GetTracks method.
func (service *DefaultTrackService) GetTracks(input models.GetTracksRequest) (*models.Tracks, error) {
	// Validate the input
	if input.Ids == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdsRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"ids": input.Ids, "market": input.Market}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointTracks, params)
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
	var tracks models.Tracks
	if err := json.Unmarshal(data, &tracks); err != nil {
		return nil, err
	}

	// Return the Tracks
	return &tracks, nil
}

// GetSavedTracks implements the TrackService's interface GetSavedTracks method.
func (service *DefaultTrackService) GetSavedTracks(input models.GetSavedTracksRequest) (*models.SavedTracks, error) {
	// Add inputs to the query parameters
	params := map[string]string{"market": input.Market, "limit": strconv.Itoa(input.Limit), "offset": strconv.Itoa(input.Offset)}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointSaveTracks, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetSavedTracks, Err: err}
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

	// Unmarshal the response data into SavedTracks struct
	var savedTracks models.SavedTracks
	if err := json.Unmarshal(data, &savedTracks); err != nil {
		return nil, err
	}

	// Return the SavedTracks
	return &savedTracks, nil
}

// SaveTracks implements the TrackService's interface SaveTracks method.
func (service *DefaultTrackService) SaveTracks(input models.SaveTracksRequest) error {
	// Validate the input
	if input.Ids == "" {
		return &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdsRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"ids": input.Ids}

	// Make an API call
	res, err := service.client.Put(context.Background(), consts.EndpointSaveTracks, nil, params, input.Body)
	if err != nil {
		return &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToSaveTracks, Err: err}
	}

	// Handle Spotify API error
	if res.StatusCode != http.StatusOK {
		return utils.ParseSpotifyError(res, utils.RegErrorType)
	}

	// Return the empty response
	return nil
}

// RemoveSavedTracks implements the TrackService's interface RemoveSavedTracks method.
func (service *DefaultTrackService) RemoveSavedTracks(input models.RemoveTracksRequest) error {
	// Validate the input
	if input.Ids == "" {
		return &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdsRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"ids": input.Ids}

	// Make an API call
	res, err := service.client.Delete(context.Background(), consts.EndpointSaveTracks, nil, params, input.Body)
	if err != nil {
		return &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToRemoveSavedTracks, Err: err}
	}

	// Handle Spotify API error
	if res.StatusCode != http.StatusOK {
		return utils.ParseSpotifyError(res, utils.RegErrorType)
	}

	// Return the empty response
	return nil
}

// CheckSavedTracks implements the TrackService's interface CheckSavedTracks method.
func (service *DefaultTrackService) CheckSavedTracks(input models.CheckSavedTracksRequest) (*models.CheckSavedTracks, error) {
	// Validate the input
	if input.Ids == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdsRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"ids": input.Ids}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointCheckSavedTracks, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToCheckSavedTracks, Err: err}
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

	// Unmarshal the response data into CheckSavedTracks struct
	var checkSavedTracks models.CheckSavedTracks
	if err := json.Unmarshal(data, &checkSavedTracks); err != nil {
		return nil, err
	}

	// Return the CheckSavedTracks
	return &checkSavedTracks, nil
}

// CheckSeveralTracksAudioFeatures implements the TrackService's interface CheckSeveralTracksAudioFeatures method.
func (service *DefaultTrackService) CheckSeveralTracksAudioFeatures(input models.GetSeveralTracksAudioFeaturesRequest) (*models.SeveralTracksAudioFeatures, error) {
	// Validate the input
	if input.Ids == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdsRequired}
	}

	// Add inputs to the query parameters
	params := map[string]string{"ids": input.Ids}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointSeveralTracksAudioFeatures, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetSeveralTracksAudioFeatures, Err: err}
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

	// Unmarshal the response data into SeveralTracksAudioFeatures struct
	var severalTracksAudioFeatures models.SeveralTracksAudioFeatures
	if err := json.Unmarshal(data, &severalTracksAudioFeatures); err != nil {
		return nil, err
	}

	// Return the SeveralTracksAudioFeatures
	return &severalTracksAudioFeatures, nil
}

// CheckTracksAudioFeatures implements the TrackService's interface CheckTracksAudioFeatures method.
func (service *DefaultTrackService) CheckTracksAudioFeatures(input models.GetTracksAudioFeaturesRequest) (*models.TracksAudioFeatures, error) {
	// Validate the input
	if input.Id == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointTracksTracksAudioFeatures, input.Id)

	// Add inputs to the query parameters
	params := map[string]string{"id": input.Id}

	// Make an API call
	res, err := service.client.Get(context.Background(), endpoint, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetTracksAudioFeatures, Err: err}
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

	// Unmarshal the response data into TracksAudioFeatures struct
	var tracksAudioFeatures models.TracksAudioFeatures
	if err := json.Unmarshal(data, &tracksAudioFeatures); err != nil {
		return nil, err
	}

	// Return the TracksAudioFeatures
	return &tracksAudioFeatures, nil
}

// CheckTracksAudioAnalysis implements the TrackService's interface CheckTracksAudioAnalysis method.
func (service *DefaultTrackService) CheckTracksAudioAnalysis(input models.GetTracksAudioAnalysisRequest) (*models.TracksAudioAnalysis, error) {
	// Validate the input
	if input.Id == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgIdRequired}
	}

	// Substitute id in the endpoint
	endpoint := fmt.Sprintf(consts.EndpointTracksAudioAnalysis, input.Id)

	// Add inputs to the query parameters
	params := map[string]string{"id": input.Id}

	// Make an API call
	res, err := service.client.Get(context.Background(), endpoint, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetTracksAudioAnalysis, Err: err}
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

	// Unmarshal the response data into TracksAudioAnalysis struct
	var tracksAudioAnalysis models.TracksAudioAnalysis
	if err := json.Unmarshal(data, &tracksAudioAnalysis); err != nil {
		return nil, err
	}

	// Return the TracksAudioAnalysis
	return &tracksAudioAnalysis, nil
}

// GetRecommendations implements the TrackService's interface GetRecommendations method.
func (service *DefaultTrackService) GetRecommendations(input models.GetRecommendationsRequest) (*models.GetRecommendations, error) {
	// Validate the input
	if input.SeedArtists == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgSeedArtistsRequired}
	}
	if input.SeedGenres == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgSeedGenresRequired}
	}
	if input.SeedTracks == "" {
		return nil, &utils.AppError{Status: http.StatusBadRequest, Message: consts.MsgSeedTracksRequired}
	}

	// Add inputs to the query parameters
	const format = 'f'
	const precision = 1
	const bitSize = 64
	params := map[string]string{
		"limit":  strconv.Itoa(input.Limit),
		"market": input.Market,

		"seed_artists": input.SeedArtists,
		"seed_genres":  input.SeedGenres,
		"seed_tracks":  input.SeedTracks,

		"min_acousticness":    strconv.FormatFloat(input.MinAcousticness, format, precision, bitSize),
		"max_acousticness":    strconv.FormatFloat(input.MaxAcousticness, format, precision, bitSize),
		"target_acousticness": strconv.FormatFloat(input.TargetAcousticness, format, precision, bitSize),

		"min_danceability":    strconv.FormatFloat(input.MinDanceability, format, precision, bitSize),
		"max_danceability":    strconv.FormatFloat(input.MaxDanceability, format, precision, bitSize),
		"target_danceability": strconv.FormatFloat(input.TargetDanceability, format, precision, bitSize),

		"min_duration_ms":    strconv.Itoa(input.MinDurationMs),
		"max_duration_ms":    strconv.Itoa(input.MaxDurationMs),
		"target_duration_ms": strconv.Itoa(input.TargetDurationMs),

		"min_energy":    strconv.FormatFloat(input.MinEnergy, format, precision, bitSize),
		"max_energy":    strconv.FormatFloat(input.MaxEnergy, format, precision, bitSize),
		"target_energy": strconv.FormatFloat(input.TargetEnergy, format, precision, bitSize),

		"min_instrumentalness":    strconv.FormatFloat(input.MinInstrumentalness, format, precision, bitSize),
		"max_instrumentalness":    strconv.FormatFloat(input.MaxInstrumentalness, format, precision, bitSize),
		"target_instrumentalness": strconv.FormatFloat(input.TargetInstrumentalness, format, precision, bitSize),

		"min_key":    strconv.Itoa(input.MinKey),
		"max_key":    strconv.Itoa(input.MaxKey),
		"target_key": strconv.Itoa(input.TargetKey),

		"min_liveness":    strconv.FormatFloat(input.MinLiveness, format, precision, bitSize),
		"max_liveness":    strconv.FormatFloat(input.MaxLiveness, format, precision, bitSize),
		"target_liveness": strconv.FormatFloat(input.TargetLiveness, format, precision, bitSize),

		"min_loudness":    strconv.FormatFloat(input.MinLoudness, format, precision, bitSize),
		"max_loudness":    strconv.FormatFloat(input.MaxLoudness, format, precision, bitSize),
		"target_loudness": strconv.FormatFloat(input.TargetLoudness, format, precision, bitSize),

		"min_mode":    strconv.Itoa(input.MinMode),
		"max_mode":    strconv.Itoa(input.MaxMode),
		"target_mode": strconv.Itoa(input.TargetMode),

		"min_popularity":    strconv.Itoa(input.MinPopularity),
		"max_popularity":    strconv.Itoa(input.MaxPopularity),
		"target_popularity": strconv.Itoa(input.TargetPopularity),

		"min_speechiness":    strconv.FormatFloat(input.MinSpeechiness, format, precision, bitSize),
		"max_speechiness":    strconv.FormatFloat(input.MaxSpeechiness, format, precision, bitSize),
		"target_speechiness": strconv.FormatFloat(input.TargetSpeechiness, format, precision, bitSize),

		"min_tempo":    strconv.FormatFloat(input.MinTempo, format, precision, bitSize),
		"max_tempo":    strconv.FormatFloat(input.MaxTempo, format, precision, bitSize),
		"target_tempo": strconv.FormatFloat(input.TargetTempo, format, precision, bitSize),

		"min_time_signature":    strconv.Itoa(input.MinTimeSignature),
		"max_time_signature":    strconv.Itoa(input.MaxTimeSignature),
		"target_time_signature": strconv.Itoa(input.TargetTimeSignature),

		"min_valence":    strconv.FormatFloat(input.MinValence, format, precision, bitSize),
		"max_valence":    strconv.FormatFloat(input.MaxValence, format, precision, bitSize),
		"target_valence": strconv.FormatFloat(input.TargetValence, format, precision, bitSize),
	}

	// Make an API call
	res, err := service.client.Get(context.Background(), consts.EndpointRecommendations, params)
	if err != nil {
		return nil, &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgFailedToGetRecommendations, Err: err}
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

	// Unmarshal the response data into GetRecommendations struct
	var getRecommendations models.GetRecommendations
	if err := json.Unmarshal(data, &getRecommendations); err != nil {
		return nil, err
	}

	// Return the GetRecommendations
	return &getRecommendations, nil
}
