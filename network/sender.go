package network

import (
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func SendSYNPacket(target_ip string, port int) {
	//set src and dst ip to send the SYN packet
	var src_ip = net.ParseIP("127.0.0.1")
	var dst_ip = net.ParseIP(target_ip)

	//set src and dst ports to send the SYN packet
	var src_port = layers.TCPPort(51234)
	var dst_port = layers.TCPPort(port)

	//set layers of the packet
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

	//create a RAW socket
	conn, err := net.ListenPacket("ip4:tcp", "0.0.0.0")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//ignore the other header of the TCP packet
	dataToSend := buf.Bytes()[20:]
	//log.Println("writing request")

	//send the packet
	if _, err := conn.WriteTo(dataToSend, &net.IPAddr{IP: dst_ip}); err != nil {
		log.Fatal(err)

	}
	//wait 1 seconds to not lose the packet
	//time.Sleep(500 * time.Millisecond)

}
