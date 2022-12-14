/*
   Ghost Legs

   Ghost Legs is a kind of lottery game common in Asia. It starts with a number of vertical lines.
   Between the lines there are random horizontal connectors binding all lines into a connected diagram, like the one below.

   A  B  C
   |  |  |
   |--|  |
   |  |--|
   |  |--|
   |  |  |
   1  2  3

   To play the game, a player chooses a line in the top and follow it downwards.
   When a horizontal connector is encountered, he must follow the connector to turn to another vertical line and continue downwards.
   Repeat this until reaching the bottom of the diagram.

   In the example diagram, when you start from A, you will end up in 2. Starting from B will end up in 1.
   Starting from C will end up in 3. It is guaranteed that every top label will map to a unique bottom label.

   Given a Ghost Legs diagram, find out which top label is connected to which bottom label. List all connected pairs.

   Input
       • W: an integer that represents the width of the diagram.
       • H: an integer that represents the height of the diagram.
       • diagram: an array of string that represents the diagram.
           The diagram itself is composed of characters: '|' and '-', (and space).
           The top line in the diagram has a number of labels T.
           The bottom line contains labels B.

           Each T and B is a single visible ASCII character that can be of any random value. Do not assume they will always be ABC or 123.
           As a rule of the game, left and right horizontal connectors will never appear at the same point.
           All diagrams are having the same style as the test cases.

   Output
   An array of strings where each string represents the pairs between top and bottom labels, TB, in the order of the top labels from Left to Right.

   Constraints:
   • 3 < W ≤ 100
   • 3 < H ≤ 100
   • ASCII characters used in the top and bottom labels are in range of Hex 21 to Hex 7E, inclusive

   Example 1:
       Input:
           W = 7,
           H = 7,
           diagram = [
               "A  B  C",
               "|  |  |",
               "|--|  |",
               "|  |--|",
               "|  |--|",
               "|  |  |",
               "1  2  3",
           ]
       Output: ["A2", "B1", "C3"]

   source: codingame
*/

package challenge_030

import (
	"errors"
	"strings"

	"code-challenges-go/utils"
)

const (
	INVALID_DIAGRAM = "the width and height of input 'diagram' should be equal to input 'W' and input 'H'"
	OUT_OF_RANGE_H  = "the value of input 'H' should be between 4 and 100"
	OUT_OF_RANGE_W  = "the value of input 'W' should be between 4 and 100"
)

type TPosition = []int

func isValid(W, H int, diagram []string) error {
	switch {
	case W <= 3 || W > 100:
		return errors.New(OUT_OF_RANGE_W)
	case H <= 3 || H > 100:
		return errors.New(OUT_OF_RANGE_H)
	case len(diagram) != H ||
		!utils.Reduce(diagram, true, func(valid bool, row string, i int, diagram []string) bool {
			return valid && len(row) == W
		}):
		return errors.New(INVALID_DIAGRAM)
	default:
		return nil
	}
}

func Solution(W, H int, diagram []string) ([]string, error) {
	err := isValid(W, H, diagram)
	if err != nil {
		return nil, err
	}

	tops := strings.Split(diagram[0], "")
	tops = utils.Filter(tops, func(el string) bool {
		return el != " "
	})

	pairs := utils.Map(tops, func(top string, i int, tops []string) string {
		bottom := ""

		currentPosition := TPosition{strings.Index(diagram[0], top), 1}

		for len(bottom) == 0 {
			x := currentPosition[0]
			y := currentPosition[1]

			if string([]rune(diagram[y])[x]) != "|" {
				bottom = string([]rune(diagram[y])[x])
				break
			}

			if x != 0 && string([]rune(diagram[y])[x-1]) == "-" {
				currentPosition = TPosition{x - 3, y + 1}
			} else if x != W-1 && string([]rune(diagram[y])[x+1]) == "-" {
				currentPosition = TPosition{x + 3, y + 1}
			} else {
				currentPosition = TPosition{x, y + 1}
			}
		}

		return top + bottom
	})

	return pairs, nil
}
