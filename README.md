# Package `hal-glfw`

Gives GoLang applications a Window with an OpenGL (glfw) context with callbacks for Mouse and Keyboard.

# Installation

`go get -u github.com/bellina-ui/hal-glfw`

# why does hal-glfw exist?

`hal-glfw` is used by the `Bellina UI Library`

# usage

Pay careful attention to the init() function in the example.
runtime.LockOSThread() must be called in init().

```go
package main

import (
	"github.com/bellina-ui/hal_glfw"
	"runtime"
)

func on_loop() {
	// this gets called every loop...use it wisely
}

func on_after_gl_initialized() {
	// This callback gets called after the OpenGL context has been
	// initialized.
	// All OpenGL calls are valid now
}

func on_before_window_delete() {
	// Right before the window is deleted, this callback can be used
	// to free up resources.
}

func on_resize(width, height int) {
	// This callback gets called for every resize of the window.
}

func on_mouse_move(x, y int) {
	// This callback gets called for every movement of the mouse.
}

func on_mouse_button(button xel.MouseButton, action xel.Button_Action) {
	// This callback gets called for every press of the mouse button.
}

func on_key(key xel.KeyboardKey, action xel.Button_Action, alt, ctrl, shift bool) {
	// This callback gets called for every press of the keyboard.
}

func init() {
	// This bit is required for GLFW.
	runtime.LockOSThread()
}

func main() {

	// Create an 800 by 600 pixel window (is this 1997 again!)
	xel.Init(100, 100, 800, 600)

	// Setup the callbacks
	xel.Set_Callbacks(on_after_gl_initialized, on_loop, on_before_window_delete, on_resize, on_mouse_move, on_mouse_button, on_key)

	// Start the loop!
	xel.Loop("Hello, World!")
}
```

# running the examples
```
 go run .\hal-glfw\example\simple\example-simple.go
 go run .\hal-glfw\example\opengl\example-opengl.go

 // this WILL NOT run on your machine until you read the comment on line 87 in "example-svg.go"
 go run .\hal-glfw\example\svg\example-svg.go
```
