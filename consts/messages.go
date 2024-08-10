package consts

// Constants for env variables
const (
	EnvClientId     = "SPOTIFY_CLIENT_ID"
	EnvClientSecret = "SPOTIFY_CLIENT_SECRET"
	EnvRedirectUrl  = "SPOTIFY_REDIRECT_URL"
)

// Constants for Spotify API credentials
const (
	MsgClientIdNotFound     = "SPOTIFY_CLIENT_ID not found in environment variables"
	MsgClientSecretNotFound = "SPOTIFY_CLIENT_SECRET not found in environment variables"
	MsgRedirectUrlNotFound  = "SPOTIFY_REDIRECT_URL not found in environment variables"
)

// Constants for others
const (
	MsgStateGenerationFailure        = "State generation failure"
	MsgGettingAuthUrlFailure         = "Getting auth url failure"
	MsgOpeningBrowserFailure         = "Opening browser failure"
	MsgCredentialsNotInitialized     = "Credentials not initialized"
	MsgAuthTokenNotInitialized       = "Auth token not initialized"
	MsgHttpClientNotInitialized      = "Http client not initialized"
	MsgRefreshTokensFailure          = "Refresh tokens failure"
	MsgFailedToReadResponseBody      = "Failed to read response body"
	MsgFailedToUnmarshalResponseData = "Failed to unmarshal response data"

	MsgFailedToGetCurrentUserProfile = "Failed to get current user profile"

	MsgUnsupportedPlatform = "unsupported platform"
	MsgUknownErrorType     = "Unknown error type"

	MsgCodeReceived = "Okay! You can close this window now."
)
