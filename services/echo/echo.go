package main

import (
	"io"
	"log"
	"net"
)

// Handler should
func handleConnection(conn net.Conn) {
	io.Copy(conn, conn)
	conn.Close()
}

//TCP Based Echo Service
//
//   One echo service is defined as a connection based application on TCP.
//   A server listens for TCP connections on TCP port 7.  Once a
//   connection is established any data received is sent back.  This
//   continues until the calling user terminates the connection.
//

// Main loop for our echo service
func main() {
	ln, err := net.Listen("tcp", "0.0.0.0:7")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		defer conn.Close()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConnection(conn)
	}
}
