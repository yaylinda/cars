import csv

LINDA_WEIGHTS = {
    "Styling": 2.0,
    "Acceleration": 0.0,
    "Handling": 3.0,
    "Fun Factor": 0.0,
    "Cool Factor": 1.0,
    "Features": 9.0,
    "Comfort": 8.0,
    "Quality": 5.0,
    "Practicality": 3.0,
    "Value": 10.0,
}

SEAN_WEIGHTS = {
    "Styling": 6.5,
    "Acceleration": 5.2,
    "Handling": 3.1,
    "Fun Factor": 8.1,
    "Cool Factor": 8.5,
    "Features": 7.5,
    "Comfort": 2.3,
    "Quality": 3.7,
    "Practicality": 0,
    "Value": 0.2,
}


def combine_weights():
    combined = {}
    for key in LINDA_WEIGHTS:
        combined[key] = ( LINDA_WEIGHTS[key] + SEAN_WEIGHTS[key] ) / 2.0
    return combined


def read_data():
    rows = []
    with open('dougscores.csv', 'r') as f:
        reader = csv.DictReader(f)
        for row in reader:
            rows.append(row)
    return rows


def weight_scores(scores, weights):
    result = {}

    for car in scores:



    return result


def main():
    scores = read_data()

    combined_weights = combine_weights()

    scores_linda = weight_scores(scores, LINDA_WEIGHTS)
    scores_sean = weight_scores(scores, SEAN_WEIGHTS)
    scores_combined = weight_scores(scores, combined_weights)


if __name__ == '__main__':
    main()

