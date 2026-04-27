package network

import (
	"log"
	"net"

	"port-analyzer/utils"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func SendSYNPacket(target_ip string, port int, conn net.PacketConn) {

	//set src and dst ip to send the SYN packet
	var src_ip = utils.GetIP()
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

	tcp.Options = []layers.TCPOption{
		{
			OptionType:   layers.TCPOptionKindMSS,
			OptionLength: 4,
			OptionData:   []byte{0x05, 0xb4}, // 1460 Bytes
		},
		{
			OptionType:   layers.TCPOptionKindSACKPermitted,
			OptionLength: 2,
		},
		{
			OptionType:   layers.TCPOptionKindWindowScale,
			OptionLength: 3,
			OptionData:   []byte{0x07}, // Multiplikator
		},
		{
			OptionType: layers.TCPOptionKindNop, // NOP für das Alignment
		},
		{
			OptionType:   layers.TCPOptionKindTimestamps,
			OptionLength: 10,
			OptionData:   make([]byte, 8), // 8 Bytes für TS Value und Echo Reply
		},
	}

	//compute checksum
	tcp.SetNetworkLayerForChecksum(ip)

	//build the packet
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		ComputeChecksums: true,
		FixLengths:       true,
	}

	if err := gopacket.SerializeLayers(buf, opts, tcp); err != nil {
		log.Fatal(err)
	}

	//create a RAW socket

	dataToSend := buf.Bytes()

	//send the packet
	if _, err := conn.WriteTo(dataToSend, &net.IPAddr{IP: dst_ip}); err != nil {
		log.Fatal(err)

	}

}
