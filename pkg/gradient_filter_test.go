package pkg

import (
	"image"
	"image/color"
	"testing"
)

func TestGradientFilter(t *testing.T) {
	// Create a simple 3x3 test image
	img := image.NewGray(image.Rect(0, 0, 3, 3))
	// Set up a simple gradient pattern
	pixels := [][]uint8{
		{0, 128, 255},
		{0, 128, 255},
		{0, 128, 255},
	}
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			img.Set(x, y, color.Gray{Y: pixels[y][x]})
		}
	}

	tests := []struct {
		name    string
		mask    GradientMask
		wantErr bool
	}{
		{
			name:    "Sobel Operator 1 - Horizontal Edge Detection",
			mask:    SobelOperator1,
			wantErr: false,
		},
		{
			name:    "Sobel Operator 2 - Vertical Edge Detection",
			mask:    SobelOperator2,
			wantErr: false,
		},
		{
			name:    "Roberts Cross Operator 1",
			mask:    RobertsCrossOperator1,
			wantErr: false,
		},
		{
			name:    "Roberts Cross Operator 2",
			mask:    RobertsCrossOperator2,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GradientFilter(img, tt.mask)
			if (err != nil) != tt.wantErr {
				t.Errorf("GradientFilter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Error("GradientFilter() returned nil image")
				return
			}

			// Verify the output image dimensions
			if got.Bounds() != img.Bounds() {
				t.Errorf("GradientFilter() output dimensions = %v, want %v", got.Bounds(), img.Bounds())
			}
		})
	}
}

func TestGradientMasks(t *testing.T) {
	tests := []struct {
		name     string
		maskFn   GradientMask
		expected [][]int
	}{
		{
			name:   "Sobel Operator 1",
			maskFn: SobelOperator1,
			expected: [][]int{
				{-1, -2, -1},
				{0, 0, 0},
				{1, 2, 1},
			},
		},
		{
			name:   "Sobel Operator 2",
			maskFn: SobelOperator2,
			expected: [][]int{
				{-1, 0, 1},
				{-2, 0, 2},
				{-1, 0, 1},
			},
		},
		{
			name:   "Roberts Cross Operator 1",
			maskFn: RobertsCrossOperator1,
			expected: [][]int{
				{-1, 0},
				{0, 1},
			},
		},
		{
			name:   "Roberts Cross Operator 2",
			maskFn: RobertsCrossOperator2,
			expected: [][]int{
				{0, -1},
				{1, 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.maskFn()
			if !compareMatrices(got, tt.expected) {
				t.Errorf("%s = %v, want %v", tt.name, got, tt.expected)
			}
		})
	}
}

// Helper function to compare two 2D integer slices
func compareMatrices(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
