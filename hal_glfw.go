package hal_glfw

import (
	"fmt"
	"github.com/bellina-ui/hal"
	"github.com/goxjs/gl"
	"github.com/goxjs/glfw"
	"time"
)

var g_window_left, g_window_top int
var g_user_on_tick func()

// hal_glfw exports these variables for use by the user of the library
var Window_Width, Window_Height int
var Mouse_X, Mouse_Y int

func Init(left, top, width, height int) {

	if err := glfw.Init(gl.ContextWatcher); err != nil {
		panic("failed to initialize glfw")
	}

	g_window_left, g_window_top = left, top

	glfw.WindowHint(glfw.Resizable, int(glfw.Resizable))
	glfw.WindowHint(glfw.Samples, 4)

	Window_Width, Window_Height = width, height
}

func Set_Callbacks(
	on_after_gl_initialized,
	on_tick func(),
	on_before_window_delete func(),
	on_resize func(width, height int),
	on_mouse_move func(x, y int),
	on_mouse_button func(button hal.Mouse_Button, action hal.Button_Action),
	on_key func(key hal.Keyboard_Key, action hal.Button_Action, alt, ctrl, shift bool)) {

	g_user_on_after_gl_initialized = on_after_gl_initialized
	g_user_on_tick = on_tick
	g_user_on_before_window_delete = on_before_window_delete
	g_user_on_resize = on_resize
	g_user_on_mouse_move = on_mouse_move
	g_user_on_mouse_button = on_mouse_button
	g_user_on_key = on_key
}

func Loop(title string) {

	create_window(title)

	if g_user_on_after_gl_initialized != nil {
		g_user_on_after_gl_initialized()
	}

	glfw.SwapInterval(1)

	for !g_window.ShouldClose() {

		then := time.Now().UnixNano()

		if g_user_on_tick != nil {
			g_user_on_tick()
		}

		g_window.SwapBuffers()

		glfw.PollEvents()

		// todo come up with way to limit to 60fps
		for time.Now().UnixNano() - then < 10 * int64(time.Millisecond) { 
			time.Sleep(10 * time.Millisecond) // 2ms
		}
	}

	if g_user_on_before_window_delete != nil {
		g_user_on_before_window_delete()
	}

	glfw.Terminate()

	fmt.Println("(-) GLFW Uninitialized (hal_glfw)")
}


