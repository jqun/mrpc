package serializer

type ProtoSerializer struct{}

var Proto = ProtoSerializer{}

func (_ ProtoSerializer) Marshall(message interface{}) ([]byte, error) {
	return nil, nil
}

func (_ ProtoSerializer) Unmarshall(data []byte, message interface{}) error {
	// todo
	return nil
}