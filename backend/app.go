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

func (a *App) Add(website, email, password string) string {
	return AddToJSON(website, email, password)
}

func (a *App) Delete(websiteToDelete string) string {
	return DeleteFromJSON(websiteToDelete)
}

func (a *App) Search(webToSearch string) ([]interface{}, error) {
	return SearchWebsite(webToSearch)
}

func (a *App) AllData() []UserData {
	return ShowAll()
}

// func (a *App) Edit(web, email, password, editOption string) string {
// 	spacesPattern := `^\s+$`
// 	spacesRe := regexp.MustCompile(spacesPattern)

// validateInput := func(value, errorMessage string) string {
// 	if value == "" || spacesRe.MatchString(value) {
// 		return errorMessage
// 	}
// 	return ""
// }

// 	websiteError := validateInput(web, "Enter a website")
// 	emailError := ""
// 	passwordError := ""

// 	switch editOption {
// 	case "email":
// 		emailError = validateInput(email, "Enter a valid email")
// 	case "password":
// 		passwordError = validateInput(password, "Enter a valid password")
// 	case "both":
// 		emailError = validateInput(email, "Enter a valid email")
// 		passwordError = validateInput(password, "Enter a valid password")
// 		if emailError != "" || passwordError != "" {
// 			return "Enter valid values"
// 		}
// 		if email == "" || password == "" {
// 			return "Enter both email and password"
// 		}
// 	default:
// 		return "Invalid edit option"
// 	}

// 	if websiteError != "" {
// 		return websiteError
// 	}

// 	data := EditConfig{
// 		WebsiteToEdit: web,
// 		NewEmail:      email,
// 		NewPassword:   password,
// 		EditOption:    editOption,
// 	}

// 	return EditJSON(data)
// }

// ? ====================================================================

func (a *App) Edit(web, email, password, editOption string) string {
	spacesPattern := `^\s+$`
	spacesRe := regexp.MustCompile(spacesPattern)

	if web == "" ||
		spacesRe.MatchString(web) {
		return "Enter a website"
	}

	if editOption == "email" && (email == "" || spacesRe.MatchString(email)) {
		return "Enter a valid email"
	}

	if editOption == "password" && (password == "" || spacesRe.MatchString(password)) {
		return "Enter a valid password"
	}

	if editOption == "both" && (password == "" ||
		spacesRe.MatchString(password) ||
		email == "" ||
		spacesRe.MatchString(email)) {
		return "Enter valid values"
	}

	if !(editOption == "email" || editOption == "password" || editOption == "both") {
		return "Invalid edit option"
	}

	if editOption == "both" {
		data := EditConfig{
			WebsiteToEdit: web,
			NewEmail:      email,
			NewPassword:   password,
			EditOption:    editOption,
		}
		return EditJSON(data)
	}

	if editOption == "email" {
		data := EditConfig{
			WebsiteToEdit: web,
			NewEmail:      email,
			EditOption:    editOption,
		}

		return EditJSON(data)
	} else if editOption == "password" {
		data := EditConfig{
			WebsiteToEdit: web,
			NewPassword:   password,
			EditOption:    editOption,
		}

		return EditJSON(data)
	} else if editOption == "both" {
		data := EditConfig{
			WebsiteToEdit: web,
			NewEmail:      email,
			NewPassword:   password,
			EditOption:    editOption,
		}

		return EditJSON(data)
	} else {
		return "Unexpected edit option"
	}
}
