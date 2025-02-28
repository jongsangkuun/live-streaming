package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	envStatus        string
	debug            bool
	allowedOrigins   []string
	postgresUser     string
	postgresPassword string
	postgresDB       string
	postgresHost     string
	postgresPort     string
	redisHost        string
	redisPort        string
	redisDatabase    int
	redisPassword    string
}

func NewConfig() *Config {
	return &Config{}
}

func getEnvWithPrefix(prefix, key string) (string, error) {
	value := os.Getenv(prefix + key)
	if value == "" {
		return "", fmt.Errorf("environment variable %s%s not set", prefix, key)
	}
	return value, nil
}

func getEnvIntWithPrefix(prefix, key string) (int, error) {
	valueStr, err := getEnvWithPrefix(prefix, key)
	if err != nil {
		return 0, err
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, fmt.Errorf("invalid %s%s value: %s", prefix, key, valueStr)
	}
	return value, nil
}

func EnvLoad() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := NewConfig()

	envStatus, err := getEnvWithPrefix("", "ENV_STATUS")
	if err != nil {
		return nil, err
	}
	config.envStatus = strings.ToUpper(envStatus)

	prefix := ""
	switch config.envStatus {
	case "DEV":
		prefix = "DEV_"
	case "PROD":
		prefix = "PROD_"
	default:
		return nil, fmt.Errorf("invalid env status: %s", config.envStatus)
	}

	debugStr, err := getEnvWithPrefix(prefix, "DEBUG")
	if err != nil {
		return nil, err
	}
	config.debug = debugStr == "true"

	allowedOriginsStr, err := getEnvWithPrefix(prefix, "ALLOWED_ORIGINS")
	if err != nil {
		return nil, err
	}
	config.allowedOrigins = strings.Split(allowedOriginsStr, ",")

	config.postgresUser, err = getEnvWithPrefix(prefix, "POSTGRES_USER")
	if err != nil {
		return nil, err
	}
	// ... (나머지 postgres 설정도 동일하게 적용)

	config.redisHost, err = getEnvWithPrefix(prefix, "REDIS_HOST")
	if err != nil {
		return nil, err
	}

	config.redisPort, err = getEnvWithPrefix(prefix, "REDIS_PORT")
	if err != nil {
		return nil, err
	}

	config.redisDatabase, err = getEnvIntWithPrefix(prefix, "REDIS_DATABASE") // prefix 없이 사용
	if err != nil {
		return nil, err
	}

	config.redisPassword, err = getEnvWithPrefix(prefix, "REDIS_PASSWORD")
	if err != nil {
		return nil, err
	}

	return config, nil
}
