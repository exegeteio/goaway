package main

import (
	"log"

	"github.com/exegeteio/goaway/pkg/config"
)

func main() {
	c := config.Read()
	log.Printf("Here:  %v", c)
}
