package main

import (
	"math"
	"math/rand"
)

// Camera defines the frame boundaries and an origin for all rays
// for the 2d image that will be created from a 3d world
type Camera struct {
	lowerLeft  *Vec3
	horizontal *Vec3
	vertical   *Vec3
	origin     *Vec3
	u, v, w    *Vec3
	lensRadius float64
}

// GetRay returns a Ray going from a random point in the Camera's lens
// to the pixel at (s, t) in the virtual film plane.
func (c Camera) GetRay(s, t float64) Ray {
	point := randomInUnitDisk().Times(c.lensRadius)
	offset := c.u.Times(point.x()).Plus(c.v.Times(point.y()))
	direction := c.lowerLeft.
		Plus(c.horizontal.Times(s)).
		Plus(c.vertical.Times(t)).
		Minus(c.origin).
		Minus(offset)
	return Ray{c.origin.Plus(offset), direction}
}

// NewCamera returns a Camera object from:
// lookFrom: The source of Rays from the Camera
// lookAt: The center of the scene to send Rays at
// viewUp: Defines the 'up' direction of the camera, allows roll / tilt
// vFOV: Vertical field of view, 0-180 degrees. How 'big' an image feels.
// aspect: (Width / Height) ratio of the image size
func NewCamera(lookFrom, lookAt, viewUp *Vec3, vFOV float64, aspect float64, aperature float64, focusDist float64) Camera {
	lensRadius := aperature / 2
	theta := vFOV * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight

	w := lookFrom.Minus(lookAt).AsUnit()
	u := cross(viewUp, w).AsUnit()
	v := cross(w, u)

	origin := lookFrom
	lowerLeft := origin.
		Minus(u.Times(halfWidth * focusDist)).
		Minus(v.Times(halfHeight * focusDist)).
		Minus(w.Times(focusDist))
	horizontal := u.Times(halfWidth * 2 * focusDist)
	vertical := v.Times(halfHeight * 2 * focusDist)
	return Camera{
		lowerLeft,
		horizontal,
		vertical,
		origin,
		u,
		v,
		w,
		lensRadius,
	}
}

// randomInUnitDisk returns a random point within a virtual 3d disk.
// Useful for approximating an aperture / depth of field from a Camera.
func randomInUnitDisk() *Vec3 {
	var p *Vec3
	for {
		p = NewVec3(rand.Float64(), rand.Float64(), 0).Times(2).
			Minus(NewVec3(1, 1, 0))
		if dot(p, p) < 1 {
			return p
		}
	}
}
