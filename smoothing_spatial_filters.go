package main

import (
	"fmt"
	"image"
	"image/color"
)

var ThreeByThreeUniform = [][]uint8{
	{1, 1, 1},
	{1, 1, 1},
	{1, 1, 1},
}

var ThreeByThreeWeighted = [][]uint8{
	{1, 2, 1},
	{2, 4, 2},
	{1, 2, 1},
}

func SmoothingSpatialFilter(img image.Image) image.Image {
	// This is from Section 3.5.1 of DIP book
	bounds := img.Bounds()
	var pixels [][]color.Gray

	var mask [][]uint8 = ThreeByThreeUniform
	var maskSum uint8 = 0

	for _, row := range mask {
		for _, cell := range row {
			maskSum += cell
		}
	}
	fmt.Printf("Mask: %v\n", mask)
	fmt.Printf("Mask sum: %v\n", maskSum)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		var xPixels []color.Gray
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			// c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)

			// TODO - better done with loops? (allows for larger masks)
			var level int = 0
			if x == 0 {
				if y == 0 {
					level = int(mask[1][1]*(color.GrayModel.Convert(img.At(x, y)).(color.Gray)).Y) +
						int(mask[1][2]*(color.GrayModel.Convert(img.At(x+1, y)).(color.Gray)).Y) +
						int(mask[2][1]*(color.GrayModel.Convert(img.At(x, y+1)).(color.Gray)).Y) +
						int(mask[2][2]*(color.GrayModel.Convert(img.At(x+1, y+1)).(color.Gray)).Y)
					level = level / 4
				} else {
					level = int(mask[0][1]*(color.GrayModel.Convert(img.At(x, y-1)).(color.Gray)).Y) +
						int(mask[0][2]*(color.GrayModel.Convert(img.At(x+1, y-1)).(color.Gray)).Y) +
						int(mask[1][1]*(color.GrayModel.Convert(img.At(x, y)).(color.Gray)).Y) +
						int(mask[1][2]*(color.GrayModel.Convert(img.At(x+1, y)).(color.Gray)).Y) +
						int(mask[2][1]*(color.GrayModel.Convert(img.At(x, y+1)).(color.Gray)).Y) +
						int(mask[2][2]*(color.GrayModel.Convert(img.At(x+1, y+1)).(color.Gray)).Y)
					level = level / 6
				}
			} else if y == 0 {
				level = int(mask[1][0]*(color.GrayModel.Convert(img.At(x-1, y)).(color.Gray)).Y) +
					int(mask[1][1]*(color.GrayModel.Convert(img.At(x, y)).(color.Gray)).Y) +
					int(mask[1][2]*(color.GrayModel.Convert(img.At(x+1, y)).(color.Gray)).Y) +
					int(mask[2][0]*(color.GrayModel.Convert(img.At(x-1, y+1)).(color.Gray)).Y) +
					int(mask[2][1]*(color.GrayModel.Convert(img.At(x, y+1)).(color.Gray)).Y) +
					int(mask[2][2]*(color.GrayModel.Convert(img.At(x+1, y+1)).(color.Gray)).Y)
				level = level / 6

			} else {
				// int casting to avoid overflowing
				level = int(mask[0][0]*(color.GrayModel.Convert(img.At(x-1, y-1)).(color.Gray)).Y) +
					int(mask[0][1]*(color.GrayModel.Convert(img.At(x, y-1)).(color.Gray)).Y) +
					int(mask[0][2]*(color.GrayModel.Convert(img.At(x+1, y-1)).(color.Gray)).Y) +
					int(mask[1][0]*(color.GrayModel.Convert(img.At(x-1, y)).(color.Gray)).Y) +
					int(mask[1][1]*(color.GrayModel.Convert(img.At(x, y)).(color.Gray)).Y) +
					int(mask[1][2]*(color.GrayModel.Convert(img.At(x+1, y)).(color.Gray)).Y) +
					int(mask[2][0]*(color.GrayModel.Convert(img.At(x-1, y+1)).(color.Gray)).Y) +
					int(mask[2][1]*(color.GrayModel.Convert(img.At(x, y+1)).(color.Gray)).Y) +
					int(mask[2][2]*(color.GrayModel.Convert(img.At(x+1, y+1)).(color.Gray)).Y)

				// fmt.Printf("x-1, y-1: %v:%v\n", mask[0][0], (color.GrayModel.Convert(img.At(x-1, y-1)).(color.Gray)).Y)
				// fmt.Printf("x, y-1: %v:%v\n", mask[0][1], (color.GrayModel.Convert(img.At(x, y-1)).(color.Gray)).Y)
				// fmt.Printf("x+1, y-1: %v:%v\n", mask[0][2], (color.GrayModel.Convert(img.At(x+1, y-1)).(color.Gray)).Y)
				// fmt.Printf("x-1, y: %v:%v\n", mask[1][0], (color.GrayModel.Convert(img.At(x-1, y)).(color.Gray)).Y)
				// fmt.Printf("x, y: %v:%v\n", mask[1][1], (color.GrayModel.Convert(img.At(x, y)).(color.Gray)).Y)
				// fmt.Printf("x+1, y: %v:%v\n", mask[1][2], (color.GrayModel.Convert(img.At(x+1, y)).(color.Gray)).Y)
				// fmt.Printf("x-1, y+1: %v:%v\n", mask[2][0], (color.GrayModel.Convert(img.At(x-1, y+1)).(color.Gray)).Y)
				// fmt.Printf("x, y+1: %v:%v\n", mask[2][1], (color.GrayModel.Convert(img.At(x, y+1)).(color.Gray)).Y)
				// fmt.Printf("x+1, y+1: %v:%v\n", mask[2][2], (color.GrayModel.Convert(img.At(x+1, y+1)).(color.Gray)).Y)
				// fmt.Printf("Sum levels: %v\n", level)
				level = level / int(maskSum)
				// fmt.Printf("Old value :%v, New value: %v\n", mask[1][1]*(color.GrayModel.Convert(img.At(x, y)).(color.Gray)).Y, level)
			}

			xPixels = append(xPixels, color.Gray{uint8(level)})
		}
		pixels = append(pixels, xPixels)
	}

	return PixelsToImage(pixels)
}
