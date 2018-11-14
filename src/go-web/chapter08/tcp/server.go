package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	service := "127.0.0.1:1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError1(err)

	listen, err := net.ListenTCP("tcp", tcpAddr)
	checkError1(err)

	for {
		// 进行监听
		conn, err := listen.Accept()
		if err != nil {
			continue
		}

		go func(conn net.Conn) {
			// set 2 minutes timeout
			conn.SetReadDeadline(time.Now().Add(2 * time.Minute))

			request := make([]byte, 128)
			defer conn.Close()
			for {
				readLen, err := conn.Read(request)
				if err != nil {
					fmt.Println(err)
					break
				}

				if readLen == 0 {
					break
				} else if strings.TrimSpace(string(request[:readLen])) == "timestamp" {
					daytime := strconv.FormatInt(time.Now().Unix(), 10)
					conn.Write([]byte(daytime))
				} else {
					daytime := time.Now().String()
					conn.Write([]byte(daytime))
				}

				request = make([]byte, 128)
			}
		}(conn)
	}
}

func checkError1(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
