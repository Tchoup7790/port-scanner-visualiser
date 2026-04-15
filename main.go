package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"sort"
	"strings"
	"sync"
	"time"
)

func scanPort(host string, port int) (bool, string) {
	dialer := &net.Dialer{
		Timeout: 1 * time.Second,
	}

	address := net.JoinHostPort(host, fmt.Sprintf("%d", port))

	var conn net.Conn
	var err error

	conn, err = tls.DialWithDialer(dialer, "tcp", address, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		conn, err = net.DialTimeout("tcp", address, 1*time.Second)
	}
	if err != nil {
		return false, ""
	}

	_ = conn.SetReadDeadline(time.Now().Add(2 * time.Second))

	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)

	protocole := string(buf[:n])

	if protocole == "" {
		_, _ = conn.Write([]byte("HEAD / HTTP/1.0\r\nHost: " + host + "\r\n\r\n"))

		_ = conn.SetReadDeadline(time.Now().Add(2 * time.Second))

		buf := make([]byte, 1024)
		n, _ := conn.Read(buf)

		protocole = strings.Split(string(buf[:n]), " ")[0]
	}

	if protocole == "" {
		protocole = "undefined"
	}

	_, isTLS := conn.(*tls.Conn)
	if isTLS && protocole != "undefined" {
		protocole = "HTTPS"
	}

	_ = conn.Close()

	return true, string(protocole)
}

func main() {
	hostPtr := flag.String("host", "scanme.nmap.org", "host to scan")
	flag.Parse()
	host := *hostPtr

	var wg sync.WaitGroup
	sem := make(chan struct{}, 100)

	var mu sync.Mutex

	openPortsMap := map[int]string{}

	for port := 1; port <= 1024; port++ {
		wg.Add(1)
		sem <- struct{}{}

		go func(p int) {
			defer wg.Done()
			defer func() { <-sem }()

			isOpen, protocole := scanPort(host, p)
			if isOpen {
				mu.Lock()
				openPortsMap[p] = protocole
				mu.Unlock()
			}
		}(port)
	}

	wg.Wait()

	keys := make([]int, 0, len(openPortsMap))
	for k := range openPortsMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	fmt.Println("PORT", "PROTOCOLE")
	for _, k := range keys {
		fmt.Println(k, openPortsMap[k])
	}
}
