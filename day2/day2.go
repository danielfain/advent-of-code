package day2

import (
	"strconv"
	"strings"
)

type GameCubes struct {
	red   int
	green int
	blue  int
}

func SumOfPossibleGameIds(games string, cubes GameCubes) (sum int) {
	for _, game := range strings.Split(games, "\n") {
		gameValid := true

		id, _ := strconv.Atoi(game[5:strings.Index(game, ":")])

		hands := strings.Split(game, ": ")[1]

		for _, hand := range strings.Split(hands, ";") {
			hand = strings.Trim(hand, " ")
			handCubes := strings.Split(hand, ", ")

			for _, cube := range handCubes {
				vals := strings.Split(cube, " ")

				num, _ := strconv.Atoi(vals[0])
				color := vals[1]

				switch color {
				case "red":
					if num > cubes.red {
						gameValid = false
					}
				case "green":
					if num > cubes.green {
						gameValid = false
					}
				case "blue":
					if num > cubes.green {
						gameValid = false
					}
				}
			}
		}

		if gameValid {
			sum += id
		}
	}
	return sum
}
