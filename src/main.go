package main

import (
	"fmt"
)

// Main outputs a PPM image of a red-green color gradient
func main() {
	nx := 200
	ny := 100
	// Print PPM file header / metadata
	fmt.Printf("P3\n%d %d\n255\n", nx, ny)
	// Iterates through pixels from top-left (max Y) to
	// bottom-right (max X) pixels. Left-to-right, top-to-botton.
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			// Scale amount of red/green based on pixel coordinates
			r := float64(i) / float64(nx)
			g := float64(j) / float64(ny)
			b := 0.2

			// Print (Red, Green, Blue) as integers 0-255
			ir := int32(255.99 * r)
			ig := int32(255.99 * g)
			ib := int32(255.99 * b)
			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
}
