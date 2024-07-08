package main

import (
	"fmt"
	"net"
	"time"
)

func Check(port string, destination string) string {
	address := destination + ":" + port
	timeout := time.Duration(15 * time.Second)
	conn, err := net.DialTimeout("tcp", address, timeout)

	var status string

	if err != nil {
		status = fmt.Sprintf("[DOWN]%v is unreachable, \n Error: %v", destination, err)

	} else {
		status = fmt.Sprintf("[CLEAR], %v is reachable, \n From: %v\n To:%v", destination, conn.LocalAddr(), conn.RemoteAddr())
	}
	return status
}
