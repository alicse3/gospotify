package gospotify

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/alicse3/gospotify/apis"
	"github.com/alicse3/gospotify/consts"
	"github.com/alicse3/gospotify/models"
	"github.com/alicse3/gospotify/utils"
)

// Client represents the Spotify API client.
type Client struct {
	HttpClient  *utils.HttpClient // Client for making HTTP requests
	AuthToken   *models.AuthToken // Auth token for authenticating requests
	credentials *Credentials

	// Services to interact with Spotify api
	userService      apis.UserService
	albumService     apis.AlbumService
	artistService    apis.ArtistService
	audiobookService apis.AudiobookService
	categoryService  apis.CategoryService
	chapterService   apis.ChapterService
	episodeService   apis.EpisodeService
	genreService     apis.GenreService
	marketService    apis.MarketService
	playerService    apis.PlayerService
}

// GetCredentialsFromEnv reads the credentials(SPOTIFY_CLIENT_ID, SPOTIFY_CLIENT_SECRET, SPOTIFY_REDIRECT_URL) from environment variables and returns them.
// It throws an error if there are any.
func GetCredentialsFromEnv() (*Credentials, error) {
	// Get SPOTIFY_CLIENT_ID value from env
	cliendId := os.Getenv(consts.EnvClientId)
	if cliendId == "" {
		return nil, &utils.Error{Type: utils.AppErrorType, AppError: &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgClientIdNotFound}}
	}

	// Get SPOTIFY_CLIENT_SECRET value from env
	cliendSecret := os.Getenv(consts.EnvClientSecret)
	if cliendSecret == "" {
		return nil, &utils.Error{Type: utils.AppErrorType, AppError: &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgClientSecretNotFound}}
	}

	// Get SPOTIFY_REDIRECT_URL value from env
	redirectUrl := os.Getenv(consts.EnvRedirectUrl)
	if redirectUrl == "" {
		return nil, &utils.Error{Type: utils.AppErrorType, AppError: &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgRedirectUrlNotFound}}
	}

	return &Credentials{ClientId: cliendId, ClientSecret: cliendSecret, RedirectUrl: redirectUrl}, nil
}

// DefaultClient initializes and returns a new Spotify client.
func DefaultClient() (*Client, error) {
	credentials, err := GetCredentialsFromEnv()
	if err != nil {
		return nil, err
	}

	return NewClient(credentials)
}

// DefaultClientWithCustomScopes initializes the client with given custom scopes and returns a new Spotify client.
func DefaultClientWithCustomScopes(scopes []string) (*Client, error) {
	credentials, err := GetCredentialsFromEnv()
	if err != nil {
		return nil, err
	}

	return NewClientWithCustomScopes(credentials, scopes)
}

// NewClient initializes and returns a new Spotify client.
func NewClient(credentials *Credentials) (*Client, error) {
	return NewClientWithDependencies(credentials, &utils.DefaultStateGenerator{}, &utils.DefaultHttpServer{}, utils.NewDefaultBrowserOpener(&utils.DefaultCommandExectutor{}), []string{})
}

// NewClientWithCustomScopes initializes the client with given custom scopes and returns a new Spotify client.
// For example, to initialize client with all the default scopes:
//
//		 // Initialize the client with custom scopes
//		client, err := gospotify.NewClientWithCustomScopes(
//			&gospotify.Credentials{
//				ClientId:     "your_client_id",
//				ClientSecret: "your_client_secret",
//				RedirectUrl:  "your_redirect_uri",
//			},
//	       gospotify.AllScopes, // Passing all scopes
//		)
func NewClientWithCustomScopes(credentials *Credentials, scopes []string) (*Client, error) {
	return NewClientWithDependencies(credentials, &utils.DefaultStateGenerator{}, &utils.DefaultHttpServer{}, utils.NewDefaultBrowserOpener(&utils.DefaultCommandExectutor{}), scopes)
}

