# Loomis

A Golang implementation of digital image processing algorithms based on "Digital Image Processing" by Gonzalez and Woods. The project offers both sequential and concurrent implementations of various image processing techniques.

## Features

- Intensity Transformations
  - Grayscale conversion
  - Log transformation
  - Power-law (Gamma) transformation
  - Bit plane slicing
  - Bit number slicing
  - Histogram equalization

- Spatial Filtering
  - Linear smoothing filters
  - Non-linear smoothing filters
  - Gaussian smoothing
  - Laplacian sharpening
  - Unsharp masking
  - Gradient filters (Sobel, Roberts Cross)

- Frequency Domain
  - Discrete Fourier Transform

- Statistical Functions
  - Gaussian PDF
  - Rayleigh PDF
  - Histogram analysis

## Installation

Running:
$cd pkg
$go build
$cd ..
$go install .

References:
Rafael C. Gonzalez and Richard E. Woods. 2008. Digital Image Processing. Prentice Hall, Upper Saddle River, N.J.

Photo Credits:
Green Bee Eater photograph - https://unsplash.com/photos/L4hg5o67jdw
