package main

import (
	"noise_playground/constants"
	"noise_playground/render"
	"noise_playground/shaders"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(constants.Width, constants.Height, "Noise playground", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

func main() {
	defer glfw.Terminate()
	runtime.LockOSThread()

	vertexShader, err := shaders.NewShaderFromFile("./glsl/vertex.glsl", gl.VERTEX_SHADER)
	panic_if_error(err)

	fragmentShader, err := shaders.NewShaderFromFile("./glsl/fragment.glsl", gl.FRAGMENT_SHADER)
	panic_if_error(err)

	program, err := shaders.NewProgram(vertexShader, fragmentShader)
	panic_if_error(err)

	err = program.Link()
	panic_if_error(err)

	window := initGlfw()

	render.RenderLoop(window, program)

}
