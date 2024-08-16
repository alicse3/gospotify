package gospotify

// Scopes represent the permissions that a Spotify user grants to a third-party application.
// These permissions determine the level of access the application has to the user's data.
// For more details, visit https://developer.spotify.com/documentation/web-api/concepts/scopes

const (
	// Write access to user-provided images.
	ScopeUgcImageUpload = "ugc-image-upload"

	// Read access to a user’s player state.
	ScopeUserReadPlaybackState = "user-read-playback-state"
	// Write access to a user’s playback state
	ScopeUserModifyPlaybackState = "user-modify-playback-state"
	// Read access to a user’s currently playing content.
	ScopeUserReadCurrentlyPlaying = "user-read-currently-playing"

	// Remote control playback of Spotify. This scope is currently available to Spotify iOS and Android SDKs.
	ScopeAppRemoteControl = "app-remote-control"
	// Control playback of a Spotify track. This scope is currently available to the Web Playback SDK. The user must have a Spotify Premium account.
	ScopeStreaming = "streaming"

	// Read access to user's private playlists.
	ScopePlaylistReadPrivate = "playlist-read-private"
	// Include collaborative playlists when requesting a user's playlists.
	ScopePlaylistReadCollaborative = "playlist-read-collaborative"
	// Write access to a user's private playlists.
	ScopePlaylistModifyPrivate = "playlist-modify-private"
	// Write access to a user's public playlists.
	ScopePlaylistModifyPublic = "playlist-modify-public"

	// Write/delete access to the list of artists and other users that the user follows.
	ScopeUserFollowModify = "user-follow-modify"
	// Read access to the list of artists and other users that the user follows.
	ScopeUserFollowRead = "user-follow-read"

	// Read access to a user’s playback position in a content.
	ScopeUserReadPlaybackPosition = "user-read-playback-position"
	// Read access to a user's top artists and tracks.
	ScopeUserTopRead = "user-top-read"
	// Read access to a user’s recently played tracks.
	ScopeReadRecentlyPlayed = "user-read-recently-played"

	// Write/delete access to a user's "Your Music" library.
	ScopeUserLibraryModify = "user-library-modify"
	// Read access to a user's library.
	ScopeUserLibraryRead = "user-library-read"

	// Read access to user’s email address.
	ScopeUserReadEmail = "user-read-email"
	// Read access to user’s subscription details (type of user account).
	ScopeUserPrivateRead = "user-read-private"

	// Link a partner user account to a Spotify user account
	ScopeUserSoaLink = "user-soa-link"
	// Unlink a partner user account from a Spotify account
	ScopeUserSoaUnlink = "user-soa-unlink"
	// Modify entitlements for linked users
	ScopeSoaManageEntitlements = "soa-manage-entitlements"
	// Update partner information
	ScopeSoaManagePartner = "soa-manage-partner"
	// Create new partners, platform partners only
	ScopeSoaCreatePartner = "soa-create-partner"
)

var (
	// AllScopes contains all the Spotify scopes
	// For details, visit https://developer.spotify.com/documentation/web-api/concepts/scopes
	AllScopes = []string{
		ScopeUgcImageUpload,

		ScopeUserReadPlaybackState,
		ScopeUserModifyPlaybackState,
		ScopeUserReadCurrentlyPlaying,

		ScopeAppRemoteControl,
		ScopeStreaming,

		ScopePlaylistReadPrivate,
		ScopePlaylistReadCollaborative,
		ScopePlaylistModifyPrivate,
		ScopePlaylistModifyPublic,

		ScopeUserFollowModify,
		ScopeUserFollowRead,

		ScopeUserReadPlaybackPosition,
		ScopeUserTopRead,
		ScopeReadRecentlyPlayed,

		ScopeUserLibraryModify,
		ScopeUserLibraryRead,

		ScopeUserReadEmail,
		ScopeUserPrivateRead,

		// Add below scopes explicitly if required
		// ScopeUserSoaLink,
		// ScopeUserSoaUnlink,
		// ScopeSoaManageEntitlements,
		// ScopeSoaManagePartner,
		// ScopeSoaCreatePartner,
	}
)
