package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

const RECV_BUF_LEN = 1024

func main() {
	service := ":8888"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	fmt.Println("Starting the server")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handlerClient(conn)
	}
}

func handlerClient(conn net.Conn) {
	buf := make([]byte, RECV_BUF_LEN)
	fmt.Println("Accepted the Connection :", conn.RemoteAddr())
	defer conn.Close()

	for {
		n, err := conn.Read(buf)
		switch err {
		case nil:
			conn.Write(buf[0:n])
			fmt.Println(string(buf[0:n]))
		case io.EOF:
			fmt.Printf("Warning: End of data: %s \n", err)
			return
		default:
			fmt.Printf("Error: Reading data : %s \n", err)
			return
		}
	}

	dayTime := time.Now().String()
	fmt.Print(dayTime)
	conn.Write([]byte(dayTime))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
