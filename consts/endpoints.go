package consts

const (
	BaseUrlAccounts = "https://accounts.spotify.com" // Base address of the Spotify token API.
	BaseUrlApi      = "https://api.spotify.com"      // Base address of the Spotify Web API.
)

const (
	EndpointAuthorize = "/authorize" // Spotify authorization endpoint.
	EndpointToken     = "/api/token" // Spotify token endpoint.

	EndpointMe = "/v1/me" // Get detailed profile information about the current user (including the current user's username).

	// Albums related endpoints
	EndpointAlbum         = "/v1/albums/%s"           // Get Spotify catalog information for a single album.
	EndpointAlbums        = "/v1/albums"              // Get Spotify catalog information for multiple albums.
	EndpointAlbumTracks   = "/v1/albums/%s/tracks"    // Get Spotify catalog information about an album’s tracks.
	EndpointMyAlbums      = "/v1/me/albums"           // To work with the current Spotify user's 'Your Music' library.
	EndpointCheckMyAlbums = "/v1/me/albums/contains"  // Check if one or more albums is already saved in the current Spotify user's 'Your Music' library.
	EndpointNewReleases   = "/v1/browse/new-releases" // Get a list of new album releases featured in Spotify.

	// Artists related endpoints
	EndpointArtist          = "/v1/artists/%s"                 // Get Spotify catalog information for a single artist identified by their unique Spotify ID.
	EndpointArtists         = "/v1/artists"                    // Get Spotify catalog information for several artists based on their Spotify IDs.
	EndpointArtistAlbums    = "/v1/artists/%s/albums"          // Get Spotify catalog information about an artist's albums.
	EndpointArtistTopTracks = "/v1/artists/%s/top-tracks"      // Get Spotify catalog information about an artist's top tracks by country.
	EndpointRelatedArtists  = "/v1/artists/%s/related-artists" // Get Spotify catalog information about artists similar to a given artist. Similarity is based on analysis of the Spotify community's listening history.

	// Audiobooks related endpoints
	EndpointAudiobook         = "/v1/audiobooks/%s"          // Get Spotify catalog information for a single audiobook. Audiobooks are only available within the US, UK, Canada, Ireland, New Zealand and Australia markets.
	EndpointAudiobooks        = "/v1/audiobooks"             // Get Spotify catalog information for several audiobooks identified by their Spotify IDs. Audiobooks are only available within the US, UK, Canada, Ireland, New Zealand and Australia markets.
	EndpointAudiobookChapters = "/v1/audiobooks/%s/chapters" // Get Spotify catalog information about an audiobook's chapters. Audiobooks are only available within the US, UK, Canada, Ireland, New Zealand and Australia markets.
	EndpointMyAudiobooks      = "/v1/me/audiobooks"          // To work with the audiobooks saved in the current Spotify user's 'Your Music' library.
	EndpointMySavedAudiobooks = "/v1/me/audiobooks/contains" // Check if one or more audiobooks are already saved in the current Spotify user's library.

	// Categories related endpoints
	EndpointBrowseCategories = "/v1/browse/categories"    // Get a list of categories used to tag items in Spotify (on, for example, the Spotify player’s “Browse” tab).
	EndpointBrowseCategory   = "/v1/browse/categories/%s" // Get a single category used to tag items in Spotify (on, for example, the Spotify player’s “Browse” tab).

	// Chapters related endpoints
	EndpointGetChapter  = "/v1/chapters/%s" // Get Spotify catalog information for a single audiobook chapter. Chapters are only available within the US, UK, Canada, Ireland, New Zealand and Australia markets.
	EndpointGetChapters = "/v1/chapters"    // Get Spotify catalog information for several audiobook chapters identified by their Spotify IDs. Chapters are only available within the US, UK, Canada, Ireland, New Zealand and Australia markets.

	// Episodes related endpoints
	EndpointEpisode         = "/v1/episodes/%s"          // Get Spotify catalog information for a single episode identified by its unique Spotify ID.
	EndpointEpisodes        = "/v1/episodes"             // A comma-separated list of the Spotify IDs for the episodes. Maximum: 50 IDs.
	EndpointMyEpisodes      = "/v1/me/episodes"          // To work with the list of episodes in the current Spotify user's library.
	EndpointCheckMyEpisodes = "/v1/me/episodes/contains" // Check if one or more episodes is already saved in the current Spotify user's 'Your Episodes' library.

	// Genres related endpoints
	EndpointGetAvailableGenreSeeds = "/v1/recommendations/available-genre-seeds" // Retrieve a list of available genres seed parameter values for recommendations.

	// Markets related endpoints
	EndpointGetAvailableMarkets = "/v1/markets" // Get the list of markets where Spotify is available.

	// Player related endpoints
	EndpointPlaybackState         = "/v1/me/player"                   // Get information about the user’s current playback state, including track or episode, progress, and active device.
	EndpointAvailableDevices      = "/v1/me/player/devices"           // Get information about a user’s available Spotify Connect devices. Some device models are not supported and will not be listed in the API response.
	EndpointCurrentlyPlayingTrack = "/v1/me/player/currently-playing" // Get the object currently being played on the user's Spotify account.
	EndpointStartOrResumePlayback = "/v1/me/player/play"              // Start a new context or resume current playback on the user's active device. This API only works for users who have Spotify Premium. The order of execution is not guaranteed when you use this API with other Player API endpoints.
	EndpointPausePlayback         = "/v1/me/player/pause"             // Pause playback on the user's account. This API only works for users who have Spotify Premium. The order of execution is not guaranteed when you use this API with other Player API endpoints.
	EndpointSkipToNext            = "/v1/me/player/next"              // Skips to next track in the user’s queue. This API only works for users who have Spotify Premium. The order of execution is not guaranteed when you use this API with other Player API endpoints.
	EndpointSkipToPrevious        = "/v1/me/player/previous"          // Skips to previous track in the user’s queue. This API only works for users who have Spotify Premium. The order of execution is not guaranteed when you use this API with other Player API endpoints.
	EndpointSeekToPosition        = "/v1/me/player/seek"              // Seeks to the given position in the user’s currently playing track. This API only works for users who have Spotify Premium. The order of execution is not guaranteed when you use this API with other Player API endpoints.
	EndpointRepeatMode            = "/v1/me/player/repeat"            // Set the repeat mode for the user's playback. This API only works for users who have Spotify Premium. The order of execution is not guaranteed when you use this API with other Player API endpoints.
	EndpointPlaybackVolume        = "/v1/me/player/volume"            // Set the volume for the user’s current playback device. This API only works for users who have Spotify Premium. The order of execution is not guaranteed when you use this API with other Player API endpoints.
	EndpointTogglePlaybackShuffle = "/v1/me/player/shuffle"           // Toggle shuffle on or off for user’s playback. This API only works for users who have Spotify Premium. The order of execution is not guaranteed when you use this API with other Player API endpoints.
	EndpointRecentlyPlayedTracks  = "/v1/me/player/recently-played"   // Get tracks from the current user's recently played tracks. Note: Currently doesn't support podcast episodes.
	EndpointUsersQueue            = "/v1/me/player/queue"             // Get the list of objects that make up the user's queue.
	EndpointPlaybackQueue         = "/v1/me/player/queue"             // Add an item to the end of the user's current playback queue. This API only works for users who have Spotify Premium. The order of execution is not guaranteed when you use this API with other Player API endpoints.

	// Playlists related endpoints
	EndpointPlaylists             = "/v1/playlists/%s"                   // Get a playlist owned by a Spotify user.
	EndpointPlaylistItems         = "/v1/playlists/%s/tracks"            // Get or Update or Add or Remove the full details of the items of a playlist owned by a Spotify user.
	EndpointCurrentUsersPlaylists = "/v1/me/playlists"                   // Get a list of the playlists owned or followed by the current Spotify user.
	EndpointUsersPlaylists        = "/v1/users/%s/playlists"             // Get or Create a playlist for a Spotify user.
	EndpointFeaturedPlaylists     = "/v1/browse/featured-playlists"      // Get a list of Spotify featured playlists (shown, for example, on a Spotify player's 'Browse' tab).
	EndpointCategoryPlaylists     = "/v1/browse/categories/%s/playlists" // Get a list of Spotify playlists tagged with a particular category.
	EndpointPlaylistCoverImage    = "/v1/playlists/%s/images"            // Get the current image associated with a specific playlist. Or Replace the image used to represent a specific playlist.

	// Search related endpoints
	EndpointSearch = "/v1/search" // Get Spotify catalog information about albums, artists, playlists, tracks, shows, episodes or audiobooks that match a keyword string. Audiobooks are only available within the US, UK, Canada, Ireland, New Zealand and Australia markets.

	// Shows related endpoints
	EndpointShow            = "/v1/shows/%s"          // Get Spotify catalog information for a single show identified by its unique Spotify ID.
	EndpointShows           = "/v1/shows"             // Get Spotify catalog information for several shows based on their Spotify IDs.
	EndpointShowEpisodes    = "/v1/shows/%s/episodes" // Get Spotify catalog information about an show’s episodes. Optional parameters can be used to limit the number of episodes returned.
	EndpointSaveShows       = "/v1/me/shows"          // Get or Remove or Save shows to/from current Spotify user's library.
	EndpointCheckSavedShows = "/v1/me/shows/contains" // Check if one or more shows is already saved in the current Spotify user's library.

	// Tracks related endpoints
	EndpointTrack                      = "/v1/tracks/%s"          // Get Spotify catalog information for a single track identified by its unique Spotify ID.
	EndpointTracks                     = "/v1/tracks"             // Get Spotify catalog information for multiple tracks based on their Spotify IDs.
	EndpointSaveTracks                 = "/v1/me/tracks"          // Get or Save or Remove tracks from/to the current user's 'Your Music' library.
	EndpointCheckSavedTracks           = "/v1/me/tracks/contains" // Check if one or more tracks is already saved in the current Spotify user's 'Your Music' library.
	EndpointSeveralTracksAudioFeatures = "/v1/audio-features"     // Get audio features for multiple tracks based on their Spotify IDs.
	EndpointTracksTracksAudioFeatures  = "/v1/audio-features/%s"  // Get audio feature information for a single track identified by its unique Spotify ID.
	EndpointTracksAudioAnalysis        = "/v1/audio-analysis/%s"  // Get a low-level audio analysis for a track in the Spotify catalog. The audio analysis describes the track’s structure and musical content, including rhythm, pitch, and timbre.
	// Recommendations are generated based on the available information for a given seed entity and matched against similar artists and tracks. If there is sufficient information about the provided seeds, a list of tracks will be returned together with pool size details.
	// For artists and tracks that are very new or obscure there might not be enough data to generate a list of tracks.
	EndpointRecommendations = "/v1/recommendations"
)
