package models

// GetPlaylistRequest represents the get playlist request information.
type GetPlaylistRequest struct {
	PlaylistId      string // Required: The Spotify ID of the playlist.
	Market          string
	Fields          string
	AdditionalTypes string
}

// ChangePlaylistDetailsBody represents the change playlist details body information.
type ChangePlaylistDetailsBody struct {
	Name          string `json:"name"`
	Public        bool   `json:"public"`
	Collaborative bool   `json:"collaborative"`
	Description   string `json:"description"`
}

// ChangePlaylistDetailsRequest represents the change playlist details request information.
type ChangePlaylistDetailsRequest struct {
	PlaylistId string // Required: The Spotify ID of the playlist.
	Body       ChangePlaylistDetailsBody
}

// GetPlaylistItemsRequest represents the get playlist items request information.
type GetPlaylistItemsRequest struct {
	PlaylistId      string // Required: The Spotify ID of the playlist.
	Market          string
	Fields          string
	Limit           int
	Offset          int
	AdditionalTypes string
}

// UpdatePlaylistItemsBody represents the update playlist items body information.
type UpdatePlaylistItemsBody struct {
	Uris         []string `json:"uris"`
	RangeStart   int      `json:"range_start"`
	InsertBefore int      `json:"insert_before"`
	RangeLength  int      `json:"range_length"`
	SnapshotId   string   `json:"snapshot_id"`
}

// UpdatePlaylistItemsRequest represents the update playlist items request information.
type UpdatePlaylistItemsRequest struct {
	PlaylistId string // Required: The Spotify ID of the playlist.
	Uris       string
	Body       UpdatePlaylistItemsBody
}

// AddPlaylistItemsBody represents the add playlist items body information.
type AddPlaylistItemsBody struct {
	Uris     []string `json:"uris"`
	Position int      `json:"position"`
}

// AddPlaylistItemsRequest represents the add playlist items request information.
type AddPlaylistItemsRequest struct {
	PlaylistId string // Required: The Spotify ID of the playlist.
	Position   int
	Uris       string
	Body       AddPlaylistItemsBody
}

// RemovePlaylistItemsBody represents the remove playlist items body information.
type RemovePlaylistItemsBody struct {
	// Required: array of objects.
	// An array of objects containing Spotify URIs of the tracks or episodes to remove.
	// For example: { "tracks": [{ "uri": "spotify:track:4iV5W9uYEdYUVa79Axb7Rh" },{ "uri": "spotify:track:1301WleyT98MSxVHPZCA6M" }] }.
	// A maximum of 100 objects can be sent at once.
	Tracks []struct {
		Uri string `json:"uri"`
	} `json:"tracks"`
	SnapshotId string `json:"snapshot_id"`
}

// RemovePlaylistItemsRequest represents the remove playlist items request information.
type RemovePlaylistItemsRequest struct {
	PlaylistId string // Required: The Spotify ID of the playlist.
	Body       RemovePlaylistItemsBody
}

// GetCurrentUsersPlaylistsRequest represents the get current user's playlists request information.
type GetCurrentUsersPlaylistsRequest struct {
	Limit  int
	Offset int
}

// GetUsersPlaylistsRequest represents the get user's playlists request information.
type GetUsersPlaylistsRequest struct {
	UserId string // Required: The user's Spotify user ID.
	Limit  int
	Offset int
}

// CreatePlaylistBody represents the create playlist body information.
type CreatePlaylistBody struct {
	// Required: The name for the new playlist, for example "Your Coolest Playlist".
	// This name does not need to be unique; a user may have several playlists with the same name.
	Name          string `json:"name"`
	Public        bool   `json:"public"`
	Collaborative bool   `json:"collaborative"`
	Description   string `json:"description"`
}

// CreatePlaylistRequest represents the create playlist request information.
type CreatePlaylistRequest struct {
	UserId string // Required: The user's Spotify user ID.
	Body   CreatePlaylistBody
}

// GetFeaturedPlaylistsRequest represents the get featured playlists request information.
type GetFeaturedPlaylistsRequest struct {
	Locale string
	Limit  int
	Offset int
}

// GetCategoryPlaylistsRequest represents the get category playlists request information.
type GetCategoryPlaylistsRequest struct {
	CategoryId string // Required: The Spotify category ID for the category.
	Limit      int
	Offset     int
}

// GetPlaylistCoverImageRequest represents the get playlist cover image request information.
type GetPlaylistCoverImageRequest struct {
	PlaylistId string // Required: The Spotify ID of the playlist.
}

// GetCustomPlaylistCoverImageRequest represents the get custom playlist cover image request information.
type GetCustomPlaylistCoverImageRequest struct {
	PlaylistId string // Required: The Spotify ID of the playlist.
	Body       string
}

