package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/alexandregv/RP42/pkg/oauth"
)

const URL = "https://api.intra.42.fr"

// fetch() queries an endpoint of the API.
func fetch(ctx context.Context, endpoint string) []byte {
	client := ctx.Value("apiClient").(*oauth.Client)

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
func GetUser(ctx context.Context, login string) *User {
	resp := fetch(ctx, fmt.Sprint("/v2/users/", login))

	user := User{}
	json.Unmarshal(resp, &user)

	return &user
}

// GetUserLastLocation returns the last Location of an user.
func GetUserLastLocation(ctx context.Context, user *User) *Location {
	resp := fetch(ctx, fmt.Sprint("/v2/users/", user.Login, "/locations?filter[active]=true"))

	locations := []Location{}
	json.Unmarshal(resp, &locations)

	if len(locations) > 0 {
		return &locations[len(locations)-1]
	} else {
		return nil
	}
}

// GetUserFirstLocation returns the first Location of an user (in a day).
func GetUserFirstLocation(ctx context.Context, user *User) *Location {
	now := time.Now().UTC()
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC).Format("2006-01-02T15:04:05.000Z")
	resp := fetch(ctx, fmt.Sprint("/v2/users/"+user.Login+"/locations?range[begin_at]="+midnight+","+now.Format("2006-01-02T15:04:05.000Z"+"&sort=begin_at")))

	locations := []Location{}
	json.Unmarshal(resp, &locations)

	if len(locations) > 0 {
		return &locations[0]
	} else {
		return nil
	}
}

// GetUserCoalition() returns the Coalition of an user.
func GetUserCoalition(ctx context.Context, user *User) *Coalition {
	resp := fetch(ctx, fmt.Sprint("/v2/coalitions_users/", "?user_id=", fmt.Sprint(user.ID), "&sort=-created_at"))
	coalition_users := []CoalitionUser{}
	json.Unmarshal(resp, &coalition_users)

	resp = fetch(ctx, fmt.Sprint("/v2/users/", user.Login, "/coalitions"))
	coalitions := []Coalition{}
	json.Unmarshal(resp, &coalitions)

	if len(coalitions) > 0 {
		for i, n := range coalitions {
			if n.ID == coalition_users[0].CoalitionID {
				return &coalitions[i]
			}
		}
	}
	return nil
}
