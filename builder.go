package twitchapi

import (
	"io/ioutil"
	"log"
	"net/http"
)

const apiurl string = "https://api.twitch.tv/kraken"

// TwitchAPI will return an api struct
type TwitchAPI struct{}

// NewAPI Will return a new api struct for use
func NewAPI() *TwitchAPI {
	return &TwitchAPI{}
}

func fetchAPI(url string) []byte {
	client := http.Client{}
	req, err := http.NewRequest("GET", apiurl+url, nil)
	if err != nil {
		log.Println(err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	res.Body.Close()

	return body
}
