package network

import (
	"fmt"
	"log"
	"port-analyzer/utils"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func Sniffer(device string, results chan int, target_ip string) {

	//open handle
	handle, err := pcap.OpenLive(device, 1600, false, pcap.BlockForever)
	if err != nil {
		log.Panicln("unable to open the handle")
	}
	defer handle.Close()

	//set tcp filter -> we wan't to catch tcp packets
	//filter := fmt.Sprintf("tcp and src host %s", target_ip)
	filter := fmt.Sprintf("tcp and dst host %s", utils.GetIP().String())
	handle.SetBPFFilter(filter)
	var os_detected bool = false
	var os string

	packets := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packets.Packets() {
		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		ipLayer := packet.Layer(layers.LayerTypeIPv4)

		if tcpLayer != nil && ipLayer != nil {
			tcp, _ := tcpLayer.(*layers.TCP)
			ip, _ := ipLayer.(*layers.IPv4)
			//fmt.Printf("FROM PORT: %d\n", tcp.SrcPort)

			//detect OS (ttl)
			if os_detected == false && ip.TTL > 0 {
				fmt.Print("\r              \r")
				fmt.Printf("\nTTL is: %d\n", ip.TTL)
				os = utils.DetectOs(ip.TTL)
				fmt.Printf("\nOperating system is probably: %s\n", os)
				os_detected = true
			}

			//if we have syn-ack -> write to results
			if tcp.SYN && tcp.ACK {
				openPort := tcp.SrcPort
				results <- int(openPort)
				fmt.Printf("\nWindow size is: %d\n", tcp.Window)

			}

		}

	}
}
