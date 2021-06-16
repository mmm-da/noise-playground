package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	width  = 500
	height = 500
)

func main() {
	defer glfw.Terminate()
	runtime.LockOSThread()

	window := initGlfw()
	program := initOpenGL()

	for !window.ShouldClose() {
		draw(window, program)
	}
}

func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	prog := gl.CreateProgram()
	gl.LinkProgram(prog)
	return prog
}

func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "Noise playground", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

func draw(window *glfw.Window, program uint32) {
	gl.ClearColor(1.0, 0.0, 0.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	var VAO, VBO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.GenBuffers(1, &VBO)
	gl.BindVertexArray(VAO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)

	var pixelData [height][width]Pixel = generateRawPixels()
	fmt.Println(pixelData)

	glfw.PollEvents()
	window.SwapBuffers()
}

type Pixel struct {
	x     int
	y     int
	color float32
}

func generateRawPixels() [height][width]Pixel {
	var buffer [height][width]Pixel
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			buffer[i][j] = Pixel{i, j, 0.0}
		}
	}
	return buffer
}
