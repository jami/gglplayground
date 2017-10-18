package graphic

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/jami/gglplayground/lib/system"
)

// Window helper struct
type Window struct {
	Size struct {
		Width  int
		Height int
	}

	currentScene *Scene
	timer        system.Timer
}

// PushScene performes the transition to another scene
func (w *Window) PushScene(scene *Scene) {
	w.currentScene = scene
}

// updateScene updates the current scene
func (w *Window) updateScene() {
	w.timer.Update()

	if w.currentScene != nil {
		w.currentScene.Update(w.timer.GetTick())
	}
}

func (w *Window) setupScene() {
	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.LIGHTING)

	gl.ClearColor(0.5, 0.5, 0.5, 0.0)
	gl.ClearDepth(1)
	gl.DepthFunc(gl.LEQUAL)

	ambient := []float32{0.5, 0.5, 0.5, 1}
	diffuse := []float32{1, 1, 1, 1}
	lightPosition := []float32{-5, 5, 10, 0}
	gl.Lightfv(gl.LIGHT0, gl.AMBIENT, &ambient[0])
	gl.Lightfv(gl.LIGHT0, gl.DIFFUSE, &diffuse[0])
	gl.Lightfv(gl.LIGHT0, gl.POSITION, &lightPosition[0])
	gl.Enable(gl.LIGHT0)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Frustum(-1, 1, -1, 1, 1.0, 10.0)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
}

// Run main loop
func (w *Window) Run() {
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)

	window, err := glfw.CreateWindow(w.Size.Width, w.Size.Height, "GGL Playground", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	w.setupScene()
	w.timer.Reset()

	for !window.ShouldClose() {
		w.updateScene()
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

// NewWindow creates a new window
func NewWindow() *Window {
	w := &Window{}
	w.Size.Width = 800
	w.Size.Height = 600
	return w
}
