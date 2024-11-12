package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
}

// LoadConfig reads a YAML config file and unmarshals it into a Config struct
func LoadConfig(path string) (*Config, error) {
	var config Config
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	if err = yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &config, nil
}
