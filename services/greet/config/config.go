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
	LOCAL_PORT            string
	AWS_DEFAULT_REGION    string
	AWS_ACCESS_KEY_ID     string
	AWS_SECRET_ACCESS_KEY string
}

var Config = Conf{
	AWS_DEFAULT_REGION:    GetEnvOrDefault("AWS_DEFAULT_REGION", "ap-southeast-1"),
	AWS_ACCESS_KEY_ID:     GetEnvOrDefault("AWS_ACCESS_KEY_ID", "123"),
	AWS_SECRET_ACCESS_KEY: GetEnvOrDefault("AWS_SECRET_ACCESS_KEY", "123"),
	LOCAL_PORT:            GetEnvOrDefault("FUNCTIONS_CUSTOMHANDLER_PORT", "8080"),
}
