package main

import (
	"math"
	"sort"
	"strconv"
)

// function to convert slice of strings to []float64
func Atoi(values []string) ([]float64, error) {
	result := []float64{}
	for _, decVal := range values {
		numVal, err := strconv.ParseFloat(decVal, 64)
		if err != nil {
			return nil, err
		}
		result = append(result, numVal)
	}
	return result, nil
}

// Median: Middle value when all values are sorted
// If N is odd: Median = Middle value
// If N is even: Median = Average of the two middle values

func Median(population []float64) float64 {
	length := len(population)
	sort.Slice(population, func(i, j int) bool {
		return population[i] < population[j]
	})
	if len(population)%2 == 0 {
		return (population[(length/2)] + population[(length/2)-1]) / 2
	}
	return population[(length / 2)]
}

// Variance = Σ((x - mean)²) / N
func Variance(population []float64, average float64) float64 {
	var sumSquares float64
	for _, val := range population {
		diff := val - average
		sumSquares += diff * diff
	}
	return sumSquares / float64(len(population))
}

// Standard Deviation = √Variance

func StandardDeviation(variance float64) float64 {
	return math.Round(math.Sqrt(variance))
}

func sum(nums []float64) float64 {
	var total float64
	for _, num := range nums {
		total += num
	}
	return total
}
