package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
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
			rayColor := getRayColorRatio(Ray{origin, lowerLeftCorner.Add(horizontal.Mul(u), vertical.Mul(v))})
			r := uint8(255 * rayColor.x)
			g := uint8(255 * rayColor.y)
			b := uint8(255 * rayColor.z)
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
