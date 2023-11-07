package util

import (
	"encoding/json"
	"os"
)

// ReadJSONFile reads a JSON file and decodes it into the provided interface.
func ReadJSONFile(filePath string, v interface{}) error {
	configFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	if err := decoder.Decode(v); err != nil {
		return err
	}

	return nil
}
