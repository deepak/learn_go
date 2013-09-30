// needed to create a CLI ?
package main

import (
	"net"
	"io"
	"log"
)

// NOTE:
// log.Fatal will log the message and calls `os.Exit(1)`

// TODO:
// SERVER_URL can be an ENV variable
// optionally can echo to a file/stream
// TDD ?
// io.copy is not safe for concurrency ?

const (
	SERVER_URL = "0.0.0.0:7777"
)

func main() {
	listener, err := net.Listen("tcp", SERVER_URL)
	log.Printf("listening on " + SERVER_URL)
	
	if err != nil {
		log.Fatal(err)
	}
	
	// http://golang.org/doc/effective_go.html#defer
	// `listener.Close()` will be called finally
	defer func(listener net.Listener) {
		log.Printf("closing connection")
		listener.Close()
	}(listener)
	
	for {
		conn, err := listener.Accept()
		log.Printf("accepted connection")
		if err != nil {
			log.Fatal(err)
		}
		
		go echoInput(conn)
	}
}

// created once per telnet session
// http://golang.org/pkg/net/#Conn
func echoInput(conn net.Conn) {
	log.Printf("handeling connection")
	// Echo all incoming data.
	io.Copy(conn, conn)
	// Shut down the connection.
	conn.Close()
}
