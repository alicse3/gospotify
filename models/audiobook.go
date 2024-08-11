package models

// GetAudiobookRequest represents the get audio book's request information.
type GetAudiobookRequest struct {
	Id string // Required: The Spotify ID for the audiobook.
}

// GetAudiobooksRequest represents the get audio books request information.
type GetAudiobooksRequest struct {
	Ids    string // Required: A comma-separated list of the Spotify IDs. Maximum: 50 IDs.
	Market string
}

// GetAudiobookChaptersRequest represents the get audio book chapters request information.
type GetAudiobookChaptersRequest struct {
	Id     string // Required: The Spotify ID for the audiobook.
	Market string
	Limit  int
	Offset int
}

// GetSavedAudiobooksRequest represents the get saved audio books request information.
type GetSavedAudiobooksRequest struct {
	Limit  int
	Offset int
}

// SaveAudiobooksRequest represents the save audio books request information.
type SaveAudiobooksRequest struct {
	Ids string // Required: A comma-separated list of the Spotify IDs. Maximum: 50 IDs.
}

// RemoveAudiobooksRequest represents the remove audio books request information.
type RemoveAudiobooksRequest struct {
	Ids string // Required: A comma-separated list of the Spotify IDs. Maximum: 50 IDs.
}

// CheckSavedAudiobooksRequest represents the remove audio books request information.
type CheckSavedAudiobooksRequest struct {
	Ids string // Required: A comma-separated list of the Spotify IDs. Maximum: 50 IDs.
}

// Audiobook represents the audiobook's information retrieved from the Spotify API.
type Audiobook struct {
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
	Chapters      struct {
		Href     string `json:"href"`
		Limit    int    `json:"limit"`
		Next     string `json:"next"`
		Offset   int    `json:"offset"`
		Previous string `json:"previous"`
		Total    int    `json:"total"`
		Items    []struct {
			AudioPreviewUrl  any      `json:"audio_preview_url"`
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
			IsPlayable           bool     `json:"is_playable"`
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
	} `json:"chapters"`
}

// Audiobooks represents the audiobooks information retrieved from the Spotify API.
type Audiobooks struct {
	Audiobooks []Audiobook `json:"audiobooks"`
}

// AudiobookChapters represents the audiobook chapters information retrieved from the Spotify API.
type AudiobookChapters struct {
	Href     string `json:"href"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
	Items    []struct {
		AudioPreviewUrl  any      `json:"audio_preview_url"`
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
		IsPlayable           bool     `json:"is_playable"`
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

// SavedAudiobooks represents the saved audiobook's information retrieved from the Spotify API.
type SavedAudiobooks struct {
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
}

// CheckSavedAudiobooks represents the check saved audiobooks information retrieved from the Spotify API.
type CheckSavedAudiobooks []bool
