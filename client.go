package gospotify

import (
	"context"
	"net/http"
	"os"

	"github.com/alicse3/gospotify/apis"
	"github.com/alicse3/gospotify/consts"
	"github.com/alicse3/gospotify/models"
	"github.com/alicse3/gospotify/utils"
)

// Client represents the Spotify API client.
type Client struct {
	// Services to interact with Spotify api
	UserService      apis.UserService
	AlbumService     apis.AlbumService
	ArtistService    apis.ArtistService
	AudiobookService apis.AudiobookService
	CategoryService  apis.CategoryService
	ChapterService   apis.ChapterService
	EpisodeService   apis.EpisodeService
	GenreService     apis.GenreService
	MarketService    apis.MarketService
	PlayerService    apis.PlayerService
	PlaylistService  apis.PlaylistService
}

// GetCredentialsFromEnv reads the credentials(SPOTIFY_CLIENT_ID, SPOTIFY_CLIENT_SECRET, SPOTIFY_REDIRECT_URL) from environment variables and returns them.
// It throws an error if there are any.
func GetCredentialsFromEnv() (*utils.Credentials, error) {
	// Get SPOTIFY_CLIENT_ID value from env
	clientId := os.Getenv(consts.EnvClientId)
	if clientId == "" {
		return nil, &utils.Error{Type: utils.AppErrorType, AppError: &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgClientIdNotFound}}
	}

	// Get SPOTIFY_CLIENT_SECRET value from env
	clientSecret := os.Getenv(consts.EnvClientSecret)
	if clientSecret == "" {
		return nil, &utils.Error{Type: utils.AppErrorType, AppError: &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgClientSecretNotFound}}
	}

	// Get SPOTIFY_REDIRECT_URL value from env
	redirectUrl := os.Getenv(consts.EnvRedirectUrl)
	if redirectUrl == "" {
		return nil, &utils.Error{Type: utils.AppErrorType, AppError: &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgRedirectUrlNotFound}}
	}

	return &utils.Credentials{ClientId: clientId, ClientSecret: clientSecret, RedirectUrl: redirectUrl}, nil
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
func NewClient(credentials *utils.Credentials) (*Client, error) {
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
func NewClientWithCustomScopes(credentials *utils.Credentials, scopes []string) (*Client, error) {
	return NewClientWithDependencies(credentials, &utils.DefaultStateGenerator{}, &utils.DefaultHttpServer{}, utils.NewDefaultBrowserOpener(&utils.DefaultCommandExectutor{}), scopes)
}

// NewClientWithDependencies initializes and returns a new Spotify client.
func NewClientWithDependencies(
	credentials utils.CredentialsExchanger,
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
	return initClient(authToken, credentials.(*utils.Credentials)), nil
}

// initClient is a re-usable method to create a client with provided dependencies.
func initClient(authToken *models.AuthToken, credentials *utils.Credentials) *Client {
	// Create an HTTP httpClient with access token
	httpClient := utils.NewHttpClientWithToken(consts.BaseUrlApi, authToken, credentials)

	// Intialize services and return the Client instance
	return &Client{
		UserService:      apis.NewDefultUserService(httpClient),
		AlbumService:     apis.NewDefaultAlbumService(httpClient),
		ArtistService:    apis.NewDefaultArtistService(httpClient),
		AudiobookService: apis.NewDefaultAudiobookService(httpClient),
		CategoryService:  apis.NewDefaultCategoryService(httpClient),
		ChapterService:   apis.NewDefaultChapterService(httpClient),
		EpisodeService:   apis.NewDefaultEpisodeService(httpClient),
		GenreService:     apis.NewDefaultGenreService(httpClient),
		MarketService:    apis.NewDefaultMarketService(httpClient),
		PlayerService:    apis.NewDefaultPlayerService(httpClient),
		PlaylistService:  apis.NewDefaultPlaylistService(httpClient),
	}
}

// NewClientWithToken initializes and returns a new Spotify client with the provided token.
// This is useful when you have a valid token and want to create a client with that token.
// For example, you can use this method when you want to set the permanent token.
// It doesn't support the token refresh functionality. Error will be thrown when the access token is expired.
func NewClientWithToken(token string) (*Client, error) {
	return initClient(&models.AuthToken{AccessToken: token}, nil), nil
}
