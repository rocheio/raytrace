package main

import (
	"fmt"
	"math"
)

// Vec3 is a 3-dimensional vector
type Vec3 struct {
	e [3]float64
}

// Vec3 can represent space (X, Y, Z)
func (v Vec3) x() float64 { return v.e[0] }
func (v Vec3) y() float64 { return v.e[1] }
func (v Vec3) z() float64 { return v.e[2] }

// Vec3 can represent color (R, G, B)
func (v Vec3) r() float64 { return v.e[0] }
func (v Vec3) g() float64 { return v.e[1] }
func (v Vec3) b() float64 { return v.e[2] }

// Length returns the mathematical length of a 3d vector
func (v *Vec3) Length() float64 {
	return math.Sqrt(v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2])
}

// PrintInts prints the Vec3 values as a line of integers
func (v Vec3) PrintInts() {
	fmt.Printf("%d %d %d\n", int64(v.e[0]), int64(v.e[1]), int64(v.e[2]))
}

// Times returns a new vector multiplied by a float
func (v Vec3) Times(f float64) Vec3 {
	return NewVec3(v.e[0]*f, v.e[1]*f, v.e[2]*f)
}

// Divide returns a new vector divided by a float
func (v Vec3) Divide(f float64) Vec3 {
	return NewVec3(v.e[0]/f, v.e[1]/f, v.e[2]/f)
}

// Pow returns a new vector raised to a giver power
func (v Vec3) Pow(f float64) Vec3 {
	return NewVec3(
		math.Pow(v.e[0], f),
		math.Pow(v.e[1], f),
		math.Pow(v.e[2], f),
	)
}

// Minus returns a new vector equal to one vec subtracted from another
func (v Vec3) Minus(other Vec3) Vec3 {
	return NewVec3(
		v.e[0]-other.e[0],
		v.e[1]-other.e[1],
		v.e[2]-other.e[2],
	)
}

// Plus returns a new vector from two added vectors
func (v Vec3) Plus(other Vec3) Vec3 {
	return NewVec3(
		v.e[0]+other.e[0],
		v.e[1]+other.e[1],
		v.e[2]+other.e[2],
	)
}

// Cross returns a Vec3 from two vectors multiplied together
func (v Vec3) Cross(other Vec3) Vec3 {
	return NewVec3(
		v.e[0]*other.e[0],
		v.e[1]*other.e[1],
		v.e[2]*other.e[2],
	)
}

// AsUnit returns a new, similar, Vec3 scaled between -1 and 1
func (v Vec3) AsUnit() Vec3 {
	return v.Divide(v.Length())
}

// NewVec3 returns a Vec3 with arguments transposed to a Vec3 slice
func NewVec3(e0, e1, e2 float64) Vec3 {
	return Vec3{
		e: [3]float64{e0, e1, e2},
	}
}

// dot returns the sum of two vectors multiplied together
func dot(v1, v2 Vec3) float64 {
	return v1.e[0]*v2.e[0] + v1.e[1]*v2.e[1] + v1.e[2]*v2.e[2]
}

// reflect returns the Vec3 reflected from an input Vec3 and a surface normal
func reflect(v, n Vec3) Vec3 {
	return v.Minus(n.Times(2 * dot(v, n)))
}
