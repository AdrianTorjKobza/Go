package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80" // Listen port 80.
type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting Broker service on port %s.", webPort)

	// Define HTTP server.
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// Start the server and listen to incoming requests.
	// Print the error, if any.
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
