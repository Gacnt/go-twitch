package twitchapi

import "time"

// StreamResponse is the response from the twitch API
type StreamResponse struct {
	Stream Stream `json:"stream,omitempty"`
	Links  Links  `json:"links"`
}

// Stream is the layout for the Stream response from the
// twitch api
type Stream struct {
	Game        string    `json:"game"`
	Viewers     uint      `json:"viewers"`
	AverageFPS  float64   `json:"average_fps"`
	Delay       int       `json:"delay"`
	VideoHeight int       `json:"video_height"`
	IsPlaylist  int       `json:"is_playlist"`
	CreatedAt   time.Time `json:"created_at"`
	ID          uint      `json:"_id"`
	Channel     Channel   `json:"channel"`
	Preview     Preview   `json:"preview"`
	Links       Links     `json:"_links"`
}

// Channel is the layout for the channel response from the
// twitch api
type Channel struct {
	Mature                       bool      `json:"mature"`
	Status                       string    `json:"status"`
	BroadcasterLanguage          string    `json:"broadcaster_language"`
	DisplayName                  string    `json:"display_name"`
	Game                         string    `json:"game"`
	Delay                        int       `json:"delay"`
	Language                     string    `json:"language"`
	ID                           uint      `json:"_id"`
	Name                         string    `json:"name"`
	CreatedAt                    time.Time `json:"created_at"`
	UpdatedAt                    time.Time `json:"updated_at"`
	Logo                         string    `json:"logo"`
	Banner                       string    `json:"banner"`
	VideoBanner                  string    `json:"video_banner"`
	Background                   string    `json:"background"`
	ProfileBanner                string    `json:"profile_banner"`
	ProfileBannerBackgroundColor string    `json:"profile_banner_background_color"`
	Partner                      bool      `json:"partner"`
	URL                          string    `json:"url"`
	Views                        uint      `json:"views"`
	Followers                    uint      `json:"followers"`
	Links                        Links     `json:"links"`
}

// Preview is the stream image preview picture returned by
// the twitch api
type Preview struct {
	Small    string `json:"small"`
	Medium   string `json:"medium"`
	Large    string `json:"large"`
	Template string `json:"template"`
}

// Links is the generic struct for the type of links a response
// can send back, not all will be filled and whatever isn't will be omitted as empty
type Links struct {
	Channel       string `json:"channel,omityempty"`
	Self          string `json:"self,omitempty"`
	Follows       string `json:"follows,omitempty"`
	Commercial    string `json:"commercial,omitempty"`
	StreamKey     string `json:"stream_key,omitempty"`
	Chat          string `json:"chat,omitempty"`
	Features      string `json:"features,omitempty"`
	Subscriptions string `json:"subscriptions,omitempty"`
	Editors       string `json:"editors,omitempty"`
	Teams         string `json:"teams,omitempty"`
	Videos        string `json:"videos,omitempty"`
}
