package serializer

type Serializer interface {
	Marshall(message interface{}) ([]byte, error)
	Unmarshall(data []byte, message interface{}) error
}
