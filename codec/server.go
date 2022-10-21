package codec

import (
	"bufio"
	"io"
	"mrpc/compressor"
	"mrpc/header"
	"mrpc/serializer"
	"net/rpc"
	"sync"
)

type reqCtx struct {
	requestId    uint64
	compressType compressor.CompressType
}

type serverCodec struct {
	r io.Reader
	w io.Writer
	c io.Closer

	request    header.RequestHeader
	serializer serializer.Serializer
	mutex      sync.Mutex
	seq        uint64
	pending    map[uint64]*reqCtx
}

func NewServerCodec(conn io.ReadWriteCloser, serializer serializer.Serializer) rpc.ServerCodec {
	return &serverCodec{
		r:          bufio.NewReader(conn),
		w:          bufio.NewWriter(conn),
		c:          conn,
		serializer: serializer,
		pending:    make(map[uint64]*reqCtx),
	}
}

func (s *serverCodec) ReadRequestHeader(r *rpc.Request) error {
	// todo
	return nil
}

func (s *serverCodec) ReadRequestBody(interface{}) error {
	// todo
	return nil
}

func (s *serverCodec) WriteResponse(r *rpc.Response, data interface{}) error {
	// todo
	return nil
}

func (s *serverCodec) Close() error {
	// todo
	return nil
}
