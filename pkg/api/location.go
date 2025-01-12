package api

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// Location represents a connection location from the 42's API.
// Truncated to keep only useful entries.
type Location struct {
	//	ID       int         `json:"id"`
	BeginAt time.Time `json:"begin_at"`
	//	EndAt    time.Time   `json:"end_at"`
	//	Primary  bool        `json:"primary"`
	//	Floor    interface{} `json:"floor"`
	//	Row      interface{} `json:"row"`
	//	Post     interface{} `json:"post"`
	Host     string `json:"host"`
	CampusID int    `json:"campus_id"`
	//	User     User        `json:"user"`
}

// GetUserFirstLocation returns the first Location of an user (in a day).
func GetUserFirstLocation(ctx context.Context, user *User) (loc *Location, err error) {
	now := time.Now().UTC()
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC).Format("2006-01-02T15:04:05.000Z")
	resp, err := fetch(ctx, fmt.Sprint("/v2/users/"+user.Login+"/locations?range[begin_at]="+midnight+","+now.Format("2006-01-02T15:04:05.000Z"+"&sort=begin_at")))
	if err != nil {
		return nil, err
	}

	locations := []Location{}
	err = json.Unmarshal(resp, &locations)
	if err != nil {
		return nil, err
	}

	if len(locations) > 0 {
		return &locations[0], nil
	} else {
		return nil, nil
	}
}

// GetUserLastLocation returns the last Location of an user.
func GetUserLastLocation(ctx context.Context, user *User) (loc *Location, err error) {
	resp, err := fetch(ctx, fmt.Sprint("/v2/users/", user.Login, "/locations?filter[active]=true"))
	if err != nil {
		return nil, err
	}

	locations := []Location{}
	err = json.Unmarshal(resp, &locations)
	if err != nil {
		return nil, err
	}

	if len(locations) > 0 {
		return &locations[len(locations)-1], nil
	} else {
		return nil, nil
	}
}
