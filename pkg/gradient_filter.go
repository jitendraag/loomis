// First order derivatives / gradient / Sobel operators / Roberts cross gradient operators
package pkg

import (
	"image"
	"image/color"
)

type GradientMask func() [][]int

func GradientFilter(img image.Image, maskFn GradientMask) image.Image {
	// TODO: Some of this is similar to smoothing spatial filter
	// This is from Section 3.5.1 of DIP book
	bounds := img.Bounds()
	var pixels [][]color.Gray

	var mask [][]int = maskFn()

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		var xPixels []color.Gray
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {

			var level int = 0
			for rowIndex, row := range mask {
				for colIndex, cell := range row {
					if x+rowIndex-1 < 0 || y+colIndex-1 < 0 {
						continue
					}
					level += cell * int(color.GrayModel.Convert(img.At(x+rowIndex-1, y+colIndex-1)).(color.Gray).Y)
				}
			}

			xPixels = append(xPixels, color.Gray{uint8(level)})
		}
		pixels = append(pixels, xPixels)
	}

	return PixelsToImage(pixels)
}

func SobelOperator1() [][]int {
	return [][]int{
		{-1, -2, -1},
		{0, 0, 0},
		{1, 2, 1},
	}
}

func SobelOperator2() [][]int {
	return [][]int{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}
}

func RobertsCrossOperator1() [][]int {
	return [][]int{
		{-1, 0},
		{0, 1},
	}
}

func RobertsCrossOperator2() [][]int {
	return [][]int{
		{0, -1},
		{1, 0},
	}
}
