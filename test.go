package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"image"
	"log"
	"strings"

	"os"

	"image/color"
	"image/draw"
	"image/jpeg"
	_ "image/jpeg"

	"github.com/jitendraag/loomis/pkg"
)

func main() {
	// TODO: ideally all new commands should self document
	// TODO: we are ignoring input/output completely right now
	var command = flag.String("command", "histgray", "Command to execute, possible options: histgray, intensity_levels, log_transformation")
	// var inputFileName = flag.String("i", "testdata/green-bee-eater-grayscale.jpg", "Input file name")
	// var outputFileName = flag.String("o", "testdata/output.jpg", "Output file name")
	var levels = flag.Float64("levels", 4, "Intensity levels (1 to 256) / constant")
	var gamma = flag.Float64("gamma", 2.5, "Gamma for power law")
	var numberOfBits = flag.Uint("bits", 2, "Number of bits to set to zero")
	var bitNumber = flag.Uint("bit", 2, "Exact bit to set to zero")

	var help = flag.Bool("help", false, "Show help")

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	// fmt.Printf("%v, %v, %v", *command, *inputFileName, *outputFileName)

	switch *command {
	case "histgray":
		testFile()
	case "intensity_levels":
		testIntensityLevels(int(*levels))
	case "log_transformation":
		// Levels being used as the constant
		testLogTransformation(int(*levels))
	case "power_law":
		testPowerLawTransformation(*levels, *gamma)
	case "bitplane_slicing":
		testBitPlaneSlicing(uint8(*numberOfBits))
	case "bitnumber_slicing":
		testBitPlaneSlicingBitNumber(uint8(*bitNumber))
	case "histequalisation":
		testHistogramEqualisation()
	case "histnormal":
		testNormalisedHistogram()
	case "gray":
		testConvertToGrayscale()
	case "smooth_spatial":
		testSmoothingSpatialFilter()
	case "nonlinear_smooth_spatial":
		testNonlinearSmoothingSpatialFilter()
	case "gaussian_spatial":
		testGaussianSpatialFilter()
	case "laplacian":
		testLaplacian()
	case "scaled_laplacian":
		testScaledLaplacian()
	case "scaled_laplacian_mask":
		testScaledLaplacianMaskAddition()
	case "unsharp_masking":
		testUnsharpMasking()
	case "unsharp_masking_scaled":
		testUnsharpMaskingScaled()
	default:
		flag.Usage()
	}

	// testBase64String()
	// writeImage()
}

func testFile() {
	// Decode the JPEG data. If reading from file, create a reader with
	img := pkg.FileNameToImage("testdata/green-bee-eater-grayscale.jpg")

	var levels []int = pkg.HistogramGrayscale(img, 2)

	fmt.Printf("%v\n", levels)
}

