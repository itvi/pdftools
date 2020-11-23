package util

import (
	"log"
	"net"
)

// GetLocalIP get ip of the server
func GetLocalIP() string {
	cnn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Println(err)
		return ""
	}
	defer cnn.Close()

	localAddr := cnn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}
