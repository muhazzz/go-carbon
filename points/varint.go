package points

import (
	"encoding/binary"
)

var varintCacheSize int = 1000 // < should be lower then 16k
var varintCache []byte

func init() {
	varintCache = make([]byte, varintCacheSize*2-128)
	for i := 0; i < varintCacheSize; i++ {
		if i < 128 {
			binary.PutUvarint(varintCache[i:], uint64(i))
		} else {
			binary.PutUvarint(varintCache[i*2-128:], uint64(i))
		}
	}
}

func encodeVarint(value int) []byte {
	if value < varintCacheSize {
		if value < 128 {
			return varintCache[value : value+1]
		}
		return varintCache[value*2-128 : value+2]
	}

	var buf [10]byte
	l := binary.PutUvarint(buf[:], uint64(value))
	return buf[:l]
}
