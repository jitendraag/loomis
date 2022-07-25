package pkg

import (
	"image"
	"image/color"
	"sort"
)

type FilterMask func() [][]uint8
type orderStatistic func([][]uint8) uint8

func SmoothingSpatialFilter(img image.Image, maskFn FilterMask) image.Image {
	// This is from Section 3.5.1 of DIP book
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
					if x+rowIndex-1 < 0 || y+colIndex-1 < 0 {
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

func NonlinearSmoothingSpatialFilter(img image.Image, windowSize uint, statisticFn orderStatistic) image.Image {
	// This is from Section 3.5.2 of DIP book
	bounds := img.Bounds()
	var pixels [][]color.Gray

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		var xPixels []color.Gray
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {

			var level uint8 = 0
			var windowLevels [][]uint8
			var windowLevel uint8 = 0
			for i := 1; i < int(windowSize); i++ {
				var xLevels []uint8

				for j := 1; j < int(windowSize); j++ {
					if x+i-1 < 0 || y+j-1 < 0 {
						continue
					}
					windowLevel = color.GrayModel.Convert(img.At(x+i-1, y+j-1)).(color.Gray).Y
					xLevels = append(xLevels, windowLevel)
				}
				windowLevels = append(windowLevels, xLevels)
			}

			level = statisticFn(windowLevels)

			xPixels = append(xPixels, color.Gray{uint8(level)})
		}
		pixels = append(pixels, xPixels)
	}

	return PixelsToImage(pixels)
}

// Sample functions for linear transformations
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

// Sample functions for non-linear / order-statistic transformations
func MinOrder(levels [][]uint8) uint8 {
	var mininimum = levels[0][0]

	for _, row := range levels {
		for _, value := range row {
			if value < mininimum {
				mininimum = value
			}
		}
	}

	return mininimum
}

func MaxOrder(levels [][]uint8) uint8 {
	var maximum = levels[0][0]

	for _, row := range levels {
		for _, value := range row {
			if value < maximum {
				maximum = value
			}
		}
	}

	return maximum
}

func MedianOrder(levels [][]uint8) uint8 {
	var values []int

	for _, row := range levels {
		for _, value := range row {
			values = append(values, int(value))
		}
	}
	sort.Ints(values)
	var medianIndex = len(values) / 2

	return uint8(values[medianIndex])
}
