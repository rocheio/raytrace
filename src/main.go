package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// randomInUnitSphere returns a random vector within a unit sphere
func randomInUnitSphere() Vec3 {
	var point Vec3
	for {
		// Calculate a random point in a unit cube
		point = NewVec3(
			rand.Float64(),
			rand.Float64(),
			rand.Float64(),
		).Times(2).Minus(NewVec3(1, 1, 1))

		// If the point is within a unit sphere, return it
		if point.Length()*point.Length() < 1 {
			return point
		}
	}
}

// colorFromRay returns the (R, G, B) value of a pixel based on
// how a Ray interacts with objects in a HitableList
func colorFromRay(r *Ray, world HitableList) Vec3 {
	var rec HitRecord

	// Use 0.001 tMin to avoid hits near zero causing shadow acne
	if world.Hit(r, 0.001, math.MaxFloat64, &rec) {
		// Absorb half the color of the Ray and reflect it randomly.
		// Recurse with the reflected Ray until it does not hit an object.
		target := rec.p.Plus(rec.normal).Plus(randomInUnitSphere())
		ray := Ray{rec.p, target.Minus(rec.p)}
		return colorFromRay(&ray, world).Times(0.5)
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

// init seeds the random library for true randomness
func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// main outputs a PPM image of the scene to stdout of size (nx, ny) pixels
func main() {
	nx := 200
	ny := 100
	numSamples := 100

	// Define boundaries of and objects in the scene
	camera := Camera{
		NewVec3(-2, -1, -1),
		NewVec3(4, 0, 0),
		NewVec3(0, 2, 0),
		NewVec3(0, 0, 0),
	}
	hitables := []Hitable{
		// little sphere on top (focus of scene)
		Sphere{NewVec3(0, 0, -1), 0.5},
		// huge sphere on bottom (terrain/ground)
		Sphere{NewVec3(0, -100.5, -1), 100},
	}
	world := HitableList{hitables, 2}

	// Print PPM file header / metadata
	fmt.Printf("P3\n%d %d\n255\n", nx, ny)

	// Iterates through pixels from top-left (max Y) to
	// bottom-right (max X) pixels. Left-to-right, top-to-botton.
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			// Start with an empty color vector for each pixel
			color := NewVec3(0, 0, 0)

			// Take many samples for each pixel, randomly using pixels
			// adjacent to the target pixel to provide anti-aliasing
			for s := 0; s < numSamples; s++ {
				// Calculate (X, Y) where this sample intersects the background
				u := (float64(i) + rand.Float64()) / float64(nx)
				v := (float64(j) + rand.Float64()) / float64(ny)
				// Aggregate the colors of each sample
				ray := camera.GetRay(u, v)
				color = color.Plus(colorFromRay(&ray, world))
			}

			// Average the color from all the samples
			color = color.Divide(float64(numSamples))
			// Correct color to gamma 2
			color = color.Pow(0.5)
			// Print (Red, Green, Blue) as integers 0-255
			color.Times(255.99).PrintInts()
		}
	}
}
