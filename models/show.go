package models

// GetShowRequest represents the get show's request information.
type GetShowRequest struct {
	Id     string // Required: The Spotify ID for the show.
	Market string
}

// GetShowsRequest represents the get shows request information.
type GetShowsRequest struct {
	Ids    string // Required: A comma-separated list of the Spotify IDs for the shows. Maximum: 50 IDs.
	Market string
}

// GetShowEpisodesRequest represents the get shows request information.
type GetShowEpisodesRequest struct {
	Id     string // Required: The Spotify ID for the show.
	Market string
	Limit  int
	Offset int
}

// GetSavedShowsRequest represents the get saved shows request information.
type GetSavedShowsRequest struct {
	Limit  int
	Offset int
}

// SaveShowsRequest represents the save shows request information.
type SaveShowsRequest struct {
	Ids string // Required: A comma-separated list of the Spotify IDs for the shows. Maximum: 50 IDs.
}

// RemoveShowsRequest represents the remove shows request information.
type RemoveShowsRequest struct {
	Ids    string // Required: A comma-separated list of the Spotify IDs for the shows. Maximum: 50 IDs.
	Market string
}

// CheckSavedShowsRequest represents the check saved shows request information.
type CheckSavedShowsRequest struct {
	Ids string // Required: A comma-separated list of the Spotify IDs for the shows. Maximum: 50 IDs.
}

// Show represents the show's information retrieved from the Spotify API.
type Show struct {
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
	IsExternallyHosted bool         `json:"is_externally_hosted"`
	Languages          []string     `json:"languages"`
	MediaType          string       `json:"media_type"`
	Name               string       `json:"name"`
	Publisher          string       `json:"publisher"`
	Type               string       `json:"type"`
	Uri                string       `json:"uri"`
	TotalEpisodes      int          `json:"total_episodes"`
	Episodes           ShowEpisodes `json:"episodes"`
}

// Shows represents the shows information retrieved from the Spotify API.
type Shows struct {
	Shows struct {
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
	} `json:"shows"`
}

// ShowEpisodes represents the show episodes information retrieved from the Spotify API.
type ShowEpisodes struct {
	Href     string `json:"href"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
	Items    []struct {
		AudioPreviewUrl string `json:"audio_preview_url"`
		Description     string `json:"description"`
		HtmlDescription string `json:"html_description"`
		DurationMs      int    `json:"duration_ms"`
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
		IsExternallyHosted   bool `json:"is_externally_hosted"`
		IsPlayable           bool `json:"is_playable"`
		Language             string
		Languages            []string `json:"languages"`
		Name                 string   `json:"name"`
		ReleaseDate          string   `json:"release_date"`
		ReleaseDatePrecision string   `json:"release_date_precision"`
		ResumePoint          struct {
			FullyPlayed      bool `json:"fully_played"`
			ResumePositionMs int  `json:"resume_position_ms"`
		} `json:"resume_point"`
		Type         string `json:"type"`
		Uri          string `json:"uri"`
		Restrictions struct {
			Reason string `json:"reason"`
		} `json:"restrictions"`
	} `json:"items"`
}

// SavedShows represents the saved shows information retrieved from the Spotify API.
type SavedShows struct {
	Href     string `json:"href"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
	Items    []struct {
		AddedAt string `json:"added_at"`
		Show    struct {
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
	} `json:"items"`
}

// CheckSavedShows represents the check saved shows information retrieved from the Spotify API.
type CheckSavedShows []bool
