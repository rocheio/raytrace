package main

import (
	"fmt"
)

// Vec3 is a 3-dimensional vector
// It can represent space (X, Y, Z) or color (R, G, B)
type Vec3 struct {
	e [3]float64
}

// PrintInts prints the Vec3 values as a line of integers
func (v Vec3) PrintInts() {
	fmt.Printf("%d %d %d\n", int64(v.e[0]), int64(v.e[1]), int64(v.e[2]))
}

// Times returns a new vector multiplied by a float
func (v Vec3) Times(f float64) Vec3 {
	return NewVec3(v.e[0]*f, v.e[1]*f, v.e[2]*f)
}

// NewVec3 returns a Vec3 with arguments transposed to a Vec3 slice
func NewVec3(e0, e1, e2 float64) Vec3 {
	return Vec3{
		e: [3]float64{e0, e1, e2},
	}
}
