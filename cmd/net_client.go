package cmd

import (
	"net"
	"strconv"
)

var (
	UDPConn, UDPConnErr = net.DialUDP("udp", nil, UDPAddr)
	UDPAddr             = ""
)

func SetServerAddr(UDPHost string, UDPPort int) {
	UDPAddr = UDPHost + ":" + strconv.Itoa(UDPPort)
}
