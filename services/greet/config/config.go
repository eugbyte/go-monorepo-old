package config

import "os"

func GetEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

type Conf struct {
	LOCAL_PORT string
}

var Config = Conf{
	LOCAL_PORT: GetEnvOrDefault("PORT", "8080"),
}
