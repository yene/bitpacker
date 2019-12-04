package main

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/yene/bitpacker"
)

func main() {
	// | Flag 1bit | HeaderLength 3bit | Data 4bit |
	type BitStruct struct {
		Flag         bitpacker.Uint1
		HeaderLength bitpacker.Uint3
		Data         bitpacker.Uint4
	}

	x := BitStruct{1, 2, 255}
	number := bitpacker.Pack(x)
	fmt.Printf("%08b\n", number)

	var y BitStruct
	bitpacker.Unpack(&y, number)
	fmt.Println(y)

	// Write it into a buffer like this:
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, number)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%08b\n", buf)
}