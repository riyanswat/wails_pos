package backend

// File path:
// var FileName string = "./backend/embed/data.json"
var FileName string = "./embed/data.json"

// STRUCTS:
// user json struct
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
