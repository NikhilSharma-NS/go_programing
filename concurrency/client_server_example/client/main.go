package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {

	connetion, err := net.Dial("tcp", "localhost:8081")

	if err != nil {
		log.Println(err)
	}
	defer connetion.Close()

	mustCopy(os.Stdout, connetion)

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Panicln(err)
	}
}
