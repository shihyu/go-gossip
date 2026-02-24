package main

import (
    "fmt"
    "math"
    "reflect"
)

func main() {
    fmt.Printf("uint8  : 0 ~ %d\n", math.MaxUint8)
    fmt.Printf("uint16 : 0 ~ %d\n", math.MaxUint16)
    fmt.Printf("uint32 : 0 ~ %d\n", math.MaxUint32)
    fmt.Printf("uint64 : 0 ~ %d\n", uint64(math.MaxUint64))
    fmt.Printf("int8   : %d ~ %d\n", math.MinInt8, math.MaxInt8)
    fmt.Printf("int16  : %d ~ %d\n", math.MinInt16, math.MaxInt16)
    fmt.Printf("int32  : %d ~ %d\n", math.MinInt32, math.MaxInt32)
    fmt.Printf("int64  : %d ~ %d\n", math.MinInt64, math.MaxInt64)
    fmt.Printf("整數預設型態: %s\n", reflect.TypeOf(1))
}
