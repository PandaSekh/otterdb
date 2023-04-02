package utils

import (
	"fmt"
	"net"
)

func PrintLocalIp() {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println(err)
	}

	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)
	ipAddress := conn.LocalAddr().(*net.UDPAddr)
	fmt.Printf("Local IP is: %s\n", ipAddress)
}
