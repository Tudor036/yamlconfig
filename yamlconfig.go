package yamlconfig

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadYaml[Config interface{}](path string) (*Config, error) {
	if err := validateConfigPath(path); err != nil {
		return nil, err
	}

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	config := new(Config)
	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func validateConfigPath(path string) error {
	info, err := os.Stat(path)

	if err != nil {
		return err
	}

	if info.IsDir() {
		return errors.New("The given path doesn't lead to a file.")
	}

	return nil
}