// NewClientWithDependencies initializes and returns a new Spotify client.
func NewClientWithDependencies(
	credentials CredentialsExchanger,
	stateGenerator utils.StateGenerator,
	httpServer utils.HttpServer,
	browserOpener utils.BrowserOpener,
	scopes []string,
) (*Client, error) {
	// Generate a random state string for security
	state, err := stateGenerator.GetRandomState(16)
	if err != nil {
		return nil, &utils.Error{Type: utils.AppErrorType, AppError: &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgStateGenerationFailure, Err: err}}
	}

	// Use channel for obtaining the authorization code
	ch := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start an HTTP server to listen for the authentication callback
	go httpServer.StartServer(ctx, ch)

	// TODO: Handle http server errors through channel?

	// Generate authorization url with provided state and scopes
	authUrl, err := credentials.GetAuthorizationUrl(scopes, state)
	if err != nil {
		return nil, &utils.Error{Type: utils.AppErrorType, AppError: &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgGettingAuthUrlFailure, Err: err}}
	}

	// Open the authUrl in default browser
	if err := browserOpener.Open(authUrl); err != nil {
		return nil, &utils.Error{Type: utils.AppErrorType, AppError: &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgOpeningBrowserFailure, Err: err}}
	}

	// Wait for the authentication callback and get the code
	code := <-ch

	// Create an HTTP client and get an access token
	httpClient := utils.NewHttpClient(consts.BaseUrlAccounts)
	authToken, err := credentials.ExchangeCodeForTokens(httpClient, code)
	if err != nil {
		return nil, err
	}

	// Init and return the Client instance
	return initClient(httpClient, authToken, credentials.(*Credentials), authToken.AccessToken), nil
}

// initClient is a re-usable method to create a client with provided dependencies.
func initClient(httpClient *utils.HttpClient, authToken *models.AuthToken, credentials *Credentials, token string) *Client {
	// Create an HTTP client with access token
	httpClientWithToken := utils.NewHttpClientWithToken(consts.BaseUrlApi, token)

	// Create and return the Client instance
	return &Client{
		HttpClient:  httpClient,
		AuthToken:   authToken,
		credentials: credentials,

		// Intialize all services with dependencies
		userService:      apis.NewDefultUserService(httpClientWithToken),
		albumService:     apis.NewDefaultAlbumService(httpClientWithToken),
		artistService:    apis.NewDefaultArtistService(httpClientWithToken),
		audiobookService: apis.NewDefaultAudiobookService(httpClientWithToken),
		categoryService:  apis.NewDefaultCategoryService(httpClientWithToken),
		chapterService:   apis.NewDefaultChapterService(httpClientWithToken),
		episodeService:   apis.NewDefaultEpisodeService(httpClientWithToken),
		genreService:     apis.NewDefaultGenreService(httpClientWithToken),
		marketService:    apis.NewDefaultMarketService(httpClientWithToken),
		playerService:    apis.NewDefaultPlayerService(httpClientWithToken),
	}
}

// NewClientWithToken initializes and returns a new Spotify client with the provided token.
// This is useful when you have a valid token and want to create a client with that token.
// For example, you can use this method when you want to set the permanent token. It doesn't support token refresh functionality.
func NewClientWithToken(token string) (*Client, error) {
	return initClient(nil, nil, nil, token), nil
}

// RefreshTokens refreshes the tokens.
func (c *Client) RefreshTokens() error {
	// To make sure the dependencies have initialized before refreshing the tokens
	if c.credentials == nil {
		return &utils.Error{Type: utils.AppErrorType, AppError: &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgCredentialsNotInitialized}}
	}
	if c.AuthToken == nil {
		return &utils.Error{Type: utils.AppErrorType, AppError: &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgAuthTokenNotInitialized}}
	}
	if c.HttpClient == nil {
		return &utils.Error{Type: utils.AppErrorType, AppError: &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgHttpClientNotInitialized}}
	}

	// Set the required headers
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	// Set the form values for the token refresh request
	formValues := map[string]string{
		"grant_type":    "refresh_token",
		"client_id":     c.credentials.ClientId,
		"client_secret": c.credentials.ClientSecret,
		"refresh_token": c.AuthToken.RefreshToken,
	}

	// Make a POST request to the token endpoint
	res, err := c.HttpClient.Post(context.Background(), consts.EndpointRefresh, headers, nil, formValues, nil)
	if err != nil {
		return err
	}

	// Read the response body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// Unmarshal the response data into an AuthToken struct
	var authToken models.AuthToken
	if err := json.Unmarshal(data, &authToken); err != nil {
		return err
	}

	// Update the client's authToken
	c.AuthToken = &authToken

	return nil
}