func testBase64String() {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(pkg.ImageData1))
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	var levels []int = pkg.HistogramGrayscale(img, 25)

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
	img := pkg.FileNameToImage("testdata/green-bee-eater-grayscale.jpg")

	newImage := pkg.ReduceIntensityLevels(img, levelCount)

	out, _ := os.Create(fmt.Sprintf("intensity%d.jpg", levelCount))
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testLogTransformation(constant int) {
	img := pkg.FileNameToImage("testdata/green-bee-eater-grayscale.jpg")

	newImage := pkg.LogTransformation(img, constant)

	out, _ := os.Create(fmt.Sprintf("intensity_log_transformation%f.jpg", constant))
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testPowerLawTransformation(constant float64, gamma float64) {
	img := pkg.FileNameToImage("testdata/green-bee-eater-grayscale.jpg")

	newImage := pkg.PowerLawTransformation(img, constant, gamma)

	out, _ := os.Create(fmt.Sprintf("intensity_gamma_transformation%d.jpg", constant))
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testBitPlaneSlicing(numberOfBits uint8) {
	img := pkg.FileNameToImage("testdata/green-bee-eater-grayscale.jpg")

	newImage := pkg.BitPlaneSlicing(img, numberOfBits)

	out, _ := os.Create(fmt.Sprintf("bit_plane_splicing%d.jpg", numberOfBits))
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testBitPlaneSlicingBitNumber(bitNumber uint8) {
	img := pkg.FileNameToImage("testdata/green-bee-eater-grayscale.jpg")

	newImage := pkg.BitPlaneSlicingBitNumber(img, bitNumber)

	out, _ := os.Create(fmt.Sprintf("bit_plane_splicing%d.jpg", bitNumber))
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testHistogramEqualisation() {
	img := pkg.FileNameToImage("testdata/green-bee-eater-grayscale.jpg")

	newImage := pkg.HistogramEqualisation(img)

	out, _ := os.Create("histogram_equalisation.jpg")
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testNormalisedHistogram() {
	img := pkg.FileNameToImage("testdata/green-bee-eater-grayscale.jpg")

	var probabilities []float64 = pkg.NormalisedHistogramGrayscale(img, 0)
	var meanIntensity float64 = pkg.MeanIntensity(img, 0)

	fmt.Printf("Probabilities: %v\n", probabilities)
	fmt.Printf("Mean intensity: %v\n", meanIntensity)
}

func testConvertToGrayscale() {
	img := pkg.FileNameToImage("testdata/green-bee-eater-color.jpg")

	newImage := pkg.ConvertToGrayscale(img)

	out, _ := os.Create("grayscale.jpg")
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testSmoothingSpatialFilter() {
	img := pkg.FileNameToImage("testdata/green-bee-eater-grayscale.jpg")

	newImage := pkg.SmoothingSpatialFilter(img, pkg.FiveByFiveUniform)

	out, _ := os.Create("smoothing_spatial.jpg")
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testNonlinearSmoothingSpatialFilter() {
	img := pkg.FileNameToImage("testdata/green-bee-eater-grayscale.jpg")

	newImage := pkg.NonlinearSmoothingSpatialFilter(img, 5, pkg.MedianOrder)

	out, _ := os.Create("nonlinear_smoothing_spatial.jpg")
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testGaussianSpatialFilter() {
	img := pkg.FileNameToImage("testdata/green-bee-eater-grayscale.jpg")

	newImage := pkg.GaussianSpatialFilter(img, pkg.GaussianFiveByFiveSigmaOne)

	out, _ := os.Create("gaussian_spatial.jpg")
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testLaplacian() {
	img := pkg.FileNameToImage("testdata/green-bee-eater-grayscale.jpg")

	newImage := pkg.Laplacian(img, pkg.LaplacianMask4)

	out, _ := os.Create("laplacian_1.jpg")
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testScaledLaplacian() {
	img := pkg.FileNameToImage("testdata/green-bee-eater-grayscale.jpg")

	newImage := pkg.ScaledLaplacian(img, pkg.LaplacianMask1)

	out, _ := os.Create("scaled_laplacian_1.jpg")
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testScaledLaplacianMaskAddition() {
	img := pkg.FileNameToImage("testdata/green-bee-eater-grayscale.jpg")

	newImage := pkg.ScaledLaplacianMaskAddition(img, pkg.LaplacianMask1)

	out, _ := os.Create("scaled_laplacian_mask_1.jpg")
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testUnsharpMasking() {
	img := pkg.FileNameToImage("testdata/green-bee-eater-grayscale.jpg")

	newImage := pkg.UnsharpMasking(img, 1.0)

	out, _ := os.Create("unsharp_masking.jpg")
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testUnsharpMaskingScaled() {
	img := pkg.FileNameToImage("testdata/green-bee-eater-grayscale.jpg")

	newImage := pkg.UnsharpMaskingScaled(img, 1.0)

	out, _ := os.Create("unsharp_masking_scaled.jpg")
	jpeg.Encode(out, newImage, nil)
	out.Close()
}
