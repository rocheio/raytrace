package main

import (
	"fmt"
)

// Color returns white to teal based on the Y value of the direction of a Ray
func colorFromRay(r *Ray) Vec3 {
	white := NewVec3(1.0, 1.0, 1.0)
	teal := NewVec3(0.5, 0.7, 1.0)
	// Scale Y value to between 0 and 1
	t := 0.5 * (r.Direction.AsUnit().y() + 1.0)
	// Blend white and teal in proportion to normalized Y value
	return white.Times(1.0 - t).Plus(teal.Times(t))
}

// Main outputs a PPM image of a red-green color gradient
func main() {
	nx := 200
	ny := 100

	// Define the bounds of the scene and the Ray within it
	lowerLeft := NewVec3(-2.0, -1.0, -1.0)
	horizontal := NewVec3(4.0, 0.0, 0.0)
	vertical := NewVec3(0.0, 2.0, 0.0)
	origin := NewVec3(0.0, 0.0, 0.0)

	// Print PPM file header / metadata
	fmt.Printf("P3\n%d %d\n255\n", nx, ny)

	// Iterates through pixels from top-left (max Y) to
	// bottom-right (max X) pixels. Left-to-right, top-to-botton.
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {

			// Calculate (X, Y) where Ray intersects the background
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)

			direction := lowerLeft.Plus(horizontal.Times(u)).Plus(vertical.Times(v))
			r := Ray{origin, direction}

			color := colorFromRay(&r)

			// Print (Red, Green, Blue) as integers 0-255
			color.Times(255.99).PrintInts()
		}
	}
}
