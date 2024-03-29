package backend

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func DeleteFromJSON(itemToDelete string) string {
	// FileName := "./backend/embed/data.json"

	data, err := os.ReadFile(FileName)
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
	itemFound := false

	// Check if the item == itemToDelete
	for _, user := range users {

		if strings.ToLower(user.Item) == strings.ToLower(itemToDelete) {
			itemFound = true
		} else {
			// Append to updatedUsers only if the item doesn't match
			updatedUsers = append(updatedUsers, user)
		}
	}

	// If the item is not found, return an error
	if !itemFound {
		return "Item not found"
	}

	// Marshal the updated slice back to JSON
	updatedData, err := json.Marshal(updatedUsers)
	if err != nil {
		return "Failed to marshal updated data"
	}

	if err := os.WriteFile(FileName, updatedData, os.ModePerm); err != nil {
		fmt.Printf("Error updating the file: %v\n", err)
		return "Failed to update the file"
	}

	return "Deleted successfully"
}
