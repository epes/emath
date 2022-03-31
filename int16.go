package emath

import "encoding/binary"

func Int16ToBytes(i int16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(i))
	return b
}
