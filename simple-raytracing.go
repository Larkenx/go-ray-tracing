package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func gradient() {
	// Create a blank image 100x200 pixels
	width := 200
	height := 100
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	bounds := img.Bounds()

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r := uint8(255 * x / width)
			g := uint8(255 * y / height)
			b := uint8(255 * 0.2)
			img.SetRGBA(x, y, color.RGBA{r, g, b, 255})
		}
	}

	// fmt.Println(img.Pix)

	outputFile, err := os.Create("output.png")
	if err != nil {
		fmt.Println("Uh oh, could not save image!")
		return
	}

	png.Encode(outputFile, img)
	outputFile.Close()
}

func getRayColorRatio(r Ray) Vector3D {
	directionUnitVector := r.direction.Unit()
	t := float64(0.5 * (directionUnitVector.y + 1.0))
	return Vector3D{1, 1, 1}.Mul(1 - t).Add(Vector3D{0.5, 0.7, 1}.Mul(t))
}

func sub(p float64, q float64) float64 {
	return p - q
}

func hitsSphere(center Vector3D, radius float64, r Ray) float64 {
	oc := r.origin.Sub(center)
	a := r.direction.Dot(r.direction)
	b := float64(oc.Dot(r.direction) * 2)
	c := float64(oc.Dot(oc) - radius*radius)
	discriminant := float64(b*b - 4*a*c)
	if discriminant < 0 {
		return -1.0
	}
	return (-b - math.Sqrt(discriminant)) / (2.0 * a)
}

func colorSphere(r Ray) Vector3D {
	center := Vector3D{0, 0, -1}
	t := hitsSphere(center, 0.5, r)
	if t > 0.0 {
		n := r.PositionAt(t).Sub(Vector3D{0, 0, -1}).Unit()
		return Vector3D{n.x + 1, n.y + 1, n.z + 1}.Mul(0.5)
	}
	unitDirection := r.direction.Unit()
	t = 0.5 * (unitDirection.y + 1.0)
	return Vector3D{1, 1, 1}.Mul(1 - t).Add(Vector3D{0.5, 0.7, 1}.Mul(t))
}

func main() {
	// Create a blank image 100x200 pixels
	width := 200
	height := 100
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	bounds := img.Bounds()

	lowerLeftCorner := Vector3D{-2, -1, -1}
	horizontal := Vector3D{4, 0, 0}
	vertical := Vector3D{0, 2, 0}
	origin := Vector3D{0, 0, 0}

	fmt.Println(bounds)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			u := float64(x) / float64(width)
			v := float64(y) / float64(height)
			ray := Ray{origin: origin, direction: lowerLeftCorner.Add(horizontal.Mul(u), vertical.Mul(v))}
			// colorRatios := getRayColorRatio(ray)
			colorRatios := colorSphere(ray)
			r := uint8(255 * colorRatios.x)
			g := uint8(255 * colorRatios.y)
			b := uint8(255 * colorRatios.z)
			img.SetRGBA(x, y, color.RGBA{r, g, b, 255})
		}
	}

	// fmt.Println(img.Pix)

	outputFile, err := os.Create("output.png")
	if err != nil {
		fmt.Println("Uh oh, could not save image!")
		return
	}

	png.Encode(outputFile, img)
	outputFile.Close()
}
