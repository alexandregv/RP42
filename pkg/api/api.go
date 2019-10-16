package api

import (
	"encoding/json"
	"fmt"
	"github.com/alexandregv/RP42/pkg/oauth"
	"io/ioutil"
)

const URL = "https://api.intra.42.fr"

// fetch() queries an endpoint of the API.
func fetch(endpoint string) []byte {
	client := oauth.GetClient()

	resp, err := client.Get(fmt.Sprint(URL, endpoint))
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

// GetUser() returns an User, based on his login.
func GetUser(login string) *User {
	resp := fetch(fmt.Sprint("/v2/users/", login))

	user := User{}
	json.Unmarshal(resp, &user)

	return &user
}

// GetUserLastLocation returns the last Location of an user.
func GetUserLastLocation(login string) *Location {
	resp := fetch(fmt.Sprint("/v2/users/", login, "/locations?filter[active]=true"))

	locations := []Location{}
	json.Unmarshal(resp, &locations)

	if len(locations) > 0 {
		return &locations[len(locations)-1]
	} else {
		return nil
	}
}

// GetUserCoalition() returns the Coalition of an user.
func GetUserCoalition(login string) *Coalition {
	resp := fetch(fmt.Sprint("/v2/users/", login, "/coalitions"))

	coalitions := []Coalition{}
	json.Unmarshal(resp, &coalitions)

	if len(coalitions) > 0 {
		return &coalitions[len(coalitions)-1]
	} else {
		return nil
	}
}
