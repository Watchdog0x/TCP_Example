package main

import (
	"fmt"
	"net"
)

func main() {
	// create a listener on localhost:8080
	addr, err := net.ResolveTCPAddr("tcp", ":8080")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	fmt.Printf("%#v, %v\n", addr, addr)

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println("Error creating listener:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Listening on", listener.Addr())

	for {
		// accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// handle the connection in a separate goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println(conn.RemoteAddr().String())
	conn.Write([]byte("Hello Darkness\n"))

	// read data from the connection
	// and do something with it
	// ...
}
