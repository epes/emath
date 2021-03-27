package emath

import "encoding/binary"

func Int32ToBytes(i int32) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint32(b, uint32(i))
	return b
}

func Int32Max(a int32, b int32) int32 {
	if a > b {
		return a
	}

	return b
}

func Int32Min(a int32, b int32) int32 {
	if a > b {
		return a
	}

	return b
}
