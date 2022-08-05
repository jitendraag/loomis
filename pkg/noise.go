package pkg

import (
	"math"
)


func GaussianPdf(mean int, standardDeviation int) []float64 {
	// From DIP 5.2.2
	var pdf []float64 = make([]float64, MaxGrayscaleLevels)
	
	for i := 0; i < MaxGrayscaleLevels; i++ {
		pdf[i] = (float64)((1 / ( math.Sqrt(2.0 * math.Pi) * float64(standardDeviation))) * math.Exp(-math.Pow(float64(i - mean), 2) / (2 * math.Pow(float64(standardDeviation), 2))))
	}

	pdf = Normalize(pdf, 0, 1)

	return pdf
}

func RayleighPdf(constant_a float64, constant_b float64) []float64 {
	// From DIP 5.2.2
	var pdf []float64 = make([]float64, MaxGrayscaleLevels)
	
	for i := 0; i < MaxGrayscaleLevels; i++ {
		rayleigh_step := float64(i) * 0.01
		if rayleigh_step < constant_a {
			pdf[i] = 0
			continue
		}
		pdf[i] = (2.0/constant_b) * (rayleigh_step - constant_a) * math.Exp(-math.Pow((rayleigh_step-constant_a),2) / constant_b)
	}

	pdf = Normalize(pdf, 0, 1)

	return pdf
}

func Normalize(input []float64, range_minimum int, range_maximum int) []float64{
	// Scale the given input array to values between minimum and maximum range
	// https://en.wikipedia.org/wiki/Feature_scaling
	var minimum float64 = input[0]
	var maximum float64 = input[0]
	var output []float64 = make([]float64, len(input))

	for _, value := range input {
		if value < minimum {
			minimum = value
		}
		if value > maximum {
			maximum = value
		}
	}

	if minimum == maximum {
		// TODO, return error
		return input
	}


	for index, value := range input {
		output[index] = float64(range_minimum) + ( (value - minimum) * float64(range_maximum - range_minimum) / (maximum - minimum))
	}
	return output
}