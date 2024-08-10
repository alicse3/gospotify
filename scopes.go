package gospotify

// Scopes represent the permissions that a Spotify user grants to a third-party application.
// These permissions determine the level of access the application has to the user's data.
// For more details, visit https://developer.spotify.com/documentation/web-api/concepts/scopes

const (
	ScopeUserPrivateRead = "user-read-private" // Read access to user’s subscription details (type of user account).
	ScopeUserReadEmail   = "user-read-email"   // Read access to user’s email address.
)

var (
	// AllScopes contains all the Spotify scopes
	// For details, visit https://developer.spotify.com/documentation/web-api/concepts/scopes
	AllScopes = []string{
		ScopeUserPrivateRead,
		ScopeUserReadEmail,
	}
)
