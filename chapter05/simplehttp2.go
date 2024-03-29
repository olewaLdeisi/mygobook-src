package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	// 解析地址和端口号
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	// 建立连接
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))

	os.Exit(0)
}
