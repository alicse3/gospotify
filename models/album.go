package models

// GetAlbumRequest represents the get album's request information.
type GetAlbumRequest struct {
	Id     string // Required: The Spotify ID of the album.
	Market string
}

// GetAlbumsRequest represents the get albums request information.
type GetAlbumsRequest struct {
	Ids    string // Required: A comma-separated list of the Spotify IDs for the albums. Maximum: 20 IDs.
	Market string
}

// GetAlbumTracksRequest represents the get album tracks request information.
type GetAlbumTracksRequest struct {
	Id     string // Required: The Spotify ID of the album.
	Market string
	Limit  int
	Offset int
}

// GetSavedAlbumsRequest represents the get saved albums request information.
type GetSavedAlbumsRequest struct {
	Limit  int
	Offset int
	Market string
}

// SaveAlbumsRequest represents the save albums request information.
type SaveAlbumsRequest struct {
	Ids  string // Required: A comma-separated list of the Spotify IDs for the albums. Maximum: 20 IDs.
	Body struct {
		Ids []string `json:"ids"`
	}
}

// RemoveAlbumsRequest represents the remove albums request information.
type RemoveAlbumsRequest struct {
	Ids  string // Required: A comma-separated list of the Spotify IDs for the albums. Maximum: 20 IDs.
	Body struct {
		Ids []string `json:"ids"`
	}
}

// CheckSavedAlbumsRequest represents the check saved albums request information.
type CheckSavedAlbumsRequest struct {
	Ids string // Required: A comma-separated list of the Spotify IDs for the albums. Maximum: 20 IDs.
}

// GetNewReleasesRequest represents the get new releases request information.
type GetNewReleasesRequest struct {
	Limit  int
	Offset int
}

// AlbumTracks represents the track's information retrieved from the Spotify API.
type AlbumTracks struct {
	Href     string `json:"href"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
	Items    []struct {
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
		ExternalUrls     struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href       string `json:"href"`
		Id         string `json:"id"`
		IsPlayable bool   `json:"is_playable"`
		LinkedFrom struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href string `json:"href"`
			Id   string `json:"id"`
			Type string `json:"type"`
			Uri  string `json:"uri"`
		} `json:"linked_from"`
		Restrictions struct {
			Reason string `json:"reason"`
		} `json:"restrictions"`
		Name        string `json:"name"`
		PreviewUrl  string `json:"preview_url"`
		TrackNumber int    `json:"track_number"`
		Type        string `json:"type"`
		Uri         string `json:"uri"`
		IsLocal     bool   `json:"is_local"`
	} `json:"items"`
}

// Album represents the album's information retrieved from the Spotify API.
type Album struct {
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
	Genres     []string `json:"genres"`
	Label      string   `json:"label"`
	Popularity int      `json:"popularity"`
}

// Albums represents the albums information retrieved from the Spotify API.
type Albums struct {
	Albums []Album `json:"albums"`
}

// SavedAlbums represents the saved albums information retrieved from the Spotify API.
type SavedAlbums struct {
	Href     string `json:"href"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
	Items    []struct {
		AddedAt string `json:"added_at"`
		Album   Album  `json:"album"`
	} `json:"items"`
}

// CheckSavedAlbums represents the check saved albums information retrieved from the Spotify API.
type CheckSavedAlbums []bool

// NewlyReleasedAlbums represents the newly released albums information retrieved from the Spotify API.
type NewlyReleasedAlbums struct {
	Albums struct {
		Href     string  `json:"href"`
		Limit    int     `json:"limit"`
		Next     string  `json:"next"`
		Offset   int     `json:"offset"`
		Previous string  `json:"previous"`
		Total    int     `json:"total"`
		Items    []Album `json:"items"`
	} `json:"albums"`
}
