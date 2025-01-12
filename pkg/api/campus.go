package api

import (
	"context"
	"encoding/json"
	"fmt"
)

// Campus represents a campus from the 42's API.
// Truncated to keep only useful entries.
type Campus struct {
	// ID       int    `json:"id"`
	Name string `json:"name"`
	// TimeZone string `json:"time_zone"`
	// Language struct {
	// 	ID         int    `json:"id"`
	// 	Name       string `json:"name"`
	// 	Identifier string `json:"identifier"`
	// } `json:"language"`
	// UsersCount         int    `json:"users_count"`
	// VogsphereID        int    `json:"vogsphere_id"`
	// Country            string `json:"country"`
	// Address            string `json:"address"`
	// Zip                string `json:"zip"`
	// City               string `json:"city"`
	// Website            string `json:"website"`
	// Facebook           string `json:"facebook"`
	// Twitter            string `json:"twitter"`
	// Active             bool   `json:"active"`
	// Public             bool   `json:"public"`
	// EmailExtension     string `json:"email_extension"`
	// DefaultHiddenPhone bool   `json:"default_hidden_phone"`
	// Endpoint           struct {
	// 	ID          int       `json:"id"`
	// 	URL         string    `json:"url"`
	// 	Description string    `json:"description"`
	// 	CreatedAt   time.Time `json:"created_at"`
	// 	UpdatedAt   time.Time `json:"updated_at"`
	// } `json:"endpoint"`
}

// GetCampus() returns a Campus, based on its id.
func GetCampus(ctx context.Context, id int) *Campus {
	resp := fetch(ctx, fmt.Sprint("/v2/campus/", id))

	campus := Campus{}
	json.Unmarshal(resp, &campus)

	return &campus
}
