package mrpc

import (
	"mrpc/serializer"
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
