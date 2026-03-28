package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"port-analyzer/network"
	"sync"
	"time"
)

var device = "any"

const (
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorReset  = "\033[0m"
)

func worker(target string, jobs chan int, wg *sync.WaitGroup, conn net.PacketConn) {
	defer wg.Done()

	for port := range jobs {
		network.SendSYNPacket(target, port, conn)

		fmt.Printf("... scanner is on port %d ...\n", port)

		time.Sleep(100 * time.Millisecond)
	}
}

func main() {

	var portNames = map[int]string{
		20:    "FTP Data",
		21:    "FTP",
		22:    "SSH",
		23:    "Telnet",
		25:    "SMTP",
		53:    "DNS",
		80:    "HTTP",
		110:   "POP3",
		119:   "NNTP",
		123:   "NTP",
		137:   "NetBIOS",
		139:   "NetBIOS",
		143:   "IMAP",
		161:   "SNMP",
		179:   "BGP",
		389:   "LDAP",
		443:   "HTTPS",
		445:   "SMB",
		465:   "SMTPS",
		500:   "ISAKMP",
		587:   "SMTP",
		631:   "CUPS",
		993:   "IMAPS",
		995:   "POP3S",
		1433:  "MSSQL",
		1521:  "Oracle",
		1723:  "PPTP",
		1883:  "MQTT",
		2049:  "NFS",
		2083:  "cPanel",
		2181:  "Zookeeper",
		2375:  "Docker",
		3000:  "Dev Server",
		3306:  "MySQL",
		3389:  "RDP",
		4444:  "Metasploit",
		5000:  "Flask",
		5432:  "PostgreSQL",
		5601:  "Kibana",
		5672:  "RabbitMQ",
		5900:  "VNC",
		6379:  "Redis",
		6443:  "Kubernetes",
		7001:  "WebLogic",
		7474:  "Neo4j",
		7687:  "Neo4j Bolt",
		8000:  "HTTP",
		8080:  "HTTP",
		8443:  "HTTPS",
		9000:  "Dev Server",
		9090:  "Prometheus",
		9200:  "Elasticsearch",
		27017: "MongoDB",
	}
	var target string

	if len(os.Args) < 2 {
		target = "127.0.0.1"
	} else {
		target = os.Args[1]
	}
	conn, err := net.ListenPacket("ip4:tcp", "0.0.0.0")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	jobs := make(chan int)
	results := make(chan int)

	var wg sync.WaitGroup

	go network.Sniffer(device, results, target)
	time.Sleep(1 * time.Second)

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go worker(target, jobs, &wg, conn)
	}

	for port := range portNames {
		jobs <- port
	}
	close(jobs)

	go func() {
		wg.Wait()
		time.Sleep(2 * time.Second)
		close(results)
	}()

	seen := make(map[int]bool)

	time.Sleep(2 * time.Second)

	for res := range results {
		name, exists := portNames[res]
		if seen[res] {
			continue
		}
		seen[res] = true
		if exists {
			fmt.Printf("\n%s┌──────────────────────────────────────────┐%s\n", ColorGreen, ColorReset)
			fmt.Printf("│  %s[FOUND]%s Port: %-5d  Name: %-13s │\n", ColorGreen, ColorReset, res, name)
			fmt.Printf("%s└──────────────────────────────────────────┐%s\n", ColorGreen, ColorReset)
		} else {
			fmt.Printf("\n%s┌──────────────────────────────────────────┐%s\n", ColorGreen, ColorReset)
			fmt.Printf("│  %s[FOUND]%s Port: %-5d              │\n", ColorGreen, ColorReset, res)
			fmt.Printf("%s└──────────────────────────────────────────┐%s\n", ColorGreen, ColorReset)
		}

	}
	fmt.Print("--- scan finished ---\n")
	os.Exit(0)
}
