/*
   Surface

   "The wars of the 21st century will be fought over water."
   Although freshwater is available in limited quantity, it’s not actually scarce.
   There’s more than enough to satisfy the current needs of the global population,
   but only if it were possible to locate and measure the bodies of water available in a geographical area!

   Your mission is to pinpoint the surface areas of water.
   You have a map which describes the contents of each square meter of a geographical zone.
   One square meter is composed of either land or water. One map can contain several bodies of water.

   Your program receives as input a list of coordinates.
   For each one you must determine the surface area of the lake which is located there.
   If there is no lake, then the surface area equals 0.

   A map in ASCII format is provided as input.
   The character # represents land and the letter O (uppercase) represents water.

   ####
   ##O#
   #OO#
   ####

   A lake is made up of a set of water squares which are horizontally or vertically adjacent.
   Two squares which are only diagonally adjacent are not part of the same lake.

   In this example, the lake which is located in coordinates (1, 2) has a surface area of 3 square meters.

   Input
       • L: an integer that represents the width of the map.
       • H: an integer that represents the height of the map.
       • rows: an array of H string where each string represents L squares of the map.
       • N: an integer that represents the number of coordinates to be tested.
       • coordinates: an array of string where each string represents the coordinate to be tested (X and Y separated by a space)

   Output
   An array of N integer where each integer represents the amount of surface areas of the lake located at the coordinates given in input.

   Constraints:
   • 0 < L < 10000
   • 0 < H < 10000
   • 0 ≤ X < L
   • 0 ≤Y < H
   • 0 < N < 1000

   Example 1:
       Input:
           L = 4,
           H = 4,
           rows = [
               "####",
               "##O#",F
               "#OO#",
               "####"
           ]
           N = 3,
           coordinates = ["0 0", "1 2", "2 1"]
       Output: [0, 3, 3]

   source: codingame
*/

package challenge_028

import (
	"errors"
	"strconv"
	"strings"

	"code-challenges-go/utils"
)

const (
	INVALID_COORDINATES = "the length of input 'coordinates' should be equal to input 'N' and each element inside it should represent a valid coordinate"
	INVALID_ROWS        = "the length of input 'rows' should be equal to input 'H' and each element inside it should be a string with the length of 'L' and consists of only '#' or 'O' character"
	OUT_OF_RANGE_L      = "the length of input 'L' should be between 1 and 9999"
	OUT_OF_RANGE_H      = "the length of input 'H' should be between 1 and 9990"
	OUT_OF_RANGE_N      = "the length of input 'N' should be between 1 and 999"
)

type TPosition = []int

func isValid(L, H int, rows []string, N int, coordinates []string) error {
	switch {
	case L < 1 || L > 9999:
		return errors.New(OUT_OF_RANGE_L)
	case H < 1 || H > 9999:
		return errors.New(OUT_OF_RANGE_H)
	case N < 1 || N > 999:
		return errors.New(OUT_OF_RANGE_N)
	case len(rows) != H || !utils.Reduce(rows, true, func(valid bool, row string, i int, rows []string) bool {
		splitted := strings.Split(row, "")

		validRow := len(splitted) == L && utils.Reduce(splitted, true, func(valid bool, char string, i int, chars []string) bool {
			return valid && (char == "#" || char == "O")
		})

		return valid && validRow
	}):
		return errors.New(INVALID_ROWS)
	case len(coordinates) != N || !utils.Reduce(coordinates, true, func(valid bool, coordinate string, i int, coordinates []string) bool {
		splitted := strings.Split(coordinate, " ")

		_, err := strconv.Atoi(splitted[0])
		if err != nil {
			return valid && false
		}

		_, err = strconv.Atoi(splitted[1])
		if err != nil {
			return valid && false
		}

		return valid && true
	}):
		return errors.New(INVALID_COORDINATES)
	default:
		return nil
	}
}

func checkNodeVisited(position TPosition, visited []TPosition) bool {
	x := position[0]
	y := position[1]

	for i := 0; i < len(visited); i++ {
		if visited[i][0] == x && visited[i][1] == y {
			return true
		}
	}

	return false
}

