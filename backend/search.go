package backend

import (
	"encoding/json"
	"fmt"
	"os"

	// "os"
	"strings"
)

//! ========================================

// * ------------------------------------

func SearchItem(itemToSearch string) ([]interface{}, error) {
	// FileName := "./backend/embed/data.json"

	data, err := os.ReadFile(FileName)
	if err != nil {
		return nil, fmt.Errorf("Failed to read JSON file: %v", err)
	}

	var users []UserData
	if err := json.Unmarshal(data, &users); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal JSON data: %v", err)
	}

	for _, user := range users {
		if strings.ToLower(user.Item) == strings.ToLower(itemToSearch) {
			return []interface{}{user, "yes"}, nil
		}
	}

	// If the loop completes and the item is not found, then return "Not found"
	return []interface{}{UserData{}, "no"}, nil
}

// * ------------------------------------

// func SearchItem(itemToSearch string) (UserData, string) {
// 	FileName := "./backend/embed/data.json"

// 	data, err := os.ReadFile(FileName)
// 	if err != nil {
// 		return UserData{}, fmt.Sprintf("Failed to read JSON file: %v", err)
// 	}

// 	var users []UserData
// 	if err := json.Unmarshal(data, &users); err != nil {
// 		return UserData{}, "Failed to unmarshal JSON data"
// 	}

// 	for _, user := range users {
// 		if strings.ToLower(user.Item) == strings.ToLower(itemToSearch) {
// 			return user, "yes"
// 		}
// 	}

// 	// If the loop completes and the item is not found, then return "Not found"
// 	return UserData{}, "no"
// }

//! ========================================

// func SearchItem(itemToSearch string) string {
// 	FileName := "./backend/embed/data.json"
// 	//TODO: return multiple values from this function
// 	// the first value should be a string or bool indicating that the value was found
// 	// the second should be a UserData struct exposing the data to the frontend
// 	data, err := os.ReadFile(FileName)
// 	if err != nil {
// 		return fmt.Sprintf("Failed to read JSON file: %v", err)
// 	}

// 	var users []UserData
// 	if err := json.Unmarshal(data, &users); err != nil {
// 		return "Failed to unmarshal JSON data"
// 	}

// 	for _, user := range users {
// 		if strings.ToLower(user.Item) == strings.ToLower(itemToSearch) {
// 			return "yes"
// 		}
// 	}

// 	// If the loop completes and the item is not found, then return "Not found"
// 	return "no"
// }
