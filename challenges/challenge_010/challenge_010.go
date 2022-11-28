/*
   Fire Control

   You need to cut down some trees in a forest fire to stop the fire from spreading.
   Write code to help you determine the least amount of trees to cut to contain the fire.

   The size of the forest is a 6 by 6 grid.
   Fire can spread sideways as well as diagonally.
   To stop the fire, you need to remove two layers of trees.

   # Tree
   = Cut-down tree
   o Empty space
   * Fire

   If there are no trees that can be saved by blocking the fire, output: 'JUST RUN'.
   if there is no fire, output: 'RELAX'.

   Input
       • forest: an array of strings which represents the 6 x 6 grid forest.

   Output
   The number of trees to be cut or 'JUST RUN' or 'RELAX'

   Constraints:
       • forest length === 6
       • forest width === 6

   Example 1:
       Input:
           forest = [
               "*#####",
               "######".
               "######",
               "######",
               "######",
               "######",
           ]
       To be cut:
           forest = [
               "*12###",
               "345###".
               "678###",
               "######",
               "######",
               "######",
           ]
       Output: 8

   Example 2:
       Input:
           forest = [
               "######",
               "######".
               "######",
               "######",
               "======",
               "******",
           ]
       To be cut:
           forest = [
               "######",
               "######".
               "######",
               "123456",
               "======",
               "******",
           ]
       Output: 6

   source: codingame
*/

package challenge_010

import (
	"errors"
	"strconv"
	"strings"

	"code-challenges-go/utils"
)

const (
	INVALID_FOREST = "the value of input 'forest' should be an array of strings which represents the 6 x 6 grid forest an contain either '#', '=', 'o', or '*'"
)

type TFirePoint struct {
	x, y int
}

func isValid(forest []string) error {
	switch {
	case len(forest) != 6 || !utils.Reduce(forest, true, func(valid bool, forestLine string, i int, forest []string) bool {
		splitted := strings.Split(forestLine, "")

		validContent := utils.Reduce(splitted, true, func(validContent bool, content string, i int, splitted []string) bool {
			return validContent && (content == "#" || content == "=" || content == "o" || content == "*")
		})

		return valid && len(forestLine) == 6 && validContent
	}):
		return errors.New(INVALID_FOREST)
	default:
		return nil
	}
}

func Solution(forest []string) (string, error) {
	err := isValid(forest)
	if err != nil {
		return "", err
	}

	removedTress := 0
	firePoints := []TFirePoint{}

	for j := 0; j < 6; j++ {
		for i := 0; i < 6; i++ {
			if string([]rune(forest[j])[i]) == "*" {
				firePoints = append(firePoints, TFirePoint{x: i, y: j})
			}
		}
	}

	if len(firePoints) == 0 {
		return "RELAX", nil
	}

	resultForest := []string{}

	for j := 0; j < 6; j++ {
		resultSquares := strings.Split(forest[j], "")

		for i := 0; i < 6; i++ {
			if string([]rune(forest[j])[i]) == "#" {
				fires := utils.Filter(firePoints, func(point TFirePoint) bool {
					x := point.x
					y := point.y

					return (i == x-2 || i == x-1 || i == x || i == x+1 || i == x+2) && (j == y-2 || j == y-1 || j == y || j == y+1 || j == y+2)
				})

				if len(fires) > 0 {
					resultSquares[i] = "="
					removedTress += 1
				}
			}
		}

		resultLine := strings.Join(resultSquares, "")

		resultForest = append(resultForest, resultLine)
	}

	anyTreesRemain := utils.Reduce(resultForest, false, func(anyTreesRemain bool, forestLine string, i int, resultForest []string) bool {
		splitted := strings.Split(forestLine, "")

		return anyTreesRemain || utils.Includes(splitted, "#")
	})

	if anyTreesRemain {
		return strconv.Itoa(removedTress), nil
	} else {
		return "JUST RUN", nil
	}
}
