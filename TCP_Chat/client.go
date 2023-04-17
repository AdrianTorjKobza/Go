package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client struct {
	conn net.Conn
	nick string
	room *room
	// chan<- is a channel operator that specifies that the channel is write-only.
	// This means that the function or method can only send values to the channel,
	// but cannot receive values from it.
	// Sends a command value to a channel called commands.
	commands chan<- command
}

// Read input messages from the client connection represented by c.conn.
func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')

		if err != nil {
			return
		}

		// Remove any trailing \r and \n characters from the message.
		msg = strings.Trim(msg, "\r\n")

		args := strings.Split(msg, " ")   // Split the message by space.
		cmd := strings.TrimSpace(args[0]) // The first element of the message is the command.

		switch cmd {
		case "/nick":
			c.commands <- command{
				id:     CMD_NICK,
				client: c,
				args:   args,
			}
		case "/join":
			c.commands <- command{
				id:     CMD_JOIN,
				client: c,
				args:   args,
			}
		case "/rooms":
			c.commands <- command{
				id:     CMD_ROOMS,
				client: c,
				args:   args,
			}
		case "/msg":
			c.commands <- command{
				id:     CMD_MSG,
				client: c,
				args:   args,
			}
		case "/quit":
			c.commands <- command{
				id:     CMD_QUIT,
				client: c,
				args:   args,
			}
		default:
			c.err(fmt.Errorf("Unknown command: %s", cmd))
		}
	}
}

// Write a byte slice containing an error message, that will be sent to the client.
func (c *client) err(err error) {
	c.conn.Write([]byte("ERR: " + err.Error() + "\n"))
}

// Write a message to the client connection.
func (c *client) msg(msg string) {
	c.conn.Write([]byte("> " + msg + "\n"))
}
