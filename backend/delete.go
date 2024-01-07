package backend

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func DeleteFromJSON(websiteToDelete string) string {
	filename := "./backend/embed/data.json"

	data, err := os.ReadFile(filename)
	if err != nil {
		return "Failed to read JSON file"
	}

	// Unmarshal JSON data into a slice of UserData
	var users []UserData
	if err := json.Unmarshal(data, &users); err != nil {
		return "Failed to unmarshal JSON data"
	}

	// Create a new slice to store updated data
	var updatedUsers []UserData
	websiteFound := false

	// Check if the website == websiteToDelete
	for _, user := range users {

		if strings.ToLower(user.Website) == strings.ToLower(websiteToDelete) {
			websiteFound = true
		} else {
			// Append to updatedUsers only if the website doesn't match
			updatedUsers = append(updatedUsers, user)
		}
	}

	// If the website is not found, return an error
	if !websiteFound {
		return "Website not found"
	}

	// Marshal the updated slice back to JSON
	updatedData, err := json.Marshal(updatedUsers)
	if err != nil {
		return "Failed to marshal updated data"
	}

	if err := os.WriteFile(filename, updatedData, os.ModePerm); err != nil {
		fmt.Printf("Error updating the file: %v\n", err)
		return "Failed to update the file"
	}

	return "Deleted successfully"
}
