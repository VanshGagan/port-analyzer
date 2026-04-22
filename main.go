package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"port-analyzer/banner"
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

// the work of each worker
func worker(target string, jobs chan int, wg *sync.WaitGroup, conn net.PacketConn) {
	defer wg.Done()

	for port := range jobs {
		network.SendSYNPacket(target, port, conn)

		//fmt.Printf("... scanner is on port %d ...\n", port)

		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	//start timer and loading animation
	s := time.Now()
	go utils.Loading()

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

	var workerWg sync.WaitGroup

	go network.Sniffer(device, results, target)
	time.Sleep(1 * time.Second)

	for i := 1; i <= 10; i++ {
		workerWg.Add(1)
		go worker(target, jobs, &workerWg, conn)
	}

	for port := range portNames {
		jobs <- port
	}
	close(jobs)

	go func() {
		workerWg.Wait()
		time.Sleep(4 * time.Second)
		close(results)
	}()

	var bannerWG sync.WaitGroup
	seen := make(map[int]bool)

	time.Sleep(2 * time.Second)

	utils.KeepLoading = false
	time.Sleep(200 * time.Millisecond)
	fmt.Print("\r                          \r")

	for res := range results {
		name, exists := portNames[res]
		if seen[res] {
			continue
		}
		seen[res] = true
		if exists {
			fmt.Printf("\n%s┌──────────────────────────────────────────┐%s\n", ColorGreen, ColorReset)
			fmt.Printf("│  %s[FOUND]%s Port: %-5d  Name: %-13s│\n", ColorGreen, ColorReset, res, name)
			fmt.Printf("%s└──────────────────────────────────────────┐%s\n", ColorGreen, ColorReset)
		} else {
			fmt.Printf("\n%s┌──────────────────────────────────────────┐%s\n", ColorGreen, ColorReset)
			fmt.Printf("│  %s[FOUND]%s Port: %-5d              │\n", ColorGreen, ColorReset, res)
			fmt.Printf("%s└──────────────────────────────────────────┐%s\n", ColorGreen, ColorReset)
		}
		if len(os.Args) > 2 && os.Args[2] == "-b" {
			bannerWG.Add(1)
			go banner.Banner_grabber(res, target, &bannerWG)
		}

	}
	bannerWG.Wait()

	utils.KeepLoading = false
	time.Sleep(250 * time.Millisecond)
	fmt.Print("\r          \r")

	t := time.Now()
	elapsed := t.Sub(s)
	if len(seen) < 1 {
		fmt.Printf("\n%s┌──────────────────────────┐%s\n", ColorGreen, ColorReset)
		fmt.Printf("│  %sNo open ports found!    %s│\n", ColorYellow, ColorReset)
		fmt.Printf("%s└──────────────────────────┘%s\n", ColorGreen, ColorReset)

	}
	fmt.Print("--- scan finished ---\n")
	fmt.Printf("\nscanned in %.2f seconds\n", elapsed.Seconds())

	os.Exit(0)
}
