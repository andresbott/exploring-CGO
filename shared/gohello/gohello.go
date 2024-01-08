package gohello

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -L./libhello -lhello
#include "./libhello/hello.h"

*/
import "C"

//-g enable debug symbols
//-Wall enable all warnings

func Hello() string {
	result := C.hello() // calling the C function
	parsed := C.GoString(result)
	return parsed
}
