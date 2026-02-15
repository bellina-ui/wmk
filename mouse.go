package wmk

import (
	"fmt"

	"github.com/bellina-ui/hal"
	"github.com/goxjs/glfw"
)

var g_user_on_mouse_move func(x, y int)
var g_user_on_mouse_button func(button hal.Mouse_Button, action hal.Button_Action)

func glfw_on_mouse_move(
	window *glfw.Window,
	x, y float64) {

	Mouse_X, Mouse_Y = int(x), int(y)

	if g_user_on_mouse_move != nil {
		g_user_on_mouse_move(int(x), int(y))
	}
}

func glfw_on_mouse_button(
	window *glfw.Window,
	button glfw.MouseButton,
	action glfw.Action,
	mods glfw.ModifierKey) {

	var _button hal.Mouse_Button
	var _action hal.Button_Action

	if button == glfw.MouseButtonLeft {
		_button = hal.Mouse_Button_LEFT

	} else if button == glfw.MouseButtonRight {
		_button = hal.Mouse_Button_RIGHT

	} else {

		fmt.Println("Unrecognized mouse button %i in wmk/mouse.go", button)
		return
	}

	if action == glfw.Press {
		_action = hal.Button_Action_DOWN

	} else if action == glfw.Release {
		_action = hal.Button_Action_UP

	} else {

		fmt.Println("Unrecognized action %i in wmk/mouse.go", action)
		return
	}

	if g_user_on_mouse_button != nil {
		g_user_on_mouse_button(_button, _action)
	}
}
