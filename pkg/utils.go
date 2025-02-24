package pkg

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
)

func FileNameToImage(filename string) image.Image {
	reader, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func PixelsToImage(pixels [][]color.Gray) (image.Image, error) {
	// Validate input
	if len(pixels) == 0 {
		return nil, fmt.Errorf("empty pixel array")
	}

	height := len(pixels)
	width := len(pixels[0])

	// Validate consistent dimensions
	for i := range pixels {
		if len(pixels[i]) != width {
			return nil, fmt.Errorf("inconsistent pixel array dimensions at row %d", i)
		}
	}

	// Create grayscale image instead of RGBA
	rect := image.Rect(0, 0, height, width)
	newImage := image.NewGray(rect)

	// Loop through y first for better performance
	for y := 0; y < width; y++ {
		for x := 0; x < height; x++ {
			newImage.Set(x, y, pixels[x][y])
		}
	}

	return newImage, nil
}
