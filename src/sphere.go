package main

import (
	"math"
)

// Sphere is a spherical Hitable
type Sphere struct {
	center   Vec3
	radius   float64
	material Material
}

// Hit returns True and updates HitRecord if it intersects the Sphere
func (s Sphere) Hit(r *Ray, tMin, tMax float64, rec *HitRecord) bool {
	oc := r.Origin.Minus(s.center)
	a := dot(r.Direction, r.Direction)
	b := dot(oc, r.Direction)
	c := dot(oc, oc) - s.radius*s.radius
	discriminant := b*b - a*c
	if discriminant > 0 {
		// Check both possible intersection points on the Sphere.
		// Update HitRecord if either point is between time threshold.
		temp := (-b - math.Sqrt(b*b-a*c)) / a
		if temp < tMax && temp > tMin {
			rec.t = temp
			rec.p = r.PointAtParameter(rec.t)
			rec.normal = rec.p.Minus(s.center).Divide(s.radius)
			rec.material = &s.material
			return true
		}
		temp = (-b + math.Sqrt(b*b-a*c)) / a
		if temp < tMax && temp > tMin {
			rec.t = temp
			rec.p = r.PointAtParameter(rec.t)
			rec.normal = rec.p.Minus(s.center).Divide(s.radius)
			rec.material = &s.material
			return true
		}
	}
	return false
}
