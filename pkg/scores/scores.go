package scores

import (
	"fmt"
	"strconv"
)

type CarScore struct {
	car    string
	scores map[string]float64
}

// ConvertScores converts the raw rows read from the csv file to a list of
// CarScores
func ConvertScores(rawScores []map[string]string) []CarScore {
	var scores []CarScore

	for _, r := range rawScores {
		row := make(map[string]float64)
		for key, value := range r {
			s, err := strconv.Atoi(value)
			if err != nil {
				break
			}
			row[key] = float64(s)
		}

		scores = append(
			scores, CarScore{
				car: fmt.Sprintf(
					"%s %s (%s)",
					r["Make"],
					r["Model"],
					r["Year"],
				),
				scores: row,
			},
		)
	}

	return scores
}

// AverageWeights given a list of weights, averages the weights across each category
func AverageWeights(weights []map[string]float64) map[string]float64 {
	numWeights := len(weights)

	combined := make(map[string]float64)

	for key, _ := range weights[0] {
		sum := float64(0)
		for i, _ := range weights {
			sum += weights[i][key]
		}
		combined[key] = sum / float64(numWeights)
	}

	return combined
}

// WeightScores apply the weights to the original scores for each row
func WeightScores(
	original []CarScore,
	weights map[string]float64,
	maxWeight float64,
) []CarScore {
	var results []CarScore

	for _, row := range original {
		weightedScores := make(map[string]float64)
		for key, value := range row.scores {
			if weight, exists := weights[key]; exists {
				weightedScores[key] = value * (weight / maxWeight)
			}
		}

		results = append(
			results,
			CarScore{
				car:    row.car,
				scores: weightedScores,
			},
		)
	}

	return results
}
