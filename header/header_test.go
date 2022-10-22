package header

import (
	"fmt"
	"testing"
)

func TestResponseHeader_Marshal(t *testing.T) {
	r := &ResponseHeader{
		CompressType: 1,
		Id:           1,
		Error:        "aaa",
		ResponseLen:  2,
		Checksum:     100,
	}

	data := r.Marshal()

	r2 := &ResponseHeader{}
	_ = r2.Unmarshal(data)
	fmt.Println(r2)
}
