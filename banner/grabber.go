package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("couldn't do 3way handshake")
	}

	fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: localhost\r\n\r\n")

	d := make([]byte, 1024)

	n, err := conn.Read(d)
	if err != nil {
		log.Fatal("couldn't read")
	}
	data := string(d[:n])

	lines := strings.Split(data, "\n")
	fmt.Println(lines[0])
}
