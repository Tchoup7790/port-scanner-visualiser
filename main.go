package main

import (
	"sync"

	"port-scanner-visualiser/config"
	"port-scanner-visualiser/scanner"
	"port-scanner-visualiser/ui"
)

func main() {
	host := config.GetFlags()

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

			isOpen, protocole := scanner.ScanPort(host, p)
			if isOpen {
				mu.Lock()
				openPortsMap[p] = protocole
				mu.Unlock()
			}
		}(port)
	}

	wg.Wait()

	ui.PrintResult(openPortsMap)
}
