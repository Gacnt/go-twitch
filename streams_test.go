package twitchapi

import "testing"

func TestGetStream(t *testing.T) {
	tapi := NewAPI()

	stream := tapi.GetStream("adasdad")
	
	t.Log(stream)
}
