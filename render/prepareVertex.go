package render

import (
	"math/rand"
	"noise_playground/constants"
)

type Pixel struct {
	x     int
	y     int
	color float32
}

func (p Pixel) getVec3(screen_width, screen_height int) (float32, float32, float32) {
	new_x := fromScalarToOpenGlValue(p.x, 0, screen_width, -1.0, 1.0)
	new_y := fromScalarToOpenGlValue(p.y, 0, screen_height, -1.0, 1.0)
	return new_x, new_y, 0.0
}

func generatePixelArray() [constants.Height][constants.Width]Pixel {
	var buffer [constants.Height][constants.Width]Pixel
	for i := 0; i < constants.Height; i++ {
		for j := 0; j < constants.Width; j++ {
			buffer[i][j] = Pixel{i, j, rand.Float32()}
		}
	}
	return buffer
}

func generateVertexArray(pixels [constants.Height][constants.Width]Pixel) *[]float32 {
	vertexArray := make([]float32, constants.Height*constants.Width)
	for _, row := range pixels {
		for _, el := range row {
			pixel_x, pixel_y, pixel_z := el.getVec3(constants.Width, constants.Height)
			vertexArray = append(vertexArray, pixel_x)
			vertexArray = append(vertexArray, pixel_y)
			vertexArray = append(vertexArray, pixel_z)
		}
	}
	return &vertexArray
}

func fromScalarToOpenGlValue(x, in_min, in_max int, out_min, out_max float32) float32 {
	return (float32(x)-float32(in_min))*(out_max-out_min)/(float32(in_max)-float32(in_min)) + out_min
}
