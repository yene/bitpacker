package bitpacker_test

import (
	"testing"

	"github.com/yene/bitpacker"
)

func TestBasicPackUnpack(t *testing.T) {
	type BitStruct struct {
		Flag         bitpacker.Uint1
		HeaderLength bitpacker.Uint3
		Data         bitpacker.Uint4
	}

	x := BitStruct{1, 2, 255}
	number := bitpacker.Pack(x)
	var y BitStruct
	bitpacker.Unpack(&y, number)
	if y.Flag != 1 || y.HeaderLength != 2 || y.Data != 15 {
		t.Errorf("Unpacked data did not match.")
	}
}

func TestBool(t *testing.T) {
	type BitStruct struct {
		Flag         bool
		HeaderLength bitpacker.Uint3
		Data         bitpacker.Uint4
	}

	x := BitStruct{true, 2, 15}
	number := bitpacker.Pack(x)
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
	number := bitpacker.Pack(x)
	var y BitStruct
	bitpacker.Unpack(&y, number)
	if y.Flag != 1 || y.HeaderLength != 2 || y.Data != 15 {
		t.Errorf("Unpacked data did not match.")
	}
}
