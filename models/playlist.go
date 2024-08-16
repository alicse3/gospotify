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

// TracksBody represents the remove playlist items, tracks body information.
type TracksBody struct {
	Uri string `json:"uri"`
}

// RemovePlaylistItemsBody represents the remove playlist items body information.
type RemovePlaylistItemsBody struct {
	// Required: array of objects.
	// An array of objects containing Spotify URIs of the tracks or episodes to remove.
	// For example: { "tracks": [{ "uri": "spotify:track:4iV5W9uYEdYUVa79Axb7Rh" },{ "uri": "spotify:track:1301WleyT98MSxVHPZCA6M" }] }.
	// A maximum of 100 objects can be sent at once.
	Tracks     []TracksBody `json:"tracks"`
	SnapshotId string       `json:"snapshot_id"`
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

// GetCustomPlaylistCoverImageBody represents the get custom playlist cover image body information.
type GetCustomPlaylistCoverImageBody struct {
	Url    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

// GetCustomPlaylistCoverImageRequest represents the get custom playlist cover image request information.
type GetCustomPlaylistCoverImageRequest struct {
	PlaylistId string // Required: The Spotify ID of the playlist.
	Body       []GetCustomPlaylistCoverImageBody
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
				AlbumType        string   `json:"album_type"`
				TotalTracks      int      `json:"total_tracks"`
				AvailableMarkets []string `json:"available_markets"`
				ExternalUrls     struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href   string `json:"href"`
				Id     string `json:"id"`
				Images []struct {
					Url    string `json:"url"`
					Height int    `json:"height"`
					Width  int    `json:"width"`
				} `json:"images"`
				Name                 string `json:"name"`
				ReleaseDate          string `json:"release_date"`
				ReleaseDatePrecision string `json:"release_date_precision"`
				Restrictions         struct {
					Reason string `json:"reason"`
				} `json:"restrictions"`
				Type    string `json:"type"`
				Uri     string `json:"uri"`
				Artists []struct {
					ExternalUrls struct {
						Spotify string `json:"spotify"`
					} `json:"external_urls"`
					Href string `json:"href"`
					Id   string `json:"id"`
					Name string `json:"name"`
					Type string `json:"type"`
					Uri  string `json:"uri"`
				} `json:"artists"`
				Tracks     AlbumTracks `json:"tracks"`
				Copyrights []struct {
					Text string `json:"text"`
					Type string `json:"type"`
				} `json:"copyrights"`
				ExternalIds struct {
					Isrc string `json:"isrc"`
					Ean  string `json:"ean"`
					Upc  string `json:"upc"`
				} `json:"external_ids"`
				Genres             []string `json:"genres"`
				Label              string   `json:"label"`
				Popularity         int      `json:"popularity"`
				AudioPreviewUrl    string   `json:"audio_preview_url"`
				Description        string   `json:"description"`
				HtmlDescription    string   `json:"html_description"`
				DurationMs         int      `json:"duration_ms"`
				Explicit           bool     `json:"explicit"`
				IsExternallyHosted bool     `json:"is_externally_hosted"`
				IsPlayable         bool     `json:"is_playable"`
				Language           string
				Languages          []string `json:"languages"`
				Show               struct {
					AvailableMarkets []string `json:"available_markets"`
					Copyrights       []struct {
						Text string `json:"text"`
						Type string `json:"type"`
					} `json:"copyrights"`
					Description     string `json:"description"`
					HtmlDescription string `json:"html_description"`
					Explicit        bool   `json:"explicit"`
					ExternalUrls    struct {
						Spotify string `json:"spotify"`
					} `json:"external_urls"`
					Href   string `json:"href"`
					Id     string `json:"id"`
					Images []struct {
						Url    string `json:"url"`
						Height int    `json:"height"`
						Width  int    `json:"width"`
					} `json:"images"`
					IsExternallyHosted bool     `json:"is_externally_hosted"`
					Languages          []string `json:"languages"`
					MediaType          string   `json:"media_type"`
					Name               string   `json:"name"`
					Publisher          string   `json:"publisher"`
					Type               string   `json:"type"`
					Uri                string   `json:"uri"`
					TotalEpisodes      int      `json:"total_episodes"`
				} `json:"show"`
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
			AlbumType        string   `json:"album_type"`
			TotalTracks      int      `json:"total_tracks"`
			AvailableMarkets []string `json:"available_markets"`
			ExternalUrls     struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href   string `json:"href"`
			Id     string `json:"id"`
			Images []struct {
				Url    string `json:"url"`
				Height int    `json:"height"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name                 string `json:"name"`
			ReleaseDate          string `json:"release_date"`
			ReleaseDatePrecision string `json:"release_date_precision"`
			Restrictions         struct {
				Reason string `json:"reason"`
			} `json:"restrictions"`
			Type    string `json:"type"`
			Uri     string `json:"uri"`
			Artists []struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				Id   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				Uri  string `json:"uri"`
			} `json:"artists"`
			Tracks     AlbumTracks `json:"tracks"`
			Copyrights []struct {
				Text string `json:"text"`
				Type string `json:"type"`
			} `json:"copyrights"`
			ExternalIds struct {
				Isrc string `json:"isrc"`
				Ean  string `json:"ean"`
				Upc  string `json:"upc"`
			} `json:"external_ids"`
			Genres             []string `json:"genres"`
			Label              string   `json:"label"`
			Popularity         int      `json:"popularity"`
			AudioPreviewUrl    string   `json:"audio_preview_url"`
			Description        string   `json:"description"`
			HtmlDescription    string   `json:"html_description"`
			DurationMs         int      `json:"duration_ms"`
			Explicit           bool     `json:"explicit"`
			IsExternallyHosted bool     `json:"is_externally_hosted"`
			IsPlayable         bool     `json:"is_playable"`
			Language           string
			Languages          []string `json:"languages"`
			Show               struct {
				AvailableMarkets []string `json:"available_markets"`
				Copyrights       []struct {
					Text string `json:"text"`
					Type string `json:"type"`
				} `json:"copyrights"`
				Description     string `json:"description"`
				HtmlDescription string `json:"html_description"`
				Explicit        bool   `json:"explicit"`
				ExternalUrls    struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href   string `json:"href"`
				Id     string `json:"id"`
				Images []struct {
					Url    string `json:"url"`
					Height int    `json:"height"`
					Width  int    `json:"width"`
				} `json:"images"`
				IsExternallyHosted bool     `json:"is_externally_hosted"`
				Languages          []string `json:"languages"`
				MediaType          string   `json:"media_type"`
				Name               string   `json:"name"`
				Publisher          string   `json:"publisher"`
				Type               string   `json:"type"`
				Uri                string   `json:"uri"`
				TotalEpisodes      int      `json:"total_episodes"`
			} `json:"show"`
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
type PlaylistCoverImage []struct {
	Url    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}
