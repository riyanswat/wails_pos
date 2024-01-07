package main

import (
	"context"
	// "embed"
	"pos/backend"
	"regexp"
)

// var content embed.FS

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

// Expose backend apis to frontend
func (a *App) Generate(length int) string {
	return backend.GenerateRandomPassword(length)
}

func (a *App) Add(website, email, password string) string {
	return backend.AddToJSON(website, email, password)
}

func (a *App) Delete(websiteToDelete string) string {
	return backend.DeleteFromJSON(websiteToDelete)
}

func (a *App) Search(webToSearch string) ([]interface{}, error) {
	return backend.SearchWebsite(webToSearch)
}

func (a *App) AllData() []backend.UserData {
	return backend.ShowAll()
}

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
		data := backend.EditConfig{
			WebsiteToEdit: web,
			NewEmail:      email,
			NewPassword:   password,
			EditOption:    editOption,
		}
		return backend.EditJSON(data)
	}

	if editOption == "email" {
		data := backend.EditConfig{
			WebsiteToEdit: web,
			NewEmail:      email,
			EditOption:    editOption,
		}

		return backend.EditJSON(data)
	} else if editOption == "password" {
		data := backend.EditConfig{
			WebsiteToEdit: web,
			NewPassword:   password,
			EditOption:    editOption,
		}

		return backend.EditJSON(data)
	} else if editOption == "both" {
		data := backend.EditConfig{
			WebsiteToEdit: web,
			NewEmail:      email,
			NewPassword:   password,
			EditOption:    editOption,
		}

		return backend.EditJSON(data)
	} else {
		return "Unexpected edit option"
	}
}
