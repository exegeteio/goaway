package config

import (
	"errors"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type route struct {
	Key         string `yaml:"key"`
	Domain      string `yaml:"domain"`
	Destination string `yaml:"destination"`
}

type fixed_length struct {
	Length int     `yaml:"length"`
	Routes []route `yaml:"route"`
}

type Config struct {
	Defaults     route          `yaml:"defaults"`
	Routes       []route        `yaml:"route"`
	FixedLengths []fixed_length `yaml:"fixed_length"`
}

func findRoute(routes []route, key string) (string, string, error) {
	for _, route := range routes {
		if route.Key == key {
			return route.Domain, route.Destination, nil
		}
	}
	return "", "", errors.New("Route not found")
}

// Find Route for a given key.
func (c *Config) Route(key string) (string, string) {
	domain, destination, err := findRoute(c.Routes, key)
	if err != nil {
		return c.Defaults.Domain, c.Defaults.Destination
	}
	return domain, destination
}

// Find Fixed Length route for a given length.
func (c *Config) FixedLength(length int, key string) (string, string, error) {
	for _, fixed_length := range c.FixedLengths {
		if fixed_length.Length == length {
			// Check the list of routes for a match.
			domain, destination, err := findRoute(fixed_length.Routes, key)
			if err == nil {
				return domain, destination, nil
			}
		}
	}
	return "", "", errors.New("Fixed Length route not found")
}

// Read config file from disk and marshal into a Config struct.
func Read() Config {
	config_data := readConfig()

	c := Config{}
	err := yaml.Unmarshal(config_data, &c)

	if err != nil {
		log.Fatalf("Unable to unmarshal config.yml: %v", err)
	}

	return c
}

// Read config file from disk.
func readConfig() []byte {
	content, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("Unable to open config.yml:  %v", err)
	}
	return content
}
