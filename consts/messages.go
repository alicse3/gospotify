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

	MsgFailedToGetAlbum         = "Failed to get an Album"
	MsgFailedToGetAlbums        = "Failed to get Albums"
	MsgFailedToGetTracks        = "Failed to get Tracks"
	MsgFailedToSaveAlbums       = "Failed to save Albums"
	MsgFailedToRemoveAlbums     = "Failed to remove Albums"
	MsgFailedToGetSavedAlbums   = "Failed to get Saved Albums"
	MsgFailedToCheckSavedAlbums = "Failed to check Saved Albums"
	MsgFailedToGetNewReleases   = "Failed to get New Releases"

	MsgFailedToGetArtist          = "Failed to get an Artist"
	MsgFailedToGetArtists         = "Failed to get an Artists"
	MsgFailedToGetArtistAlbums    = "Failed to get Artist Albums"
	MsgFailedToGetArtistTopTracks = "Failed to get Artist Top Tracks"
	MsgFailedToGetRelatedArtists  = "Failed to get Related Artists"

	MsgFailedToGetAudiobook         = "Failed to get an Audiobook"
	MsgFailedToGetAudiobooks        = "Failed to get Audiobooks"
	MsgFailedToGetAudiobookChapters = "Failed to get Audiobook chapters"
	MsgFailedToGetSavedAudiobooks   = "Failed to get saves Audiobooks"
	MsgFailedToSaveAudiobooks       = "Failed to save Audiobooks"
	MsgFailedToCheckSavedAudiobooks = "Failed to check saved Audiobooks"

	MsgFailedToGetBrowseCategories = "Failed to get Browse Categories"
	MsgFailedToGetBrowseCategory   = "Failed to get Browse Category"

	MsgFailedToGetChapter  = "Failed to get a chapter"
	MsgFailedToGetChapters = "Failed to get chapters"

	MsgIdRequired         = "ID is required"
	MsgIdsRequired        = "IDs are required"
	MsgCategoryIdRequired = "Category ID is required"

	MsgFailedToGetEpisode         = "Failed to get an Episode"
	MsgFailedToGetEpisodes        = "Failed to get Episodes"
	MsgFailedToGetSavedEpisodes   = "Failed to get saved Episodes"
	MsgFailedToSaveEpisodes       = "Failed to save Episodes"
	MsgFailedToRemoveEpisodes     = "Failed to remove Episodes"
	MsgFailedToCheckSavedEpisodes = "Failed to check saved Episodes"

	MsgFailedToGetAvailableGenreSeeds = "Failed to get Available Genre Seeds"

	MsgFailedToGetAvailableMarkets = "Failed to get Available Markets"

	MsgUnsupportedPlatform = "unsupported platform"
	MsgUknownErrorType     = "Unknown error type"

	MsgCodeReceived = "Okay! You can close this window now."
)
