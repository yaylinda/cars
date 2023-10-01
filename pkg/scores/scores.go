package scores

import (
	"fmt"
	"sort"
	"strconv"
)

type CarScore struct {
	car    string
	scores map[string]float64
	total  float64
}

// ConvertScores converts the raw rows read from the csv file to a list of
// CarScores
func ConvertScores(rawScores []map[string]string) []CarScore {
	var scores []CarScore

	for _, r := range rawScores {
		total := float64(0)
		row := make(map[string]float64)

		for key, value := range r {
			s, err := strconv.Atoi(value)
			if err != nil {
				break
			}
			row[key] = float64(s)
			total += float64(s)
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
				total:  total,
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

// WeightScores apply the weights to the original scores for each row.
// By default, returns the cars in descending order, based on the total score.
func WeightScores(
	original []CarScore,
	weights map[string]float64,
	maxWeight float64,
) []CarScore {
	var results []CarScore

	for _, row := range original {
		total := float64(0)
		weightedScores := make(map[string]float64)

		for key, value := range row.scores {
			if weight, exists := weights[key]; exists {
				weightedScores[key] = value * (weight / maxWeight)
				total += weightedScores[key]
			}
		}

		results = append(
			results,
			CarScore{
				car:    row.car,
				scores: weightedScores,
				total:  total,
			},
		)
	}

	sortByTotal(results, true)

	return results
}

// SortByTotal performs an in-place sort on the given CarScores array based on
// the total score.
// desc=true -> highest score first
// desc=false -> lowest score first
func sortByTotal(scores []CarScore, desc bool) {
	sort.Slice(
		scores, func(i, j int) bool {
			if desc {
				return scores[i].total > scores[j].total
			} else {
				return scores[i].total < scores[j].total
			}
		},
	)
}
