package gohello

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -L./lib
#include <stdlib.h>
#include "./hello.h"
*/
import "C"
import "unsafe"

//-g enable debug symbols
//-Wall enable all warnings

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
