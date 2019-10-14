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

func GetUser(login string) *User {
	resp := fetch(fmt.Sprint("/v2/users/", login))

	user := User{}
	json.Unmarshal(resp, &user)
	return &user
}

func GetUserLastLocation(login string) *Location {
	resp := fetch(fmt.Sprint("/v2/users/", login, "/locations"/*?filter[active]=true"*/))

	locations := []Location{}
	json.Unmarshal(resp, &locations)
	return &locs[0]
}
