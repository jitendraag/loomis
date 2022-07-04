package main

import (
	"fmt"
	"image"
	"image/color"
	"math"
)

func LogTransformation(img image.Image, constant int) image.Image {
	// This is from Section 3.2.2 of DIP book
	// TODO, should the constant be a float?
	bounds := img.Bounds()
	var pixels [][]color.Gray
	var maxIntensity color.Gray

	maxIntensity = color.GrayModel.Convert(img.At(bounds.Min.X, bounds.Min.Y)).(color.Gray)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		var xPixels []color.Gray
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			level := float64(constant) * math.Log(1+float64(c.Y))

			xPixels = append(xPixels, color.Gray{uint8(level)})
		}
		pixels = append(pixels, xPixels)
	}

	normaliser := int(maxIntensity.Y) / MaxGrayscaleLevels
	if normaliser != 0 {
		for rowIndex, row := range pixels {
			for colIndex, pixel := range row {
				pixels[rowIndex][colIndex] = color.Gray{pixel.Y / uint8(normaliser)}
			}
		}
	}

	return PixelsToImage(pixels)
}

func GammaTransformation(img image.Image, constant float64, gamma float64) image.Image {
	return PowerLawTransformation(img, constant, gamma)
}

func PowerLawTransformation(img image.Image, constant float64, gamma float64) image.Image {
	// This is from Section 3.2.3 of DIP book
	if constant <= 0 || gamma <= 0 {
		// XXX: return error in this case?
		return img
	}
	bounds := img.Bounds()
	var pixels [][]color.Gray
	var maxIntensity color.Gray

	maxIntensity = color.GrayModel.Convert(img.At(bounds.Min.X, bounds.Min.Y)).(color.Gray)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		var xPixels []color.Gray
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			level := float64(constant) * math.Log(1+float64(c.Y))

			if level > float64(maxIntensity.Y) {
				maxIntensity = color.Gray{uint8(level)}
			}

			xPixels = append(xPixels, color.Gray{uint8(level)})
		}
		pixels = append(pixels, xPixels)
	}

	normaliser := int(maxIntensity.Y) / MaxGrayscaleLevels
	if normaliser != 0 {
		for rowIndex, row := range pixels {
			for colIndex, pixel := range row {
				pixels[rowIndex][colIndex] = color.Gray{pixel.Y / uint8(normaliser)}
			}
		}
	}

	return PixelsToImage(pixels)
}

func ContrastStretching(img image.Image) image.Image {
	// This is from Section 3.2.4 of DIP book
	/* TODO - Need to understand more
	The example given uses three equations
	y1 = m1x + c1 (in example c1 is zero)
	y2 = m2x + c2
	y3 = m2x + c3

	Slope of y1 > y2
	Slope of y3 > y1

	Need to read more to decide if this piece wise contrast stretching function often has more than 3 pieces.
	If m2 = 0 then we get a tresholding function with two intensities (binary image)
	*/

	// TODO - Similar understanding is needed for intensity level slicing where piece wise graph definition is an input for transformation function
	return img
}

func BitPlaneSlicing(img image.Image, numberOfBits uint8) image.Image {
	// This is from Section 3.2.4 of DIP book
	// We set given number of least signficant bits to zero
	bounds := img.Bounds()
	var pixels [][]color.Gray

	if numberOfBits >= 8 || numberOfBits < 1 {
		return img
	}

	var bitMask uint8 = 1<<8 - 1
	bitMask = bitMask << uint8(numberOfBits)

	// fmt.Printf("%v\n", bitMask)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		var xPixels []color.Gray
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			level := c.Y & bitMask
			xPixels = append(xPixels, color.Gray{uint8(level)})
		}
		pixels = append(pixels, xPixels)
	}

	return PixelsToImage(pixels)
}

func BitPlaneSlicingBitNumber(img image.Image, bitNumber uint8) image.Image {
	// This is from Section 3.2.4 of DIP book
	// We set the given bit to zero
	bounds := img.Bounds()
	var pixels [][]color.Gray

	if bitNumber >= 8 {
		return img
	}

	var bitMask uint8 = 1
	bitMask = bitMask << uint8(bitNumber)
	bitMask = ^bitMask

	fmt.Printf("%v\n", bitMask)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		var xPixels []color.Gray
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			level := c.Y & bitMask
			xPixels = append(xPixels, color.Gray{uint8(level)})
		}
		pixels = append(pixels, xPixels)
	}

	return PixelsToImage(pixels)
}

func ConvertToGrayscale(img image.Image) image.Image {
	bounds := img.Bounds()
	var pixels [][]color.Gray

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		var xPixels []color.Gray
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			xPixels = append(xPixels, color.Gray{uint8(c.Y)})
		}
		pixels = append(pixels, xPixels)
	}

	return PixelsToImage(pixels)
}
