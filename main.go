package main

import (
	"flag"
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
	hostPtr := flag.String("host", "scanme.nmap.org", "a string")
	flag.Parse()
	host := *hostPtr

	for port := 1; port <= 1024; port++ {
		fmt.Printf("Scanning port %d...\r", port)
		if scanPort(host, port) {
			fmt.Printf("Port %d : OPEN\n", port)
		}
	}
}
