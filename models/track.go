package models

// GetTrackRequest represents the get track's request information.
type GetTrackRequest struct {
	Id     string // Required: The Spotify ID for the track.
	Market string
}

// GetTracksRequest represents the get tracks request information.
type GetTracksRequest struct {
	Ids    string // Required: A comma-separated list of the Spotify IDs. For example: ids=4iV5W9uYEdYUVa79Axb7Rh,1301WleyT98MSxVHPZCA6M. Maximum: 50 IDs.
	Market string
}

// GetSavedTracksRequest represents the saved tracks request information.
type GetSavedTracksRequest struct {
	Market string
	Limit  int
	Offset int
}

// SaveTracksBody represents the save tracks body information.
type SaveTracksBody struct {
	Ids string `json:"ids"`
}

// SaveTracksRequest represents the save tracks request information.
type SaveTracksRequest struct {
	Ids  string // Required: A comma-separated list of the Spotify IDs. For example: ids=4iV5W9uYEdYUVa79Axb7Rh,1301WleyT98MSxVHPZCA6M. Maximum: 50 IDs.
	Body SaveTracksBody
}

// RemoveTracksBody represents the remove tracks body information.
type RemoveTracksBody struct {
	Ids string `json:"ids"`
}

// RemoveTracksRequest represents the remove tracks request information.
type RemoveTracksRequest struct {
	Ids  string // Required: A comma-separated list of the Spotify IDs. For example: ids=4iV5W9uYEdYUVa79Axb7Rh,1301WleyT98MSxVHPZCA6M. Maximum: 50 IDs.
	Body RemoveTracksBody
}

// CheckSavedTracksRequest represents the check saved tracks request information.
type CheckSavedTracksRequest struct {
	Ids string // Required: A comma-separated list of the Spotify IDs. For example: ids=4iV5W9uYEdYUVa79Axb7Rh,1301WleyT98MSxVHPZCA6M. Maximum: 50 IDs.
}

// GetSeveralTracksAudioFeaturesRequest represents the several tracks audio features request information.
type GetSeveralTracksAudioFeaturesRequest struct {
	Ids string // Required: A comma-separated list of the Spotify IDs for the tracks. Maximum: 100 IDs.
}

// GetTracksAudioFeaturesRequest represents the tracks audio features request information.
type GetTracksAudioFeaturesRequest struct {
	Id string // Required: The Spotify ID for the track.
}

// GetTracksAudioAnalysisRequest represents the tracks audio analysis request information.
type GetTracksAudioAnalysisRequest struct {
	Id string // Required: The Spotify ID for the track.
}

// GetRecommendationsRequest represents the recommendations request information.
type GetRecommendationsRequest struct {
	Limit                  int
	Market                 string
	SeedArtists            string // Required: A comma separated list of Spotify IDs for seed artists. Up to 5 seed values may be provided in any combination of seed_artists, seed_tracks and seed_genres.
	SeedGenres             string // Required: A comma separated list of any genres in the set of available genre seeds. Up to 5 seed values may be provided in any combination of seed_artists, seed_tracks and seed_genres.
	SeedTracks             string // Required: A comma separated list of Spotify IDs for a seed track. Up to 5 seed values may be provided in any combination of seed_artists, seed_tracks and seed_genres.
	MinAcousticness        float64
	MaxAcousticness        float64
	TargetAcousticness     float64
	MinDanceability        float64
	MaxDanceability        float64
	TargetDanceability     float64
	MinDurationMs          int
	MaxDurationMs          int
	TargetDurationMs       int
	MinEnergy              float64
	MaxEnergy              float64
	TargetEnergy           float64
	MinInstrumentalness    float64
	MaxInstrumentalness    float64
	TargetInstrumentalness float64
	MinKey                 int
	MaxKey                 int
	TargetKey              int
	MinLiveness            float64
	MaxLiveness            float64
	TargetLiveness         float64
	MinLoudness            float64
	MaxLoudness            float64
	TargetLoudness         float64
	MinMode                int
	MaxMode                int
	TargetMode             int
	MinPopularity          int
	MaxPopularity          int
	TargetPopularity       int
	MinSpeechiness         float64
	MaxSpeechiness         float64
	TargetSpeechiness      float64
	MinTempo               float64
	MaxTempo               float64
	TargetTempo            float64
	MinTimeSignature       int
	MaxTimeSignature       int
	TargetTimeSignature    int
	MinValence             float64
	MaxValence             float64
	TargetValence          float64
}

