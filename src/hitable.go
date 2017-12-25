package main

// HitRecord stores the time, hit point, surface normal, and material
// of the last object a Ray hit during tracing
type HitRecord struct {
	t        float64
	p        Vec3
	normal   Vec3
	material *Material
}

// Hitable is any object in the scene that could be hit by a Ray
type Hitable interface {
	Hit(r *Ray, tMin, tMax float64, rec *HitRecord) bool
}

// HitableList contains many objects in the scene
type HitableList struct {
	hitables []Hitable
	length   int
}

// Hit checks each object in the list and updates HitRecord
// with the closest object (if any) hit by a Ray
func (h HitableList) Hit(r *Ray, tMin, tMax float64, rec *HitRecord) bool {
	var tempRec HitRecord
	hitAnything := false
	closestSoFar := tMax
	for i := 0; i < h.length; i++ {
		if h.hitables[i].Hit(r, tMin, closestSoFar, &tempRec) {
			hitAnything = true
			closestSoFar = tempRec.t
			*rec = tempRec
		}
	}
	return hitAnything
}
