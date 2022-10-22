package codec

import (
	"bufio"
	"hash/crc32"
	"io"
	"mrpc/compressor"
	"mrpc/header"
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
	response   header.ResponseHeader
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

func (c *clientCodec) WriteRequest(r *rpc.Request, param interface{}) error {
	c.mutex.Lock()
	c.pending[r.Seq] = r.ServiceMethod
	c.mutex.Unlock()

	if _, ok := compressor.Compressors[c.compressor]; !ok {
		return NotFoundCompressorError
	}

	reqBody, err := c.serializer.Marshall(param)
	if err != nil {
		return err
	}

	compressedReqBody, err := compressor.Compressors[c.compressor].Zip(reqBody)
	if err != nil {
		return err
	}

	h := header.RequestPool.Get().(*header.RequestHeader)
	defer func() {
		h.ResetHeader()
		header.RequestPool.Put(h)
	}()

	h.Id = r.Seq
	h.Method = r.ServiceMethod
	h.RequestLen = uint32(len(compressedReqBody))
	h.CompressType = c.compressor
	h.Checksum = crc32.ChecksumIEEE(compressedReqBody)

	if err = sendFrame(c.w, h.Marshal()); err != nil {
		return err
	}

	if err = write(c.w, compressedReqBody); err != nil {
		return err
	}

	c.w.(*bufio.Writer).Flush()
	return nil
}

func (c *clientCodec) ReadResponseHeader(r *rpc.Response) error {
	c.response.ResetHeader()
	data, err := recvFrame(c.r)
	if err != nil {
		return err
	}

	err = c.response.Unmarshal(data)
	if err != nil {
		return err
	}

	c.mutex.Lock()
	r.Seq = c.response.Id
	r.Error = c.response.Error
	r.ServiceMethod = c.pending[r.Seq]
	delete(c.pending, r.Seq)
	c.mutex.Unlock()
	return nil
}

func (c *clientCodec) ReadResponseBody(param interface{}) error {
	if param == nil {
		if c.response.ResponseLen != 0 {
			if err := read(c.r, make([]byte, c.response.ResponseLen)); err != nil {
				return err
			}
		}
		return nil
	}

	respBody := make([]byte, c.response.ResponseLen)
	err := read(c.r, respBody)
	if err != nil {
		return err
	}

	if c.response.Checksum != 0 {
		if crc32.ChecksumIEEE(respBody) != c.response.Checksum {
			return UnexpectedChecksumError
		}
	}

	if c.response.GetCompressType() != c.compressor {
		return CompressorTypeMismatchError
	}

	resp, err := compressor.Compressors[c.response.GetCompressType()].Unzip(respBody)
	if err != nil {
		return err
	}
	return c.serializer.Unmarshall(resp, param)
}

func (c *clientCodec) Close() error {
	return c.c.Close()
}
