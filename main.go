package main

import (
	"Functions/Functions" // Importing user-defined package Functions
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go data.txt")
	} else {
		// Open the file named "data.txt"
		file, err := os.Open("data.txt")
		if err != nil {
			fmt.Println("Error:", err)
		}
		defer file.Close() // Close the file when main function exits
		//create a scanner to read from the file
		scanner := bufio.NewScanner(file)
		var floats []float64 // Slice to store float64 values read from the file

		// Read lines from the entire file
		for scanner.Scan() {
			line := scanner.Text()                   //read the current line
			num, err := strconv.ParseFloat(line, 64) // Convert the line to a float64
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			floats = append(floats, num) // Append the float64 value to the floats slice
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error:", err)
		}

		//Calculate the statistics using Functions package and rounding them to integers, i.e Average,Median,Variance and Standard deviation
		Average := Functions.Average(floats)
		roundAverage := int(math.Round(Average))
		Median := Functions.Median(floats)
		roundMedian := int(math.Round(Median))
		Variance := Functions.Variance(floats, Average)
		roundVariance := int(math.Round(Variance))
		StandardDeviation := Functions.StandardDev(Variance)
		roundStandardDeviation := int(math.Round(StandardDeviation))

		//Print the rounded integer values
		fmt.Println("Average:", roundAverage)
		fmt.Println("Median:", roundMedian)
		fmt.Println("Variance:", roundVariance)
		fmt.Println("Standard Deviation:", roundStandardDeviation)
	}
}
