package main

import (
	"log"
	"net/http"

	"github.com/exegeteio/goaway/pkg/config"
	"github.com/exegeteio/goaway/pkg/router"
)

func main() {
	c := config.Read()
	router := router.Router{
		Config: c,
	}
	log.Printf("Here:  %v", c.Routes)

	// How to pass config into the handler?
	http.HandleFunc("/", router.Handler)

	log.Println("Listening on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
