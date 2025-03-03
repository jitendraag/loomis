package pkg

import (
	"math"
	"testing"
)

func TestGaussianPdf(t *testing.T) {
	tests := []struct {
		name              string
		mean              int
		standardDeviation int
		wantSum           float64 // PDF should sum to approximately 1
		wantPeakAt        int     // Peak should be at mean
	}{
		{
			name:              "standard normal distribution",
			mean:              128,
			standardDeviation: 30,
			wantSum:           1.0,
			wantPeakAt:        128,
		},
		{
			name:              "narrow distribution",
			mean:              100,
			standardDeviation: 10,
			wantSum:           1.0,
			wantPeakAt:        100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GaussianPdf(tt.mean, tt.standardDeviation)

			// Check if length is correct
			if len(got) != MaxGrayscaleLevels {
				t.Errorf("GaussianPdf() length = %v, want %v", len(got), MaxGrayscaleLevels)
			}

			// Check if peak is at mean
			maxIdx := 0
			for i := 1; i < len(got); i++ {
				if got[i] > got[maxIdx] {
					maxIdx = i
				}
			}
			if maxIdx != tt.wantPeakAt {
				t.Errorf("Peak at = %v, want %v", maxIdx, tt.wantPeakAt)
			}

			// Check if values are normalized between 0 and 1
			for _, v := range got {
				if v < 0 || v > 1 {
					t.Errorf("Found value outside [0,1] range: %v", v)
				}
			}
		})
	}
}

func TestRayleighPdf(t *testing.T) {
	tests := []struct {
		name            string
		a               float64
		b               float64
		wantZeroBeforeA bool
	}{
		{
			name:            "standard rayleigh",
			a:               2.0,
			b:               1.0,
			wantZeroBeforeA: true,
		},
		{
			name:            "shifted rayleigh",
			a:               5.0,
			b:               2.0,
			wantZeroBeforeA: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RayleighPdf(tt.a, tt.b)

			// Check if length is correct
			if len(got) != MaxGrayscaleLevels {
				t.Errorf("RayleighPdf() length = %v, want %v", len(got), MaxGrayscaleLevels)
			}

			// Check if values before a are zero
			if tt.wantZeroBeforeA {
				aIndex := int(tt.a / 0.01)
				for i := 0; i < aIndex && i < len(got); i++ {
					if got[i] != 0 {
						t.Errorf("Expected zero before a, got %v at index %v", got[i], i)
					}
				}
			}

			// Check if values are normalized between 0 and 1
			for _, v := range got {
				if v < 0 || v > 1 {
					t.Errorf("Found value outside [0,1] range: %v", v)
				}
			}
		})
	}
}

func TestNormalize(t *testing.T) {
	tests := []struct {
		name  string
		input []float64
		min   int
		max   int
		want  []float64
	}{
		{
			name:  "simple normalization",
			input: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
			min:   0,
			max:   1,
			want:  []float64{0.0, 0.25, 0.5, 0.75, 1.0},
		},
		{
			name:  "same value array",
			input: []float64{2.0, 2.0, 2.0},
			min:   0,
			max:   1,
			want:  []float64{2.0, 2.0, 2.0}, // Should return input unchanged
		},
		{
			name:  "negative to positive range",
			input: []float64{-2.0, -1.0, 0.0, 1.0, 2.0},
			min:   -1,
			max:   1,
			want:  []float64{-1.0, -0.5, 0.0, 0.5, 1.0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Normalize(tt.input, tt.min, tt.max)

			if len(got) != len(tt.want) {
				t.Errorf("Normalize() length = %v, want %v", len(got), len(tt.want))
			}

			// Check if values match expected (within floating-point precision)
			for i := range got {
				if math.Abs(got[i]-tt.want[i]) > 1e-10 {
					t.Errorf("Normalize()[%d] = %v, want %v", i, got[i], tt.want[i])
				}
			}

			// For non-constant arrays, check if values are within specified range
			if len(tt.input) > 1 && tt.input[0] != tt.input[1] {
				for _, v := range got {
					if v < float64(tt.min) || v > float64(tt.max) {
						t.Errorf("Found value outside [%v,%v] range: %v", tt.min, tt.max, v)
					}
				}
			}
		})
	}
}
