package render

import (
	"noise_playground/shaders"
	"unsafe"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.0/glfw"
)

func RenderLoop(window *glfw.Window, program *shaders.Program) {

	program.Use()

	var VAO, VBO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.GenBuffers(1, &VBO)
	gl.BindVertexArray(VAO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 32, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	for !window.ShouldClose() {
		RenderTick(window, program, VBO, VAO)
	}
}

func RenderTick(window *glfw.Window, program *shaders.Program, VBO, VAO uint32) {
	gl.ClearColor(1.0, 0.0, 0.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	pixelArray := generatePixelArray()
	vertexArray := generateVertexArray(pixelArray)

	gl.BindVertexArray(VAO)
	gl.BufferData(gl.ARRAY_BUFFER, 3, unsafe.Pointer(vertexArray), gl.DYNAMIC_DRAW)
	gl.DrawArrays(gl.POINTS, 0, 3)

	glfw.PollEvents()
	window.SwapBuffers()
}
