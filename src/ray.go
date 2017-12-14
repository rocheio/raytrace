package main

// Ray points from one Vec3 in 3d space to another
type Ray struct {
	Origin, Direction Vec3
}

// PointAtParameter returns the Vec3 located along a Ray at a time
func (r *Ray) PointAtParameter(t float64) Vec3 {
	return r.Origin.Plus(r.Direction.Times(t))
}
