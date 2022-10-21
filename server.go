package mrpc

import (
	"log"
	"mrpc/codec"
	"mrpc/serializer"
	"net"
	"net/rpc"
)

type Server struct {
	*rpc.Server
	serializer.Serializer
}

func NewServer(opts ...Option) *Server {
	options := options{
		serializer: serializer.Proto,
	}

	for _, option := range opts {
		option(&options)
	}

	return &Server{
		Server:     &rpc.Server{},
		Serializer: options.serializer,
	}
}

func (s *Server) Register(rcvr interface{}) error {
	return s.Server.Register(rcvr)
}

func (s *Server) RegisterName(name string, rcvr interface{}) error {
	return s.Server.RegisterName(name, rcvr)
}

func (s *Server) Serve(ln net.Listener) {
	log.Printf("mrpc started on: %v", ln.Addr().String())
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue // 此处可以考虑休眠Duration
		}
		go s.Server.ServeCodec(codec.NewServerCodec(conn, s.Serializer))
	}
}
