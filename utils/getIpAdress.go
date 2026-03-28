package utils

import (
	"log"
	"net"
	"os"
)

func GetIP() net.IP {
	hostname, _ := os.Hostname()
	addrs, err := net.LookupIP(hostname)

	if err != nil {
		log.Println("Failed to detect machine host name. ", err.Error())
		return nil
	}
	//log.Println("All Addrs: ", addrs)

	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			return ipv4
		}
	}
	log.Print("Couldn't find ipv4 address.")
	return nil
}
