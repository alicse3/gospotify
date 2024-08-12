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
	httpClient  *utils.HttpClient // Client for making HTTP requests
	authToken   *models.AuthToken // Auth token for authenticating requests
	credentials *Credentials      // For refreshing the tokens

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
		httpClient:  httpClient,
		authToken:   authToken,
		credentials: credentials,

		// Intialize all services with dependencies
		UserService:      apis.NewDefultUserService(httpClientWithToken),
		AlbumService:     apis.NewDefaultAlbumService(httpClientWithToken),
		ArtistService:    apis.NewDefaultArtistService(httpClientWithToken),
		AudiobookService: apis.NewDefaultAudiobookService(httpClientWithToken),
		CategoryService:  apis.NewDefaultCategoryService(httpClientWithToken),
		ChapterService:   apis.NewDefaultChapterService(httpClientWithToken),
		EpisodeService:   apis.NewDefaultEpisodeService(httpClientWithToken),
		GenreService:     apis.NewDefaultGenreService(httpClientWithToken),
		MarketService:    apis.NewDefaultMarketService(httpClientWithToken),
		PlayerService:    apis.NewDefaultPlayerService(httpClientWithToken),
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
	if c.authToken == nil {
		return &utils.Error{Type: utils.AppErrorType, AppError: &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgAuthTokenNotInitialized}}
	}
	if c.httpClient == nil {
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
		"refresh_token": c.authToken.RefreshToken,
	}

	// Make a POST request to the token endpoint
	res, err := c.httpClient.Post(context.Background(), consts.EndpointRefresh, headers, nil, formValues, nil)
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
	c.authToken = &authToken

	return nil
}

// CheckAndRefreshTokens checks for the AuthToken expiry and then triggers refresh tokens call if needed.
func (c *Client) CheckAndRefreshTokens() error {
	// To make sure the dependencies have initialized before checking the expiry
	if c.authToken == nil {
		return &utils.Error{Type: utils.AppErrorType, AppError: &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgAuthTokenNotInitialized}}
	}

	// Check if the token has expired
	if time.Now().After(c.authToken.ExpiryTime) {
		if err := c.RefreshTokens(); err != nil {
			return &utils.Error{Type: utils.AppErrorType, AppError: &utils.AppError{Status: http.StatusInternalServerError, Message: consts.MsgRefreshTokensFailure, Err: err}}
		}
	}

	return nil
}
