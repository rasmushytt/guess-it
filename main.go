package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	totalX, totalY, total, deviation, devSqrX, devSqrY float64
	data                                               []float64
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	//Set some random numbers
	data = append(data, 159, 164)

	//Get new values with scanner
	for scanner.Scan() {
		inp, _ := strconv.ParseFloat(scanner.Text(), 64)

		//Count values
		total++
		totalX += total
		totalY += inp
		data = append(data, inp)

		//Calc average
		meanX := totalX / total
		meanY := totalY / total

		//Find devation total, deviation of X & Y and their squared values
		for x, y := range data {
			deviation += (float64(x) - meanX) * (y - meanY)
			devSqrX += math.Pow(float64(x)-meanX, 2)
			devSqrY += math.Pow(y-meanY, 2)
		}
		//Find more values
		stdX, stdY := math.Pow(devSqrX/total, 0.5), math.Pow(devSqrY/total, 0.5)
		rise := deviation / (math.Pow(devSqrX, 0.5) * math.Pow(devSqrY, 0.5))
		slope := rise * (stdY / stdX)
		intercept := meanY - slope*meanX

		//Predict values. Adding 6 to increase chance of correct value
		lower_bounds := slope*(total-1) + intercept - 6
		upper_bounds := slope*(total-1) + intercept + 6
		fmt.Printf("%.0f %.0f\n", lower_bounds, upper_bounds)
	}
}
