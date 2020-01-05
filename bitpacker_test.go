package bitpacker_test

import (
	"testing"

	"github.com/yene/bitpacker"
)

/* WIP: custom size
func TestNByteStruct(t *testing.T) {
	// Total length is 11
	type BitStruct struct {
		Flag         int `uint1`
		HeaderLength int `uint3`
		Data1        int `uint4`
		Data2        int `uint2`
		Data3        int `uint1`
	}

	x := BitStruct{1, 2, 3, 2, 7}
	number := bitpacker.Pack(x, 11)
	var y BitStruct
	bitpacker.Unpack(&y, number)
	if y.Flag != 1 || y.HeaderLength != 2 || y.Data1 != 3 || y.Data3 != 7 {
		t.Errorf("Unpacked data did not match.")
	}
}*/

// bitpacker uspports all number types to make it more convinient for the user
func TestNumberTypes(t *testing.T) {
	type BitStruct struct {
		NakedBool bool
		Bool      bool   `uint1`
		IntBool   uint   `bool`
		Int       int    `uint1`
		Uint      uint   `uint1`
		Uint8     uint8  `uint1`
		Uint16    uint16 `uint1`
		Uint32    uint32 `uint1`
		Uint64    uint64 `uint1`
	}

	x := BitStruct{true, true, 1, 1, 1, 1, 1, 1, 1}
	number := bitpacker.Pack(x, 9)
	// e := fmt.Sprintf("%08b\n", number)
	if number != 511 {
		t.Errorf("Unpacked data did not match.")
	}
}

func TestBasicPackUnpack(t *testing.T) {
	type BitStruct struct {
		Flag         int `uint1`
		HeaderLength int `uint3`
		Data         int `uint4`
	}

	x := BitStruct{1, 2, 255}
	number := bitpacker.PackByte(x)
	var y BitStruct
	bitpacker.Unpack(&y, number)
	if y.Flag != 1 || y.HeaderLength != 2 || y.Data != 15 {
		t.Errorf("Unpacked data did not match.")
	}
}

func TestBool(t *testing.T) {
	type BitStruct struct {
		Flag         bool
		HeaderLength int `uint3`
		Data         int `uint4`
	}

	x := BitStruct{true, 2, 15}
	number := bitpacker.PackByte(x)

	var y BitStruct
	bitpacker.Unpack(&y, number)
	if y.Flag != true {
		t.Errorf("Unpacked data did not match.")
	}
}

func TestTags(t *testing.T) {
	type BitStruct struct {
		Flag         int `uint1`
		HeaderLength int `uint3`
		Data         int `uint4`
	}

	x := BitStruct{1, 2, 255}
	number := bitpacker.PackByte(x)
	var y BitStruct
	bitpacker.Unpack(&y, number)
	if y.Flag != 1 || y.HeaderLength != 2 || y.Data != 15 {
		t.Errorf("Unpacked data did not match.")
	}
}
