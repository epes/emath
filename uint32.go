package emath

import (
	"encoding/binary"
)

func Uint32ToBytes(u uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, u)
	return b
}

func Uint32Max(a uint32, b uint32) uint32 {
	if a > b {
		return a
	}

	return b
}

func Uint32Min(a uint32, b uint32) uint32 {
	if a > b {
		return b
	}

	return a
}
