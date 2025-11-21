package config

import (
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig() map[string]interface{} {
	f, err := os.Open("config.yaml")
	if err != nil {
		log.Fatalf("Error Reading config file %v", err)
	}

	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	var config map[string]interface{}

	// Unmarshal YAML into map
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error unmarshalling YAML: %v", err)
	}
	return config
}
