package main

import (
	"log"
	"net"
)

func main() {
	s := newServer() // Initialize the server.
	go s.run()       // Launch a new go routine.

	// Create a TCP network listener on port 8888.
	listener, err := net.Listen("tcp", ":8888")

	if err != nil {
		log.Fatalf("Failed to start the server: %s", err.Error())
	}

	// Defers the closing of a network listener until the end of the current function.
	defer listener.Close()
	log.Printf("Server started on port 8888...")

	// Infinite loop that waits for incoming client connections on a network listener,
	// and accepts them as they arrive.
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %s", err.Error())
			continue
		}

		// Launch a new go routine.
		go s.newClient(conn)
	}
}
