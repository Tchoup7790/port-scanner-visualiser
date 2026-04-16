// Package config
package config

import "flag"

func GetFlags() string {
	hostPtr := flag.String("host", "google.com", "host to scan")
	flag.Parse()
	return *hostPtr
}
