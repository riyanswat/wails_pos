package backend

// File path:
// var FileName string = "./backend/embed/data.json"
var FileName string = "./embed/data.json"

// App password
var AppPassword string = "riyan"

// STRUCTS:
// user json struct
type UserData struct {
	Item     string `json:"item"`
	Quantity string `json:"quantity"`
	Price    string `json:"price"`
}

// Editing configuration struct
type EditConfig struct {
	ItemToEdit  string
	NewQuantity string
	NewPrice    string
	EditOption  string // "quantity", "price", "both"
}
