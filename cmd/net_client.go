package cmd

import (
	"net"
	"strconv"
)

var (
	UDPConn, UDPConnErr = net.Dial("udp", UDPAddr)
	UDPAddr             = ""
)

func SetServerAddr(UDPHost string, UDPPort int) {
	UDPAddr = UDPHost + ":" + strconv.Itoa(UDPPort)
}
