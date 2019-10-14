package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	url = "https://api.intra.42.fr"
)

func fetch(endpoint string) []byte {
	client := GetClient()

	resp, err := client.Get(fmt.Sprint(url, endpoint))
	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return body
	} else {
		panic(fmt.Sprintf("The API responded with a bad status code (%d): %s", resp.StatusCode, string(body)))
	}
}
