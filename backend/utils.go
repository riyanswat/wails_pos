package backend

// File path:
// var FileName string = "./backend/embed/data.json"
var FileName string = "./embed/data.json"

// STRUCTS:
// user json struct
type UserData struct {
	Item     string `json:"item"`
	Quantity string `json:"quantity"`
	Password string `json:"password"`
}

// Editing configuration struct
type EditConfig struct {
	ItemToEdit  string
	NewQuantity string
	NewPassword string
	EditOption  string // "quantity", "password", "both"
}
