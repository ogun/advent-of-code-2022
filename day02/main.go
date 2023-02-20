package main

func Part1() int {
	sum := 0
	for _, game := range Data {
		response := normalizeResponse(game.response)
		result := playGame(game.opponent, response)
		additional := additionalScore(response)

		sum += result + additional
	}

	return sum
}

func Part2() int {
	sum := 0
	for _, game := range Data {
		response := findResponse(game.opponent, game.response)
		result := playGame(game.opponent, response)
		additional := additionalScore(response)

		sum += result + additional
	}

	return sum
}

func findResponse(opponent, response string) string {
	// X lose, Y draw, and Z win.
	if response == "X" {
		switch opponent {
		case "A":
			return "C"
		case "B":
			return "A"
		case "C":
			return "B"

		}
	}

	if response == "Z" {
		switch opponent {
		case "A":
			return "B"
		case "B":
			return "C"
		case "C":
			return "A"

		}
	}

	return opponent
}

func normalizeResponse(response string) string {
	if response == "X" {
		return "A"
	}
	if response == "Y" {
		return "B"
	}

	return "C"
}

func additionalScore(response string) int {
	//1 for Rock, 2 for Paper, and 3 for Scissors

	switch response {
	case "A":
		return 1
	case "B":
		return 2
	case "C":
		return 3
	}

	return 0
}

func playGame(opponent, response string) int {
	// A for Rock, B for Paper, and C for Scissors
	if response == opponent {
		return 3
	}

	switch opponent {
	case "A":
		switch response {
		case "B":
			return 6
		case "C":
			return 0
		}
	case "B":
		switch response {
		case "A":
			return 0
		case "C":
			return 6
		}
	case "C":
		switch response {
		case "A":
			return 6
		case "B":
			return 0
		}
	}

	return 0
}
