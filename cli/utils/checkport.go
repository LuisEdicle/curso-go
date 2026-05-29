package utils

import (
	"fmt"
	"net"
	"time"
)

func CheckPort(host string, ports []string) {

	for _, port := range ports {
		timeout := time.Second
		conn, err := net.DialTimeout(
			"tcp",
			net.JoinHostPort(host, port),
			timeout,
		)

		if err != nil {
			fmt.Printf("A porta %s está fechada. \n", port)
			continue
		}

		if conn != nil {
			fmt.Printf("A Porta %s está aberta\n", port)
			defer conn.Close()
		}
	}
}
