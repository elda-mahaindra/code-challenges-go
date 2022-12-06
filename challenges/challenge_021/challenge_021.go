/*
   Substitution Encoding

   You want to easily encode and decode messages with a simple and personalized method. To do so, you will use a substitution method.

   The principle is simple, you have a comparison table like this one:

   A B
   C D

   and a message to encode written with the characters available in your table:

   CBA

   You are going to take each of the characters that compose the message and replace them by its position in the table:

   C => 10 (row 1 column 0)
   B => 01 (row 0 column 1)
   A => 00 (row 0 column 0)

   The message becomes: 100100

   Input
       • n: an integer n represents the number of rows in comparison table.
       • rows: an array of n string which each string represents the row of comparison table and composed of characters separated by a space.
       • message: a string message to encode.

   Output
   The encoded message.

   Constraints:
   • 1 ≤ n ≤ 10

   Example 1:
       Input:
           n = 2,
           rows = [
               "A B",
               "C D",
           ],
           message = "CBA",
       Output: "100100"

   source: codingame
*/

package challenge_021

import (
	"errors"
	"fmt"
	"strings"

	"code-challenges-go/utils"
)

const (
	INVALID_ROWS   = "the length of input 'rows' should equal to the input 'n' and each row should be a valid row which composed by some single characters separated by a space"
	OUT_OF_RANGE_N = "the value of input 'N' should be between 1 and 10"
)

type TCharMap struct {
	char     string
	position []int
}

func isValid(n int, rows []string) error {
	switch {
	case n < 1 || n > 10:
		return errors.New(OUT_OF_RANGE_N)
	case len(rows) != n || !utils.Reduce(rows, true, func(valid bool, row string, i int, rows []string) bool {
		splitted := strings.Split(row, " ")

		return valid && utils.Reduce(splitted, true, func(valid bool, char string, i int, splitted []string) bool {
			return valid && len(char) == 1
		})
	}):
		return errors.New(INVALID_ROWS)
	default:
		return nil
	}
}

func Solution(n int, rows []string, message string) (string, error) {
	err := isValid(n, rows)
	if err != nil {
		return "", err
	}

	charMaps := utils.Reduce(rows, []TCharMap{}, func(charMaps []TCharMap, row string, y int, rows []string) []TCharMap {
		splitted := strings.Split(row, " ")

		maps := []TCharMap{}
		for x := 0; x < len(splitted); x++ {
			maps = append(maps, TCharMap{char: splitted[x], position: []int{x, y}})
		}

		return append(charMaps, maps...)
	})

	encoded := strings.Join(utils.Map(strings.Split(message, ""), func(letter string, i int, letters []string) string {
		found := TCharMap{}
		for _, charMap := range charMaps {
			if charMap.char == letter {
				found = charMap
			}
		}

		x := found.position[0]
		y := found.position[1]

		return fmt.Sprintf("%d%d", y, x)
	}), "")

	return encoded, nil
}
