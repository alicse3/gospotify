package models

// GetPlaybackStateRequest represents the get playback state request information.
type GetPlaybackStateRequest struct {
	Market          string
	AdditionalTypes string
}

// TransferPlaybackRequestBody represents the transfer playback request's body information.
type TransferPlaybackRequestBody struct {
	DeviceIds []string `json:"device_ids"` // Required: A JSON array containing the ID of the device on which playback should be started/transferred.
	Play      bool     `json:"play"`
}

// TransferPlaybackRequest represents the transfer playback request information.
type TransferPlaybackRequest struct {
	Body TransferPlaybackRequestBody
}

// GetCurrentlyPlayingTrackRequest represents the currently playing track request information.
type GetCurrentlyPlayingTrackRequest struct {
	Market          string
	AdditionalTypes string
}

// StartOrResumePlaybackRequestOffset represents the start or resume playback request's offset information.
type StartOrResumePlaybackRequestOffset struct {
	Position int `json:"position"`
}

// StartOrResumePlaybackRequestBody represents the start or resume playback request's body information.
type StartOrResumePlaybackRequestBody struct {
	ContextUri string                             `json:"context_uri"`
	Uris       []string                           `json:"uris"`
	Offset     StartOrResumePlaybackRequestOffset `json:"offset"`
	PositionMs int                                `json:"position_ms"`
}

// StartOrResumePlaybackRequest represents the start or resume playback request information.
type StartOrResumePlaybackRequest struct {
	DeviceId string
	Body     StartOrResumePlaybackRequestBody
}

// PausePlaybackRequest represents the pause playback request information.
type PausePlaybackRequest struct {
	DeviceId string
}

// SkipToNextRequest represents the skip to next request information.
type SkipToNextRequest struct {
	DeviceId string
}

// SkipToPreviousRequest represents the skip to previous request information.
type SkipToPreviousRequest struct {
	DeviceId string
}

// SeekToPositionRequest represents the seek to position request information.
type SeekToPositionRequest struct {
	PositionMs int // Required: The position in milliseconds to seek to. Must be a positive number. Passing in a position that is greater than the length of the track will cause the player to start playing the next song.
	DeviceId   string
}

// SetRepeatModeRequest represents the set repeat mode request information.
type SetRepeatModeRequest struct {
	// Required: track, context or off.
	// track will repeat the current track.
	// context will repeat the current context.
	// off will turn repeat off.
	State    string
	DeviceId string
}

// SetPlaybackVolumeRequest represents the set playback volume request information.
type SetPlaybackVolumeRequest struct {
	VolumePercent int // Required: The volume to set. Must be a value from 0 to 100 inclusive.
	DeviceId      string
}

// TogglePlaybackShuffleRequest represents the toggle playback shuffle request information.
type TogglePlaybackShuffleRequest struct {
	// Required: true or false.
	// true : Shuffle user's playback.
	// false : Do not shuffle user's playback.
	State    bool
	DeviceId string
}

// GetRecentlyPlayedTracksRequest represents the recently played tracks request information.
type GetRecentlyPlayedTracksRequest struct {
	Limit  int
	After  int
	Before int
}

// AddItemToPlaybackQueueRequest represents the add item to playback queue request information.
type AddItemToPlaybackQueueRequest struct {
	Uri      string // Required: The uri of the item to add to the queue. Must be a track or an episode uri.
	DeviceId string
}

// Device represents the device information retrieved from the Spotify API.
type Device struct {
	Id               string `json:"id"`
	IsActive         bool   `json:"is_active"`
	IsPrivateSession bool   `json:"is_private_session"`
	IsRestricted     bool   `json:"is_restricted"`
	Name             string `json:"name"`
	Type             string `json:"type"`
	VolumePercent    int    `json:"volume_percent"`
	SupportsVolume   bool   `json:"supports_volume"`
}

// PlaybackState represents the playback state information retrieved from the Spotify API.
type PlaybackState struct {
	Device       Device `json:"device"`
	RepeatState  string `json:"repeat_state"`
	ShuffleState bool   `json:"shuffle_state"`
	Context      struct {
		Type         string `json:"type"`
		Href         string `json:"href"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Uri string `json:"uri"`
	} `json:"context"`
	Timestamp  int  `json:"timestamp"`
	ProgressMs int  `json:"progress_ms"`
	IsPlaying  bool `json:"is_playing"`
	Item       struct {
		Track struct {
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
		} `json:"track"`
		Episode struct {
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
			IsExternallyHosted   bool     `json:"is_externally_hosted"`
			IsPlayable           bool     `json:"is_playable"`
			Language             string   `json:"language"`
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
		} `json:"episode"`
	} `json:"item"`
	CurrentlyPlayingType string `json:"currently_playing_type"`
	Actions              struct {
		InterruptingPlayback  bool `json:"interrupting_playback"`
		Pausing               bool `json:"pausing"`
		Resuming              bool `json:"resuming"`
		Seeking               bool `json:"seeking"`
		SkippingNext          bool `json:"skipping_next"`
		SkippingPrev          bool `json:"skipping_prev"`
		TogglingRepeatContext bool `json:"toggling_repeat_context"`
		TogglingShuffle       bool `json:"toggling_shuffle"`
		ToggleRepeatTrack     bool `json:"toggling_repeat_track"`
		TransferringPlayback  bool `json:"transferring_playback"`
	} `json:"actions"`
}

// AvailableDevices represents the available devices information retrieved from the Spotify API.
type AvailableDevices struct {
	Devices []Device `json:"devices"`
}

// RecentlyPlayedTracks represents the recently played tracks information retrieved from the Spotify API.
type RecentlyPlayedTracks struct {
	Href    string `json:"href"`
	Limit   int    `json:"limit"`
	Next    string `json:"next"`
	Cursors struct {
		After  string `json:"after"`
		Before string `json:"before"`
	} `json:"cursors"`
	Total int `json:"total"`
	Items []struct {
		Track struct {
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
		} `json:"track"`
		PlayedAt string `json:"played_at"`
		Context  struct {
			Type         string `json:"type"`
			Href         string `json:"href"`
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Uri string `json:"uri"`
		} `json:"context"`
	} `json:"items"`
}

// UsersQueue represents the users queue information retrieved from the Spotify API.
type UsersQueue struct {
	CurrentlyPlaying struct {
		Track struct {
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
		} `json:"track"`
		Episode struct {
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
			IsExternallyHosted   bool     `json:"is_externally_hosted"`
			IsPlayable           bool     `json:"is_playable"`
			Language             string   `json:"language"`
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
		} `json:"episode"`
	} `json:"currently_playing"`
	Queue []struct {
		Track struct {
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
		} `json:"track"`
		Episode struct {
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
			IsExternallyHosted   bool     `json:"is_externally_hosted"`
			IsPlayable           bool     `json:"is_playable"`
			Language             string   `json:"language"`
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
		} `json:"episode"`
	} `json:"queue"`
}
