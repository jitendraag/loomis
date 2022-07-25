package pkg

import (
	"image"
	"image/color"
)

type laplacianFilterMask func() [][]int

func Laplacian(img image.Image, maskFn laplacianFilterMask) image.Image {
	// This is from Section 3.6.2 of DIP book
	// TODO, not using smoothing filters because Laplacian mask has negative values, combine both methods
	bounds := img.Bounds()
	var pixels [][]color.Gray

	var mask [][]int = maskFn()
	var maskSum int = 0

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
					level += cell * int(color.GrayModel.Convert(img.At(x+rowIndex-1, y+colIndex-1)).(color.Gray).Y)
				}
			}
			if maskSum > 0 {
				level = level / int(maskSum)
			}
			if level < 0 {
				// This is why Laplacian needs to be scaled
				level = 0
			}

			xPixels = append(xPixels, color.Gray{uint8(level)})
		}
		pixels = append(pixels, xPixels)
	}

	return PixelsToImage(pixels)
}

func ScaledLaplacian(img image.Image, maskFn laplacianFilterMask) image.Image {
	// This is from Section 3.6.2 of DIP book
	// TODO, not using smoothing filters because Laplacian mask has negative values, combine both methods
	bounds := img.Bounds()
	var pixels [][]color.Gray
	var levels [][]int

	var mask [][]int = maskFn()
	var maskSum int = 0

	for _, row := range mask {
		for _, cell := range row {
			maskSum += cell
		}
	}

	var minimumLevel int = 255

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		var xLevels []int

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
			if maskSum > 0 {
				level = level / int(maskSum)
			}
			if level < minimumLevel {
				minimumLevel = level
			}
			xLevels = append(xLevels, level)
		}
		levels = append(levels, xLevels)
	}

	for _, row := range levels {
		var xPixels []color.Gray
		for _, value := range row {
			var newLevel int = value + minimumLevel
			if newLevel > MaxGrayscaleLevels {
				newLevel = MaxGrayscaleLevels
			}
			xPixels = append(xPixels, color.Gray{uint8(newLevel)})
		}
		pixels = append(pixels, xPixels)
	}

	return PixelsToImage(pixels)
}

func ScaledLaplacianMaskAddition(img image.Image, maskFn laplacianFilterMask) image.Image {
	// This is from Section 3.6.2 of DIP book
	maskImg := ScaledLaplacian(img, maskFn)

	bounds := img.Bounds()
	var pixels [][]color.Gray

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		var xPixels []color.Gray
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			// TODO: Can this overflow?
			level := color.GrayModel.Convert(img.At(x, y)).(color.Gray).Y + color.GrayModel.Convert(maskImg.At(x, y)).(color.Gray).Y
			xPixels = append(xPixels, color.Gray{uint8(level)})
		}
		pixels = append(pixels, xPixels)
	}

	return PixelsToImage(pixels)
}

// Sample functions for laplacian transformations
func LaplacianMask1() [][]int {
	return [][]int{
		{0, 1, 0},
		{1, -4, 1},
		{0, 1, 0},
	}
}

func LaplacianMask2() [][]int {
	return [][]int{
		{1, 1, 1},
		{1, -8, 1},
		{1, 1, 1},
	}
}

func LaplacianMask3() [][]int {
	return [][]int{
		{0, -1, 0},
		{-1, -4, -1},
		{0, -1, 0},
	}
}

func LaplacianMask4() [][]int {
	return [][]int{
		{-1, -1, -1},
		{-1, -8, -1},
		{-1, -1, -1},
	}
}
