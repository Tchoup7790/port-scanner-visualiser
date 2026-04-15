package main

import (
	"flag"
	"fmt"
	"net"
	"sort"
	"sync"
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

			if scanPort(host, p) {
				mu.Lock()
				openPortsMap[p] = "OPEN"
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

	for _, k := range keys {
		fmt.Println(k, openPortsMap[k])
	}
}
