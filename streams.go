package twitchapi

import (
	"encoding/json"
	"log"
)

// GetStream will retrieve the stream information "/streams/test_channel"
func (t *TwitchAPI) GetStream(streamName string) *StreamResponse {
	body := fetchAPI("/streams/" + streamName)
	resp := &StreamResponse{}
	err := json.Unmarshal(body, &resp)
	if err != nil {	
		log.Println(err)
	}
	return resp
}
