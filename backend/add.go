package backend

import (
	"encoding/json"
	"os"
	"regexp"
	"strings"
)

func AddToJSON(website, email, password string) string {
	filename := "./backend/embed/data.json"

	// create data dir
	if err := os.MkdirAll("embed", os.ModePerm); err != nil {
		return "Error creating data dir: " + err.Error()
	}

	// Convert input args to UserData struct
	data := UserData{
		Website:  strings.ToLower(website),
		Email:    email,
		Password: password,
	}

	spacesRegex := `^\s+$`
	spacesReg := regexp.MustCompile(spacesRegex)

	if password == "" ||
		website == "" ||
		email == "" ||
		spacesReg.MatchString(website) ||
		spacesReg.MatchString(email) ||
		spacesReg.MatchString(password) {
		return "Fill all the fields"
	}

	// email validation
	// emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	// re := regexp.MustCompile(emailRegex)

	// if !re.MatchString(email) {
	// 	return "Invalid email"
	// }

	// Read existing JSON file
	var existingData []UserData
	file, err := os.Open(filename)
	if err != nil {
		existingData = make([]UserData, 0)
	} else {
		defer file.Close()
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&existingData); err != nil {
			return "Error decoding json"
		}
	}

	// check if website already exists
	for _, entry := range existingData {
		if entry.Website == data.Website {
			return "Website already exists"
		}
	}

	// Append new data
	existingData = append(existingData, data)

	// Create/open the JSON for writing
	file, err = os.Create(filename)
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