// Track represents the track's information retrieved from the Spotify API.
type Track struct {
	Album struct {
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
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Href       string `json:"href"`
	Id         string `json:"id"`
	IsPlayable bool   `json:"is_playable"`
	LinkedFrom struct {
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
}

// Tracks represents the tracks information retrieved from the Spotify API.
type Tracks struct {
	Tracks []Track `json:"tracks"`
}

// SavedTracks represents the saved tracks information retrieved from the Spotify API.
type SavedTracks struct {
	Href     string `json:"href"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
	Items    []struct {
		AddedAt string `json:"added_at"`
		Track   Track  `json:"track"`
	} `json:"items"`
}

// CheckSavedTracks represents the check saved tracks information retrieved from the Spotify API.
type CheckSavedTracks []bool

// SeveralTracksAudioFeatures represents the several tracks audio features information retrieved from the Spotify API.
type SeveralTracksAudioFeatures struct {
	AudioFeatures TracksAudioFeatures `json:"audio_features"`
}

// TracksAudioFeatures represents the tracks audio features information retrieved from the Spotify API.
type TracksAudioFeatures struct {
	Acousticness     float64 `json:"acousticness"`
	AnalysisUrl      string  `json:"analysis_url"`
	Danceability     float64 `json:"danceability"`
	DurationMs       int     `json:"duration_ms"`
	Energy           float64 `json:"energy"`
	Id               string  `json:"id"`
	Instrumentalness float64 `json:"instrumentalness"`
	Key              int     `json:"key"`
	Liveness         float64 `json:"liveness"`
	Loudness         float64 `json:"loudness"`
	Mode             int     `json:"mode"`
	Speechiness      float64 `json:"speechiness"`
	Tempo            float64 `json:"tempo"`
	TimeSignature    int     `json:"time_signature"`
	TrackHref        string  `json:"track_href"`
	Type             string  `json:"type"`
	Uri              string  `json:"uri"`
	Valence          float64 `json:"valence"`
}

// TracksAudioAnalysis represents the tracks audio analysis information retrieved from the Spotify API.
type TracksAudioAnalysis struct {
	Meta struct {
		AnalyzerVersion string  `json:"analyzer_version"`
		Platform        string  `json:"platform"`
		DetailedStatus  string  `json:"detailed_status"`
		StatusCode      int     `json:"status_code"`
		Timestamp       int     `json:"timestamp"`
		AnalysisTime    float64 `json:"analysis_time"`
		InputProcess    string  `json:"input_process"`
	} `json:"meta"`
	Track struct {
		NumSamples              int     `json:"num_samples"`
		Duration                float64 `json:"duration"`
		SampleMd5               string  `json:"sample_md5"`
		OffsetSeconds           int     `json:"offset_seconds"`
		WindowSeconds           int     `json:"window_seconds"`
		AnalysisSampleRate      int     `json:"analysis_sample_rate"`
		AnalysisChannels        int     `json:"analysis_channels"`
		EndOfFadeIn             float64 `json:"end_of_fade_in"`
		StartOfFadeOut          float64 `json:"start_of_fade_out"`
		Loudness                float64 `json:"loudness"`
		Tempo                   float64 `json:"tempo"`
		TempoConfidence         float64 `json:"tempo_confidence"`
		TimeSignature           int     `json:"time_signature"`
		TimeSignatureConfidence float64 `json:"time_signature_confidence"`
		Key                     int     `json:"key"`
		KeyConfidence           float64 `json:"key_confidence"`
		Mode                    int     `json:"mode"`
		ModeConfidence          float64 `json:"mode_confidence"`
		Codestring              string  `json:"codestring"`
		CodeVersion             float64 `json:"code_version"`
		Echoprintstring         string  `json:"echoprintstring"`
		EchoprintVersion        float64 `json:"echoprint_version"`
		Synchstring             string  `json:"synchstring"`
		SynchVersion            float64 `json:"synch_version"`
		Rhythmstring            string  `json:"rhythmstring"`
		RhythmVersion           float64 `json:"rhythm_version"`
	} `json:"track"`
	Bars []struct {
		Start      float64 `json:"start"`
		Duration   float64 `json:"duration"`
		Confidence float64 `json:"confidence"`
	} `json:"bars"`
	Beats []struct {
		Start      float64 `json:"start"`
		Duration   float64 `json:"duration"`
		Confidence float64 `json:"confidence"`
	} `json:"beats"`
	Sections []struct {
		Start                   float64 `json:"start"`
		Duration                float64 `json:"duration"`
		Confidence              float64 `json:"confidence"`
		Loudness                float64 `json:"loudness"`
		Tempo                   float64 `json:"tempo"`
		TempoConfidence         float64 `json:"tempo_confidence"`
		Key                     int     `json:"key"`
		KeyConfidence           float64 `json:"key_confidence"`
		Mode                    int     `json:"mode"`
		ModeConfidence          float64 `json:"mode_confidence"`
		TimeSignature           int     `json:"time_signature"`
		TimeSignatureConfidence float64 `json:"time_signature_confidence"`
	} `json:"sections"`
	Segments []struct {
		Start           float64   `json:"start"`
		Duration        float64   `json:"duration"`
		Confidence      float64   `json:"confidence"`
		LoudnessStart   float64   `json:"loudness_start"`
		LoudnessMax     float64   `json:"loudness_max"`
		LoudnessMaxTime float64   `json:"loudness_max_time"`
		LoudnessEnd     float64   `json:"loudness_end"`
		Pitches         []float64 `json:"pitches"`
		Timbre          []float64 `json:"timbre"`
	} `json:"segments"`
	Tatums []struct {
		Start      float64 `json:"start"`
		Duration   float64 `json:"duration"`
		Confidence float64 `json:"confidence"`
	} `json:"tatums"`
}

// GetRecommendations represents the recommendations information retrieved from the Spotify API.
type GetRecommendations struct {
	Seeds []struct {
		AfterFilteringSize int    `json:"after_filtering_size"`
		AfterRelinkingSize int    `json:"after_relinking_size"`
		Href               string `json:"href"`
		Id                 string `json:"id"`
		InitialPoolSize    int    `json:"initial_pool_size"`
		Type               string `json:"type"`
	} `json:"seeds"`
	Tracks []Track `json:"tracks"`
}
