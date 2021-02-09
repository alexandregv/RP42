package api

//import "time"

// User represents an user from the 42's API.
// Truncated to keep only useful entries.
type User struct {
	ID int `json:"id"`
	//		Email           string        `json:"email"`
	Login string `json:"login"`
	//		FirstName       string        `json:"first_name"`
	//		LastName        string        `json:"last_name"`
	//		URL             string        `json:"url"`
	//		Phone           interface{}   `json:"phone"`
	//		Displayname     string        `json:"displayname"`
	//		ImageURL        string        `json:"image_url"`
	//		Staff           bool          `json:"staff?"`
	//		CorrectionPoint int           `json:"correction_point"`
	//		PoolMonth       string        `json:"pool_month"`
	//		PoolYear        string        `json:"pool_year"`
	Location string `json:"location"`
	//		Wallet          int           `json:"wallet"`
	//		Groups          []interface{} `json:"groups"`
	CursusUsers []CursusUser `json:"cursus_users"`
	//		ProjectsUsers  []interface{} `json:"projects_users"`
	//		LanguagesUsers []struct {
	//			ID         int       `json:"id"`
	//			LanguageID int       `json:"language_id"`
	//			UserID     int       `json:"user_id"`
	//			Position   int       `json:"position"`
	//			CreatedAt  time.Time `json:"created_at"`
	//		} `json:"languages_users"`
	//		Achievements []interface{} `json:"achievements"`
	//		Titles       []interface{} `json:"titles"`
	//		TitlesUsers  []interface{} `json:"titles_users"`
	//		Partnerships []interface{} `json:"partnerships"`
	//		Patroned     []struct {
	//			ID          int       `json:"id"`
	//			UserID      int       `json:"user_id"`
	//			GodfatherID int       `json:"godfather_id"`
	//			Ongoing     bool      `json:"ongoing"`
	//			CreatedAt   time.Time `json:"created_at"`
	//			UpdatedAt   time.Time `json:"updated_at"`
	//		} `json:"patroned"`
	//		Patroning       []interface{} `json:"patroning"`
	//		ExpertisesUsers []struct {
	//			ID          int       `json:"id"`
	//			ExpertiseID int       `json:"expertise_id"`
	//			Interested  bool      `json:"interested"`
	//			Value       int       `json:"value"`
	//			ContactMe   bool      `json:"contact_me"`
	//			CreatedAt   time.Time `json:"created_at"`
	//			UserID      int       `json:"user_id"`
	//		} `json:"expertises_users"`
	Campus []struct {
		//			ID       int    `json:"id"`
		Name string `json:"name"`
		//			TimeZone string `json:"time_zone"`
		//			Language struct {
		//				ID         int       `json:"id"`
		//				Name       string    `json:"name"`
		//				Identifier string    `json:"identifier"`
		//				CreatedAt  time.Time `json:"created_at"`
		//				UpdatedAt  time.Time `json:"updated_at"`
		//			} `json:"language"`
		//			UsersCount  int `json:"users_count"`
		//			VogsphereID int `json:"vogsphere_id"`
	} `json:"campus"`
	//		CampusUsers []struct {
	//			ID        int  `json:"id"`
	//			UserID    int  `json:"user_id"`
	//			CampusID  int  `json:"campus_id"`
	//			IsPrimary bool `json:"is_primary"`
	//		} `json:"campus_users"`
}

type CursusUser struct {
	//			ID           int           `json:"id"`
	//			BeginAt      time.Time     `json:"begin_at"`
	//			EndAt interface{} `json:"end_at"`
	//			Grade        interface{}   `json:"grade"`
	Level float64 `json:"level"`
	//			Skills       []interface{} `json:"skills"`
	//			CursusID     int           `json:"cursus_id"`
	//			HasCoalition bool          `json:"has_coalition"`
	//			User         User          `json:"user"`
	Cursus struct {
		//				ID        int       `json:"id"`
		//				CreatedAt time.Time `json:"created_at"`
		//				Name      string    `json:"name"`
		Slug string `json:"slug"`
	} `json:"cursus"`
}
