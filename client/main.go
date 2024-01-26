package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()

	fmt.Fprintln(conn, "I dialed you!")
}
