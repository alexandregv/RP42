package main

import (
	"time"
)

type User struct {

	ID    int    `json:"id"`
	Login string `json:"login"`
	URL   string `json:"url"`
}

type Coalition struct {

	ID       int    `json:"id"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	ImageURL string `json:"image_url"`
	Color    string `json:"color"`
	Score    int    `json:"score"`
	UserID   int    `json:"user_id"`
}

type Location struct {
	ID       int         `json:"id"`
	BeginAt  time.Time   `json:"begin_at"`
	EndAt    time.Time   `json:"end_at"`
	Primary  bool        `json:"primary"`
	Floor    interface{} `json:"floor"`
	Row      interface{} `json:"row"`
	Post     interface{} `json:"post"`
	Host     string      `json:"host"`
	CampusID int         `json:"campus_id"`
	User     User        `json:"user"`
}


