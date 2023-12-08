package helpers

import (
	"encoding/json"
	"os"
)

func SaveJSONFile(fileName string, data any) error {
	// convert data into JSON format
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// save the converted data in our json file
	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadFromJSONFile(fileName string, targetData any) error {
	//get the data from the JSON file
	jsonData, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	//convert from JSON format into readable type
	err = json.Unmarshal(jsonData, targetData)
	if err != nil {
		return err
	}
	return nil
}
