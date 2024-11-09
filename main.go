package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatalln("Invalid number of arguments\nUsage : ./math-skills <filename>")
	}
	fileContent, err := os.ReadFile(args[0])
	if err != nil {
		log.Fatalln(err)
	}
	population, err := Atoi(strings.Fields(string(fileContent)))
	if err != nil {
		log.Fatalln(err)
	}
	if len(population) <= 0 {
		log.Fatalln("Population data is empty")
	}
	average := math.Round(sum(population) / float64(len(population)))
	median := math.Round(Median(population))
	variance := math.Round(Variance(population, average))
	standardDeviation := StandardDeviation(variance)
	fmt.Printf("Average: %.0f\nMedian: %.0f\nVariance: %.0f\nStandard Deviation: %.0f\n", average, median, variance, standardDeviation)

}
