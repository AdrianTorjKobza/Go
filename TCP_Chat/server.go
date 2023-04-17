package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

type server struct {
	rooms    map[string]*room
	commands chan command
}

// Return a pointer to a newly created server struct.
// Create a new server instance with empty map of rooms and an unbuffered channel of commands.
func newServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NICK:
			s.nick(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.listRooms(cmd.client, cmd.args)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client, cmd.args)
		}
	}
}

// Creates a new client object to represent an incoming connection.
func (s *server) newClient(conn net.Conn) {
	log.Printf("New client connected: %s", conn.RemoteAddr().String())

	// Initialize the new client object with default values for the nickname and commands channel.
	c := &client{
		conn:     conn,
		nick:     "anonymous",
		commands: s.commands,
	}

	c.readInput()
}

func (s *server) nick(c *client, args []string) {
	c.nick = args[1]
	c.msg(fmt.Sprintf("You name is %s", c.nick))
}

// Add the user to the room.
func (s *server) join(c *client, args []string) {
	roomName := args[1]
	r, ok := s.rooms[roomName]

	// Create the room, if the room doesn't exist.
	if !ok {
		r = &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
		}

		s.rooms[roomName] = r
	}

	r.members[c.conn.RemoteAddr()] = c // Associate the user with the room.
	s.quitCurrentRoom(c)               // Leave the current room, if the user joins a new room.

	c.room = r
	r.broadcast(c, fmt.Sprintf("%s has joined the room", c.nick)) // Inform the room members, that a new user has joined the room.
	c.msg(fmt.Sprintf("Welcome to %s", r.name))
}

// List the available rooms.
func (s *server) listRooms(c *client, args []string) {
	var rooms []string

	for name := range s.rooms {
		rooms = append(rooms, name)
	}

	c.msg(fmt.Sprintf("List of available rooms: %s", strings.Join(rooms, ", ")))
}

// Inform the user to join a room.
func (s *server) msg(c *client, args []string) {
	if c.room == nil {
		c.err(errors.New("You must join a room first."))
		return
	}

	c.room.broadcast(c, c.nick+": "+strings.Join(args[1:len(args)], " "))
}

// Disconnect the user from the server.
func (s *server) quit(c *client, args []string) {
	log.Printf("Client disconnected: %s", c.conn.RemoteAddr().String())
	s.quitCurrentRoom(c)
	c.msg("Bye!")
	c.conn.Close()
}

// Remove the user from the room.
func (s *server) quitCurrentRoom(c *client) {
	if c.room != nil {
		delete(c.room.members, c.conn.RemoteAddr()) // Removes the client's connection address from the members map of the current room.
		c.room.broadcast(c, fmt.Sprintf("%s has left the room.", c.nick))
	}
}
