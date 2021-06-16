package main

func remap(x, in_min, in_max int, out_min, out_max float32) float32 {
	return (float32(x)-float32(in_min))*(out_max-out_min)/(float32(in_max)-float32(in_min)) + out_min
}
