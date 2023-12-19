package config

import "os"

type GreetingConfig struct {
	URL string
}

func NewGreetingConfig() GreetingConfig {
	url, ok := os.LookupEnv("GREETING_URL")
	if !ok {
		return GreetingConfig{URL: "http://localhost:8080"}
	}
	return GreetingConfig{URL: url}
}