// CheckAndRefreshTokens checks for the AuthToken expiry and then triggers refresh tokens call if needed.
func (c *Client) CheckAndRefreshTokens() error {
	// To make sure the dependencies have initialized before checking the expiry
	if c.AuthToken == nil {
		return &utils.Error{Type: utils.AppErrorType, AppError: &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgAuthTokenNotInitialized}}
	}

	// Check if the token has expired
	if time.Now().After(c.AuthToken.ExpiryTime) {
		if err := c.RefreshTokens(); err != nil {
			return &utils.Error{Type: utils.AppErrorType, AppError: &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgRefreshTokensFailure, Err: err}}
		}
	}

	return nil
}

// GetCurrentUserProfile calls user service to get current user profile.
func (c *Client) GetCurrentUserProfile() (*models.User, error) {
	return c.userService.GetCurrentUserProfile()
}

// GetAlbum gets the album information.
func (c *Client) GetAlbum(input models.GetAlbumRequest) (*models.Album, error) {
	return c.albumService.GetAlbum(input)
}

// GetAlbums gets the albums information.
func (c *Client) GetAlbums(input models.GetAlbumsRequest) (*models.Albums, error) {
	return c.albumService.GetAlbums(input)
}

// GetAlbumTracks gets the album tracks information.
func (c *Client) GetAlbumTracks(input models.GetAlbumTracksRequest) (*models.Tracks, error) {
	return c.albumService.GetAlbumTracks(input)
}

// GetSavedAlbums gets the saved albums information.
// Required authorization scopes: user-library-read
func (c *Client) GetSavedAlbums(input models.GetSavedAlbumsRequest) (*models.SavedAlbums, error) {
	return c.albumService.GetSavedAlbums(input)
}

// SaveAlbums saves the albums information.
// Required authorization scopes: user-library-modify
func (c *Client) SaveAlbums(input models.SaveAlbumsRequest) error {
	return c.albumService.SaveAlbums(input)
}

// CheckSavedAlbums checks and returns the saved albums information.
// Required authorization scopes: user-library-read
func (c *Client) CheckSavedAlbums(input models.CheckSavedAlbumsRequest) (*models.CheckSavedAlbums, error) {
	return c.albumService.CheckSavedAlbums(input)
}

// GetNewReleases returns the new releases information.
func (c *Client) GetNewReleases(input models.GetNewReleasesRequest) (*models.NewlyReleasedAlbums, error) {
	return c.albumService.GetNewReleases(input)
}

// GetArtist returns the artist information.
func (c *Client) GetArtist(input models.GetArtistRequest) (*models.Artist, error) {
	return c.artistService.GetArtist(input)
}

// GetArtists returns the artists information.
func (c *Client) GetArtists(input models.GetArtistsRequest) (*models.Artists, error) {
	return c.artistService.GetArtists(input)
}

// GetArtistAlbums returns the artist albums information.
func (c *Client) GetArtistAlbums(input models.GetArtistAlbumsRequest) (*models.ArtistAlbums, error) {
	return c.artistService.GetArtistAlbums(input)
}

// GetArtistTopTracks returns the artists top tracks information.
func (c *Client) GetArtistTopTracks(input models.GetArtistTopTracksRequest) (*models.ArtistTopTracks, error) {
	return c.artistService.GetArtistTopTracks(input)
}

// GetRelatedArtists returns the related artists information.
func (c *Client) GetRelatedArtists(input models.GetRelatedArtistsRequest) (*models.Artists, error) {
	return c.artistService.GetRelatedArtists(input)
}

// GetAudiobook returns the audiobook information.
func (c *Client) GetAudiobook(input models.GetAudiobookRequest) (*models.Audiobook, error) {
	return c.audiobookService.GetAudiobook(input)
}

