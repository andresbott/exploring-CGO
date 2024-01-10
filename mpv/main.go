package main

import "C"
import (
	"github.com/andresbott/exploring-CGO/mpv/mpv"
	"log"
	"runtime"

	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	// ????
	runtime.LockOSThread()
}
func main() {
	//Init gl
	err := glfw.Init()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	mpvHandle, err := mpv.Create()
	if err != nil {
		panic(err)
	}

	err = mpv.Init(mpvHandle)
	if err != nil {
		panic(err)
	}

	err = mpv.RenderContextCreate(mpvHandle)
	if err != nil {
		panic(err)
	}

}
