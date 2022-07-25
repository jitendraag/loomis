package pkg

import "image"

func GaussianThreeByTreeSigmaOne() [][]uint8 {
	// 3x3 Gaussian kernel
	// Standard deviation 1
	return [][]uint8{
		{1, 2, 1},
		{2, 4, 2},
		{1, 2, 1},
	}
}

func GaussianFiveByFiveSigmaOne() [][]uint8 {
	// 5x5 Gaussian kernel
	// Standard deviation 1
	return [][]uint8{
		{1, 4, 7, 4, 1},
		{4, 16, 26, 16, 4},
		{7, 26, 41, 26, 7},
		{4, 16, 26, 16, 4},
		{1, 4, 7, 4, 1},
	}
}

func GaussianSpatialFilter(img image.Image, maskFn FilterMask) image.Image {
	// TODO: how to generate these filters for sigma 2, 3 and for 7x7 kernels?
	// NOTE: Still keeping a separate entry point for Gaussian because in future kernel size and standard deviation will become parameters
	return SmoothingSpatialFilter(img, maskFn)
}
