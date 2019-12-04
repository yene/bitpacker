# bitpacker

A small package that uses custom types to pack a byte.

```go
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

```

Will result in

```
Warning: Uint4 contains a number higher than what can be encoded, number will be reduced to 4 bits.
10101111
{1 2 15}

```



## Notes
* We panic if something is invalid.
* TODO: Add support for tags.
* TODO: Add tests. 🤨
* https://github.com/lunixbochs/struc/issues/7
