package helper

import (
	"encoding/json"
	"os"
)

func SaveJSONFile(fileName string, data any) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadFromJSONFile(fileName string, targetData any) error {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return err 
	}

err = json.Unmarshal(file, targetData)
if err != nil {
	return err 
}

	return nil
}