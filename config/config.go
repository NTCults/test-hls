package config

import (
	"os"
	"strings"
)

type Config struct {
	Port         string
	SecretKey    string
	AccessKeyUrl string
}

func NewConfig() *Config {
	var conf Config
	conf.Port = readEnvOrDefault("PORT", "8090")

	conf.SecretKey = readEnvOrDefault("SECRET_KEY", "some_secret_key")

	conf.AccessKeyUrl = readEnvOrDefault("ACCESS_KEY_URL", "/access/{eventID}/{clientID}/{keyName}")

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
