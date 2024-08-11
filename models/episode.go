package models

// GetEpisodeRequest represents the get episode request information.
type GetEpisodeRequest struct {
	Id     string // Required: The Spotify ID for the episode.
	Market string
}

// GetEpisodesRequest represents the get episode's request information.
type GetEpisodesRequest struct {
	Ids    string // Required: A comma-separated list of the Spotify IDs for the episodes. Maximum: 50 IDs.
	Market string
}

// GetSavedEpisodesRequest represents the get saved episode's request information.
type GetSavedEpisodesRequest struct {
	Market string
	Limit  int
	Offset int
}

// SaveEpisodesRequest represents the save episode's request information.
type SaveEpisodesRequest struct {
	Ids  string // Required: A comma-separated list of the Spotify IDs. Maximum: 50 IDs.
	Body struct {
		Ids []string `json:"ids"`
	}
}

// RemoveEpisodesRequest represents the remove episode's request information.
type RemoveEpisodesRequest struct {
	Ids  string // Required: A comma-separated list of the Spotify IDs. Maximum: 50 IDs.
	Body struct {
		Ids []string `json:"ids"`
	}
}

// CheckSavedEpisodesRequest represents the check saved episode's request information.
type CheckSavedEpisodesRequest struct {
	Ids string // Required: A comma-separated list of the Spotify IDs for the episodes. Maximum: 50 IDs.
}

// Episode represents the episode's information retrieved from the Spotify API.
type Episode struct {
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
	Show struct {
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
}

// Episodes represents the episodes information retrieved from the Spotify API.
type Episodes struct {
	Episodes []Episode `json:"episodes"`
}

// SavedEpisodes represents the saved episodes information retrieved from the Spotify API.
type SavedEpisodes struct {
	Href     string `json:"href"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
	Items    []struct {
		AddedAt string  `json:"added_at"`
		Episode Episode `json:"episode"`
	} `json:"items"`
}

// CheckSavedEpisodes represents the check saved episodes information retrieved from the Spotify API.
type CheckSavedEpisodes []bool
