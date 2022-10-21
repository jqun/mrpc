package codec

import (
	"bufio"
	"io"
	"mrpc/compressor"
	"mrpc/serializer"
	"net/rpc"
	"sync"
)

// 通过官方提供的codec接口，实现rpc序列化协议
type clientCodec struct {
	r io.Reader
	w io.Writer
	c io.Closer

	compressor compressor.CompressType // 数据如何压缩
	serializer serializer.Serializer   // 数据如何编解码
	mutex      sync.Mutex
	pending    map[uint64]string
}

func NewClientCodec(conn io.ReadWriteCloser, compressType compressor.CompressType, serializer serializer.Serializer) rpc.ClientCodec {
	return &clientCodec{
		r:          bufio.NewReader(conn),
		w:          bufio.NewWriter(conn),
		c:          conn,
		compressor: compressType,
		serializer: serializer,
		pending:    make(map[uint64]string),
	}
}

func (c *clientCodec) WriteRequest(r *rpc.Request, i interface{}) error {
	// todo
	return nil
}

func (c *clientCodec) ReadResponseHeader(response *rpc.Response) error {
	// todo
	return nil
}

func (c *clientCodec) ReadResponseBody(_ interface{}) error {
	// todo
	return nil
}

func (c *clientCodec) Close() error {
	// todo
	return nil
}
