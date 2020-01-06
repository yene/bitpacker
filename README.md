# bitpacker

A small package that uses custom types to pack a byte.

```go
// | Flag 1bit | HeaderLength 3bit | Data 4bit |
type BitStruct struct {
	Flag         int `uint1`
	HeaderLength int `uint3`
	Data         int `uint4`
}

x := BitStruct{1, 2, 255}
number := bitpacker.PackByte(x)
fmt.Printf("%08b\n", number)

var y BitStruct
bitpacker.UnpackByte(&y, number)
fmt.Println(y)
```

Will result in

```bash
10101111
{1 2 15} # 255 is silently converted into uint4
```

## Notes
* TODO: Replace panic with return error.
* TODO: Improve support for custom length.
* https://github.com/lunixbochs/struc/issues/7
