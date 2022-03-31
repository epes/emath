package emath

import "encoding/binary"

func Int32ToBytes(i int32) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint32(b, uint32(i))
	return b
}
