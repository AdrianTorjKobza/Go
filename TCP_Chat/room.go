package main

import "net"

type room struct {
	name string
	// Define variable members, as a map with keys of type net.Addr and values of type *client.
	// net.Addr is a type that represents a network address.
	members map[net.Addr]*client
}

// Send a message to all the members of the room, except for the sender.
func (r *room) broadcast(sender *client, msg string) {
	for addr, m := range r.members {
		if addr != sender.conn.RemoteAddr() {
			m.msg(msg)
		}
	}
}
