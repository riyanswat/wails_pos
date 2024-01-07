package backend

import (
	"encoding/json"
	"fmt"
	"os"
)

func ShowAll() []UserData {
	filename := "./backend/embed/data.json"

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(fmt.Errorf("Failed to read JSON file: %v", err))
	}

	var users []UserData
	if err := json.Unmarshal(data, &users); err != nil {
		fmt.Println(fmt.Errorf("Failed to unmarshal JSON data: %v", err))
	}

	return users
}
