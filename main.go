package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// TODO:  Create a Config module???
type route struct {
	Key         string `yaml:"key"`
	Domain      string `yaml:"domain"`
	Destination string `yaml:"destination"`
}

// TODO:  Create a Config module???
type fixed_length struct {
	Length int
	Route  []route
}

// TODO:  Create a Config module???
type Config struct {
	Defaults    route
	Route       []route
	FixedLength []fixed_length `yaml:"fixed_length"`
}

func main() {
	c := unmarshalConfig()
	log.Printf("Here:  %v", c)
}

// TODO:  Create a Config module???
func unmarshalConfig() Config {
	config_data := readConfig()

	c := Config{}
	err := yaml.Unmarshal(config_data, &c)

	if err != nil {
		log.Fatalf("Unable to unmarshal config.yml: %v", err)
	}

	return c
}

// TODO:  Create a Config module???
func readConfig() []byte {
	content, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("Unable to open config.yml:  %v", err)
	}
	return content
}
