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
		Flag         uint  `uint1`
		HeaderLength uint8 `uint3`
		Data         int   `uint4`
	}

	x := BitStruct{1, 2, 255}
	number := bitpacker.PackByte(x)
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
