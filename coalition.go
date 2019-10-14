package main

type Coalition struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	ImageURL string `json:"image_url"`
	Color    string `json:"color"`
	Score    int    `json:"score"`
	UserID   int    `json:"user_id"`
}
