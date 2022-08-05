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
	var command = flag.String("command", "histgray", "Command to execute, possible options: histgray, intensity_levels, log_transformation")
	var inputFileName = flag.String("i", "testdata/green-bee-eater-grayscale.jpg", "Input file name")
	var outputFileName = flag.String("o", "testdata/output.jpg", "Output file name")
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
		testFile(int(*levels), *inputFileName)
	case "intensity_levels":
		testIntensityLevels(int(*levels), *inputFileName, *outputFileName)
	case "log_transformation":
		// Levels being used as the constant
		testLogTransformation(int(*levels), *inputFileName, *outputFileName)
	case "power_law":
		testPowerLawTransformation(*levels, *gamma, *inputFileName, *outputFileName)
	case "bitplane_slicing":
		testBitPlaneSlicing(uint8(*numberOfBits), *inputFileName, *outputFileName)
	case "bitnumber_slicing":
		testBitPlaneSlicingBitNumber(uint8(*bitNumber), *inputFileName, *outputFileName)
	case "histequalisation":
		testHistogramEqualisation(*inputFileName, *outputFileName)
	case "histnormal":
		testNormalisedHistogram(*inputFileName)
	case "gray":
		testConvertToGrayscale(*inputFileName, *outputFileName)
	case "smooth_spatial":
		testSmoothingSpatialFilter(*inputFileName, *outputFileName)
	case "nonlinear_smooth_spatial":
		testNonlinearSmoothingSpatialFilter(*inputFileName, *outputFileName)
	case "gaussian_spatial":
		testGaussianSpatialFilter(*inputFileName, *outputFileName)
	case "laplacian":
		testLaplacian(*inputFileName, *outputFileName)
	case "scaled_laplacian":
		testScaledLaplacian(*inputFileName, *outputFileName)
	case "scaled_laplacian_mask":
		testScaledLaplacianMaskAddition(*inputFileName, *outputFileName)
	case "unsharp_masking":
		testUnsharpMasking(*inputFileName, *outputFileName)
	case "unsharp_masking_scaled":
		testUnsharpMaskingScaled(*inputFileName, *outputFileName)
	case "gradient":
		testGradientFilter(int(*levels), *inputFileName, *outputFileName)
	case "dft":
		testDiscreetFourierTransform(*inputFileName, *outputFileName)
	case "gaussian_pdf":
		testGaussianPdf()
	case "rayleigh_pdf":
		testRayleighPdf()
	default:
		flag.Usage()
	}

	// testBase64String()
	// writeImage()
}

