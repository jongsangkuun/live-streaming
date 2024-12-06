package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type Config struct {
	Develop struct {
		DBHost     string `ini:"db_host"`
		DBPort     int    `ini:"db_port"`
		DBID       string `ini:"db_id"`
		DBPassword string `ini:"db_pw"`
		ServerPort int    `ini:"server_port"`
	} `ini:"DEVELOP"`
	Product struct {
		DBHost     string `ini:"db_host"`
		DBPort     int    `ini:"db_port"`
		DBID       string `ini:"db_id"`
		DBPassword string `ini:"db_pw"`
		ServerPort int    `ini:"server_port"`
	} `ini:"PRODUCT"`
}

// LoadConfig reads an INI config file and unmarshals it into a Config struct
func LoadConfig(path string) (*Config, error) {
	cfg, err := ini.Load(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config

	err = cfg.Section("DEVELOP").MapTo(&config.Develop)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DEVELOP section: %w", err)
	}

	err = cfg.Section("PRODUCT").MapTo(&config.Product)
	if err != nil {
		return nil, fmt.Errorf("failed to parse PRODUCT section: %w", err)
	}

	return &config, nil
}
