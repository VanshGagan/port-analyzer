package main

import (
	"fmt"
	"net"
	"strings"
	"port-analyzer/utils"
	"time"
	"os"
)


//read the banner
func main() {


	var portNames map[int]string
	var target string
	portNames = utils.PortNames

	if len(os.Args) < 2 {
		target = "127.0.0.1"
	} else {
		target = os.Args[1]
	}

	for port, name := range portNames {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", target, port), 1*time.Second)
		if err != nil {
			continue
		}
		if port == 80 || port == 443{
			fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: %s\r\nConnection: close\r\n\r\n", target)
		}
		d:= make([]byte, 1024)

		conn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))

		n, err := conn.Read(d)
		if err != nil {
			continue
		}
		data := string(d[:n])

		lines := strings.Split(data, "\n")
		fmt.Printf("Banner of %s\n", name)
		fmt.Printf("%s\n\n", lines[0])
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
