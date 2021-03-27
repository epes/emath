package emath

import "encoding/binary"

func Int16ToBytes(i int16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(i))
	return b
}

func Int16Max(a int16, b int16) int16 {
	if a > b {
		return a
	}

	return b
}

func Int16Min(a int16, b int16) int16 {
	if a > b {
		return a
	}

	return b
}
