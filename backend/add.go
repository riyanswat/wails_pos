package backend

import (
	"encoding/json"
	"os"
	"regexp"
	"strings"
)

func AddToJSON(item, quantity, password string) string {
	// FileName := "./backend/embed/data.json"

	// create data dir
	if err := os.MkdirAll("embed", os.ModePerm); err != nil {
		return "Error creating data dir: " + err.Error()
	}

	// Convert input args to UserData struct
	data := UserData{
		Item:     strings.ToLower(item),
		Quantity: quantity,
		Password: password,
	}

	spacesRegex := `^\s+$`
	spacesReg := regexp.MustCompile(spacesRegex)

	if password == "" ||
		item == "" ||
		quantity == "" ||
		spacesReg.MatchString(item) ||
		spacesReg.MatchString(quantity) ||
		spacesReg.MatchString(password) {
		return "Fill all the fields"
	}

	// quantity validation
	// quantityRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	// re := regexp.MustCompile(quantityRegex)

	// if !re.MatchString(quantity) {
	// 	return "Invalid quantity"
	// }

	// Read existing JSON file
	var existingData []UserData
	file, err := os.Open(FileName)
	if err != nil {
		existingData = make([]UserData, 0)
	} else {
		defer file.Close()
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&existingData); err != nil {
			return "Error decoding json"
		}
	}

	// check if item already exists
	for _, entry := range existingData {
		if entry.Item == data.Item {
			return "Item already exists"
		}
	}

	// Append new data
	existingData = append(existingData, data)

	// Create/open the JSON for writing
	file, err = os.Create(FileName)
	if err != nil {
		return "Error creating data file"
	}
	defer file.Close()

	// Encode and write updated data to the file
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(existingData); err != nil {
		return "Error writing to json"
	}

	return "Successful"
}
