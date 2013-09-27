package main
 
import (
	"net"
	"os"
)

func main() {
	println("Starting the server on 0.0.0.0:7777")
	listener, err := net.Listen("tcp", "0.0.0.0:7777")
	println("listening")
	if err != nil {
		os.Exit(1)
	}
	
	conn, err := listener.Accept()
	println("accepting connection")
	if err != nil {
		os.Exit(1)
	}
	println(conn)
}
 
