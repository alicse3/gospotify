package models

// GetArtistRequest represents the get artist's request information.
type GetArtistRequest struct {
	Id string // Required: The Spotify ID of the artist.
}

// GetArtistsRequest represents the get artists request information.
type GetArtistsRequest struct {
	Ids string // Required: A comma-separated list of the Spotify IDs for the artists. Maximum: 50 IDs.
}

// GetArtistAlbumsRequest represents the get artists albums request information.
type GetArtistAlbumsRequest struct {
	Id            string // Required: The Spotify ID of the artist.
	IncludeGroups string
	Market        string
	Limit         int
	Offset        int
}

// GetArtistTopTracksRequest represents the get artists top tracks request information.
type GetArtistTopTracksRequest struct {
	Id     string // Required: The Spotify ID of the artist.
	Market string
}

// GetRelatedArtistsRequest represents the get related artists request information.
type GetRelatedArtistsRequest struct {
	Id string // Required: The Spotify ID of the artist.
}

// Artist represents the artist's information retrieved from the Spotify API.
type Artist struct {
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
}

// Artists represents the artist's information retrieved from the Spotify API.
type Artists struct {
	Artists []Artist `json:"artists"`
}

// ArtistAlbum represents the artist's album information retrieved from the Spotify API.
type ArtistAlbum struct {
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
	AlbumGroup string `json:"album_group"`
}

// ArtistAlbums represents the artist's albums information retrieved from the Spotify API.
type ArtistAlbums struct {
	Href     string `json:"href"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
	Items    []struct {
		ArtistAlbum ArtistAlbum
	} `json:"items"`
}

// ArtistTopTracks represents the artist's track information retrieved from the Spotify API.
type ArtistTopTracks struct {
	Tracks []struct {
		Album   ArtistAlbum `json:"album"`
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
		ExternalUrls struct {
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
		Popularity  int    `json:"popularity"`
		PreviewUrl  string `json:"preview_url"`
		TrackNumber int    `json:"track_number"`
		Type        string `json:"type"`
		Uri         string `json:"uri"`
		IsLocal     bool   `json:"is_local"`
	} `json:"tracks"`
}
