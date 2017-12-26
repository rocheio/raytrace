package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// colorFromRay returns the (R, G, B) value of a pixel based on
// how a Ray interacts with objects in a HitableList.
func colorFromRay(r *Ray, world HitableList, depth int32) Vec3 {
	var rec HitRecord

	// Use 0.001 tMin to avoid hits near zero causing shadow acne
	if world.Hit(r, 0.001, math.MaxFloat64, &rec) {
		// Scatter the Ray and absorb some of its flux based on the Material.
		// Recurse with the scattered Ray until it does not hit an object.
		// If a Ray scatters over 50 times, just return pure black.
		var scattered Ray
		var attenuation Vec3
		m := *rec.material

		if depth < 50 && m.Scatter(r, &rec, &attenuation, &scattered) {
			return colorFromRay(&scattered, world, depth+1).Cross(attenuation)
		}

		return NewVec3(0, 0, 0)
	}

	// If no object is hit, draw a white to teal gradient background
	// based on the scaled Y value of the direction of a Ray
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
		// Huge green matte sphere on bottom (terrain/ground)
		Sphere{NewVec3(0, -100.5, -1), 100, Lambertian{NewVec3(0.8, 0.8, 0)}},
		// Hollow glass sphere on left
		Sphere{NewVec3(-1, 0, -1), 0.5, Dielectric{1.5}},
		Sphere{NewVec3(-1, 0, -1), -0.45, Dielectric{1.5}},
		// Blue matte sphere in middle
		Sphere{NewVec3(0, 0, -1), 0.5, Lambertian{NewVec3(0.1, 0.2, 0.5)}},
		// Shiny gold metal sphere on right
		Sphere{NewVec3(1, 0, -1), 0.5, Metal{NewVec3(0.8, 0.6, 0.2), 0}},
	}
	world := HitableList{hitables, len(hitables)}

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
				color = color.Plus(colorFromRay(&ray, world, 0))
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
