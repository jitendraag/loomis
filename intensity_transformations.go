package main

import (
	"image"
	"image/color"
	"math"
)

func LogTransformation(img image.Image, constant int) image.Image {
	// This is from Section 3.2.2 of DIP book
	// TODO, should the constant be a float?
	bounds := img.Bounds()
	var pixels [][]color.Gray
	// maxGrayscaleLevels
	var maxIntensity color.Gray

	maxIntensity = color.GrayModel.Convert(img.At(bounds.Min.X, bounds.Min.Y)).(color.Gray)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		var xPixels []color.Gray
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			level := float64(constant) * math.Log(1+float64(c.Y))

			xPixels = append(xPixels, color.Gray{uint8(level)})
		}
		pixels = append(pixels, xPixels)
	}

	normaliser := int(maxIntensity.Y) / maxGrayscaleLevels
	if normaliser != 0 {
		for rowIndex, row := range pixels {
			for colIndex, pixel := range row {
				pixels[rowIndex][colIndex] = color.Gray{pixel.Y / uint8(normaliser)}
			}
		}
	}

	return PixelsToImage(pixels)
}

func GammaTransformation(img image.Image, constant float64, gamma float64) image.Image {
	return PowerLawTransformation(img, constant, gamma)
}

func PowerLawTransformation(img image.Image, constant float64, gamma float64) image.Image {
	// This is from Section 3.2.3 of DIP book
	if constant <= 0 || gamma <= 0 {
		// XXX: return error in this case?
		return img
	}
	bounds := img.Bounds()
	var pixels [][]color.Gray
	// maxGrayscaleLevels
	var maxIntensity color.Gray

	maxIntensity = color.GrayModel.Convert(img.At(bounds.Min.X, bounds.Min.Y)).(color.Gray)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		var xPixels []color.Gray
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			level := float64(constant) * math.Log(1+float64(c.Y))

			xPixels = append(xPixels, color.Gray{uint8(level)})
		}
		pixels = append(pixels, xPixels)
	}

	normaliser := int(maxIntensity.Y) / maxGrayscaleLevels
	if normaliser != 0 {
		for rowIndex, row := range pixels {
			for colIndex, pixel := range row {
				pixels[rowIndex][colIndex] = color.Gray{pixel.Y / uint8(normaliser)}
			}
		}
	}

	return PixelsToImage(pixels)
}
