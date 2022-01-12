package gconv

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

func Encode(values ...interface{}) []byte {
	buf := new(bytes.Buffer)
	for i, _ := range values {
		if values[i] == nil {
			return buf.Bytes()
		}

		switch val := values[i].(type) {
		case int:
			buf.Write(EncodeInt(val))
		case int8:
			buf.Write(EncodeInt8(val))
		case int16:
			buf.Write(EncodeInt16(val))
		case int32:
			buf.Write(EncodeInt32(val))
		case int64:
			buf.Write(EncodeInt64(val))
		case uint:
			buf.Write(EncodeUint(val))
		case uint8:
			buf.Write(EncodeUint8(val))
		case uint16:
			buf.Write(EncodeUint16(val))
		case uint32:
			buf.Write(EncodeUint32(val))
		case uint64:
			buf.Write(EncodeUint64(val))
		case bool:
			buf.Write(EncodeBool(val))
		case string:
			buf.Write(EncodeString(val))
		case []byte:
			buf.Write(val)
		case float32:
			buf.Write(EncodeFloat32(val))
		case float64:
			buf.Write(EncodeFloat64(val))
		default:
			if err := binary.Write(buf, binary.LittleEndian, val); err != nil {
				buf.Write(EncodeString(fmt.Sprintf("%v", val)))
			}
		}

	}
	return buf.Bytes()
}

func EncodeFloat64(f float64) []byte {
	bits := math.Float64bits(f)
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, bits)
	return b
}

func EncodeFloat32(f float32) []byte {
	bits := math.Float32bits(f)
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, bits)
	return b
}

func EncodeString(b string) []byte {
	return []byte(b)
}

func EncodeBool(b bool) []byte {
	if b {
		return []byte{1}
	}

	return []byte{0}
}

func EncodeInt(i int) []byte {
	if i <= math.MaxInt8 {
		return EncodeInt8(int8(i))
	}
	if i <= math.MaxInt16 {
		return EncodeInt16(int16(i))
	}

	if i <= math.MaxInt32 {
		return EncodeInt32(int32(i))
	}

	return EncodeInt64(int64(i))
}

func EncodeUint(i uint) []byte {
	if i <= math.MaxUint8 {
		return EncodeUint8(uint8(i))
	}
	if i <= math.MaxUint16 {
		return EncodeUint16(uint16(i))
	}
	if i <= math.MaxUint32 {
		return EncodeUint32(uint32(i))
	}
	return EncodeUint64(uint64(i))
}

func EncodeInt8(i int8) []byte {
	return []byte{byte(i)}
}

func EncodeUint8(i uint8) []byte {
	return []byte{i}
}

func EncodeInt16(i int16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, uint16(i))
	return b
}

func EncodeUint16(i uint16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, i)
	return b
}

func EncodeInt32(i int32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(i))
	return b
}

func EncodeUint32(i uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, i)
	return b
}

func EncodeInt64(i int64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(i))
	return b
}

func EncodeUint64(i uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, i)
	return b
}
