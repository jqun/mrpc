package main

import (
	"log"
	"mrpc"
	"mrpc/test/message"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":1000")
	if err != nil {
		log.Fatal(err)
	}
	server := mrpc.NewServer()
	server.RegisterName("ArithService", new(message.ArithService))
	server.Serve(ln)
}


