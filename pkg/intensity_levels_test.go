package pkg

import (
	"image"
	"image/color"
	"testing"
)

func TestReduceIntensityLevels(t *testing.T) {
	tests := []struct {
		name       string
		input      image.Image
		levels     int
		wantError  bool
		checkPixel func(t *testing.T, img image.Image)
	}{
		{
			name:      "2x2 image with 2 levels",
			input:     createTestImage(2, 2, []uint8{0, 85, 170, 255}),
			levels:    2,
			wantError: false,
			checkPixel: func(t *testing.T, img image.Image) {
				// With 2 levels (normalizer = 128 + 1 = 129), values are reduced
				checkPixelValue(t, img, 0, 0, 0) // 0/129 = 0
				checkPixelValue(t, img, 1, 1, 1) // 255/129 = 1
			},
		},
		{
			name:      "1x1 image with 4 levels",
			input:     createTestImage(1, 1, []uint8{128}),
			levels:    4,
			wantError: false,
			checkPixel: func(t *testing.T, img image.Image) {
				// With 4 levels (normalizer = 64 + 1 = 65), 128/65 = 1
				checkPixelValue(t, img, 0, 0, 1)
			},
		},
		{
			name:      "3x3 image with 8 levels",
			input:     createTestImage(3, 3, []uint8{0, 32, 64, 96, 128, 160, 192, 224, 255}),
			levels:    8,
			wantError: false,
			checkPixel: func(t *testing.T, img image.Image) {
				// With 8 levels (normalizer = 32 + 1 = 33)
				checkPixelValue(t, img, 0, 0, 0) // 0/33 = 0
				checkPixelValue(t, img, 1, 1, 3) // 128/33 ≈ 3
				checkPixelValue(t, img, 2, 2, 7) // 255/33 ≈ 7
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ReduceIntensityLevels(tt.input, tt.levels)

			if (err != nil) != tt.wantError {
				t.Errorf("ReduceIntensityLevels() error = %v, wantError %v", err, tt.wantError)
				return
			}

			if err == nil {
				tt.checkPixel(t, result)
			}
		})
	}
}

// Helper function to create a test image with specific pixel values
func createTestImage(width, height int, values []uint8) image.Image {
	img := image.NewGray(image.Rect(0, 0, width, height))
	idx := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if idx < len(values) {
				img.Set(x, y, color.Gray{Y: values[idx]})
				idx++
			}
		}
	}
	return img
}

// Helper function to check if a pixel has the expected grayscale value
func checkPixelValue(t *testing.T, img image.Image, x, y int, expected uint8) {
	got := color.GrayModel.Convert(img.At(x, y)).(color.Gray).Y
	if got != expected {
		t.Errorf("Pixel at (%d,%d) = %d, want %d", x, y, got, expected)
	}
}
