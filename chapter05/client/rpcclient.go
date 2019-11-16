package client

import (
	"fmt"
	"log"
	"mygobook-src/chapter05/server"
	"net/rpc"
)

func SychCall(serverAddress string) {
	// 建立连接
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := &server.Args{7, 8}
	var reply int
	// 同步调用
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
}

func AsychCall(serverAddress string) {
	// 建立连接
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := &server.Args{7, 8}
	quotient := new(server.Quotient)
	// 异步调用
	divCall := client.Go("Arith.Divide", args, &quotient, nil)
	<-divCall.Done
	fmt.Printf("Arith: %d/%d=%d, %d%%%d=%d", args.A, args.B, quotient.Quo, args.A, args.B, quotient.Rem)
}
