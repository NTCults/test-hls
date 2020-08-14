package config

import (
	"os"
	"strings"
)

type Config struct {
	Port      string
	SecretKey string
	// AccessKeyUrl parameter allows you to set url to get access key
	// line must start with a slash character, also
	// it must contain segments: /{eventID}/{clientID}/{keyName}
	AccessKeyURL string
}

func NewConfig() *Config {
	var conf Config
	conf.Port = readEnvOrDefault("PORT", "8090")

	conf.SecretKey = readEnvOrDefault("SECRET_KEY", "some_secret_key")

	conf.AccessKeyURL = readEnvOrDefault("ACCESS_KEY_URL", "/access/{eventID}/{clientID}/{keyName}")

	return &conf
}

func (c *Config) GetSecretKeyBytes() []byte {
	return []byte(c.SecretKey)
}

func readEnvOrDefault(envParam, defaultValue string) string {
	v := strings.TrimSpace(os.Getenv(envParam))
	if v != "" {
		return v
	}
	return defaultValue
}
