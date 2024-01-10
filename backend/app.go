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

func (a *App) Add(item, quantity, price string) string {
	return AddToJSON(item, quantity, price)
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

// func (a *App) Edit(web, quantity, price, editOption string) string {
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
// 	priceError := ""

// 	switch editOption {
// 	case "quantity":
// 		quantityError = validateInput(quantity, "Enter a valid quantity")
// 	case "price":
// 		priceError = validateInput(price, "Enter a valid price")
// 	case "both":
// 		quantityError = validateInput(quantity, "Enter a valid quantity")
// 		priceError = validateInput(price, "Enter a valid price")
// 		if quantityError != "" || priceError != "" {
// 			return "Enter valid values"
// 		}
// 		if quantity == "" || price == "" {
// 			return "Enter both quantity and price"
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
// 		NewPrice:   price,
// 		EditOption:    editOption,
// 	}

// 	return EditJSON(data)
// }

// ? ====================================================================

func (a *App) Edit(web, quantity, price, editOption string) string {
	spacesPattern := `^\s+$`
	spacesRe := regexp.MustCompile(spacesPattern)

	if web == "" ||
		spacesRe.MatchString(web) {
		return "Enter an item"
	}

	if editOption == "quantity" && (quantity == "" || spacesRe.MatchString(quantity)) {
		return "Enter a valid quantity"
	}

	if editOption == "price" && (price == "" || spacesRe.MatchString(price)) {
		return "Enter a valid price"
	}

	if editOption == "both" && (price == "" ||
		spacesRe.MatchString(price) ||
		quantity == "" ||
		spacesRe.MatchString(quantity)) {
		return "Enter valid values"
	}

	if !(editOption == "quantity" || editOption == "price" || editOption == "both") {
		return "Invalid edit option"
	}

	if editOption == "both" {
		data := EditConfig{
			ItemToEdit:  web,
			NewQuantity: quantity,
			NewPrice:    price,
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
	} else if editOption == "price" {
		data := EditConfig{
			ItemToEdit: web,
			NewPrice:   price,
			EditOption: editOption,
		}

		return EditJSON(data)
	} else if editOption == "both" {
		data := EditConfig{
			ItemToEdit:  web,
			NewQuantity: quantity,
			NewPrice:    price,
			EditOption:  editOption,
		}

		return EditJSON(data)
	} else {
		return "Unexpected edit option"
	}
}
