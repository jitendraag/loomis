package pkg

import (
	"fmt"
	"image"
	"image/color"
	"math"
)

// NOTE: This method shouldn't really be used in production, use fast fourier tranform implementation instead.
func DiscreetFourierTransform(img image.Image) image.Image {
	// This is from Section 4.6.4 of DIP book
	// return img
	bounds := img.Bounds()
	var pixels [][]color.Gray
	var transformedReal [][]float64
	var transformedImaginary [][]float64

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		var transformed []float64
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			transformed = append(transformed, 0.0)
		}
		transformedReal = append(transformedReal, transformed)
		transformedImaginary = append(transformedImaginary, transformed)
	}

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for u := bounds.Min.X; u < bounds.Max.X; u++ {
				for v := bounds.Min.Y; v < bounds.Max.Y; v++ {
					w := -2.0 * math.Pi * float64(x*u/bounds.Max.X+y*v/bounds.Max.Y)
					transformedReal[u][v] += float64((color.GrayModel.Convert(img.At(x, y)).(color.Gray)).Y) * math.Cos(w)
					transformedImaginary[u][v] += float64((color.GrayModel.Convert(img.At(x, y)).(color.Gray)).Y) * math.Sin(w)
				}
			}
		}
	}

	fmt.Printf("Real: %v", transformedReal)
	fmt.Printf("Imaginary: %v", transformedImaginary)

	return PixelsToImage(pixels)
}
