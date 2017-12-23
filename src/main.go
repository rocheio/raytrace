package main

import (
	"fmt"
	"math"
)

// colorFromRay returns the (R, G, B) value of a pixel based on
// how a Ray interacts with objects in a HitableList
func colorFromRay(r *Ray, world HitableList) Vec3 {
	var rec HitRecord

	if world.Hit(r, 0, math.MaxFloat64, &rec) {
		// Return color to visualize surface normal from hit point
		return NewVec3(
			rec.normal.x()+1,
			rec.normal.y()+1,
			rec.normal.z()+1,
		).Times(0.5)
	}

	// Draw a white to teal gradient background based on the scaled
	// Y value of the direction of a Ray
	white := NewVec3(1.0, 1.0, 1.0)
	teal := NewVec3(0.5, 0.7, 1.0)
	// Scale Y value to between 0 and 1
	t := 0.5 * (r.Direction.AsUnit().y() + 1.0)
	// Blend white and teal in proportion to normalized Y value
	return white.Times(1.0 - t).Plus(teal.Times(t))
}

// Main outputs a PPM image of the scene to stdout
func main() {
	nx := 200
	ny := 100

	// Define the bounds of the scene and the origin of all Rays
	lowerLeft := NewVec3(-2.0, -1.0, -1.0)
	horizontal := NewVec3(4.0, 0.0, 0.0)
	vertical := NewVec3(0.0, 2.0, 0.0)
	origin := NewVec3(0.0, 0.0, 0.0)

	// Define objects in the scene
	hitables := []Hitable{
		// little sphere on top
		Sphere{NewVec3(0, 0, -1), 0.5},
		// huge sphere on bottom (world)
		Sphere{NewVec3(0, -100.5, -1), 100},
	}
	world := HitableList{hitables, 2}

	// Print PPM file header / metadata
	fmt.Printf("P3\n%d %d\n255\n", nx, ny)

	// Iterates through pixels from top-left (max Y) to
	// bottom-right (max X) pixels. Left-to-right, top-to-botton.
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {

			// Calculate (X, Y) where Ray intersects the background
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)

			// Create a Ray going from the static origin to this pixel in canvas
			direction := lowerLeft.Plus(horizontal.Times(u)).Plus(vertical.Times(v))
			r := Ray{origin, direction}

			// Calculate the color based on the Ray and objects in the scene
			color := colorFromRay(&r, world)

			// Print (Red, Green, Blue) as integers 0-255
			color.Times(255.99).PrintInts()
		}
	}
}
