package gconv

import (
	"math"
	"reflect"
)

func Bytes(any interface{}) []byte {
	bytes, _ := BytesE(any)
	return bytes
}

func BytesE(any interface{}) ([]byte, error) {
	switch val := any.(type) {
	case string:
		return []byte(val), nil
	case []byte:
		return val, nil
	default:
		var (
			reflectVal  = reflect.ValueOf(any)
			reflectKind = reflectVal.Kind()
		)

		if reflectKind == reflect.Ptr {
			reflectVal = reflectVal.Elem()
			reflectKind = reflectVal.Kind()
		}

		switch reflectKind {
		case reflect.Array, reflect.Slice:
			bytes := make([]byte, reflectVal.Len())
			ok := true
			for i, _ := range bytes {
				int32Val := Int32(reflectVal.Index(i).Interface())
				if int32Val < 0 || int32Val > math.MaxUint8 {
					ok = false
					break
				}
				bytes[i] = byte(int32Val)
			}

			if ok {
				return bytes, nil
			}
		}
		return Encode(any), nil
	}
}
