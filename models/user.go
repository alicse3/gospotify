package models

// GetUsersTopItemsRequest represents the get user's top items request information.
type GetUsersTopItemsRequest struct {
	Type      string // Required: The type of entity to return. Valid values: artists or tracks
	TimeRange string
	Limit     int
	Offset    int
}

// GetUsersProfileRequest represents the get user's profile request information.
type GetUsersProfileRequest struct {
	UserId string // Required: The user's Spotify user ID.
}

// FollowPlaylistBody represents the follow playlist body information.
type FollowPlaylistBody struct {
	Public bool `json:"public"`
}

// FollowPlaylistRequest represents the follow playlist request information.
type FollowPlaylistRequest struct {
	PlaylistId string // Required: The Spotify ID of the playlist.
	Body       FollowPlaylistBody
}

// UnfollowPlaylistRequest represents the unfollow playlist request information.
type UnfollowPlaylistRequest struct {
	PlaylistId string // Required: The Spotify ID of the playlist.
}

// GetFollowedArtistsRequest represents the get followed artists request information.
type GetFollowedArtistsRequest struct {
	Type  string // Required: The ID type: currently only artist is supported.
	After string
	Limit int
}

// FollowArtistsOrUsersBody represents the follow artists or users body information.
type FollowArtistsOrUsersBody struct {
	Ids []string `json:"ids"`
}

// FollowArtistsOrUsersRequest represents the follow artists or users request information.
type FollowArtistsOrUsersRequest struct {
	Type string // Required: The ID type. Allowed values: "artist", "user"
	Ids  string // Required: A comma-separated list of the artist or the user Spotify IDs. A maximum of 50 IDs can be sent in one request.
	Body FollowArtistsOrUsersBody
}

// UnfollowArtistsOrUsersBody represents the unfollow artists or users body information.
type UnfollowArtistsOrUsersBody struct {
	Ids []string `json:"ids"`
}

// FollowArtistsOrUsersRequest represents the unfollow artists or users request information.
type UnfollowArtistsOrUsersRequest struct {
	Type string // Required: The ID type: either artist or user.
	Ids  string // Required: A comma-separated list of the artist or the user Spotify IDs. For example: ids=74ASZWbe4lXaubB36ztrGX,08td7MxkoHQkXnWAYD8d6Q. A maximum of 50 IDs can be sent in one request.
	Body UnfollowArtistsOrUsersBody
}

// FollowArtistsOrUsersRequest represents the request information about, if user follows artists or users.
type UserFollowsArtistsOrUsersRequest struct {
	Type string // Required: The ID type: either artist or user.
	Ids  string // Required: A comma-separated list of the artist or the user Spotify IDs to check. For example: ids=74ASZWbe4lXaubB36ztrGX,08td7MxkoHQkXnWAYD8d6Q. A maximum of 50 IDs can be sent in one request.
}

// CurrentUserFollowsPlaylistRequest represents the request information about, if current user follows playlists.
type CurrentUserFollowsPlaylistRequest struct {
	PlaylistId string // Required: The Spotify ID of the playlist.
	Ids        string
}

// User represents the user's profile information retrieved from the Spotify API.
type User struct {
	Country         string `json:"country"`
	DisplayName     string `json:"display_name"`
	Email           string `json:"email"`
	ExplicitContent struct {
		FilterEnabled bool `json:"filter_enabled"`
		FilterLocked  bool `json:"filter_locked"`
	} `json:"explicit_content"`
	ExternalUrls struct {
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
	Product string `json:"product"`
	Type    string `json:"type"`
	Uri     string `json:"uri"`
}

// UserTopItems represents the user's top items information retrieved from the Spotify API.
type UserTopItems struct {
	Href     string `json:"href"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
	Items    []struct {
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Followers struct {
			Href  any `json:"href"`
			Total int `json:"total"`
		} `json:"followers"`
		Genres []string `json:"genres"`
		Href   string   `json:"href"`
		Id     string   `json:"id"`
		Images []struct {
			Url    string `json:"url"`
			Height int    `json:"height"`
			Width  int    `json:"width"`
		} `json:"images"`
		Name       string `json:"name"`
		Popularity int    `json:"popularity"`
		Type       string `json:"type"`
		Uri        string `json:"uri"`
		Album      struct {
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
		} `json:"album"`
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
		AvailableMarkets []string `json:"available_markets"`
		DiscNumber       int      `json:"disc_number"`
		DurationMs       int      `json:"duration_ms"`
		Explicit         bool     `json:"explicit"`
		ExternalIds      struct {
			Isrc string `json:"isrc"`
			Ean  string `json:"ean"`
			Upc  string `json:"upc"`
		} `json:"external_ids"`
		IsPlayable bool `json:"is_playable"`
		LinkedFrom struct {
		} `json:"linked_from"`
		Restrictions struct {
			Reason string `json:"reason"`
		} `json:"restrictions"`
		PreviewUrl  string `json:"preview_url"`
		TrackNumber int    `json:"track_number"`
		IsLocal     bool   `json:"is_local"`
	} `json:"items"`
}

// UserProfile represents the user profile information retrieved from the Spotify API.
type UserProfile struct {
	Country      string `json:"country"`
	DisplayName  string `json:"display_name"`
	ExternalUrls struct {
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
	Type string `json:"type"`
	Uri  string `json:"uri"`
}

// FollowedArtists represents the followed artists information retrieved from the Spotify API.
type FollowedArtists struct {
	Href    string `json:"href"`
	Limit   int    `json:"limit"`
	Next    string `json:"next"`
	Cursors struct {
		After  string `json:"after"`
		Before string `json:"before"`
	} `json:"cursors"`
	Total int      `json:"total"`
	Items []Artist `json:"items"`
}

// CheckUserFollowsArtistsOrUsers represents the information about whether a user follows artists or users, retrieved from the Spotify API.
type CheckUserFollowsArtistsOrUsers []bool

// CheckCurrentUserFollowsPlaylist represents the information about whether the current user follows playlist, retrieved from the Spotify API.
type CheckCurrentUserFollowsPlaylist []bool
