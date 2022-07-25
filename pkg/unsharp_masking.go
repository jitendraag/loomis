// Unsharp Masking and Highboost Filtering

package pkg

import (
	"image"
	"image/color"
)

func UnsharpMasking(img image.Image, maskMultiplier float64) image.Image {
	// This is from Section 3.6.3 of DIP book
	// 1. Smooth / Blur
	// 2. Subtraction
	// 3. Mask application
	var smoothImage image.Image = GaussianSpatialFilter(img, GaussianFiveByFiveSigmaOne)
	var mask [][]color.Gray = SubtractImage(img, smoothImage)
	var maskedImage image.Image = AddMask(img, mask, maskMultiplier)
	return maskedImage
}

func UnsharpMaskingScaled(img image.Image, maskMultiplier float64) image.Image {
	// This is from Section 3.6.3 of DIP book
	var smoothImage image.Image = GaussianSpatialFilter(img, GaussianFiveByFiveSigmaOne)
	var mask [][]color.Gray = SubtractImageScaled(img, smoothImage)
	var maskedImage image.Image = AddMask(img, mask, maskMultiplier)
	return maskedImage
}

func AddMask(img1 image.Image, mask [][]color.Gray, maskMultiplier float64) image.Image {
	// TODO: Can this and Laplacian operations be generalized?
	var pixels [][]color.Gray
	bounds1 := img1.Bounds()
	// NOTE: no validation for image and mask being of the same size

	for x := bounds1.Min.X; x < bounds1.Max.X; x++ {
		var xPixels []color.Gray
		for y := bounds1.Min.Y; y < bounds1.Max.Y; y++ {

			c1 := color.GrayModel.Convert(img1.At(x, y)).(color.Gray)
			var level int = int(float64(c1.Y) + maskMultiplier*float64(mask[x][y].Y))

			if level > MaxGrayscaleLevels {
				level = MaxGrayscaleLevels
			}

			xPixels = append(xPixels, color.Gray{uint8(level)})
		}
		pixels = append(pixels, xPixels)
	}

	return PixelsToImage(pixels)
}

func SubtractImage(img1 image.Image, img2 image.Image) [][]color.Gray {
	var pixels [][]color.Gray
	bounds1 := img1.Bounds()
	// NOTE: no validation for both images being of the same size

	for x := bounds1.Min.X; x < bounds1.Max.X; x++ {
		var xPixels []color.Gray
		for y := bounds1.Min.Y; y < bounds1.Max.Y; y++ {

			c1 := color.GrayModel.Convert(img1.At(x, y)).(color.Gray)
			c2 := color.GrayModel.Convert(img2.At(x, y)).(color.Gray)
			var level int = int(c1.Y - c2.Y)

			if level < 0 {
				level = 0
				// Dark Halo around edges (negative would have been worse)
			}

			xPixels = append(xPixels, color.Gray{uint8(level)})
		}
		pixels = append(pixels, xPixels)
	}

	return pixels
}

func SubtractImageScaled(img1 image.Image, img2 image.Image) [][]color.Gray {
	var pixels [][]color.Gray
	var levels [][]int
	bounds1 := img1.Bounds()
	// NOTE: no validation for both images being of the same size
	var minimumLevel int = 255

	for x := bounds1.Min.X; x < bounds1.Max.X; x++ {
		var xLevels []int
		for y := bounds1.Min.Y; y < bounds1.Max.Y; y++ {

			c1 := color.GrayModel.Convert(img1.At(x, y)).(color.Gray)
			c2 := color.GrayModel.Convert(img2.At(x, y)).(color.Gray)
			var level int = int(c1.Y - c2.Y)

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

	return pixels
}
