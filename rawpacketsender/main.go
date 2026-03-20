package main

import (
	"log"
	"net"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	var device = "lo"
	var src_ip = net.ParseIP("127.0.0.1")
	var dst_ip = net.ParseIP("127.0.0.1")

	var src_port = layers.TCPPort(51234)
	var dst_port = layers.TCPPort(631)

	handle, err := pcap.OpenLive(device, 1600, false, pcap.BlockForever)
	if err != nil {
		log.Panicln("unable to open the handle")
	}
	defer handle.Close()

	ip := &layers.IPv4{
		SrcIP:    src_ip,
		DstIP:    dst_ip,
		Protocol: layers.IPProtocolTCP,
	}

	tcp := &layers.TCP{
		SrcPort: src_port,
		DstPort: dst_port,
		Seq:     1105024978,
		SYN:     true,
		Window:  14600,
	}
	tcp.SetNetworkLayerForChecksum(ip)

	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		ComputeChecksums: true,
		FixLengths:       true,
	}

	if err := gopacket.SerializeLayers(buf, opts, ip, tcp); err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenPacket("ip4:tcp", "0.0.0.0")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	dataToSend := buf.Bytes()[20:]
	log.Println("writing request")
	if _, err := conn.WriteTo(dataToSend, &net.IPAddr{IP: dst_ip}); err != nil {
		log.Fatal(err)

	}
	time.Sleep(5 * time.Second)

}
