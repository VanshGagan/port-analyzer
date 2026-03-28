package network

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func Sniffer(device string, results chan int, target_ip string) {

	handle, err := pcap.OpenLive(device, 1600, false, pcap.BlockForever)
	if err != nil {
		log.Panicln("unable to open the handle")
	}
	defer handle.Close()

	filter := fmt.Sprintf("tcp and src host %s", target_ip)
	handle.SetBPFFilter(filter)

	packets := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packets.Packets() {
		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		ipLayer := packet.Layer(layers.LayerTypeIPv4)

		if tcpLayer != nil && ipLayer != nil {
			tcp, _ := tcpLayer.(*layers.TCP)
			//ip, _ := ipLayer.(*layers.IPv4)
			//fmt.Printf("FROM PORT: %d\n", tcp.SrcPort)

			if tcp.SYN && tcp.ACK {
				openPort := tcp.SrcPort
				results <- int(openPort)
			}
		}

	}
}
