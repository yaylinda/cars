package main

import (
	"fmt"

	"github.com/yaylinda/cars/pkg/csv"
	"github.com/yaylinda/cars/pkg/scores"
)

var (
	dougScoresFile = "dougscores.csv"

	maxWeight = float32(10.0)

	lindaWeights = map[string]float32{
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

	seanWeights = map[string]float32{
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

func main() {
	rawDougScores, err := csv.Read(dougScoresFile)
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}

	dougScores := scores.ConvertScores(rawDougScores)

	combinedWeights := scores.AverageWeights(
		[]map[string]float32{
			lindaWeights,
			seanWeights,
		},
	)

	lindaScores := scores.WeightScores(dougScores, lindaWeights, maxWeight)
	seanScores := scores.WeightScores(dougScores, seanWeights, maxWeight)
	combinedScores := scores.WeightScores(
		dougScores,
		combinedWeights,
		maxWeight,
	)

	// fmt.Println(lindaScores)
	// fmt.Println(seanScores)
	// fmt.Println(combinedScores)

	fmt.Printf("Best for Linda (by total): %v\n\n", lindaScores[0])
	fmt.Printf("Best for Sean (by total): %v\n\n", seanScores[0])
	fmt.Printf("Best Combined (by total): %v\n\n", combinedScores[0])
}
