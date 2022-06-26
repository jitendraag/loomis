package main

import (
	"image"
	"image/color"
)

const maxGrayscaleLevels int = 256

func histogramGrayscale(img image.Image, levelCount int) []int {
	var levels []int = make([]int, levelCount)
	var normaliser uint8 = uint8(maxGrayscaleLevels/levelCount) + 1
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
