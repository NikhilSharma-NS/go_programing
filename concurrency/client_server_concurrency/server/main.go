package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {

	listner, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Print(err)
	}
	for {
		connection, err := listner.Accept()
		if err != nil {
			continue
		}
		go handleConnection(connection)
	}

}

func handleConnection(c net.Conn) {
	defer c.Close()

	for {
		_, err := io.WriteString(c, "Response from server \n")
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}

}
