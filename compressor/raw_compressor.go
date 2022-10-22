package compressor

type RawCompressor struct {}

// means:do nothing

func (_ RawCompressor) Zip(data []byte) ([]byte, error) {
	return data, nil
}

func (_ RawCompressor) Unzip(data []byte) ([]byte, error) {
	return data, nil
}



