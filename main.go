package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"port-analyzer/network"
	"port-analyzer/utils"
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
	s := time.Now()
	var portNames map[int]string
	portNames = utils.PortNames
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
	t := time.Now()
	elapsed := t.Sub(s)
	fmt.Print("--- scan finished ---\n")
	fmt.Printf("\nscanned in %.2f seconds\n", elapsed.Seconds()) // <- hier
	os.Exit(0)
}
