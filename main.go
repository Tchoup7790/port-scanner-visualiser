package main

import (
	"fmt"
	"net"
	"time"
)

func scanPort(host string, port int) bool {
	address := net.JoinHostPort(host, fmt.Sprintf("%d", port))
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		return false
	}

	_ = conn.Close()
	return true
}

func main() {
	host := "scanme.nmap.org"

	for port := 1; port <= 1024; port++ {
		fmt.Printf("Scanning port %d...\r", port)
		if scanPort(host, port) {
			fmt.Printf("Port %d : OPEN\n", port)
		}
	}
}
