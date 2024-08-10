package consts

const (
	// BaseUrlAccounts is the address of the Spotify token URL.
	BaseUrlAccounts = "https://accounts.spotify.com"

	// BaseUrlApi is the base address of the Spotify Web API.
	BaseUrlApi = "https://api.spotify.com"
)

const (
	// EndpointAuthorize is the Spotify authorization endpoint.
	EndpointAuthorize = "/authorize"

	// EndpointToken is the Spotify token endpoint.
	EndpointToken = "/api/token"

	// EndpointRefresh is the Spotify refresh endpoint for refreshing the access token.
	EndpointRefresh = "/refresh"

	// EndpointMe is to get detailed profile information about the current user.
	EndpointMe = "/v1/me"
)
