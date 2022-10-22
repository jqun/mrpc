package main

import (
	"log"
	"mrpc"
	"mrpc/compressor"
	"mrpc/test/message"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":1000")
	if err != nil {
		log.Fatal(err)
	}

	client := mrpc.NewClient(conn, mrpc.WithCompress(compressor.Gzip))
	resq := message.ArithRequest{A: 20, B: 5}
	resp := message.ArithResponse{}
	err = client.Call("ArithService.Add", &resq, &resp)
	log.Printf("Arith.Add(%v, %v): %v ,Error: %v", resq.A, resq.B, resp.C, err)
}
