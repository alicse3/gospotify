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
)
