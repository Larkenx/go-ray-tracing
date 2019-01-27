package main

import (
	"math"
	"strconv"
)

// The Vector3D is a 3 dimension coordinate x, y, and z
type Vector3D struct {
	x, y, z float64
}

// String returns a string representation of p like "(3,4,5)".
func (v Vector3D) String() string {
	return "(" + strconv.FormatFloat(v.x, 'E', -1, 64) + "," + strconv.FormatFloat(v.y, 'E', -1, 64) + "," + strconv.FormatFloat(v.z, 'E', -1, 64) + ")"
}

// Add returns the vector p+q.
func (v Vector3D) Add(vectors ...Vector3D) Vector3D {
	result := Vector3D{v.x, v.y, v.z}
	for _, vec := range vectors {
		result.x += vec.x
		result.y += vec.y
		result.z += vec.z
	}
	return result
}

// Sub returns the vector p-q.
func (v Vector3D) Sub(vectors ...Vector3D) Vector3D {
	result := Vector3D{v.x, v.y, v.z}
	for _, vec := range vectors {
		result.x -= vec.x
		result.y -= vec.y
		result.z -= vec.z
	}
	return result
}

// Mul returns the vector p*k.
func (v Vector3D) Mul(k float64) Vector3D {
	return Vector3D{v.x * k, v.y * k, v.z * k}
}

// Div returns the vector p/k.
func (v Vector3D) Div(k float64) Vector3D {
	return Vector3D{v.x / k, v.y / k, v.z / k}
}

// Unit returns unit vector of v
func (v Vector3D) Unit() Vector3D {
	k := 1.0 / math.Sqrt(math.Pow(v.x, 2)+math.Pow(v.y, 2)+math.Pow(v.z, 2))
	return Vector3D{v.x * k, v.y * k, v.z * k}
}

// Dot returns the dot product of all the vectors
func (v Vector3D) Dot(vectors ...Vector3D) Vector3D {
	result := Vector3D{v.x, v.y, v.z}
	for _, vec := range vectors {
		result.x *= vec.x
		result.y *= vec.y
		result.z *= vec.z
	}
	return result
}
