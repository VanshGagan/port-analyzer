package banner

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

// read the banner
func Banner_grabber(port int, target string, wg *sync.WaitGroup) {
	defer wg.Done()

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", target, port), 1*time.Second)
	if err != nil {
		return
	}
	if port == 80 || port == 8000 || port == 8080 {
		fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: %s\r\nConnection: close\r\n\r\n", target)
	}
	d := make([]byte, 1024)

	conn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))

	n, err := conn.Read(d)
	if err != nil {
		return
	}
	data := string(d[:n])

	lines := strings.Split(data, "\n")
	fmt.Printf("Banner of %d\n", port)
	fmt.Printf("%s\n\n", lines[0])
	conn.Close()
}
