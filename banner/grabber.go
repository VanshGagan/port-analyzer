package main

import (
	"fmt"
	"net"
	"strings"
	"port-analyzer/utils"
	"time"
)


//read the banner of http
func main() {


	var portNames map[int]string
	portNames = utils.PortNames

	for port, name := range portNames {
		conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", port))
		if err != nil {
			continue
		}
		d:= make([]byte, 1024)

		conn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))

		n, err := conn.Read(d)
		if err != nil {
			continue
		}
		data := string(d[:n])

		lines := strings.Split(data, "\n")
		fmt.Println(lines[0])
		fmt.Printf("%s", name)
		conn.Close()
	}

/*
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


	start ssh port for testing
	sudo systemctl start sshd

	close ssh
	sudo systemctl status sshd

*/
}
