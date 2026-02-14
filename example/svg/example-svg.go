package main

import (
	"fmt"
	"github.com/bellina-ui/hal-glfw"
	gl3 "github.com/chsc/gogl/gl33"
	"github.com/shibukawa/nanovgo"
	"github.com/shibukawa/nanovgo/sample/demo"
	"log"
)

var ctx *nanovgo.Context
var demo_data *demo.DemoData

func on_after_gl_initialized() {
	var err error

	ctx, err = nanovgo.NewContext(0 /*nanovgo.AntiAlias | nanovgo.StencilStrokes | nanovgo.Debug*/)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("(+) Created nanovgo context")

	demo_data = load_demo_data(ctx)

	e := gl3.Init()

	if e != nil {
		panic("ok")
	}

	gl3.ClearColor(0.3, 0.3, 0.32, 1.0)
	gl3.Clear(gl3.COLOR_BUFFER_BIT | gl3.DEPTH_BUFFER_BIT | gl3.STENCIL_BUFFER_BIT)
}

func on_before_window_delete() {
	demo_data.FreeData(ctx)

	fmt.Println("(-) Deleting nanovg context")
	ctx.Delete()
}

var t float32
var w, h int = 0, 0

func on_loop() {
	t++

	gl3.Viewport(0, 0, gl3.Sizei(w), gl3.Sizei(h))
	gl3.ClearColor(0.3, 0.3, 0.32, 1.0)
	gl3.Clear(gl3.COLOR_BUFFER_BIT | gl3.DEPTH_BUFFER_BIT | gl3.STENCIL_BUFFER_BIT)
	gl3.Enable(gl3.BLEND)
	gl3.BlendFunc(gl3.SRC_ALPHA, gl3.ONE_MINUS_SRC_ALPHA)
	gl3.Enable(gl3.CULL_FACE)
	gl3.Disable(gl3.DEPTH_TEST)

	ctx.BeginFrame(w, h, 1)

	demo.RenderDemo(ctx, float32(1), float32(1), float32(w), float32(h), t, false, demo_data)

	ctx.EndFrame()

	gl3.Enable(gl3.DEPTH_TEST)
}

func on_resize(a, b int) {
	w, h = a, b

	fmt.Println(w, " ", h)
}

func main() {
	hal_glfw.Init(100, 100, 800, 600)

	hal_glfw.Set_Callbacks(on_after_gl_initialized, on_loop, on_before_window_delete, on_resize, nil, nil, nil)

	hal_glfw.Loop("Simple Jack!")
}

func load_demo_data(ctx *nanovgo.Context) *demo.DemoData {
	d := &demo.DemoData{}

	for i := 0; i < 12; i++ {
		// images will be in a different location on your machine
		// use a good search utility like "void" software's "everything" to find "image1.jpg" on your system
		// it should have been downloaded when the 'shibukawa' library was downloaded with go "get -u github.com/shibukawa/nanovgo"
		path := fmt.Sprintf("C:/goproj/pkg/mod/github.com/shibukawa/nanovgo@v0.0.0-20160822101109-9141d09b3652/sample/images/image%d.jpg", i+1)
		d.Images = append(d.Images, ctx.CreateImage(path, 0))

		if d.Images[i] == 0 {
			log.Fatalf("Could not load %s", path)
		}
	}

	d.FontIcons = ctx.CreateFont("icons", "C:/goproj/pkg/mod/github.com/shibukawa/nanovgo@v0.0.0-20160822101109-9141d09b3652/sample/entypo.ttf")

	if d.FontIcons == -1 {
		log.Fatalln("Could not add font icons.")
	}

	d.FontNormal = ctx.CreateFont("sans", "C:/goproj/pkg/mod/github.com/shibukawa/nanovgo@v0.0.0-20160822101109-9141d09b3652/sample/Roboto-Regular.ttf")

	if d.FontNormal == -1 {
		log.Fatalln("Could not add font italic.")
	}

	d.FontBold = ctx.CreateFont("sans-bold", "C:/goproj/pkg/mod/github.com/shibukawa/nanovgo@v0.0.0-20160822101109-9141d09b3652/sample/Roboto-Bold.ttf")

	if d.FontBold == -1 {
		log.Fatalln("Could not add font bold.")
	}

	return d
}
