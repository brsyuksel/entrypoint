package configuration

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigurationError struct {
	message string
}

func (c ConfigurationError) Error() string {
	return c.message
}

func newConfErr(err error) ConfigurationError {
	return ConfigurationError{message: err.Error()}
}

func Load(filePath string) (*Config, error) {
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, newConfErr(err)
	}

	conf := Config{}
	if err := yaml.Unmarshal(fileBytes, &conf); err != nil {
		return nil, newConfErr(err)
	}

	return &conf, nil
}
