package config

import "os"

type Greeting struct {
	URL string
}

func NewGreeting() Greeting {
	url, ok := os.LookupEnv("GREETING_URL")
	if !ok {
		return Greeting{URL: "http://localhost:8080"}
	}
	return Greeting{URL: url}
}
