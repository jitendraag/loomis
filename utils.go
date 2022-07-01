package main

import (
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

func PixelsToImage(pixels [][]color.Gray) image.Image {
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
