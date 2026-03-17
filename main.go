package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	target_ip := "127.0.0.1"
	if len(os.Args) < 2 {
		fmt.Print("Usage: sudo go run main.go <port> \n")
	}
	port := os.Args[1]
	address := net.JoinHostPort(target_ip, fmt.Sprintf("%s", port))

	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err != nil {
		fmt.Printf("Port %s is closed \n", port)
		return
	}
	defer conn.Close()
	fmt.Printf("Port: %s is open \n", port)

}
