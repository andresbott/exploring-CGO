package gohello

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -L./lib
#include <stdlib.h>
#include "./hello.h"
*/
import "C"

//-g enable debug symbols
//-Wall enable all warnings

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"unsafe"
)

var _ = spew.Dump // prevent the package from beeng removed

func Hello() string {
	result := C.getHello() // calling the C function
	parsed := C.GoString(result)
	return parsed
}
func SendString(in string) string {

	cIn := C.CString(in)              // Cast a string to a 'C string'
	defer C.free(unsafe.Pointer(cIn)) // free the memory of the C String

	result := C.return_string(cIn)
	defer C.free(unsafe.Pointer(result))
	parsed := C.GoString(result)

	return parsed
}

func SampleStringArray() []string {
	cStringArray := C.sample_string_array()
	count := 0
	for cStringArray != nil && *cStringArray != nil {
		count++
		cStringArray = (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(cStringArray)) + unsafe.Sizeof(*cStringArray)))
	}

	goStrings := make([]string, count)

	for i := 0; i < count; i++ {
		goStrings[i] = C.GoString((*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(cStrings)) + uintptr(i)*unsafe.Sizeof(*cStrings))))
	}

	return goStrings

	return nil

}

// TypeConversions converts different go types to C types and back
func TypeConversions() []string {
	out := []string{}

	// string
	str := "my string"
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))
	newStr := C.GoString(cStr)
	out = append(out, fmt.Sprintf("strinc conversion: %s", newStr))

	// int
	myInt := 42
	cInt := C.int(myInt)
	newInt := int(cInt)
	out = append(out, fmt.Sprintf("int conversion: %d", newInt))

	return out
}

// SampleStruct is a representation of a C struct with underlying data
type SampleStruct struct {
	Name   string
	Number int
	Data   unsafe.Pointer
}

func GetStruc() SampleStruct {
	Cst := C.get_struct()

	st := SampleStruct{
		Name:   C.GoString(&Cst.name[0]),
		Number: int(Cst.number),
		Data:   nil,
	}
	return st
}

func SendStruct(name string, number int) {
	cString := C.CString(name)
	C.free(unsafe.Pointer(cString))

}
