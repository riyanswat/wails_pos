package backend

import (
	"encoding/json"
	"os"
	"strings"
	// "regexp"
)

func EditJSON(config EditConfig) string {
	// FileName := "./backend/embed/data.json"

	data, err := os.ReadFile(FileName)
	if err != nil {
		return "Failed to read JSON file"
	}

	var users []UserData
	if err := json.Unmarshal(data, &users); err != nil {
		return "Failed to unmarshal JSON data"
	}

	for i, user := range users {
		if strings.ToLower(user.Item) == strings.ToLower(config.ItemToEdit) {
			// check edit option and update accordingly
			switch config.EditOption {
			case "quantity":
				if config.NewQuantity != "" {
					users[i].Quantity = config.NewQuantity
				}
			case "price":
				if config.NewPrice != "" {
					users[i].Price = config.NewPrice
				}
			case "both":
				if config.NewQuantity != "" {
					users[i].Quantity = config.NewQuantity
				}
				if config.NewPrice != "" {
					users[i].Price = config.NewPrice
				}
			}

			break
		}
	}

	updatedData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return "Failed to marshal updated data"
	}

	if err := os.WriteFile(FileName, updatedData, os.ModePerm); err != nil {
		return "Failed to update the file"
	}

	return "Edited successfully"
}
