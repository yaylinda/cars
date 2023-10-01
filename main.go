package main

import (
	"strconv"
)

var (
	dougScoresFile = "dougscores.csv"

	lindaWeights = map[string]float64{
		"Styling":      2.0,
		"Acceleration": 0.0,
		"Handling":     3.0,
		"Fun Factor":   0.0,
		"Cool Factor":  1.0,
		"Features":     9.0,
		"Comfort":      8.0,
		"Quality":      5.0,
		"Practicality": 3.0,
		"Value":        10.0,
	}

	seanWeights = map[string]float64{
		"Styling":      6.5,
		"Acceleration": 5.2,
		"Handling":     3.1,
		"Fun Factor":   8.1,
		"Cool Factor":  8.5,
		"Features":     7.5,
		"Comfort":      2.3,
		"Quality":      3.7,
		"Practicality": 0.0,
		"Value":        0.2,
	}
)

func combineWeights(linda, sean map[string]float64) map[string]float64 {
	combined := make(map[string]float64)
	for key, val := range linda {
		combined[key] = (val + sean[key]) / 2.0
	}
	return combined
}

func preProcess() {

}

func weightScores(scores []map[string]string, weights map[string]float64) []map[string]float64 {
	var results []map[string]float64

	for _, row := range scores {
		result := make(map[string]float64)
		result["Car"] = 0 // Placeholder to ensure "Car" key exists
		for key, value := range row {
			if weight, exists := weights[key]; exists {
				intValue, _ := strconv.Atoi(value)
				result[key] = float64(intValue) * (weight / 10.0)
			}
		}
		results = append(results, result)
	}

	return results
}

func main() {
	//dougScores, err := csv.Read()
	//if err != nil {
	//	fmt.Println("Error reading data:", err)
	//	return
	//}
	//
	//weightedScores := weightScores(dougScores, lindaWeights)
	//
	//fmt.Println(weightedScores)
}
