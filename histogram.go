package main

import (
	"fmt"
	"image"
	"image/color"
)

const MaxGrayscaleLevels int = 256

func histogramGrayscale(img image.Image, levelCount int) []int {
	var levels []int = make([]int, levelCount)
	var normaliser uint8 = uint8(MaxGrayscaleLevels/levelCount) + 1
	bounds := img.Bounds()

	// Not starting from (0,0) as per documentation
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			level := c.Y / normaliser

			levels[level]++
		}
	}
	return levels
}

func HistogramEqualisation(img image.Image) image.Image {
	// This is from Figure 3.3.1 of DIP book
	// NOTE: The operation is meant to run on grayscale image only
	var levels []int = make([]int, MaxGrayscaleLevels)
	var probabilities []float64 = make([]float64, MaxGrayscaleLevels)

	bounds := img.Bounds()
	numberOfPixels := bounds.Max.X * bounds.Max.Y
	fmt.Printf("Number of pixels: %v\n", numberOfPixels)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			// Normalisation is not needed because intensity levels will remain same
			// An 8-bit image will tranform into another 8-bit image.
			level := c.Y

			levels[level]++
		}
	}
	fmt.Printf("Original histogram counts: %v\n", levels)

	for index, count := range levels {
		probabilities[index] = float64(count) / float64(numberOfPixels)
	}
	fmt.Printf("Probabilities: %v\n", probabilities)

	numberOfIntensities := len(levels)
	var equalisedLevels []int = make([]int, numberOfIntensities)

	for index, _ := range levels {
		if index == 0 {
			equalisedLevels[index] = int(float64(numberOfIntensities-1) * probabilities[index])
		} else {
			equalisedLevels[index] = equalisedLevels[index-1] + int(float64(numberOfIntensities-1)*probabilities[index])
		}
	}
	fmt.Printf("Equalised levels: %v\n", equalisedLevels)

	var pixels [][]color.Gray

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		var xPixels []color.Gray
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			level := c.Y
			level = uint8(equalisedLevels[level])
			xPixels = append(xPixels, color.Gray{uint8(level)})
		}
		pixels = append(pixels, xPixels)
	}

	return PixelsToImage(pixels)
}
