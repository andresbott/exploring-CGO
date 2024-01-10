package mpv

/*
#include <mpv/client.h>
#include <mpv/render.h>
#include <stdlib.h>
#cgo LDFLAGS: -lmpv

*/
import "C"
import (
	"fmt"
)

func Create() (*C.mpv_handle, error) {
	handle := C.mpv_create()
	if handle == nil {
		return nil, fmt.Errorf("unable to create handle")
	}
	return handle, nil
}

func Init(handle *C.mpv_handle) error {

	ret := C.mpv_initialize(handle)
	if ret != 0 {
		return fmt.Errorf("error initializing: error code %d", int(ret))
	}
	return nil
}

func RenderContextCreate(handle *C.mpv_handle) error {

	var glCtx *C.mpv_render_context

	var params [4]C.mpv_render_param

	ret := C.mpv_render_context_create(&glCtx, handle, p)
	if ret != 0 {
		return fmt.Errorf("error creating context: error code %d", int(ret))
	}
	return nil
}
