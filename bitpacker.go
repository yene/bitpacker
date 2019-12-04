package bitpacker

import (
	"fmt"
	"reflect"
)

// PrintWarnings uses the fmt package, if you want to slimm down remove the warnings.
var PrintWarnings = true

// Types to represent bits, in reality they are still store in a byte.
// you can use Uint1 or bool to represent 1 bit

type Uint1 uint8
type Uint2 uint8
type Uint3 uint8
type Uint4 uint8
type Uint5 uint8
type Uint6 uint8
type Uint7 uint8

// Pack turns a struct with a total sum of 8 bits into a uint8, which can be used to write into a buffer.
func Pack(x interface{}) uint8 {
	v := reflect.ValueOf(x)
	if v.Kind() != reflect.Struct {
		panic("Pack expected a struct")
	}

	typeOfT := v.Type()
	bitPosition := 0
	var number uint8
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		bc := bitCountForType(f.Interface())
		if bc == 0 {
			s := fmt.Sprintf("Found invalid type %s (field %s) in struct. PLease use the exported types.", f.Type(), typeOfT.Field(i).Name)
			panic(s)
		}
		if bitPosition+bc > 8 {
			panic("The provided bit struct is bigger than 8bit.")
		}
		value := intForType(f.Interface())
		number |= value << (8 - bc - bitPosition)
		bitPosition = bitPosition + bc
		//fmt.Printf("%08b\n", value)
		//fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	return number
}

// Unpack fills the provided struct with bits
func Unpack(x interface{}, data uint8) {
	ptr := reflect.ValueOf(x)
	if ptr.Kind() != reflect.Ptr || ptr.IsNil() {
		panic("Unpack expected a pointer to a struct")
	}
	v := ptr.Elem()
	if v.Kind() != reflect.Struct {
		panic("Unpack expected a pointer to a struct")
	}

	typeOfT := v.Type()
	bitPosition := 0
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		bc := bitCountForType(f.Interface())
		if bc == 0 {
			s := fmt.Sprintf("Found invalid type %s (field %s) in struct. PLease use the exported types.", f.Type(), typeOfT.Field(i).Name)
			panic(s)
		}
		if bitPosition+bc > 8 {
			panic("The provided bit struct is bigger than 8bit.")
		}
		rightAlign := data >> (8 - (bc + bitPosition))
		numberWithType := valueForType(f.Interface(), rightAlign)
		f.Set(numberWithType)
		bitPosition = bitPosition + bc
	}

}

func bitCountForType(t interface{}) int {
	switch t.(type) {
	case bool:
		return 1
	case Uint1:
		return 1
	case Uint2:
		return 2
	case Uint3:
		return 3
	case Uint4:
		return 4
	case Uint5:
		return 5
	case Uint6:
		return 6
	case Uint7:
		return 7
	default:
		return 0
	}
}

func valueForType(t interface{}, v uint8) reflect.Value {
	switch t.(type) {
	case bool:
		v = v & 0b00000001
		if v == 1 {
			return reflect.ValueOf(true)
		}
		return reflect.ValueOf(false)
	case Uint1:
		v = v & 0b00000001
		return reflect.ValueOf(Uint1(v))
	case Uint2:
		v = v & 0b00000011
		return reflect.ValueOf(Uint2(v))
	case Uint3:
		v = v & 0b00000111
		return reflect.ValueOf(Uint3(v))
	case Uint4:
		v = v & 0b00001111
		return reflect.ValueOf(Uint4(v))
	case Uint5:
		v = v & 0b00011111
		return reflect.ValueOf(Uint5(v))
	case Uint6:
		v = v & 0b00111111
		return reflect.ValueOf(Uint6(v))
	case Uint7:
		v = v & 0b01111111
		return reflect.ValueOf(Uint7(v))
	default:
		fmt.Println("Did not handle type:", reflect.TypeOf(v))
		return reflect.ValueOf(t)
	}
}

func intForType(t interface{}) uint8 {
	switch v := t.(type) {
	case bool:
		return uint8(bool2int(v))
	case Uint1:
		if v > 1 && PrintWarnings {
			fmt.Println("Warning: Uint1 contains a number higher than what will be encoded, number will be reduced to 1 bit.")
		}
		v = v & 0b00000001
		return uint8(v)
	case Uint2:
		if v > 3 && PrintWarnings {
			fmt.Println("Warning: Uint2 contains a number higher than what can be encoded, number will be reduced to 2 bits.")
		}
		v = v & 0b00000011
		return uint8(v)
	case Uint3:
		if v > 7 && PrintWarnings {
			fmt.Println("Warning: Uint3 contains a number higher than what can be encoded, number will be reduced to 3 bits.")
		}
		v = v & 0b00000111
		return uint8(v)
	case Uint4:
		if v > 15 && PrintWarnings {
			fmt.Println("Warning: Uint4 contains a number higher than what can be encoded, number will be reduced to 4 bits.")
		}
		v = v & 0b00001111
		return uint8(v)
	case Uint5:
		if v > 31 && PrintWarnings {
			fmt.Println("Warning: Uint5 contains a number higher than what can be encoded, number will be reduced to 5 bits.")
		}
		v = v & 0b00011111
		return uint8(v)
	case Uint6:
		if v >= 63 && PrintWarnings {
			fmt.Println("Warning: Uint6 contains a number higher than what can be encoded, number will be reduced to 6 bits.")
		}
		v = v & 0b00111111
		return uint8(v)
	case Uint7:
		if v > 127 && PrintWarnings {
			fmt.Println("Warning: Uint7 contains a number higher than what can be encoded, number will be reduced to 7 bits.")
		}
		v = v & 0b01111111
		return uint8(v)
	default:
		fmt.Println("Did not handle type:", reflect.TypeOf(v))
		return 0
	}
}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}
