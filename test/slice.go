package test

import (
	"fmt"
	"unsafe"
)

func mainSlice() {
	slice := make([]int, 0)

	slice = append(slice, 1)
	fmt.Println(unsafe.Pointer(&slice[0]), len(slice), cap(slice))
	//fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&slice)).Data)

	slice = append(slice, 2)
	fmt.Println(unsafe.Pointer(&slice[0]), len(slice), cap(slice))
	//fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&slice)).Data)
}
