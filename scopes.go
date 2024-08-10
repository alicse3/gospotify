package gospotify

// Scopes represent the permissions that a Spotify user grants to a third-party application.
// These permissions determine the level of access the application has to the user's data.
// For more details, visit https://developer.spotify.com/documentation/web-api/concepts/scopes

const (
	ScopeUgcImageUpload = "ugc-image-upload" // Write access to user-provided images.

	ScopeUserReadPlaybackState    = "user-read-playback-state"    // Read access to a user’s player state.
	ScopeUserModifyPlaybackState  = "user-modify-playback-state"  // Write access to a user’s playback state
	ScopeUserReadCurrentlyPlaying = "user-read-currently-playing" // Read access to a user’s currently playing content.

	ScopeAppRemoteControl = "app-remote-control" // Remote control playback of Spotify. This scope is currently available to Spotify iOS and Android SDKs.
	ScopeStreaming        = "streaming"          // Control playback of a Spotify track. This scope is currently available to the Web Playback SDK. The user must have a Spotify Premium account.

	ScopePlaylistReadPrivate       = "playlist-read-private"       // Read access to user's private playlists.
	ScopePlaylistReadCollaborative = "playlist-read-collaborative" // Include collaborative playlists when requesting a user's playlists.
	ScopePlaylistModifyPrivate     = "playlist-modify-private"     // Write access to a user's private playlists.
	ScopePlaylistModifyPublic      = "playlist-modify-public"      // Write access to a user's public playlists.

	ScopeUserFollowModify = "user-follow-modify" // Write/delete access to the list of artists and other users that the user follows.
	ScopeUserFollowRead   = "user-follow-read"   // Read access to the list of artists and other users that the user follows.

	ScopeUserReadPlaybackPosition = "user-read-playback-position" // Read access to a user’s playback position in a content.
	ScopeUserTopRead              = "user-top-read"               // Read access to a user's top artists and tracks.
	ScopeReadRecentlyPlayed       = "user-read-recently-played"   // Read access to a user’s recently played tracks.

	ScopeUserLibraryModify = "user-library-modify" // Write/delete access to a user's "Your Music" library.
	ScopeUserLibraryRead   = "user-library-read"   // Read access to a user's library.

	ScopeUserReadEmail   = "user-read-email"   // Read access to user’s email address.
	ScopeUserPrivateRead = "user-read-private" // Read access to user’s subscription details (type of user account).

	ScopeUserSoaLink           = "user-soa-link"           // Link a partner user account to a Spotify user account
	ScopeUserSoaUnlink         = "user-soa-unlink"         // Unlink a partner user account from a Spotify account
	ScopeSoaManageEntitlements = "soa-manage-entitlements" // Modify entitlements for linked users
	ScopeSoaManagePartner      = "soa-manage-partner"      // Update partner information
	ScopeSoaCreatePartner      = "soa-create-partner"      // Create new partners, platform partners only
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