func testFile(levelCount int, inputFileName string) {
	// Decode the JPEG data. If reading from file, create a reader with
	img := pkg.FileNameToImage(inputFileName)

	var levels []int = pkg.HistogramGrayscale(img, levelCount)

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

func testIntensityLevels(levelCount int, inputFileName string, outputFileName string) {
	// Decode the JPEG data. If reading from file, create a reader with
	img := pkg.FileNameToImage(inputFileName)

	newImage := pkg.ReduceIntensityLevels(img, levelCount)

	out, _ := os.Create(outputFileName)
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testLogTransformation(constant int, inputFileName string, outputFileName string) {
	img := pkg.FileNameToImage(inputFileName)

	newImage := pkg.LogTransformation(img, constant)

	out, _ := os.Create(outputFileName)
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testPowerLawTransformation(constant float64, gamma float64, inputFileName string, outputFileName string) {
	img := pkg.FileNameToImage(inputFileName)

	newImage := pkg.PowerLawTransformation(img, constant, gamma)

	out, _ := os.Create(outputFileName)
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testBitPlaneSlicing(numberOfBits uint8, inputFileName string, outputFileName string) {
	img := pkg.FileNameToImage(inputFileName)

	newImage := pkg.BitPlaneSlicing(img, numberOfBits)

	out, _ := os.Create(outputFileName)
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testBitPlaneSlicingBitNumber(bitNumber uint8, inputFileName string, outputFileName string) {
	img := pkg.FileNameToImage(inputFileName)

	newImage := pkg.BitPlaneSlicingBitNumber(img, bitNumber)

	out, _ := os.Create(outputFileName)
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testHistogramEqualisation(inputFileName string, outputFileName string) {
	img := pkg.FileNameToImage(inputFileName)

	newImage := pkg.HistogramEqualisation(img)

	out, _ := os.Create(outputFileName)
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testNormalisedHistogram(inputFileName string) {
	img := pkg.FileNameToImage(inputFileName)

	var probabilities []float64 = pkg.NormalisedHistogramGrayscale(img, 0)
	var meanIntensity float64 = pkg.MeanIntensity(img, 0)

	fmt.Printf("Probabilities: %v\n", probabilities)
	fmt.Printf("Mean intensity: %v\n", meanIntensity)
}

func testConvertToGrayscale(inputFileName string, outputFileName string) {
	img := pkg.FileNameToImage(inputFileName)

	newImage := pkg.ConvertToGrayscale(img)

	out, _ := os.Create(outputFileName)
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testSmoothingSpatialFilter(inputFileName string, outputFileName string) {
	img := pkg.FileNameToImage(inputFileName)

	newImage := pkg.SmoothingSpatialFilter(img, pkg.ThreeByThreeWeighted)

	out, _ := os.Create(outputFileName)
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testNonlinearSmoothingSpatialFilter(inputFileName string, outputFileName string) {
	img := pkg.FileNameToImage(inputFileName)

	newImage := pkg.NonlinearSmoothingSpatialFilter(img, 5, pkg.MaxOrder)

	out, _ := os.Create(outputFileName)
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testGaussianSpatialFilter(inputFileName string, outputFileName string) {
	img := pkg.FileNameToImage(inputFileName)

	newImage := pkg.GaussianSpatialFilter(img, pkg.GaussianFiveByFiveSigmaOne)

	out, _ := os.Create(outputFileName)
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testLaplacian(inputFileName string, outputFileName string) {
	img := pkg.FileNameToImage(inputFileName)

	newImage := pkg.Laplacian(img, pkg.LaplacianMask4)

	out, _ := os.Create(outputFileName)
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testScaledLaplacian(inputFileName string, outputFileName string) {
	img := pkg.FileNameToImage(inputFileName)

	newImage := pkg.ScaledLaplacian(img, pkg.LaplacianMask1)

	out, _ := os.Create(outputFileName)
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testScaledLaplacianMaskAddition(inputFileName string, outputFileName string) {
	img := pkg.FileNameToImage(inputFileName)

	newImage := pkg.ScaledLaplacianMaskAddition(img, pkg.LaplacianMask1)

	out, _ := os.Create(outputFileName)
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testUnsharpMasking(inputFileName string, outputFileName string) {
	img := pkg.FileNameToImage(inputFileName)

	newImage := pkg.UnsharpMasking(img, 1.0)

	out, _ := os.Create(outputFileName)
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testUnsharpMaskingScaled(inputFileName string, outputFileName string) {
	img := pkg.FileNameToImage(inputFileName)

	newImage := pkg.UnsharpMaskingScaled(img, 1.0)

	out, _ := os.Create(outputFileName)
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testGradientFilter(levels int, inputFileName string, outputFileName string) {
	img := pkg.FileNameToImage(inputFileName)
	var newImage image.Image
	// TODO, find a better way to specify these operators
	switch levels {
	case 1:
		newImage = pkg.GradientFilter(img, pkg.SobelOperator1)
	case 2:

		newImage = pkg.GradientFilter(img, pkg.SobelOperator2)
	case 3:
		newImage = pkg.GradientFilter(img, pkg.RobertsCrossOperator1)
	case 4:
		newImage = pkg.GradientFilter(img, pkg.RobertsCrossOperator2)
	default:
		fmt.Printf("Only values 1-4 are supported for gradient filter.")
		newImage = img
	}

	out, _ := os.Create(outputFileName)
	jpeg.Encode(out, newImage, nil)
	out.Close()
}

func testDiscreetFourierTransform(inputFileName string, outputFileName string) {
	img := pkg.FileNameToImage(inputFileName)

	newImage := pkg.DiscreetFourierTransform(img)

	out, _ := os.Create(outputFileName)
	jpeg.Encode(out, newImage, nil)
	out.Close()
}


func testGaussianPdf() {
	pdf := pkg.GaussianPdf(128, 20)
	fmt.Printf("PDF: %v", pdf)
}


func testRayleighPdf() {
	pdf := pkg.RayleighPdf(0, 0.4)
	fmt.Printf("PDF: %v", pdf)
}