package consts

const (
	BaseUrlAccounts = "https://accounts.spotify.com" // Base address of the Spotify token API.
	BaseUrlApi      = "https://api.spotify.com"      // Base address of the Spotify Web API.
)

const (
	EndpointAuthorize = "/authorize" // Spotify authorization endpoint.
	EndpointToken     = "/api/token" // Spotify token endpoint.
	EndpointRefresh   = "/refresh"   // Spotify refresh endpoint.

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
	EndpointMyAudiobooks      = "/v1/me/audiobooks"          // To work with the the audiobooks saved in the current Spotify user's 'Your Music' library.
	EndpointMySavedAudiobooks = "/v1/me/audiobooks/contains" // Check if one or more audiobooks are already saved in the current Spotify user's library.
)
