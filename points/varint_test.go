package points

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestEncodeVarint(t *testing.T) {
	buf := make([]byte, 20)

	for i := 0; i < 75000; i++ {
		l := binary.PutUvarint(buf, uint64(i))
		b := encodeVarint(i)

		if bytes.Compare(buf[:l], b) != 0 {
			t.FailNow()
		}
	}
}
