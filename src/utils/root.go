package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// MakeRequest function to make request and return a json result.
func MakeRequest(url string) interface{} {
	
	resp, err := http.Get(url)
	
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		
		var data interface{}

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal((err))
		}
	
		json.Unmarshal(body, &data)
	
		return data
	}

	log.Println("Could not connect to server. Sleeping 10s.")
	time.Sleep(10 * time.Second)
	return MakeRequest(url)
}
