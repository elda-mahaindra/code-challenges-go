/*
   Binary Image

   You are going to write a simple program to decode some arrays of data into a black-and-write graphic.

   The graphic is composed of n lines of black and white pixels. We use . to represent a white pixel; O to represent a black pixel.

   For example, here is one line of graphic
   ....OOO.

   We shall encode it into an array 4 3 1
   because it starts with 4 whites, then 3 blacks, then 1 white.
   We assume most lines shall start with white.

   When there is a line starting with black, we add 0 at the beginning of the encoded data, to say there is no white pixel before the first black pixel.
   For example
   OO.OOOOO
   will be encoded into 0 2 1 5.

   You must output an array with one element of string "INVALID" if the graphic is not rectangular (this doesn't mean the input lines should be the same length, but the outputs lines should be).

   In this puzzle, you will be given n lines of encoded data.
   You are going to decode it into a graphic.

   Input
       • h: an integer h represents the number rows input.
       • rows: an array of strings where each string represents the encoded line of the graphic.

   Output
   an array with one element of string "INVALID" if the grid is not a rectangle, or an array of h rows where each row consists of pixels which has been represented by "." or "O".

   Constraints:
   • h < 200

   Example 1:
       Input:
           h = 4,
           rows = [
             "1 3 2 1",
             "1 3 2 1",
             "1 3 2 1",
             "1 3 2 1",
           ],
       Output: [
         ".OOO..O",
         ".OOO..O",
         ".OOO..O",
         ".OOO..O",
       ]

   source: codingame
*/

package challenge_022

import (
	"errors"
	"strconv"
	"strings"

	"code-challenges-go/utils"
)

const (
	INVALID_ROWS   = "the length of input 'rows' should equal to the input 'h'"
	OUT_OF_RANGE_H = "the value of input 'h' should be less than 200"
)

func isValid(h int, rows []string) error {
	switch {
	case h >= 200:
		return errors.New(OUT_OF_RANGE_H)
	case len(rows) != h:
		return errors.New(INVALID_ROWS)
	default:
		return nil
	}
}

func Solution(h int, rows []string) ([]string, error) {
	err := isValid(h, rows)
	if err != nil {
		return nil, err
	}

	mapped := utils.Map(rows, func(row string, i int, rows []string) string {
		splitted := strings.Split(row, " ")

		mapped := utils.Map(splitted, func(num string, i int, splitted []string) string {
			number, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}

			base := make([]string, number)
			mapped := utils.Map(base, func(el string, j int, base []string) string {
				if i%2 != 0 {
					return "O"
				} else {
					return "."
				}
			})

			return strings.Join(mapped, "")
		})

		joined := strings.Join(mapped, "")

		return joined
	})

	valid := utils.Reduce(mapped, true, func(valid bool, row string, i int, mapped []string) bool {
		return valid && len(row) == len(mapped[0])
	})

	if valid {
		return mapped, nil
	} else {
		return []string{"INVALID"}, nil
	}
}
