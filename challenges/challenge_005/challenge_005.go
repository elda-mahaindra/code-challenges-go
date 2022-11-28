/*
   ASCII Art

   ASCII art allows you to represent forms by using characters. To be precise, in our case, these forms are words.
   For example, the word "MANHATTAN" could be displayed as follows in ASCII art:

               # #  #  ### # #  #  ### ###  #  ###
               ### # # # # # # # #  #   #  # # # #
               ### ### # # ### ###  #   #  ### # #
               # # # # # # # # # #  #   #  # # # #
               # # # # # # # # # #  #   #  # # # #

    Your mission is to write a program that can display a line of text in ASCII art in a style you are given as input.

   Input
       • T: The line of text to be displayed, composed of N ASCII characters.

   Constant
       • L: the width of a letter represented in ASCII art. All letters are the same width.
           const L = 4
       • H: the height of a letter represented in ASCII art. All letters are the same height.
           const H = 5
       • rows: array of string of characters ABCDEFGHIJKLMNOPQRSTUVWXYZ? Represented in ASCII art.
           const rows = [
                   " #  ##   ## ##  ### ###  ## # # ###  ## # # #   # # ###  #  ##   #  ##   ## ### # # # # # # # # # # ### ### ",
                   "# # # # #   # # #   #   #   # #  #    # # # #   ### # # # # # # # # # # #    #  # # # # # # # # # #   #   # ",
                   "### ##  #   # # ##  ##  # # ###  #    # ##  #   ### # # # # ##  # # ##   #   #  # # # # ###  #   #   #   ## ",
                   "# # # # #   # # #   #   # # # #  #  # # # # #   # # # # # # #    ## # #   #  #  # # # # ### # #  #  #       ",
                   "# # ##   ## ##  ### #    ## # # ###  #  # # ### # # # #  #  #     # # # ##   #  ###  #  # # # #  #  ###  #  "
               ]

   Output
   The text T in ASCII art represented by an array of string.
   The characters a to z are shown in ASCII art by their equivalent in upper case.
   The characters that are not in the intervals [a-z] or [A-Z] will be shown as a question mark in ASCII art.

   Constraints:
       • 0 < N < 200

   Example 1:
       Input:
           L = 4, H = 5, T = E
       Output: [
           "### ",
           "#   ",
           "##  ",
           "#   ",
           "### ",
       ]

   Example 2:
       Input:
           L = 4, H = 5, T = MANHATTAN
       Output: [
           "# #  #  ### # #  #  ### ###  #  ### ",
           "### # # # # # # # #  #   #  # # # # ",
           "### ### # # ### ###  #   #  ### # # ",
           "# # # # # # # # # #  #   #  # # # # ",
           "# # # # # # # # # #  #   #  # # # # ",
       ]

   source: codingame
*/

package challenge_005

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"code-challenges-go/utils"
)

const (
	OUT_OF_RANGE_N = "the length of input 'T' should be between 1 and 199"
)

const L = 4
const H = 5

var rows = []string{
	" #  ##   ## ##  ### ###  ## # # ###  ## # # #   # # ###  #  ##   #  ##   ## ### # # # # # # # # # # ### ### ",
	"# # # # #   # # #   #   #   # #  #    # # # #   ### # # # # # # # # # # #    #  # # # # # # # # # #   #   # ",
	"### ##  #   # # ##  ##  # # ###  #    # ##  #   ### # # # # ##  # # ##   #   #  # # # # ###  #   #   #   ## ",
	"# # # # #   # # #   #   # # # #  #  # # # # #   # # # # # # #    ## # #   #  #  # # # # ### # #  #  #       ",
	"# # ##   ## ##  ### #    ## # # ###  #  # # ### # # # #  #  #     # # # ##   #  ###  #  # # # #  #  ###  #  ",
}

const charsAvailable = "ABCDEFGHIJKLMNOPQRSTUVWXYZ?"
const totalCharsAvailable = len(charsAvailable)

var reducedRows = utils.Reduce(rows, [][]string{}, func(reducedRows [][]string, row string, i int, rows []string) [][]string {
	reducedRow := []string{}

	for i := 0; i < totalCharsAvailable; i++ {
		sliced := row[i*L : i*L+L]

		reducedRow = append(reducedRow, sliced)
	}

	return append(reducedRows, reducedRow)
})

func isValid(T string) error {
	switch {
	case len(T) <= 0 || len(T) >= 200:
		return errors.New(OUT_OF_RANGE_N)
	default:
		return nil
	}
}

func Solution(T string) ([]string, error) {
	err := isValid(T)
	if err != nil {
		return nil, err
	}

	regex, err := regexp.Compile("[a-zA-Z]")
	if err != nil {
		return nil, err
	}

	inputs := func(T string) []string {
		upperCased := strings.ToUpper(T)
		splitted := strings.Split(upperCased, "")

		mapped := utils.Map(splitted, func(letter string, i int, splitted []string) string {
			if regex.MatchString(letter) {
				return letter
			} else {
				return "?"
			}
		})

		return mapped
	}(T)

	tIndexes := utils.Reduce(inputs, []int{}, func(tIndexes []int, tChar string, i int, inputs []string) []int {
		for i := 0; i < len(charsAvailable); i++ {
			char := string([]rune(charsAvailable)[i])

			if char == tChar {
				tIndexes = append(tIndexes, i)
			}
		}

		return tIndexes
	})

	result := []string{}
	for i := 0; i < H; i++ {
		tHArt := utils.Reduce(tIndexes, "", func(tHArt string, tIndex int, j int, tIndexes []int) string {
			tA := fmt.Sprintf("%s%s", tHArt, reducedRows[i][tIndex])

			return tA
		})

		result = append(result, tHArt)
	}

	return result, nil
}
