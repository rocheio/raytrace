package main

import (
	"math"
)

// Camera defines the frame boundaries and an origin for all rays
// for the 2d image that will be created from a 3d world
type Camera struct {
	lowerLeft  Vec3
	horizontal Vec3
	vertical   Vec3
	origin     Vec3
}

// GetRay returns a Ray going from the origin to the pixel at (u, v)
func (c Camera) GetRay(u, v float64) Ray {
	direction := c.lowerLeft.
		Plus(c.horizontal.Times(u)).
		Plus(c.vertical.Times(v)).
		Minus(c.origin)
	return Ray{c.origin, direction}
}

// NewCamera returns a Camera object from:
// lookFrom: The source of Rays from the Camera
// lookAt: The center of the scene to send Rays at
// viewUp: Defines the 'up' direction of the camera, allows roll / tilt
// vFOV: Vertical field of view, 0-180 degrees. How 'big' an image feels.
// aspect: (Width / Height) ratio of the image size
func NewCamera(lookFrom Vec3, lookAt Vec3, viewUp Vec3, vFOV float64, aspect float64) Camera {
	theta := vFOV * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight
	w := lookFrom.Minus(lookAt).AsUnit()
	u := cross(viewUp, w).AsUnit()
	v := cross(w, u)

	origin := lookFrom
	lowerLeft := origin.
		Minus(u.Times(halfWidth)).
		Minus(v.Times(halfHeight)).
		Minus(w)
	horizontal := u.Times(halfWidth * 2)
	vertical := v.Times(halfHeight * 2)
	return Camera{
		lowerLeft,
		horizontal,
		vertical,
		origin,
	}
}
