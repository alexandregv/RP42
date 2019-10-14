package api

import "time"

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