func filterPositionsByMark(mark string, positions []TPosition, rows [][]string) []TPosition {
	filtered := []TPosition{}
	for i := 0; i < len(positions); i++ {
		x := positions[i][0]
		y := positions[i][1]

		if rows[y][x] == mark {
			filtered = append(filtered, positions[i])
		}
	}

	return filtered
}

func findNeighbours(source TPosition, l, h int, rows [][]string) []TPosition {
	x := source[0]
	y := source[1]

	neighbours := []TPosition{}

	mark := rows[y][x]

	if l > 1 {
		if x == 0 {
			neighbours = append(neighbours, filterPositionsByMark(mark, []TPosition{{x + 1, y}}, rows)...)
		} else if x == l-1 {
			neighbours = append(neighbours, filterPositionsByMark(mark, []TPosition{{x - 1, y}}, rows)...)
		} else {
			neighbours = append(neighbours, filterPositionsByMark(mark, []TPosition{{x + 1, y}, {x - 1, y}}, rows)...)
		}
	}

	if h > 1 {
		if y == 0 {
			neighbours = append(neighbours, filterPositionsByMark(mark, []TPosition{{x, y + 1}}, rows)...)
		} else if y == h-1 {
			neighbours = append(neighbours, filterPositionsByMark(mark, []TPosition{{x, y - 1}}, rows)...)
		} else {
			neighbours = append(neighbours, filterPositionsByMark(mark, []TPosition{{x, y + 1}, {x, y - 1}}, rows)...)
		}
	}

	return neighbours
}

// all water 'definitely' connected when:
// - only one row has a maximum of 'l' minus one squares of land
func checkAllWaterDefinitelyConnected(l, h int, rows [][]string) (bool, int) {
	allWaterDefinitelyConnected := false
	rowsWithLand := 0
	totalLandSquares := 0

	for y := 0; y < h; y++ {
		row := rows[y]

		landSquares := 0

		for x := 0; x < l; x++ {
			mark := row[x]

			if mark == "#" {
				landSquares += 1
			}
		}

		if landSquares >= 1 {
			totalLandSquares += landSquares
			rowsWithLand += 1
		}
	}

	if rowsWithLand <= 1 && totalLandSquares < l {
		allWaterDefinitelyConnected = true
	}

	return allWaterDefinitelyConnected, totalLandSquares
}

func findSurfaceArea(source TPosition, l, h int, rows [][]string) int {
	x := source[0]
	y := source[1]

	if rows[y][x] == "#" {
		return 0
	}

	allWaterDefinitelyConnected, totalLandSquares := checkAllWaterDefinitelyConnected(l, h, rows)

	if allWaterDefinitelyConnected {
		return (l * h) - totalLandSquares
	}

	surfaceArea := 0
	visited := []TPosition{}

	queue := []TPosition{source}

	for len(queue) > 0 {
		dequeuedItem := queue[0]
		queue = queue[1:]

		visited = append(visited, dequeuedItem)
		surfaceArea += 1

		neighbours := findNeighbours(dequeuedItem, l, h, rows)

		for i := 0; i < len(neighbours); i++ {
			neighbour := neighbours[i]
			nodeVisited := checkNodeVisited(neighbour, visited)

			queued := false
			for j := 0; j < len(queue); j++ {
				if queue[j][0] == neighbour[0] && queue[j][1] == neighbour[1] {
					queued = true
				}
			}

			if !nodeVisited && !queued {
				queue = append(queue, neighbour)
			}
		}
	}

	return surfaceArea
}

func Solution(L, H int, rows []string, N int, coordinates []string) ([]int, error) {
	err := isValid(L, H, rows, N, coordinates)
	if err != nil {
		return nil, err
	}

	twoDimensionalRows := utils.Map(rows, func(row string, i int, rows []string) []string {
		splitted := strings.Split(row, "")

		return splitted
	})

	surfaceAreas := []int{}
	for i := 0; i < len(coordinates); i++ {
		splitted := strings.Split(coordinates[i], " ")

		x, err := strconv.Atoi(splitted[0])
		if err != nil {
			return nil, err
		}

		y, err := strconv.Atoi(splitted[1])
		if err != nil {
			return nil, err
		}

		source := TPosition{x, y}

		surfaceArea := findSurfaceArea(source, L, H, twoDimensionalRows)

		surfaceAreas = append(surfaceAreas, surfaceArea)
	}

	return surfaceAreas, nil
}
