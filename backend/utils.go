package backend

type UserData struct {
	Website  string `json:"website"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Editing configuration struct
type EditConfig struct {
	WebsiteToEdit string
	NewEmail      string
	NewPassword   string
	EditOption    string // "email", "password", "both"
}
