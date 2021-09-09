package main

import (
	"math"
	"sync"
)

/**
 * Calculates the value of a wavelet at a given pixel
 *
 * @param gamma the aspect ratio of the wavelet
 * @param theta the orientation of the wavelet to the horizontal (radians)
 * @param lambda the wavelength of the wavelet (pixels)
 * @param sigma the standard deviation of the wavelet (pixels)
 * @param phi the phase angle of the wavelet (radians)
 * @param x the horizontal location of the pixel on the plane
 * @param y the vertical location of the pixel on the plane
 *
 * @return the value of the pixel at (x, y)
 */
func gabor(gamma float32, theta float32, lambda float32, sigma float32, phi float32, x int, y int) float32 {

	// Cast to 64-bit floating points
	gamma64 := float64(gamma)
	theta64 := float64(theta)
	lambda64 := float64(lambda)
	sigma64 := float64(sigma)
	phi64 := float64(phi)
	x64 := float64(x)
	y64 := float64(y)

	// Calculate intermediate values
	xPrime := x64 * math.Cos(theta64) + y64 * math.Sin(theta64)
	yPrime := -1 * x64 * math.Sin(theta64) + y64 * math.Cos(theta64)

	// Calculate the gaussian component
	a := math.Exp((x64 * x64 + gamma64 * gamma64 * yPrime * yPrime) / (-2 * sigma64 * sigma64))

	// Calculate the sinusoidal component
	b := math.Cos(((2 * math.Pi * xPrime) / lambda64) + phi64)

	// Multiply the gaussian and sinusoidal components
	g := a * b

	// Cast to 32-bits and return
	return float32(g)
}

/**
 * Creates a matrix containing a wavelet
 *
 * @param beta the bandwidth of the wavelet (octaves)
 * @param gamma the aspect ratio of the wavelet
 * @param theta the orientation of the wavelet to the horizontal (radians)
 * @param lambda the wavelength of the wavelet (pixels)
 * @param sigma the standard deviation of the wavelet (pixels)
 * @param phi the phase angle of the wavelet (radians)
 *
 * @return a matrix containing the wavelet in its centre
 */
func wavelet(beta float32, gamma float32, theta float32, lambda float32, phi float32, width int, height int) [][]float32 {

	// Calculate the standard deviation of the wavelet
	twoPowBeta := math.Pow(2, float64(beta))
	sigma := float32(float64(lambda) * ((1 / math.Pi) * math.Sqrt(math.Log(2) / 2) * ((twoPowBeta + 1) / (twoPowBeta - 2))))

	// Create empty matrix to store the wavelet
	//matrix := make(make([]float32, width), height)
	matrix := make([][]float32, height)
	for row := range matrix {
		matrix[row] = make([]float32, width)
	}

	// Create a wait group
	var waitGroup sync.WaitGroup
	waitGroup.Add(height)

	// Iterate over the columns
	for j := 0; j < height; j++ {

		// Vertically centre the wavelet within the matrix
		y := j - (height / 2)

		// Process each row on its won goroutine
		go func(gamma float32, theta float32, lambda float32, sigma float32, phi float32, j int, y int) {
			defer waitGroup.Done()

			// Iterate over every pixel in the current row
			for i := 0; i < width; i++ {

				// Horizontally centre the wavelet within the matrix
				x := i - (width / 2)

				// Calculate the value of the current pixel
				matrix[i][j] = gabor(gamma, theta, lambda, sigma, phi, x, y)
			}
		} (gamma, theta, lambda, sigma, phi, j, y)
	}

	// Wait for all goroutines to finish
	waitGroup.Wait()
	return matrix
}