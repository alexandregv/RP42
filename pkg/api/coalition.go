package api

import (
	"context"
	"encoding/json"
	"fmt"
)

// Coalition represents a coalition from the 42's API.
// Truncated to keep only useful entries.
type Coalition struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	//		ImageURL string `json:"image_url"`
	//		Color    string `json:"color"`
	//		Score    int    `json:"score"`
	//		UserID   int    `json:"user_id"`
}

// CoalitionUser represents an user in coalition from the 42's API.
type CoalitionUser struct {
	ID          int `json:"id"`
	CoalitionID int `json:"coalition_id"`
	// UserID      int       `json:"user_id"`
	// CreatedAt   time.Time `json:"created_at"`
	// UpdatedAt   time.Time `json:"updated_at"`
}

// GetUserCoalition() returns the Coalition of an user.
func GetUserCoalition(ctx context.Context, user *User) (coa *Coalition, err error) {
	resp, err := fetch(ctx, fmt.Sprint("/v2/coalitions_users/", "?user_id=", fmt.Sprint(user.ID), "&sort=-created_at"))
	if err != nil {
		return nil, err
	}

	coalition_users := []CoalitionUser{}
	err = json.Unmarshal(resp, &coalition_users)
	if err != nil {
		return nil, err
	}

	resp, err = fetch(ctx, fmt.Sprint("/v2/users/", user.Login, "/coalitions"))
	if err != nil {
		return nil, err
	}

	coalitions := []Coalition{}
	err = json.Unmarshal(resp, &coalitions)
	if err != nil {
		return nil, err
	}

	if len(coalitions) > 0 {
		for i, n := range coalitions {
			if n.ID == coalition_users[0].CoalitionID {
				return &coalitions[i], nil
			}
		}
	}
	return nil, nil
}
