package main

import (
	"math"
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

// Dielectric is a material that always refracts rays when possible.
// NOTE: If you create a Dielectric sphere with a negative radius within
// a Dielectric sphere with a slightly larger positive radius, it will
// create a 'bubble' / 'hollow glass sphere' effect.
type Dielectric struct {
	// refractIndex determines how much the path of a Ray is bent
	// when entering a Material. Some common examples:
	// Air: 1
	// Water: 1.3
	// Glass: 1.4-1.7
	// Diamond: 2.4
	refractIndex float64
}

// Scatter refracts a Ray if possible and reflects it otherwise
func (d Dielectric) Scatter(rIn *Ray, rec *HitRecord, attenuation *Vec3, scattered *Ray) bool {
	var outwardNormal, refracted Vec3
	var niOverNt, reflectProb, cosine float64
	reflected := reflect(rIn.Direction, rec.normal)

	// Always 1, glass surface absorbs nothing
	*attenuation = NewVec3(1, 1, 1)

	// Check if the ray should be refracted
	if dot(rIn.Direction, rec.normal) > 0 {
		outwardNormal = rec.normal.Times(-1)
		niOverNt = d.refractIndex
		cosine = d.refractIndex * dot(rIn.Direction, rec.normal) / rIn.Direction.Length()
	} else {
		outwardNormal = rec.normal
		niOverNt = 1.0 / d.refractIndex
		cosine = -1 * dot(rIn.Direction, rec.normal) / rIn.Direction.Length()
	}

	if refract(rIn.Direction, outwardNormal, niOverNt, &refracted) {
		reflectProb = schlick(cosine, d.refractIndex)
	} else {
		reflectProb = 1.0
	}

	if rand.Float64() < reflectProb {
		*scattered = Ray{rec.p, reflected}
	} else {
		*scattered = Ray{rec.p, refracted}
	}
	return true
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

// schlick allows glass to reflect in a way that varies with angle
func schlick(cosine, refractIndex float64) float64 {
	r0 := (1 - refractIndex) / (1 + refractIndex)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}
