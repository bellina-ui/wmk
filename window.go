package wmk

import (
	"fmt"

	"github.com/goxjs/glfw"
)

var g_window *glfw.Window

var g_user_on_after_gl_initialized func()
var g_user_on_before_window_delete func()
var g_user_on_resize func(width, height int)

// cannot change the width/height types - must be "int"
func glfw_on_resize(window *glfw.Window, width int, height int) {
	Window_Width, Window_Height = width, height

	if g_user_on_resize != nil {
		g_user_on_resize(width, height)
	}
}

func create_window(title string) {
	g_window, _ = glfw.CreateWindow(int(Window_Width), int(Window_Height), title, nil, nil)

	g_window.SetPos(g_window_left, g_window_top)

	g_window.MakeContextCurrent()

	fmt.Println("(+) GLFW Initialized (wmk)")

	glfw.SwapInterval(0)

	g_window.SetCursorPosCallback(glfw_on_mouse_move)
	g_window.SetMouseButtonCallback(glfw_on_mouse_button)
	g_window.SetKeyCallback(glfw_on_key)
	g_window.SetSizeCallback(glfw_on_resize)

	width, height := g_window.GetSize()
	glfw_on_resize(g_window, width, height)
}

