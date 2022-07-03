package main

import (
	"image"
	"image/color"
)

func reduceIntensityLevels(img image.Image, levelCount int) image.Image {
	// This is from Figure 2.21 of DIP book
	// NOTE: The operation is meant to run on grayscale image only
	var normaliser uint8 = uint8(MaxGrayscaleLevels/levelCount) + 1
	bounds := img.Bounds()
	var pixels [][]color.Gray

	// Not starting from (0,0) as per documentation
	// TODO: Golang document recommends looping through y first and x later for performance
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		var xPixels []color.Gray
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			//TODO, this makes the whole image more black, I need to learn these transformations better
			level := c.Y / normaliser
			xPixels = append(xPixels, color.Gray{level})
		}
		pixels = append(pixels, xPixels)
	}

	return PixelsToImage(pixels)
}
