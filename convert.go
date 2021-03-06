package utils

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"time"
	"unsafe"
)

func SliceByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func StringToSliceByte(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func JsonBodyAsMap(body []byte) (m map[string]interface{}, err error) {
	err = json.Unmarshal(body, &m)
	//	LogErr(err)
	return
}

func StrToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		fmt.Errorf("Unsupported type StrToInt(%s) %s", s, err.Error())

	}
	return i
}

func StrToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Errorf("Unsupported type StrToInt(%s) %s", s, err.Error())
	}
	return i
}

func BoolToStr(b bool) (str string) {
	return strconv.FormatBool(b)
}

func StrToBool(str string) (b bool) {
	b, err := strconv.ParseBool(str)
	if err != nil {
		fmt.Errorf("faild to convert StrToBool(%s) with error : %s", str, err.Error())
		return false
	}
	return b //	fmt.Printf("%T, %v\n", s, s)
}

func IntToStr(i interface{}) string {
	switch i.(type) {
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%d", i)
	default:
		return "0"
	}
}

// int,int64 to bytes
func Int64ToBytes(val interface{}) (res []byte) {
	return big.NewInt(val.(int64)).Bytes()
}

/*
func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}*/

func FloatToStr(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func StrToFloat(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		//fmt.Errorf(err.Error())
		fmt.Errorf("faild to convert StrToFloat(%s) with error : %s", str, err.Error())
	}

	return f
}

// HexToBytes converts a hex string representation of bytes to a byte representation
func HexToBytes(h string) []byte {
	s, err := hex.DecodeString(h)
	if err != nil {
		fmt.Errorf("faild to convert BytesToFloat(%s) with error : %s", h, err.Error())
		return []byte("")
	}
	return s
}

// BytesToHex converts bytes to a hex string representation of bytes
func BytesToHex(b []byte) string {
	return hex.EncodeToString(b)
}

func BytesToInt64(buf []byte) int64 {
	//res, _ = strconv.ParseInt(string(val), 10, 64)

	return int64(binary.BigEndian.Uint64(buf))
}

func BytesToFloat(buf []byte) (res float64) {
	res, err := strconv.ParseFloat(string(buf), 32)
	if err != nil {
		fmt.Errorf("faild to convert BytesToFloat(%s) with error : %s", string(buf), err.Error())
	}
	return
}

// Base64ToBytes converts from a b64 string to bytes
func Base64ToBytes(h string) []byte {
	s, err := base64.URLEncoding.DecodeString(h)
	if err != nil {
		fmt.Errorf("faild to convert Base64ToBytes(%s) with error : %s", h, err.Error())
		return []byte("")
	}
	return s
}

// BytesToBase64 converts bytes to a base64 string representation
func BytesToBase64(b []byte) string {
	return base64.URLEncoding.EncodeToString(b)
}

func Itf2Bool(val interface{}) (res bool) {
	if val == nil {
		return
	}

	if value, ok := val.(bool); ok {
		return value
	} else {
		t := reflect.TypeOf(val)
		vv := reflect.Indirect(reflect.ValueOf(val))
		switch t.Kind() {
		case reflect.Bool:
			return vv.Bool()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return vv.Int() != 0
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return vv.Uint() != 0
		case reflect.String:
			if b, err := strconv.ParseBool(vv.String()); err != nil {
				fmt.Errorf("Unsupported type Itf2Bool(%v) error : %s", vv.Type().Name(), err.Error())
			} else {
				return b
			}
		default:
			fmt.Errorf("Unsupported type Itf2Bool(%v)", vv.Type().Name())
		}
	}

	return
}

func Itf2Int(val interface{}) (res int) {
	if val == nil {
		return
	}

	if value, ok := val.(int); ok {
		return value
	} else {
		t := reflect.TypeOf(val)
		vv := reflect.Indirect(reflect.ValueOf(val))
		switch t.Kind() {
		//checked
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return int(vv.Int())

			//checked
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return int(vv.Uint())

		//checked
		case reflect.Float32, reflect.Float64:
			return int(vv.Float())

			//checked
		case reflect.String:
			return StrToInt(vv.String())
			/*
				if i, err := strconv.ParseInt(vv.String(), 10, 0); err != nil {
					fmt.Errorf("Unsupported type %v", vv.Type().Name())
				} else {
					return i
				}
			*/
		case reflect.Array, reflect.Slice:
			switch t.Elem().Kind() {
			case reflect.Uint8:
				data := vv.Interface().([]byte)
				return int(binary.BigEndian.Uint32(data))
			default:
				fmt.Errorf("Unsupported struct type %v", vv.Type().Name())
			}
		default:
			fmt.Errorf("Unsupported type %v", vv.Type().Name())
		}
	}

	return
}

