package helpers

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
	if err != nil{
		return err
	}
	return nil
}

func ReadFromJSONFile(fileName string, targetData any) error {
	jsonData, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData,targetData)
	if err != nil {
		return err
	}
	return nil
}