/*
   Reverse Minesweeper

   Given a grid of mine locations (where . are empty cells and x are mines), your goal is to display the grid like it appears if you win the game.
   Each position is a digit indicating the number of mines bordering it (including diagonals).
   The mines (x) don't appear anymore. Mines and positions that do not border any mines both appear as empty cells (.).

   Input
       • w: represents the width of the grid.
       • h: represents the height of the grid.
       • lines: an array of strings which each string represents the line of the minefield, with dots (.) or mines (x).

   Output
   An h x w array of strings which represents the line of the minefield, with dots (.) or mines (x).

   Constraints:
       • 1 ≤ w ≤ 30
       • 1 ≤ h ≤ 30

   Example 1:
       Input: w = 16, h =9,
           lines = [
               "................",
               "................",
               "................",
               "................",
               "................",
               "....x...........",
               "................",
               "................",
               "................",
           ]
       Output: [
           "................",
           "................",
           "................",
           "................",
           "...111..........",
           "...1x1..........",
           "...111..........",
           "................",
           "................",
       ]

   source: codingame
*/

package challenge_009

import (
	"errors"
	"strconv"
	"strings"

	"code-challenges-go/utils"
)

const (
	INVALID_LINE   = "the length of each string in input 'lines' should equal to 'w'"
	INVALID_LINES  = "the length of input 'lines' should equal to 'h'"
	OUT_OF_RANGE_H = "the value of input 'h' should be between 1 and 30"
	OUT_OF_RANGE_W = "the value of input 'w' should be between 1 and 30"
)

type minePoint struct {
	x, y int
}

func isValid(w, h int, lines []string) error {
	switch {
	case w < 1 || w > 30:
		return errors.New(OUT_OF_RANGE_W)
	case h < 1 || h > 30:
		return errors.New(OUT_OF_RANGE_H)
	case len(lines) != h:
		return errors.New(INVALID_LINES)
	case !func(lines []string) bool {
		valid := true

		for _, line := range lines {
			validContent := true

			splitted := strings.Split(line, "")
			for _, content := range splitted {
				validContent = validContent && (content == "." || content == "x")
			}

			valid = valid && len(line) == w && validContent
		}

		return valid
	}(lines):
		return errors.New(INVALID_LINE)
	default:
		return nil
	}
}

func Solution(w, h int, lines []string) ([]string, error) {
	err := isValid(w, h, lines)
	if err != nil {
		return nil, err
	}

	minePoints := []minePoint{}

	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			if string([]rune(lines[j])[i]) == "x" {
				minePoints = append(minePoints, minePoint{x: i, y: j})
			}
		}
	}

	resultLines := []string{}

	for j := 0; j < h; j++ {
		resultContents := strings.Split(lines[j], "")

		for i := 0; i < w; i++ {
			if resultContents[i] == "x" {
				resultContents[i] = "."
			} else {
				filtered := utils.Filter(minePoints, func(point minePoint) bool {
					x := point.x
					y := point.y

					return (i == x-1 || i == x+1 || i == x) && (j == y-1 || j == y+1 || j == y)
				})

				if len(filtered) > 0 {
					resultContents[i] = strconv.Itoa(len(filtered))
				} else {
					resultContents[i] = "."
				}
			}
		}

		resultLine := strings.Join(resultContents, "")

		resultLines = append(resultLines, resultLine)
	}

	return resultLines, nil
}
