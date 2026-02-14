package main

import (
	"fmt"
	"github.com/bellina-ui/hal"
	"github.com/bellina-ui/hal-glfw"
	"runtime"
)

func on_loop() {
	// this gets called every loop...use it wisely
}

func on_after_gl_initialized() {
	// This callback gets called after the OpenGL context has been
	// initialized.
	// All OpenGL calls are valid now
	fmt.Println("OpenGL context initialized")
}

func on_before_window_delete() {
	// Right before the window is deleted, this callback can be used
	fmt.Println("Window is being deleted")
	// to free up resources.
}

func on_resize(width, height int) {
	// This callback gets called for every resize of the window.
	fmt.Printf("Window resized to: %d, %d\n", width, height)
}

func on_mouse_move(x, y int) {
	// This callback gets called for every movement of the mouse.
	fmt.Printf("Mouse moved to: %d, %d\n", x, y)
}

func on_mouse_button(button hal.Mouse_Button, action hal.Button_Action) {
	// This callback gets called for every press of the mouse button.
	fmt.Printf("Mouse button pressed: %s\n", button)
}

func on_key(key hal.Keyboard_Key, action hal.Button_Action, alt, ctrl, shift bool) {
	// This callback gets called for every press of the keyboard.
	fmt.Printf("Key pressed: %s\n", key)
}

func init() {
	// This bit is required for GLFW.
	runtime.LockOSThread()
}

func main() {

	// Create an 800 by 600 pixel window (is this 1997 again!)
	hal_glfw.Init(100, 100, 800, 600)

	// Setup the callbacks
	hal_glfw.Set_Callbacks(on_after_gl_initialized, on_loop, on_before_window_delete, on_resize, on_mouse_move, on_mouse_button, on_key)

	// Start the loop!
	hal_glfw.Loop("Hello, World!")
}
