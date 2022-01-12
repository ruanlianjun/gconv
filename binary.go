package gconv

import (
	"encoding/binary"
	"math"
)

func DecodeToInt64(b []byte) int64 {
	return int64(binary.LittleEndian.Uint64(fillZero(b, 8)))
}

func DecodeToFloat64(b []byte) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(fillZero(b, 8)))
}

// 不够的时候高位补零
func fillZero(b []byte, s int) []byte {
	if len(b) > s {
		return b[:s]
	}

	c := make([]byte, s)
	copy(c, b)
	return c
}
