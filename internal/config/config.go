package config

import (
	"log"
	"os"
)

func Add(path string, param string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	if _, err := file.WriteString(param); err != nil {
		return err
	}

	return file.Close()
}

func AddMultiple(configParams map[string]string) {
	for path, param := range configParams {
		if err := Add(path, param); err != nil {
			log.Fatalf("Failed to add param=%v to config=%v", param, path)
		}
	}
}
