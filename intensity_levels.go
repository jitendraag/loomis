package main

import (
	"image"
	"image/color"
)

func reduceIntensityLevels(img image.Image, levelCount int) image.Image {
	// This is from Figure 2.21 of DIP book
	// NOTE: The operation is meant to run on grayscale image only
	var normaliser uint8 = uint8(maxGrayscaleLevels/levelCount) + 1
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

	rect := image.Rect(0, 0, len(pixels), len(pixels[0]))
	newImage := image.NewRGBA(rect)

	// TODO, should this also loop through y first?
	// TODO, move writing a 2D pixel array to an image to another function
	for x := 0; x < len(pixels); x++ {
		for y := 0; y < len(pixels[0]); y++ {
			q := pixels[x]
			if q == nil {
				continue
			}
			p := pixels[x][y]
			// if p == nil {
			// 	continue
			// }
			original, ok := color.GrayModel.Convert(p).(color.Gray)
			if ok {
				newImage.Set(x, y, original)
			}
		}
	}
	return newImage
}
