package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/jami/gglplayground/lib/graphic"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {
	fmt.Println("Golang OpenGL Playground")
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	window := graphic.NewWindow()

	startScene := graphic.NewScene()
	window.PushScene(startScene)

	window.Run()
}
