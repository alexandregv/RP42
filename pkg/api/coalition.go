package api

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