// Playlist represents the playlist information retrieved from the Spotify API.
type Playlist struct {
	Collaborative bool   `json:"collaborative"`
	Description   string `json:"description"`
	ExternalUrls  struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Followers struct {
		Href  string `json:"href"`
		Total int    `json:"total"`
	} `json:"followers"`
	Href   string `json:"href"`
	Id     string `json:"id"`
	Images []struct {
		Url    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	} `json:"images"`
	Name  string `json:"name"`
	Owner struct {
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Followers struct {
			Href  string `json:"href"`
			Total int    `json:"total"`
		} `json:"followers"`
		Href        string `json:"href"`
		Id          string `json:"id"`
		Type        string `json:"type"`
		Uri         string `json:"uri"`
		DisplayName string `json:"display_name"`
	} `json:"owner"`
	Public     bool   `json:"public"`
	SnapshotId string `json:"snapshot_id"`
	Tracks     struct {
		Href     string `json:"href"`
		Limit    int    `json:"limit"`
		Next     string `json:"next"`
		Offset   int    `json:"offset"`
		Previous string `json:"previous"`
		Total    int    `json:"total"`
		Items    []struct {
			AddedAt string `json:"added_at"`
			AddedBy struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Followers struct {
					Href  string `json:"href"`
					Total int    `json:"total"`
				} `json:"followers"`
				Href string `json:"href"`
				Id   string `json:"id"`
				Type string `json:"type"`
				Uri  string `json:"uri"`
			} `json:"added_by"`
			IsLocal bool `json:"is_local"`
			Track   struct {
				Album   Album   `json:"album"`
				Episode Episode `json:"episode"`
			} `json:"track"`
		} `json:"items"`
	} `json:"tracks"`
	Type string `json:"type"`
	Uri  string `json:"uri"`
}

// PlaylistItems represents the playlist items information retrieved from the Spotify API.
type PlaylistItems struct {
	Href     string `json:"href"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
	Items    []struct {
		AddedAt string `json:"added_at"`
		AddedBy struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Followers struct {
				Href  string `json:"href"`
				Total int    `json:"total"`
			} `json:"followers"`
			Href string `json:"href"`
			Id   string `json:"id"`
			Type string `json:"type"`
			Uri  string `json:"uri"`
		} `json:"added_by"`
		IsLocal bool `json:"is_local"`
		Track   struct {
			Album   Album   `json:"album"`
			Episode Episode `json:"episode"`
		} `json:"track"`
	} `json:"items"`
}

// UpdatePlaylistItems represents the update playlist items information retrieved from the Spotify API.
type UpdatePlaylistItems struct {
	SnapshotId string `json:"snapshot_id"`
}

// AddPlaylistItems represents the add playlist items information retrieved from the Spotify API.
type AddPlaylistItems struct {
	SnapshotId string `json:"snapshot_id"`
}

// RemovePlaylistItems represents the remove playlist items information retrieved from the Spotify API.
type RemovePlaylistItems struct {
	SnapshotId string `json:"snapshot_id"`
}

// Playlists represents the current user's playlists information retrieved from the Spotify API.
type Playlists struct {
	Href     string `json:"href"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
	Items    []struct {
		Collaborative bool   `json:"collaborative"`
		Description   string `json:"description"`
		ExternalUrls  struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Followers struct {
			Href  string `json:"href"`
			Total int    `json:"total"`
		} `json:"followers"`
		Href   string `json:"href"`
		Id     string `json:"id"`
		Images []struct {
			Url    string `json:"url"`
			Height int    `json:"height"`
			Width  int    `json:"width"`
		} `json:"images"`
		Name  string `json:"name"`
		Owner struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Followers struct {
				Href  string `json:"href"`
				Total int    `json:"total"`
			} `json:"followers"`
			Href        string `json:"href"`
			Id          string `json:"id"`
			Type        string `json:"type"`
			Uri         string `json:"uri"`
			DisplayName string `json:"display_name"`
		} `json:"owner"`
		Public     bool   `json:"public"`
		SnapshotId string `json:"snapshot_id"`
		Tracks     struct {
			Href  string `json:"href"`
			Total int    `json:"total"`
		} `json:"tracks"`
		Type string `json:"type"`
		Uri  string `json:"uri"`
	} `json:"items"`
}

// FeaturedPlaylists represents the featured playlists information retrieved from the Spotify API.
type FeaturedPlaylists struct {
	Message   string    `json:"message"`
	Playlists Playlists `json:"playlists"`
}

// CategoryPlaylists represents the category playlists information retrieved from the Spotify API.
type CategoryPlaylists struct {
	Message   string    `json:"message"`
	Playlists Playlists `json:"playlists"`
}

// PlaylistCoverImage represents the playlist cover image information retrieved from the Spotify API.
type PlaylistCoverImage struct {
	Url    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}
