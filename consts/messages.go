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
	MsgFailedToReadResponseBody      = "Failed to read response body"
	MsgFailedToRefreshTokens         = "Failed to refresh tokens"
	MsgFailedToUnmarshalResponseData = "Failed to unmarshal response data"
	MsgFailedToMarshalRequestData    = "Failed to marshal request data"
	MsgFailedToParseUrl              = "Failed to parse URL"
	MsgFailedToCreatePostRequest     = "Failed to create post request"
	MsgFailedToCreateGetRequest      = "Failed to create get request"
	MsgFailedToCreatePutRequest      = "Failed to create put request"
	MsgFailedToCreateDeleteRequest   = "Failed to create delete request"
	MsgFailedToSendRequest           = "Failed to send request"

	MsgFailedToGetCurrentUserProfile = "Failed to get current user profile"

	MsgFailedToGetAlbum         = "Failed to get an Album"
	MsgFailedToGetAlbums        = "Failed to get Albums"
	MsgFailedToGetTracks        = "Failed to get Tracks"
	MsgFailedToSaveAlbums       = "Failed to save Albums"
	MsgFailedToRemoveAlbums     = "Failed to remove Albums"
	MsgFailedToGetSavedAlbums   = "Failed to get Saved Albums"
	MsgFailedToCheckSavedAlbums = "Failed to check Saved Albums"
	MsgFailedToGetNewReleases   = "Failed to get New Releases"
	MsgPostCallFailed           = "Failed to make a POST API call"

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

	MsgIdRequired                   = "ID is required"
	MsgIdsRequired                  = "IDs are required"
	MsgCategoryIdRequired           = "Category ID is required"
	MsgDeviceIdsRequired            = "Device ID's are required"
	MsgMustBePositiveNumber         = "Must be a positive number"
	MsgStateRequired                = "State is required"
	MsgVolumePercentMustBeInclusive = "Volume percent be a value from 0 to 100 inclusive"
	MsgUriRequired                  = "URI is required"
	MsgPlaylistIdRequired           = "Playlist ID is required"
	MsgUserIdRequired               = "User ID is required"

	MsgFailedToGetEpisode         = "Failed to get an Episode"
	MsgFailedToGetEpisodes        = "Failed to get Episodes"
	MsgFailedToGetSavedEpisodes   = "Failed to get saved Episodes"
	MsgFailedToSaveEpisodes       = "Failed to save Episodes"
	MsgFailedToRemoveEpisodes     = "Failed to remove Episodes"
	MsgFailedToCheckSavedEpisodes = "Failed to check saved Episodes"

	MsgFailedToGetAvailableGenreSeeds = "Failed to get Available Genre Seeds"

	MsgFailedToGetAvailableMarkets = "Failed to get Available Markets"

	MsgFailedToGetPlaybackState         = "Failed to get Playback State"
	MsgFailedToTransferPlayback         = "Failed to transfer Playback"
	MsgFailedToGetAvailableDevices      = "Failed to get Available Devices"
	MsgFailedToGetCurrentlyPlayingTrack = "Failed to get Currently Playing Track"
	MsgFailedToStartOrResumePlayback    = "Failed to Start or Resume the Playback"
	MsgFailedToPausePlayback            = "Failed to Pause the Playback"
	MsgFailedToSkipToNext               = "Failed to Skip to Next"
	MsgFailedToSkipToPrevious           = "Failed to Skip to Previous"
	MsgFailedToSeekToPosition           = "Failed to Seek to Position"
	MsgFailedToSetRepeatMode            = "Failed to set Repeat Mode"
	MsgFailedToSetPlaybackVolume        = "Failed to set Playback Volume"
	MsgFailedToTogglePlaybackShuffle    = "Failed to toggle playback shuffle"
	MsgFailedToGetRecentlyPlayedTracks  = "Failed to get Recently Played Tracks"
	MsgFailedToGetUsersQueue            = "Failed to get Users Queue"
	MsgFailedToAddItemToPlaybackQueue   = "Failed to add item to Playback Queue"

	MsgFailedToGetPlaylist                 = "Failed to get a Playlist"
	MsgFailedToChangePlaylistDetails       = "Failed to change Playlist Details"
	MsgFailedToGetPlaylistItems            = "Failed to get the Playlist Items"
	MsgFailedToUpdatePlaylistItems         = "Failed to update the Playlist Items"
	MsgFailedToAddPlaylistItems            = "Failed to add the Playlist Items"
	MsgFailedToRemovePlaylistItems         = "Failed to remove the Playlist Items"
	MsgFailedToGetCurrentUsersPlaylists    = "Failed to get the Current User's Playlists"
	MsgFailedToGetUsersItems               = "Failed to get the User's Playlists"
	MsgFailedToCreatePlaylist              = "Failed to create a Playlist"
	MsgFailedToGetFeaturedPlaylists        = "Failed to get the Featured Playlists"
	MsgFailedToGetCategoryPlaylists        = "Failed to get the Category Playlists"
	MsgFailedToGetPlaylistCoverImage       = "Failed to get the Playlist Cover Image"
	MsgFailedToAddCustomPlaylistCoverImage = "Failed to add the Custom Playlist Cover Image"

	MsgUnsupportedPlatform = "unsupported platform"
	MsgUknownErrorType     = "Unknown error type"

	MsgCodeReceived = "Okay! You can close this window now."
)
