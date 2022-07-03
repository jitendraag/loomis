package main

import (
	"encoding/base64"
	"fmt"
	"image"
	"log"
	"strings"

	"os"

	"image/color"
	"image/draw"
	"image/jpeg"
	_ "image/jpeg"
)

func main() {
	// testFile()
	// testBase64String()

	// writeImage()
	// testIntensityLevels(128)
	// testIntensityLevels(64)
	// testIntensityLevels(4)

	// testLogTransformation(100)

	// testPowerLawTransformation(1.0, 2.5)
	// testBitPlaneSlicing(1)
	// testBitPlaneSlicingBitNumber(0)

	testHistogramEqualisation()
}

func testFile() {
	// Decode the JPEG data. If reading from file, create a reader with
	img := FileNameToImage("testdata/green-bee-eater-grayscale.jpg")

	var levels []int = histogramGrayscale(img, 2)

	fmt.Printf("%v\n", levels)
}

func testBase64String() {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(ImageData1))
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	var levels []int = histogramGrayscale(img, 25)

	fmt.Printf("%v\n", levels)
}

func writeImage() {
	// TODO, needs error handling
	img := image.NewRGBA(image.Rect(0, 0, 1, 1))
	white := color.Gray{255}
	// black := color.Gray{0}

	draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

	out, _ := os.Create("1pixel.jpg")
	jpeg.Encode(out, img, nil)
	out.Close()
}

func testIntensityLevels(levelCount int) {
	// Decode the JPEG data. If reading from file, create a reader with
	img := FileNameToImage("testdata/green-bee-eater-grayscale.jpg")

	newImage := reduceIntensityLevels(img, levelCount)

	out, _ := os.Create(fmt.Sprintf("intensity%d.jpg", levelCount))
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testLogTransformation(constant int) {
	img := FileNameToImage("testdata/green-bee-eater-grayscale.jpg")

	newImage := LogTransformation(img, constant)

	out, _ := os.Create(fmt.Sprintf("intensity_log_transformation%f.jpg", constant))
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testPowerLawTransformation(constant float64, gamma float64) {
	img := FileNameToImage("testdata/green-bee-eater-grayscale.jpg")

	newImage := PowerLawTransformation(img, constant, gamma)

	out, _ := os.Create(fmt.Sprintf("intensity_gamma_transformation%d.jpg", constant))
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testBitPlaneSlicing(numberOfBits uint8) {
	img := FileNameToImage("testdata/green-bee-eater-grayscale.jpg")

	newImage := BitPlaneSlicing(img, numberOfBits)

	out, _ := os.Create(fmt.Sprintf("bit_plane_splicing%d.jpg", numberOfBits))
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testBitPlaneSlicingBitNumber(bitNumber uint8) {
	img := FileNameToImage("testdata/green-bee-eater-grayscale.jpg")

	newImage := BitPlaneSlicingBitNumber(img, bitNumber)

	out, _ := os.Create(fmt.Sprintf("bit_plane_splicing%d.jpg", bitNumber))
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testHistogramEqualisation() {
	img := FileNameToImage("testdata/green-bee-eater-grayscale.jpg")

	newImage := HistogramEqualisation(img)

	out, _ := os.Create("histogram_equalisation.jpg")
	jpeg.Encode(out, newImage, nil)
	out.Close()
}
