package main

import (
	"image"
	"image/color"
)

type filterMask func() [][]uint8

func ThreeByThreeUniform() [][]uint8 {
	return [][]uint8{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}
}

func FiveByFiveUniform() [][]uint8 {
	return [][]uint8{
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
	}
}

func ThreeByThreeWeighted() [][]uint8 {
	return [][]uint8{
		{1, 2, 1},
		{2, 4, 2},
		{1, 2, 1},
	}
}

func SmoothingSpatialFilter(img image.Image, maskFn filterMask) image.Image {
	// This is from Section 3.5.1 of DIP book
	// TODO, figure out how people pass the masks when they use these filters
	bounds := img.Bounds()
	var pixels [][]color.Gray

	var mask [][]uint8 = maskFn()
	var maskSum uint8 = 0

	for _, row := range mask {
		for _, cell := range row {
			maskSum += cell
		}
	}

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		var xPixels []color.Gray
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {

			var level int = 0
			for rowIndex, row := range mask {
				for colIndex, cell := range row {
					if x+rowIndex-1 < 0 {
						continue
					}
					if y+colIndex-1 < 0 {
						continue
					}
					level += int(cell * (color.GrayModel.Convert(img.At(x+rowIndex-1, y+colIndex-1)).(color.Gray)).Y)
				}
			}
			level = level / int(maskSum)

			xPixels = append(xPixels, color.Gray{uint8(level)})
		}
		pixels = append(pixels, xPixels)
	}

	return PixelsToImage(pixels)
}
