package header

import (
	"encoding/binary"
	"errors"
	"mrpc/compressor"
	"sync"
)

const (
	// MaxHeaderSize = 2 + 10 + 10 + 10 + 4 (10 refer to binary.MaxVarintLen64)
	MaxHeaderSize = 36

	Uint32Size = 4
	Uint16Size = 2
)

var UnmarshalError = errors.New("an error occurred in Unmarshal")

// RequestHeader request header structure looks like:
// +--------------+----------------+----------+------------+----------+
// | CompressType |      Method    |    ID    | RequestLen | Checksum |
// +--------------+----------------+----------+------------+----------+
// |    uint16    | uvarint+string |  uvarint |   uvarint  |  uint32  |
// +--------------+----------------+----------+------------+----------+
type RequestHeader struct {
	sync.RWMutex
	CompressType compressor.CompressType
	Method       string
	Id           uint64
	RequestLen   uint32
	Checksum     uint32
}

func (r *RequestHeader) Marshal() []byte {
	r.RLock()
	defer r.RUnlock()

	// 封装压缩类型
	idx := 0
	header := make([]byte, MaxHeaderSize+len(r.Method))
	binary.LittleEndian.PutUint16(header[idx:], uint16(r.CompressType))
	idx += Uint16Size

	idx += writeString(header[idx:], r.Method)
	idx += binary.PutUvarint(header[idx:], r.Id)
	idx += binary.PutUvarint(header[idx:], uint64(r.RequestLen))

	binary.LittleEndian.PutUint32(header[idx:], r.Checksum)
	idx += Uint32Size
	return header[:idx]
}

func (r *RequestHeader) Unmarshal(data []byte) (err error) {
	r.Lock()
	defer r.Unlock()

	if len(data) == 0 {
		return UnmarshalError
	}

	defer func() {
		if e := recover(); e == nil {
			err = UnmarshalError
		}
	}()

	idx, size := 0, 0
	r.CompressType = compressor.CompressType(binary.LittleEndian.Uint16(data[idx:]))
	idx += Uint16Size

	r.Method, size = readString(data[idx:])
	idx += size

	r.Id, size = binary.Uvarint(data[idx:])
	idx += size

	length, size := binary.Uvarint(data[idx:])
	r.RequestLen = uint32(length)
	idx += size

	r.Checksum = binary.LittleEndian.Uint32(data[idx:])
	return
}

func (r *RequestHeader) GetCompressType() compressor.CompressType {
	r.RLock()
	defer r.RUnlock()

	return r.CompressType
}

func (r *RequestHeader) ResetHeader() {
	r.Lock()
	defer r.Unlock()

	r.Id = 0
	r.Checksum = 0
	r.Method = ""
	r.CompressType = 0
	r.RequestLen = 0
}

type ResponseHeader struct {
	// todo
}

func writeString(data []byte, str string) int {
	idx := 0
	idx += binary.PutUvarint(data, uint64(len(str)))
	copy(data[idx:], str)
	idx += len(str)
	return idx
}

func readString(data []byte) (string, int) {
	idx := 0
	length, size := binary.Uvarint(data)
	idx += size
	str := string(data[idx : idx+int(length)])
	idx += len(str)
	return str, idx
}
