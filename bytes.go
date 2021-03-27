package emath

import "encoding/binary"

func IsNilBytes(b []byte) bool {
	for _, v := range b {
		if v != 0 {
			return false
		}
	}

	return true
}

func BytesToUint16(b []byte) uint16 {
	return binary.BigEndian.Uint16(b)
}

func BytesToInt16(b []byte) int16 {
	return int16(BytesToUint16(b))
}

func BytesToUint32(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}

func BytesToInt32(b []byte) int32 {
	return int32(BytesToUint32(b))
}
