package api

// CoalitionUser represents an user in coalition from the 42's API.
type CoalitionUser struct {
	ID          int `json:"id"`
	CoalitionID int `json:"coalition_id"`
	//UserID      int       `json:"user_id"`
	//CreatedAt   time.Time `json:"created_at"`
	//UpdatedAt   time.Time `json:"updated_at"`
}
