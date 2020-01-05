package bitpacker

import (
	"fmt"
	"reflect"
)

// PackByte turns a struct with a total sum of 8 bits into a uint8, which can be used to write into a buffer.
func PackByte(x interface{}) uint8 {
	res := Pack(x, 8)
	return uint8(res)
}

// Pack turns a struct with a total sum of 8 bits into a uint8, which can be used to write into a buffer.
func Pack(x interface{}, structSize int) uint {
	// TODO: check if .Elem() is appropriate here, and forcing the user to pass in a pointer, see json marshal
	v := reflect.ValueOf(x)
	if v.Kind() != reflect.Struct {
		panic("Pack expected a struct")
	}

	bitPosition := 0
	var number uint
	for i := 0; i < v.NumField(); i++ {
		valueField := v.Field(i)
		typeField := v.Type().Field(i)
		tag := typeField.Tag
		if valueField.Kind() != reflect.Bool && valueField.Kind() != reflect.Int && valueField.Kind() != reflect.Uint && valueField.Kind() != reflect.Uint8 && valueField.Kind() != reflect.Uint16 && valueField.Kind() != reflect.Uint32 && valueField.Kind() != reflect.Uint64 {
			e := fmt.Sprintf("Expected: \"%s\" to be a number type.\n", typeField.Name)
			panic(e)
		}
		// fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", typeField.Name, valueField.Interface(), tag)
		bc := bitCountForTag(tag)
		// Allowing naked bool
		if bc == 0 && valueField.Kind() == reflect.Bool {
			bc = 1
		}
		if bc == 0 {
			e := fmt.Sprintf("Field: \"%s\", has an invalid Tag: \"%s\"\n", typeField.Name, tag)
			panic(e)
		}

		if bitPosition+bc > structSize {
			panic("The provided bit struct is bigger than 8bit.")
		}

		// convert any number type to uint8
		var intValue uint = 0
		intValue = convertNumberToUint(valueField.Interface())
		intValue = ensureWidthFor(intValue, bc)

		number |= intValue << (structSize - bc - bitPosition)
		bitPosition = bitPosition + bc
	}
	return number
}

// UnpackByte fills the provided struct with bits
func UnpackByte(x interface{}, data uint8) {
	Unpack(x, data)
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

	bitPosition := 0
	for i := 0; i < v.NumField(); i++ {
		valueField := v.Field(i)
		typeField := v.Type().Field(i)
		tag := typeField.Tag
		bc := bitCountForTag(tag)
		// Allowing naked bool
		if bc == 0 && valueField.Kind() == reflect.Bool {
			bc = 1
		}
		if bc == 0 {
			e := fmt.Sprintf("Field: \"%s\", has an invalid Tag: \"%s\"\n", typeField.Name, tag)
			panic(e)
		}
		/* TODO: do we need to validate this?
		if bitPosition+bc > structSize {
			panic("The provided bit struct is bigger than 8bit.")
		}*/
		rightAlign := data >> (8 - (bc + bitPosition))
		numberWithUint := ensureWidthFor(uint(rightAlign), bc)
		reflectValue := convertValueToType(valueField.Interface(), numberWithUint)
		valueField.Set(reflectValue)
		bitPosition = bitPosition + bc
	}

}

func bitCountForTag(tag reflect.StructTag) int {
	switch tag {
	case "bool":
		return 1
	case "uint1":
		return 1
	case "uint2":
		return 2
	case "uint3":
		return 3
	case "uint4":
		return 4
	case "uint5":
		return 5
	case "uint6":
		return 6
	case "uint7":
		return 7
	default:
		return 0
	}
}

func convertNumberToUint(t interface{}) uint {
	if v, ok := t.(bool); ok {
		if v == true {
			return 1
		}
		return 0
	} else if v, ok := t.(int); ok {
		return uint(v)
	} else if v, ok := t.(uint); ok {
		return uint(v)
	} else if v, ok := t.(uint8); ok {
		return uint(v)
	} else if v, ok := t.(uint16); ok {
		return uint(v)
	} else if v, ok := t.(uint32); ok {
		return uint(v)
	} else if v, ok := t.(uint64); ok {
		return uint(v)
	} else {
		e := fmt.Sprintf("Could not convert interface to uint: %v\n", t)
		panic(e)
		// return 0
	}
}

func convertValueToType(t interface{}, v uint) reflect.Value {
	switch t.(type) {
	case bool:
		if v == 1 {
			return reflect.ValueOf(true)
		}
		return reflect.ValueOf(false)
	case int:
		return reflect.ValueOf(int(v))
	case uint:
		return reflect.ValueOf(uint(v))
	case uint8:
		return reflect.ValueOf(uint8(v))
	case uint16:
		return reflect.ValueOf(uint16(v))
	case uint32:
		return reflect.ValueOf(uint32(v))
	case uint64:
		return reflect.ValueOf(uint64(v))
	default:
		fmt.Println("Did not handle type:", t)
		return reflect.ValueOf(t)
	}
}

func ensureWidthFor(v uint, width int) uint {
	switch width {
	case 1:
		v = v & 0b00000001
		return v
	case 2:
		v = v & 0b00000011
		return v
	case 3:
		v = v & 0b00000111
		return v
	case 4:
		v = v & 0b00001111
		return v
	case 5:
		v = v & 0b00011111
		return v
	case 6:
		v = v & 0b00111111
		return v
	case 7:
		v = v & 0b01111111
		return v
	default:
		fmt.Println("ensureWidthFor not handle wdith:", width)
		return v
	}
}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}
