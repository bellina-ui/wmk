package main

import (
	"github.com/bellina-ui/wmk"
	gl3 "github.com/chsc/gogl/gl33"
	"runtime"
)

func on_after_gl_initialized() {
	e := gl3.Init()

	if e != nil {
		panic("Unable to initialize gl3")
	}

	gl3.ClearColor(0.93, 0.3, 0.32, 1.0)
	gl3.Clear(gl3.COLOR_BUFFER_BIT | gl3.DEPTH_BUFFER_BIT | gl3.STENCIL_BUFFER_BIT)
}

func on_loop() {
	gl3.Viewport(0, 0, 640, 480)

	gl3.ClearColor(0.93, 0.3, 0.32, 1.0)

	gl3.Clear(gl3.COLOR_BUFFER_BIT | gl3.DEPTH_BUFFER_BIT | gl3.STENCIL_BUFFER_BIT)
}

func on_resize(width, height int) {
}

func init() {
	// This bit is required for GLFW.
	runtime.LockOSThread()
}

func main() {
	wmk.Init(100, 100, 800, 600)

	wmk.Set_Callbacks(on_after_gl_initialized, on_loop, nil, on_resize, nil, nil, nil)

	wmk.Loop("Hello, OpenGL!")
}
