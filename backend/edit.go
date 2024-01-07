package backend

import (
	"encoding/json"
	"os"
	"strings"
	// "regexp"
)

func EditJSON(config EditConfig) string {
	filename := "./backend/embed/data.json"

	data, err := os.ReadFile(filename)
	if err != nil {
		return "Failed to read JSON file"
	}

	var users []UserData
	if err := json.Unmarshal(data, &users); err != nil {
		return "Failed to unmarshal JSON data"
	}

	for i, user := range users {
		if strings.ToLower(user.Website) == strings.ToLower(config.WebsiteToEdit) {
			// check edit option and update accordingly
			switch config.EditOption {
			case "email":
				if config.NewEmail != "" {
					users[i].Email = config.NewEmail
				}
			case "password":
				if config.NewPassword != "" {
					users[i].Password = config.NewPassword
				}
			case "both":
				if config.NewEmail != "" {
					users[i].Email = config.NewEmail
				}
				if config.NewPassword != "" {
					users[i].Password = config.NewPassword
				}
			}

			break
		}
	}

	updatedData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return "Failed to marshal updated data"
	}

	if err := os.WriteFile(filename, updatedData, os.ModePerm); err != nil {
		return "Failed to update the file"
	}

	return "Edited successfully"
}
