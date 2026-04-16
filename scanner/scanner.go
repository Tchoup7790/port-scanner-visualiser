// Package scanner
package scanner

import (
	"crypto/tls"
	"fmt"
	"net"
	"strings"
	"time"
)

func ScanPort(host string, port int) (bool, string) {
	dialer := &net.Dialer{
		Timeout: 1 * time.Second,
	}

	address := net.JoinHostPort(host, fmt.Sprintf("%d", port))

	var conn net.Conn
	var err error

	conn, err = tls.DialWithDialer(dialer, "tcp", address, &tls.Config{
		// InsecureSkipVerify permet de scanner des hosts avec des certificats auto-signés
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
