package banner

import (
	"crypto/tls"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

// read the banner
func Banner_grabber(port int, target string, wg *sync.WaitGroup) {
	defer wg.Done()
	var conn net.Conn
	var err error

	if port == 443 {
		dial_conf := &net.Dialer{Timeout: 1 * time.Second}
		tls_conf := &tls.Config{InsecureSkipVerify: true}

		conn, err = tls.DialWithDialer(dial_conf, "tcp", fmt.Sprintf("%s:%d", target, port), tls_conf)
		if err != nil {
			return
		}
	} else {

		conn, err = net.DialTimeout("tcp", fmt.Sprintf("%s:%d", target, port), 1*time.Second)
		if err != nil {
			return
		}
	}

	if port == 80 || port == 8000 || port == 8080 || port == 443 {
		fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: %s\r\nConnection: close\r\n\r\n", target)
	}

	d := make([]byte, 1024)

	conn.SetReadDeadline(time.Now().Add(1500 * time.Millisecond))

	n, err := conn.Read(d)
	if err != nil {
		return
	}
	data := string(d[:n])

	lines := strings.Split(data, "\n")
	fmt.Printf("\n\nBanner of %d\n", port)
	fmt.Printf("%s\n\n", lines[0])
	conn.Close()
}
