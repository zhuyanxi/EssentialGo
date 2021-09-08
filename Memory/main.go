package main

import (
	"fmt"
	"unsafe"
)

type smallStruct struct {
	a, b int64
	c, d float64
}

//go:noinline
func smallAllocation() *smallStruct {
	return &smallStruct{}
}

func main() {

	fmt.Println(unsafe.Sizeof(smallAllocation()))
	smallAllocation()
}
