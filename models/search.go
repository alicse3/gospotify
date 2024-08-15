package models

// SearchRequest represents the search for item request information.
type SearchRequest struct {
	// Required: Your search query.
	// You can narrow down your search using field filters.
	// The available filters are album, artist, track, year, upc, tag:hipster, tag:new, isrc, and genre.
	// Each field filter only applies to certain result types.
	// The artist and year filters can be used while searching albums, artists and tracks.
	// You can filter on a single year or a range (e.g. 1955-1960).
	// The album filter can be used while searching albums and tracks.
	// The genre filter can be used while searching artists and tracks.
	// The isrc and track filters can be used while searching tracks.
	// The upc, tag:new and tag:hipster filters can only be used while searching albums.
	// The tag:new filter will return albums released in the past two weeks and tag:hipster can be used to return only albums with the lowest 10% popularity.
	// Example: q=remaster%2520track%3ADoxy%2520artist%3AMiles%2520Davis
	Q string

	// Required: A comma-separated list of item types to search across.
	// Search results include hits from all the specified item types.
	// For example: q=abacab&type=album,track returns both albums and tracks matching "abacab".
	// Allowed values: "album", "artist", "playlist", "track", "show", "episode", "audiobook"
	Type string

	Market          string
	Limit           int
	Offset          int
	IncludeExternal string
}

// SearchResponse represents the search's information retrieved from the Spotify API.
type SearchResponse struct {
	Tracks  AlbumTracks `json:"tracks"`
	Artists struct {
		Href     string `json:"href"`
		Limit    int    `json:"limit"`
		Next     string `json:"next"`
		Offset   int    `json:"offset"`
		Previous string `json:"previous"`
		Total    int    `json:"total"`
		Items    []struct {
			Artist []Artist
		} `json:"items"`
	} `json:"artists"`
	Albums struct {
		Href     string `json:"href"`
		Limit    int    `json:"limit"`
		Next     string `json:"next"`
		Offset   int    `json:"offset"`
		Previous string `json:"previous"`
		Total    int    `json:"total"`
		Items    []struct {
			Album []Album
		} `json:"items"`
	} `json:"albums"`
	Playlists Playlists `json:"playlists"`
	Shows     struct {
		Href     string `json:"href"`
		Limit    int    `json:"limit"`
		Next     string `json:"next"`
		Offset   int    `json:"offset"`
		Previous string `json:"previous"`
		Total    int    `json:"total"`
		Items    []struct {
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
		} `json:"items"`
	} `json:"shows"`
	Episodes struct {
		Href     string `json:"href"`
		Limit    int    `json:"limit"`
		Next     string `json:"next"`
		Offset   int    `json:"offset"`
		Previous string `json:"previous"`
		Total    int    `json:"total"`
		Items    []struct {
			Episode []Episode
		} `json:"items"`
	} `json:"episodes"`
	Audiobooks struct {
		Href     string `json:"href"`
		Limit    int    `json:"limit"`
		Next     string `json:"next"`
		Offset   int    `json:"offset"`
		Previous string `json:"previous"`
		Total    int    `json:"total"`
		Items    []struct {
			Authours []struct {
				Name string `json:"name"`
			} `json:"authors"`
			AvailableMarkets []string `json:"available_markets"`
			Copyrights       []struct {
				Text string `json:"text"`
				Type string `json:"type"`
			} `json:"copyrights"`
			Description     string `json:"description"`
			HtmlDescription string `json:"html_description"`
			Edition         string `json:"edition"`
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
			Languages []string `json:"languages"`
			MediaType string   `json:"media_type"`
			Name      string   `json:"name"`
			Narrators []struct {
				Name string `json:"name"`
			} `json:"narrators"`
			Publisher     string `json:"publisher"`
			Type          string `json:"type"`
			Uri           string `json:"uri"`
			TotalChapters int    `json:"total_chapters"`
		} `json:"items"`
	} `json:"audiobooks"`
}
