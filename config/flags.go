// Package config
package config

import "flag"

func GetFlags() (string, int, int, int) {
	hostPtr := flag.String("host", "scanme.nmap.org", "host to scan")
	startPtr := flag.Int("start", 1, "start port")
	endPtr := flag.Int("end", 1024, "end port")
	concurrencyPtr := flag.Int("concurrency", 100, "number of concurrent goroutines")
	flag.Parse()
	return *hostPtr, *startPtr, *endPtr, *concurrencyPtr
}
