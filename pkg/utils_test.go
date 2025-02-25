package pkg

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"testing"
)

func TestFileNameToImage(t *testing.T) {
	// Create a temporary test image
	tmpDir := t.TempDir()
	testImagePath := filepath.Join(tmpDir, "test.png")

	// Create a small test image
	img := image.NewGray(image.Rect(0, 0, 2, 2))
	img.Set(0, 0, color.Gray{Y: 100})
	img.Set(1, 0, color.Gray{Y: 150})
	img.Set(0, 1, color.Gray{Y: 200})
	img.Set(1, 1, color.Gray{Y: 250})

	f, err := os.Create(testImagePath)
	if err != nil {
		t.Fatal(err)
	}
	if err := png.Encode(f, img); err != nil {
		f.Close()
		t.Fatal(err)
	}
	f.Close()

	// Test the function
	result := FileNameToImage(testImagePath)

	// Verify dimensions
	bounds := result.Bounds()
	if bounds.Dx() != 2 || bounds.Dy() != 2 {
		t.Errorf("Expected image dimensions 2x2, got %dx%d", bounds.Dx(), bounds.Dy())
	}

	// Verify some pixel values
	if result.At(0, 0).(color.Gray).Y != 100 {
		t.Errorf("Expected pixel at (0,0) to be 100, got %d", result.At(0, 0).(color.Gray).Y)
	}
}

func TestPixelsToImage(t *testing.T) {
	tests := []struct {
		name        string
		pixels      [][]color.Gray
		wantErr     bool
		errContains string
	}{
		{
			name:        "empty array",
			pixels:      [][]color.Gray{},
			wantErr:     true,
			errContains: "empty pixel array",
		},
		{
			name: "inconsistent dimensions",
			pixels: [][]color.Gray{
				{{Y: 100}, {Y: 150}},
				{{Y: 200}},
			},
			wantErr:     true,
			errContains: "inconsistent pixel array dimensions",
		},
		{
			name: "valid 2x2 image",
			pixels: [][]color.Gray{
				{{Y: 100}, {Y: 150}},
				{{Y: 200}, {Y: 250}},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img, err := PixelsToImage(tt.pixels)

			// Check error cases
			if tt.wantErr {
				if err == nil {
					t.Error("Expected error but got none")
				} else if tt.errContains != "" && !contains(err.Error(), tt.errContains) {
					t.Errorf("Expected error containing %q, got %v", tt.errContains, err)
				}
				return
			}

			// Check success cases
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			// Verify dimensions for valid case
			if tt.name == "valid 2x2 image" {
				bounds := img.Bounds()
				if bounds.Dx() != 2 || bounds.Dy() != 2 {
					t.Errorf("Expected image dimensions 2x2, got %dx%d", bounds.Dx(), bounds.Dy())
				}

				// Verify pixel values
				grayImg := img.(*image.Gray)
				expectedPixels := tt.pixels
				for x := 0; x < 2; x++ {
					for y := 0; y < 2; y++ {
						if grayImg.GrayAt(x, y).Y != expectedPixels[x][y].Y {
							t.Errorf("Pixel mismatch at (%d,%d): expected %d, got %d",
								x, y, expectedPixels[x][y].Y, grayImg.GrayAt(x, y).Y)
						}
					}
				}
			}
		})
	}
}

// Helper function to check if a string contains another string
func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[0:len(substr)] == substr
}
