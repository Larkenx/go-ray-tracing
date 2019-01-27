package main

// Ray has two vectors - representing the 3D origin, and the direction in which the ray is heading
type Ray struct {
	origin, direction Vector3D
}

// PositionAt returns the vector in 3D space following the ray's trajectory from origin in its direction
func (r Ray) PositionAt(t float64) Vector3D {
	return r.origin.Add(r.direction.Mul(t))
}
