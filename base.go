package gconv

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func Int64(any interface{}) int64 {
	if any == nil {
		return 0
	}

	switch val := any.(type) {
	case int:
		return int64(val)
	case int8:
		return int64(val)
	case int16:
		return int64(val)
	case int32:
		return int64(val)
	case int64:
		return val
	case uint:
		return int64(val)
	case uint8:
		return int64(val)
	case uint16:
		return int64(val)
	case uint32:
		return int64(val)
	case uint64:
		return int64(val)
	case float32:
		return int64(val)
	case float64:
		return int64(val)
	case bool:
		if val {
			return 1
		}
		return 0
	case []byte:
		return DecodeToInt64(val)
	default:
		s := String(any)
		isMinus := false
		if len(s) > 1 {
			if s[0] == '-' {
				isMinus = true
				s = s[1:]
			}

			if s[0] == '+' {
				s = s[1:]
			}
		}
		//16进制
		if len(s) > 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
			if v, err := strconv.ParseInt(s[:2], 16, 64); err == nil {
				if isMinus {
					return -v
				}
				return v
			}
		}

		//8进制
		if len(s) > 1 && s[0] == '0' {
			if v, err := strconv.ParseInt(s[1:], 8, 64); err == nil {
				if isMinus {
					return -v
				}
				return v
			}
		}

		//10进制
		if v, err := strconv.ParseInt(s, 10, 64); err == nil {
			if isMinus {
				return -v
			}
			return v
		}
		return int64(Float64(any))
	}
}

func Int32(any interface{}) int32 {
	if any == nil {
		return 0
	}

	if v, ok := any.(int32); ok {
		return v
	}

	return int32(Int64(any))
}

func String(any interface{}) string {
	if any == nil {
		return ""
	}
	switch value := any.(type) {
	case int:
		return strconv.Itoa(value)
	case int8:
		return strconv.Itoa(int(value))
	case int16:
		return strconv.Itoa(int(value))
	case int32:
		return strconv.Itoa(int(value))
	case int64:
		return strconv.FormatInt(value, 10)
	case uint:
		return strconv.FormatUint(uint64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(value, 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	case string:
		return value
	case []byte:
		return string(value)
	case time.Time:
		if value.IsZero() {
			return ""
		}
		return value.String()
	case *time.Time:
		if value == nil {
			return ""
		}
		return value.String()
	default:
		reflectValue := reflect.ValueOf(any)
		reflectKind := reflectValue.Kind()

		switch reflectKind {
		case reflect.Chan,
			reflect.Map,
			reflect.Slice,
			reflect.Func,
			reflect.Ptr,
			reflect.Interface,
			reflect.UnsafePointer:
			if reflectValue.IsZero() {
				return ""
			}
		case reflect.String:
			return reflectValue.String()
		}
		if reflectKind == reflect.Ptr {
			String(reflectValue.Elem().Interface())
		}
		if jsonContent, err := json.Marshal(value); err != nil {
			return fmt.Sprint(value)
		} else {
			return string(jsonContent)
		}
	}
}

func Float64(any interface{}) float64 {
	if any == nil {
		return 0
	}
	switch val := any.(type) {
	case float32:
		return float64(val)
	case float64:
		return val
	case []byte:
		return DecodeToFloat64(val)
	default:
		v, _ := strconv.ParseFloat(String(any), 64)
		return v
	}
}
