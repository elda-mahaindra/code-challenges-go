/*
   There is no Spoon - Episode 1

   The game is played on a rectangular grid with a given size. Some cells contain power nodes. The rest of the cells are empty.

   The goal is to find, when they exist, the horizontal and vertical neighbors of each node.

   Rules
   To do this, you must find each (x1,y1) coordinates containing a node, and display the (x2,y2) coordinates of the next node to the right, and the (x3,y3) coordinates of the next node to the bottom within the grid.

   If a neighbor does not exist, you must output the coordinates -1 -1 instead of (x2,y2) and/or (x3,y3).

   You lose if:
   • You give an incorrect neighbor for a node.
   • You give the neighbors for an empty cell.
   • You compute the same node twice.
   • You forget to compute the neighbors of a node.

   Input
       • width: one integer width for the number of cells along the x axis.
       • height: one integer height for the number of cells along the y axis.
       • lines: a height length of string array containing width characters. A dot (.) represents an empty cell. A zero (0) represents a cell containing a node.

   Output
   An array of strings where each string represents six integers on each node joined into a string (separated by a space): "x1 y1 x2 y2 x3 y3"
   Where:
   • (x1,y1) the coordinates of a node
   • (x2,y2) the coordinates of the closest neighbor on the right of the node
   • (x3,y3) the coordinates of the closest bottom neighbor
   If there is no neighbor, the coordinates should be -1 -1.

   Constraints:
       • 0 < width ≤ 30
       • 0 < height ≤ 30
       • 0 ≤ x1 < width
       • 0 ≤ y1 < height
       • -1 ≤ x2, x3 < width
       • -1 ≤ y2, y3 < height

   Example 1:
       Input: width = 2, height = 2,
           lines = [
               "00",
               "0."
           ]
       Output: ["0 0 1 0 0 1", "1 0 -1 -1 -1 -1", "0 1 -1 -1 -1 -1"]

   Example 2:
       Input: width = 5, height = 1, lines = ["0.0.0"]
       Output: ["0 0 2 0 -1 -1", "2 0 4 0 -1 -1", "4 0 -1 -1 -1 -1"]

   source: codingame
*/

package challenge_013

import (
	"errors"
	"fmt"

	"code-challenges-go/utils"
)

const (
	OUT_OF_RANGE_HEIGHT = "the length of input 'height' should be between 1 and 30"
	OUT_OF_RANGE_WIDTH  = "the length of input 'width' should be between 1 and 30"
	INVALID_GRID_HEIGHT = "the length of input 'lines' should be equal to the input 'height'"
	INVALID_GRID_WIDTH  = "the length of line in input 'lines' should be equal to the input 'width'"
)

func isValid(width, height int, lines []string) error {
	switch {
	case width < 1 || width > 30:
		return errors.New(OUT_OF_RANGE_WIDTH)
	case height < 1 || height > 30:
		return errors.New(OUT_OF_RANGE_HEIGHT)
	case len(lines) != height:
		return errors.New(INVALID_GRID_HEIGHT)
	case !utils.Reduce(lines, true, func(valid bool, line string, i int, lines []string) bool {
		return valid && len(line) == width
	}):
		return errors.New(INVALID_GRID_WIDTH)
	default:
		return nil
	}
}

func Solution(width, height int, lines []string) ([]string, error) {
	err := isValid(width, height, lines)
	if err != nil {
		return nil, err
	}

	nodes := []string{}

	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			if string([]rune(lines[j])[i]) == "0" {
				node := fmt.Sprintf("%d %d", i, j)

				if i >= width-1 {
					node += " -1 -1"
				} else {
					found := false
					ii := i + 1

					for !found && ii < width {
						if string([]rune(lines[j])[ii]) == "0" {
							node += fmt.Sprintf(" %d %d", ii, j)
							found = true
						}

						ii++
					}

					if !found {
						node += " -1 -1"
					}
				}

				if j >= height-1 {
					node += " -1 -1"
				} else {
					found := false
					jj := j + 1

					for !found && jj < height {
						if string([]rune(lines[jj])[i]) == "0" {
							node += fmt.Sprintf(" %d %d", i, jj)
							found = true
						}

						jj++
					}

					if !found {
						node += " -1 -1"
					}
				}

				nodes = append(nodes, node)
			}
		}
	}

	return nodes, nil
}