// GetAudiobooks returns the audiobooks information.
func (c *Client) GetAudiobooks(input models.GetAudiobooksRequest) (*models.Audiobooks, error) {
	return c.audiobookService.GetAudiobooks(input)
}

// GetAudiobookChapters returns the audiobooks chapters information.
func (c *Client) GetAudiobookChapters(input models.GetAudiobookChaptersRequest) (*models.AudiobookChapters, error) {
	return c.audiobookService.GetAudiobookChapters(input)
}

// GetSavedAudiobooks returns the saved audiobooks information.
// Required authorization scopes: user-library-read
func (c *Client) GetSavedAudiobooks(input models.GetSavedAudiobooksRequest) (*models.SavedAudiobooks, error) {
	return c.audiobookService.GetSavedAudiobooks(input)
}

// SaveAudiobooks saves the audiobooks information.
// Required authorization scopes: user-library-modify
func (c *Client) SaveAudiobooks(input models.SaveAudiobooksRequest) error {
	return c.audiobookService.SaveAudiobooks(input)
}

// DeleteAudiobooks deletes the audiobooks information.
// Required authorization scopes: user-library-modify
func (c *Client) DeleteAudiobooks(input models.RemoveAudiobooksRequest) error {
	return c.audiobookService.DeleteAudiobooks(input)
}

// CheckSavedAudiobooks returns the check saved audiobooks information.
// Required authorization scopes: user-library-read
func (c *Client) CheckSavedAudiobooks(input models.CheckSavedAudiobooksRequest) (*models.CheckSavedAudiobooks, error) {
	return c.audiobookService.CheckSavedAudiobooks(input)
}

// GetBrowseCategories returns the get browse categories information.
func (c *Client) GetBrowseCategories(input models.GetBrowseCategoriesRequest) (*models.Categories, error) {
	return c.categoryService.GetBrowseCategories(input)
}

// GetBrowseCategory returns the get browse category information.
func (c *Client) GetBrowseCategory(input models.GetBrowseCategoryRequest) (*models.Category, error) {
	return c.categoryService.GetBrowseCategory(input)
}

// GetChapter returns the chapter information.
func (c *Client) GetChapter(input models.GetChapterRequest) (*models.Chapter, error) {
	return c.chapterService.GetChapter(input)
}

// GetChapters returns the chapters information.
func (c *Client) GetChapters(input models.GetChaptersRequest) (*models.Chapters, error) {
	return c.chapterService.GetChapters(input)
}

// GetEpisode returns the episode information.
// Required authorization scopes: user-read-playback-position
func (c *Client) GetEpisode(input models.GetEpisodeRequest) (*models.Episode, error) {
	return c.episodeService.GetEpisode(input)
}

// GetEpisodes returns the episodes information.
// Required authorization scopes: user-read-playback-position
func (c *Client) GetEpisodes(input models.GetEpisodesRequest) (*models.Episodes, error) {
	return c.episodeService.GetEpisodes(input)
}

// GetSavedEpisodes returns the saved episodes information.
// Required authorization scopes: user-library-read, user-read-playback-position
func (c *Client) GetSavedEpisodes(input models.GetSavedEpisodesRequest) (*models.SavedEpisodes, error) {
	return c.episodeService.GetSavedEpisodes(input)
}

// SaveEpisodes saves the episodes information.
// Required authorization scopes: user-library-modify
func (c *Client) SaveEpisodes(input models.SaveEpisodesRequest) error {
	return c.episodeService.SaveEpisodes(input)
}

// RemoveEpisodes removes the episodes information.
// Required authorization scopes: user-library-modify
func (c *Client) RemoveEpisodes(input models.RemoveEpisodesRequest) error {
	return c.episodeService.RemoveEpisodes(input)
}

// CheckSavedEpisodes checks the saved episodes information.
// Required authorization scopes: user-library-read
func (c *Client) CheckSavedEpisodes(input models.CheckSavedEpisodesRequest) (*models.CheckSavedEpisodes, error) {
	return c.episodeService.CheckSavedEpisodes(input)
}

// GetAvailableGenresSeeds returns the available genres seeds information.
func (c *Client) GetAvailableGenresSeeds() (*models.Genres, error) {
	return c.genreService.GetAvailableGenresSeeds()
}