func Itf2Int64(val interface{}) (res int64) {
	if val == nil {
		return
	}

	if value, ok := val.(int64); ok {
		return value
	} else {
		t := reflect.TypeOf(val)
		vv := reflect.Indirect(reflect.ValueOf(val))
		switch t.Kind() {
		//checked
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return vv.Int()

			//checked
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return int64(vv.Uint())

		//checked
		case reflect.Float32, reflect.Float64:
			return int64(vv.Float())

			//checked
		case reflect.String:
			return StrToInt64(vv.String())
			/*if i, err := strconv.ParseInt(vv.String(), 10, 0); err != nil {
				fmt.Errorf("Unsupported type %v", vv.Type().Name())
			} else {
				return i
			}
			*/
		case reflect.Array, reflect.Slice:
			switch t.Elem().Kind() {
			case reflect.Uint8:
				data := vv.Interface().([]byte)
				return int64(binary.BigEndian.Uint64(data))
			default:
				fmt.Errorf("Unsupported struct type Itf2Int64(%v)", vv.Type().Name())
			}
		default:
			fmt.Errorf("Unsupported type Itf2Int64(%v)", vv.Type().Name())
		}
	}

	return
}

func Itf2Float(val interface{}) (res float64) {
	if val == nil {
		return
	}

	if value, ok := val.(float64); ok {
		return value
	} else {
		t := reflect.TypeOf(val)
		vv := reflect.Indirect(reflect.ValueOf(val))
		switch t.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return vv.Float()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return vv.Float()
		case reflect.String:
			if f, err := strconv.ParseFloat(vv.String(), 64); err != nil {
				fmt.Errorf("Unsupported type Itf2Float(%v) error : %s", vv.Type().Name(), err.Error())
			} else {
				return f
			}
		default:
			fmt.Errorf("Unsupported type Itf2Float(%v)", vv.Type().Name())
		}
	}

	return
}

func Itf2Float32(val interface{}) (res float32) {
	if value, ok := val.(float32); ok {
		return value
	}

	return float32(Itf2Float(val))
}

func Itf2Str(val interface{}) (res string) {
	if val == nil {
		return
	}

	if value, ok := val.(string); ok {
		return value
	} else {
		t := reflect.TypeOf(val)
		vv := reflect.Indirect(reflect.ValueOf(val))
		switch t.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return strconv.FormatInt(vv.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return strconv.FormatUint(vv.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			return strconv.FormatFloat(vv.Float(), 'f', -1, 64)
		case reflect.String:
			return vv.String()
		case reflect.Array, reflect.Slice:
			switch t.Elem().Kind() {
			case reflect.Uint8:
				data := vv.Interface().([]byte)
				return string(data)
			default:
				fmt.Errorf("Unsupported struct type Itf2Str(%v)", vv.Type().Name())
			}
		//时间类型
		case reflect.Struct:
			var c_TIME_DEFAULT time.Time
			TimeType := reflect.TypeOf(c_TIME_DEFAULT)
			if t.ConvertibleTo(TimeType) {
				return vv.Convert(TimeType).Interface().(time.Time).Format(time.RFC3339Nano)
			} else {
				fmt.Errorf("Unsupported struct type Itf2Str(%v)", vv.Type().Name())
			}
		case reflect.Bool:
			return strconv.FormatBool(vv.Bool())
		case reflect.Complex128, reflect.Complex64:
			return fmt.Sprintf("%v", vv.Complex())
		/* TODO: unsupported types below
		   case reflect.Map:
		   case reflect.Ptr:
		   case reflect.Uintptr:
		   case reflect.UnsafePointer:
		   case reflect.Chan, reflect.Func, reflect.Interface:
		*/
		default:
			fmt.Errorf("Unsupported struct type Itf2Str(%v)", vv.Type().Name())
		}
	}

	return
}

func Itf2Time(val interface{}) (res time.Time) {
	if val == nil {
		return
	}

	if value, ok := val.(time.Time); ok {
		return value
	} else {
		t := reflect.TypeOf(val)
		vv := reflect.Indirect(reflect.ValueOf(val))
		//fmt.Println("datetime21", t, t.Kind(), val, vv, vv.String())
		switch t.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return time.Unix(vv.Int(), 0)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return time.Unix(int64(vv.Uint()), 0)

		case reflect.String:

			if tm, err := time.Parse("2006-01-02 15:04:05", vv.String()); err != nil {
				fmt.Errorf("Unsupported type Itf2Time(%v) error : %s", vv.Type().Name(), err.Error())
				//fmt.Println("String1:", val, vv.String(), err.Error())
			} else {
				fmt.Errorf("Unsupported type Itf2Time(%v)", tm)
				return tm
			}
		case reflect.Struct:
			var c_TIME_DEFAULT time.Time
			TimeType := reflect.TypeOf(c_TIME_DEFAULT)
			//fmt.Println("datetime22", t, t.Kind(), t.ConvertibleTo(TimeType))

			if t.ConvertibleTo(TimeType) {
				return vv.Interface().(time.Time)
			} else {
				fmt.Errorf("Unsupported struct type Itf2Time(%v)", vv.Type().Name())
			}
		default:
			fmt.Errorf("Unsupported type Itf2Time(%v)", vv.Type().Name())
		}
	}

	return
}
