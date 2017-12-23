package main

// Camera defines the frame boundaries and an origin for all rays
// for the 2d image that will be created from a 3d world
type Camera struct {
	lowerLeft, horizontal, vertical, origin Vec3
}

// GetRay returns a Ray going from the origin to the pixel at (u, v)
func (c Camera) GetRay(u, v float64) Ray {
	direction := c.lowerLeft.Plus(c.horizontal.Times(u)).Plus(c.vertical.Times(v))
	return Ray{c.origin, direction}
}
