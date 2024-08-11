package models

// GetChapterRequest represents the get chapter request information.
type GetChapterRequest struct {
	Id     string // Required: The Spotify ID for the chapter.
	Market string
}

// GetChaptersRequest represents the get chapters request information.
type GetChaptersRequest struct {
	Ids    string // Required: A comma-separated list of the Spotify IDs. Maximum: 50 IDs.
	Market string
}

// Chapter represents the chapter information retrieved from the Spotify API.
type Chapter struct {
	AudioPreviewUrl  string   `json:"audio_preview_url"`
	AvailableMarkets []string `json:"available_markets"`
	ChapterNumber    int      `json:"chapter_number"`
	Description      string   `json:"description"`
	HtmlDescription  string   `json:"html_description"`
	DurationMs       int      `json:"duration_ms"`
	Explicit         bool     `json:"explicit"`
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
	IsPlayable           bool `json:"is_playable"`
	Languages            []string
	Name                 string `json:"name"`
	ReleaseDate          string `json:"release_date"`
	ReleaseDatePrecision string `json:"release_date_precision"`
	ResumePoint          struct {
		FullyPlayed      bool `json:"fully_played"`
		ResumePositionMs int  `json:"resume_position_ms"`
	} `json:"resume_point"`
	Type         string `json:"type"`
	Uri          string `json:"uri"`
	Restrictions struct {
		Reason string `json:"reason"`
	} `json:"restrictions"`
	Audiobook struct {
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
	} `json:"audiobook"`
}

// Chapters represents the chapters information retrieved from the Spotify API.
type Chapters struct {
	Chapters []Chapter `json:"chapters"`
}
