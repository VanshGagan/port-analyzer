package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var device = "lo"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: ./sniffer <port>")

	}
	port_parameter := os.Args[1]

	handle, err := pcap.OpenLive(device, 1600, false, pcap.BlockForever)
	if err != nil {
		log.Panicln("unable to open the handle")
	}
	defer handle.Close()

	filter := fmt.Sprintf("tcp and port %s", port_parameter)
	handle.SetBPFFilter(filter)

	packets := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packets.Packets() {
		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		ipLayer := packet.Layer(layers.LayerTypeIPv4)

		if tcpLayer != nil && ipLayer != nil {
			tcp, _ := tcpLayer.(*layers.TCP)
			ip, _ := ipLayer.(*layers.IPv4)

			flags := ""
			if tcp.SYN {
				flags += "SYN "
			}
			if tcp.ACK {
				flags += "ACK "
			}
			if tcp.FIN {
				flags += "FIN "
			}
			if tcp.RST {
				flags += "RST "
			}
			fmt.Printf("Sender: %s ----- Reciever: %s\npacket with flags: %s\n\n\n", ip.SrcIP, ip.DstIP, flags)
		}

	}
}
