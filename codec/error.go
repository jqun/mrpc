package codec

import "errors"

var (
	UnexpectedChecksumError     = errors.New("unexpected checksum")
	NotFoundCompressorError     = errors.New("not found compressor")
	InvalidSequenceError        = errors.New("invalid sequence number in response")
	CompressorTypeMismatchError = errors.New("request and response compressor type mismatch")
)