// GetAvailableMarkets returns the available markets information.
func (c *Client) GetAvailableMarkets() (*models.Markets, error) {
	return c.marketService.GetAvailableMarkets()
}

// GetPlaybackState returns the playback state information.
// Required authorization scopes: user-read-playback-state
func (c *Client) GetPlaybackState(input models.GetPlaybackStateRequest) (*models.PlaybackState, error) {
	return c.playerService.GetPlaybackState(input)
}

// TransferPlayback transfers the playback.
// Required authorization scopes: user-modify-playback-state
func (c *Client) TransferPlayback(input models.TransferPlaybackRequest) error {
	return c.playerService.TransferPlayback(input)
}

// GetAvailableDevices returns the available devices information.
// Required authorization scopes: user-read-playback-state
func (c *Client) GetAvailableDevices() (*models.AvailableDevices, error) {
	return c.playerService.GetAvailableDevices()
}

// GetCurrentlyPlayingTrack returns the currently playing track information.
// Required authorization scopes: user-read-currently-playing
func (c *Client) GetCurrentlyPlayingTrack(input models.GetCurrentlyPlayingTrackRequest) (*models.PlaybackState, error) {
	return c.playerService.GetCurrentlyPlayingTrack(input)
}

// StartOrResumePlayback starts or resumes the playback.
// Required authorization scopes: user-modify-playback-state
func (c *Client) StartOrResumePlayback(input models.StartOrResumePlaybackRequest) error {
	return c.playerService.StartOrResumePlayback(input)
}

// PausePlayback pauses the playback.
// Required authorization scopes: user-modify-playback-state
func (c *Client) PausePlayback(input models.PausePlaybackRequest) error {
	return c.playerService.PausePlayback(input)
}

// SkipToNext skips to the next.
// Required authorization scopes: user-modify-playback-state
func (c *Client) SkipToNext(input models.SkipToNextRequest) error {
	return c.playerService.SkipToNext(input)
}

// SkipToPrevious skips to the previous.
// Required authorization scopes: user-modify-playback-state
func (c *Client) SkipToPrevious(input models.SkipToPreviousRequest) error {
	return c.playerService.SkipToPrevious(input)
}

// SeekToPosition seeks to the position.
// Required authorization scopes: user-modify-playback-state
func (c *Client) SeekToPosition(input models.SeekToPositionRequest) error {
	return c.playerService.SeekToPosition(input)
}

// SetRepeatMode sets the repeat mode.
// Required authorization scopes: user-modify-playback-state
func (c *Client) SetRepeatMode(input models.SetRepeatModeRequest) error {
	return c.playerService.SetRepeatMode(input)
}

// SetPlaybackVolume sets the playback volume.
// Required authorization scopes: user-modify-playback-state
func (c *Client) SetPlaybackVolume(input models.SetPlaybackVolumeRequest) error {
	return c.playerService.SetPlaybackVolume(input)
}

// TogglePlaybackShuffle toggles the playback shuffle.
// Required authorization scopes: user-modify-playback-state
func (c *Client) TogglePlaybackShuffle(input models.TogglePlaybackShuffleRequest) error {
	return c.playerService.TogglePlaybackShuffle(input)
}

// GetRecentlyPlayedTracks returns the recently played tracks information.
// Required authorization scopes: user-read-recently-played
func (c *Client) GetRecentlyPlayedTracks(input models.GetRecentlyPlayedTracksRequest) (*models.RecentlyPlayedTracks, error) {
	return c.playerService.GetRecentlyPlayedTracks(input)
}

// GetUsersQueue returns the users queue information.
// Required authorization scopes: user-read-currently-playing, user-read-playback-state
func (c *Client) GetUsersQueue() (*models.UsersQueue, error) {
	return c.playerService.GetUsersQueue()
}

// AddItemToPlaybackQueue adds item to the playback queue.
// Required authorization scopes: user-modify-playback-state
func (c *Client) AddItemToPlaybackQueue(input models.AddItemToPlaybackQueueRequest) error {
	return c.playerService.AddItemToPlaybackQueue(input)
}
