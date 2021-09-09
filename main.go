package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	// Create command line arguments
	betaPtr := flag.Float64("beta", 2, "the bandwidth of the wavelet (octaves)")
	gammaPtr := flag.Float64("gamma", 1, "the aspect ratio of the wavelet")
	thetaPtr := flag.Float64("theta", 0, "the orientation of the wavelet to the horizontal (radians)")
	lambdaPtr := flag.Float64("lambda", 50, "the wavelength of the wavelet (pixels)")
	phiPtr := flag.Float64("phi", 0, "the phase angle of the wavelet (radians)")
	widthPtr := flag.Int("width", 500, "the width of the image (pixels)")
	heightPtr := flag.Int("height", 500, "the height of the image (pixels)")
	outputPtr := flag.String("output", "wavelet.csv", "the path of the output file (CSV format)")

	// Parse command line arguments
	flag.Parse()
	beta := float32(*betaPtr)
	gamma := float32(*gammaPtr)
	theta := float32(*thetaPtr)
	lambda := float32(*lambdaPtr)
	phi := float32(*phiPtr)
	width := *widthPtr
	height := *heightPtr
	output := *outputPtr

	// Generate the wavelet
	fmt.Println("\nGenerating wavelet...")
	fmt.Printf("β :\t %f\nγ :\t %f\nθ :\t %f\nλ :\t %f\nφ :\t %f\nw :\t %d\nh :\t %d\n", beta, gamma, theta, lambda, phi, width, height)
	wavelet := wavelet(beta, gamma, theta, lambda, phi, width, height)
	fmt.Println("Done.")

	// Save the wavelet to a CSV file
	fmt.Println("\nSaving to file...")
	fmt.Printf("o :\t%s\n", output)
	saveToFile(wavelet, output)
	fmt.Println("Done.")
}

/**
 * Saves a matrix containing a wavelet in a CSV file
 *
 * @param wavelet the matrix containing the wavelet
 * @param path the path of the file
 */
func saveToFile(wavelet [][]float32, path string) {

	// Create the file
	f, err := os.Create(path)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

	// Iterate over the matrix, writing every pixel to the file
	rows := len(wavelet)
	cols := len(wavelet[0])
	for j := 0; j < rows; j++ {
		for i := 0; i < cols; i++ {
			_, _ = f.WriteString(fmt.Sprintf("%f,", wavelet[i][j]))
		}
		_, _ = f.WriteString("\n")
	}
}