package backend

import (
	"context"
	"embed"
	"regexp"
)

var content embed.FS

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// Expose BACKEND apis to frontend
func (a *App) Generate(length int) string {
	return GenerateRandomPassword(length)
}

func (a *App) Add(item, quantity, password string) string {
	return AddToJSON(item, quantity, password)
}

func (a *App) Delete(itemToDelete string) string {
	return DeleteFromJSON(itemToDelete)
}

func (a *App) Search(itemToSearch string) ([]interface{}, error) {
	return SearchItem(itemToSearch)
}

func (a *App) AllData() []UserData {
	return ShowAll()
}

// func (a *App) Edit(web, quantity, password, editOption string) string {
// 	spacesPattern := `^\s+$`
// 	spacesRe := regexp.MustCompile(spacesPattern)

// validateInput := func(value, errorMessage string) string {
// 	if value == "" || spacesRe.MatchString(value) {
// 		return errorMessage
// 	}
// 	return ""
// }

// 	itemError := validateInput(web, "Enter a item")
// 	quantityError := ""
// 	passwordError := ""

// 	switch editOption {
// 	case "quantity":
// 		quantityError = validateInput(quantity, "Enter a valid quantity")
// 	case "password":
// 		passwordError = validateInput(password, "Enter a valid password")
// 	case "both":
// 		quantityError = validateInput(quantity, "Enter a valid quantity")
// 		passwordError = validateInput(password, "Enter a valid password")
// 		if quantityError != "" || passwordError != "" {
// 			return "Enter valid values"
// 		}
// 		if quantity == "" || password == "" {
// 			return "Enter both quantity and password"
// 		}
// 	default:
// 		return "Invalid edit option"
// 	}

// 	if itemError != "" {
// 		return itemError
// 	}

// 	data := EditConfig{
// 		ItemToEdit: web,
// 		NewQuantity:      quantity,
// 		NewPassword:   password,
// 		EditOption:    editOption,
// 	}

// 	return EditJSON(data)
// }

// ? ====================================================================

func (a *App) Edit(web, quantity, password, editOption string) string {
	spacesPattern := `^\s+$`
	spacesRe := regexp.MustCompile(spacesPattern)

	if web == "" ||
		spacesRe.MatchString(web) {
		return "Enter an item"
	}

	if editOption == "quantity" && (quantity == "" || spacesRe.MatchString(quantity)) {
		return "Enter a valid quantity"
	}

	if editOption == "password" && (password == "" || spacesRe.MatchString(password)) {
		return "Enter a valid password"
	}

	if editOption == "both" && (password == "" ||
		spacesRe.MatchString(password) ||
		quantity == "" ||
		spacesRe.MatchString(quantity)) {
		return "Enter valid values"
	}

	if !(editOption == "quantity" || editOption == "password" || editOption == "both") {
		return "Invalid edit option"
	}

	if editOption == "both" {
		data := EditConfig{
			ItemToEdit:  web,
			NewQuantity: quantity,
			NewPassword: password,
			EditOption:  editOption,
		}
		return EditJSON(data)
	}

	if editOption == "quantity" {
		data := EditConfig{
			ItemToEdit:  web,
			NewQuantity: quantity,
			EditOption:  editOption,
		}

		return EditJSON(data)
	} else if editOption == "password" {
		data := EditConfig{
			ItemToEdit:  web,
			NewPassword: password,
			EditOption:  editOption,
		}

		return EditJSON(data)
	} else if editOption == "both" {
		data := EditConfig{
			ItemToEdit:  web,
			NewQuantity: quantity,
			NewPassword: password,
			EditOption:  editOption,
		}

		return EditJSON(data)
	} else {
		return "Unexpected edit option"
	}
}
