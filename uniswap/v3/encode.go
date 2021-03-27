package v3

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// SolidityEncodePacked implements the non-standard encoding available in solidity.
// Since encoding is ambigious there is no decoding function.
//
// Using EncodePacked to pack data before hashing or signing is generally unsafe because
// the following arguments produce the same output.
// abi.SolidityEncodePacked([]Type{String,String}, []interface{}{"hello", "world01"})
// abi.SolidityEncodePacked([]Type{String,String}, []interface{}{"helloworld", "01"})
// '0x68656c6c6f776f726c643031'
func SolidityEncodePacked(args []abi.Type, values []interface{}) ([]byte, error) {
	enc := make([]byte, 0)
	var index int
	for _, arg := range args {
		switch arg.T {
		case abi.TupleTy:
			return []byte{}, errors.New("Type not supported in abi.EncodePacked()")
		case abi.ArrayTy, abi.SliceTy:
			packed, err := encodePackArray(arg, values[index:arg.Size])
			if err != nil {
				return []byte{}, err
			}
			enc = append(enc, packed...)
			index += arg.Size
		default:
			packed, err := encodePackElement(arg, reflect.ValueOf(values[index]))
			if err != nil {
				return []byte{}, err
			}
			enc = append(enc, packed...)
			index++
		}
	}
	return enc, nil
}

func encodePackArray(t abi.Type, values []interface{}) ([]byte, error) {
	encoded := make([]byte, 0, t.Size*32)
	for i := 0; i < t.Size; i++ {
		packed, err := encodePackElement(*t.Elem, reflect.ValueOf(values[i]))
		if err != nil {
			return []byte{}, err
		}
		// Array elements are packed with padding
		padded := common.LeftPadBytes(packed, 32)
		encoded = append(encoded, padded...)
	}
	return encoded, nil
}

func encodePackElement(t abi.Type, value reflect.Value) ([]byte, error) {
	value = indirect(value)

	switch t.T {
	case abi.IntTy, abi.UintTy:
		return encodePackedNum(t, value), nil
	case abi.StringTy:
		return encodePackedByteSlice(t, []byte(value.String())), nil
	case abi.AddressTy, abi.FixedBytesTy:
		if value.Kind() == reflect.Array {
			value = mustArrayToByteSlice(value)
		}
		return encodePackedByteSlice(t, value.Bytes()), nil
	case abi.BoolTy:
		if value.Bool() {
			return []byte{1}, nil
		}
		return []byte{0}, nil
	case abi.BytesTy:
		if value.Kind() == reflect.Array {
			value = mustArrayToByteSlice(value)
		}
		if value.Type() != reflect.TypeOf([]byte{}) {
			return []byte{}, errors.New("Bytes type is neither slice nor array")
		}
		return encodePackedByteSlice(t, value.Bytes()), nil
	default:
		return []byte{}, fmt.Errorf("Could not encode pack element, unknown type: %v", t.T)
	}
}

func encodePackedNum(t abi.Type, value reflect.Value) []byte {
	bytes := make([]byte, 8)
	switch kind := value.Kind(); kind {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		binary.BigEndian.PutUint64(bytes, value.Uint())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		binary.BigEndian.PutUint64(bytes, uint64(value.Int()))
	case reflect.Ptr:
		big := new(big.Int).Set(value.Interface().(*big.Int))
		bytes = big.Bytes()
	default:
		panic(fmt.Sprintf("abi: fatal error: %v", kind))
	}
	return encodePackedByteSlice(t, bytes)
}

func encodePackedByteSlice(t abi.Type, value []byte) []byte {
	size := t.Size / 8
	// If size is not set in the type, use the length of the value to pad
	if size == 0 {
		size = len(value)
	}
	padded := common.LeftPadBytes(value, size)
	return padded[len(padded)-size:]
}

// indirect recursively dereferences the value until it either gets the value
// or finds a big.Int
func indirect(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Ptr && v.Elem().Type() != reflect.TypeOf(big.Int{}) {
		return indirect(v.Elem())
	}
	return v
}

// mustArrayToByteSlice creates a new byte slice with the exact same size as value
// and copies the bytes in value to the new slice.
func mustArrayToByteSlice(value reflect.Value) reflect.Value {
	slice := reflect.MakeSlice(reflect.TypeOf([]byte{}), value.Len(), value.Len())
	reflect.Copy(slice, value)
	return slice
}
