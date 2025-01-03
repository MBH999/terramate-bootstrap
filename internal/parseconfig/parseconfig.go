package parseconfig

import (
	"io"
	"log"
	"os"
	"terramate-bootstrap/types"

	"gopkg.in/yaml.v3"
)

func ParseConfigFile(config_file *string) (*types.Terramate, error) {
	file, err := os.Open(*config_file)
	if err != nil {
		log.Fatalf("Error opening configuration file: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading configuration file: %v", err)
	}

	var config types.Terramate
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error parsing configuration file %v", err)
	}

	// fmt.Printf("%#v\n", config.Backend.AzureRM.ContainerName)

	return &config, nil
}
