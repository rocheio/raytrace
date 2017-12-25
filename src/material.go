package main

import (
	"math/rand"
)

// Material determines how Rays interact with a surface
type Material interface {
	Scatter(rIn *Ray, rec *HitRecord, attenuation *Vec3, scattered *Ray) bool
}

// Lambertian is the basic diffuse (matte) material.
// albedo determines how much flux is abosorbed from a Ray.
type Lambertian struct {
	albedo Vec3
}

// Scatter randomly reflects a Ray and absorbs some of its flux
func (l Lambertian) Scatter(rIn *Ray, rec *HitRecord, attenuation *Vec3, scattered *Ray) bool {
	target := randomInUnitSphere().Plus(rec.p).Plus(rec.normal)
	*scattered = Ray{rec.p, target.Minus(rec.p)}
	*attenuation = l.albedo
	return true
}

// Metal is a material that mathematically reflects Rays.
// fuzz can be between 0 and 1, where 0 is perfect reflectiveness
// and 1 is near-matte.
type Metal struct {
	albedo Vec3
	fuzz   float64
}

// Scatter mathematically reflects a Ray and absorbs some of its flux.
// The direction reflected is randomly adjusted based on the metal's fuzz.
func (m Metal) Scatter(rIn *Ray, rec *HitRecord, attenuation *Vec3, scattered *Ray) bool {
	reflected := reflect(rIn.Direction.AsUnit(), rec.normal)
	scatterDirection := randomInUnitSphere().Times(m.fuzz).Plus(reflected)
	*scattered = Ray{rec.p, scatterDirection}
	*attenuation = m.albedo
	return dot(scattered.Direction, rec.normal) > 0
}

// NewMetal returns a Metal object, ensuring 0 <= fuzz <= 1
func NewMetal(a Vec3, f float64) Metal {
	if f < 0 {
		f = 0
	} else if f > 1 {
		f = 1
	}
	return Metal{a, f}
}

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
