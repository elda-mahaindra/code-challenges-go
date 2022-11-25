/*
   Valid Sudoku

   Determine if a 9 x 9 Sudoku board is valid according to the following rules:
     • Each row must contain the digits 1-9 without repetition.
     • Each column must contain the digits 1-9 without repetition.
     • Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.

   Example 1:
     Input: board = [
           [7, 9, 2, 1, 5, 4, 3, 8, 6],
           [6, 4, 3, 8, 2, 7, 1, 5, 9],
           [8, 5, 1, 3, 9, 6, 7, 2, 4],
           [2, 6, 5, 9, 7, 3, 8, 4, 1],
           [4, 8, 9, 5, 6, 1, 2, 7, 3],
           [3, 1, 7, 4, 8, 2, 9, 6, 5],
           [1, 3, 6, 7, 4, 8, 5, 9, 2],
           [9, 7, 4, 2, 1, 5, 6, 3, 8],
           [5, 2, 8, 6, 3, 9, 4, 1, 7]
     ]
     Output: 'valid'

   Example 2:
     Input: board = [
           [5, 5, 5, 5, 5, 5, 5, 5, 5],
           [5, 5, 5, 5, 5, 5, 5, 5, 5],
           [5, 5, 5, 5, 5, 5, 5, 5, 5],
           [5, 5, 5, 5, 5, 5, 5, 5, 5],
           [5, 5, 5, 5, 5, 5, 5, 5, 5],
           [5, 5, 5, 5, 5, 5, 5, 5, 5],
           [5, 5, 5, 5, 5, 5, 5, 5, 5],
           [5, 5, 5, 5, 5, 5, 5, 5, 5],
           [5, 5, 5, 5, 5, 5, 5, 5, 5]
     ]
     Output: 'not valid'

   source: codr
*/

package challenge_002

import (
	"errors"
	"sort"

	"code-challenges-go/utils"
)

const (
	INVALID_INPUT_DIMENSION = "the input 'board' should be a 9 x 9 sudoku board"
	INVALID_INPUT_RANGE     = "the input 'board' should contain only numbers between 1 to 9"
)

func isValid(board [][]int) error {
	checkValidDimension := func(board [][]int, expectedXLength, expectedYLength int) bool {
		validXLength := utils.Reduce(board, true, func(valid bool, row []int, i int, board [][]int) bool {
			return valid && len(row) == expectedXLength
		})

		validYLength := len(board) == expectedYLength

		return validXLength && validYLength
	}

	checkValidRange := func(board [][]int, start, end int) bool {
		if end <= start {
			return false
		}

		valid := utils.Reduce(board, true, func(valid bool, row []int, i int, board [][]int) bool {
			validRow := utils.Reduce(row, true, func(valid bool, n int, i int, row []int) bool {
				return valid && n >= start && n <= end
			})

			return valid && validRow
		})

		return valid
	}

	switch {
	case !checkValidDimension(board, 9, 9):
		return errors.New(INVALID_INPUT_DIMENSION)
	case !checkValidRange(board, 1, 9):
		return errors.New(INVALID_INPUT_RANGE)
	default:
		return nil
	}
}

func checkDuplication(numbers []int) bool {
	numbersToBeChecked := []int{}
	numbersToBeChecked = append(numbersToBeChecked, numbers...)

	sort.Ints(numbersToBeChecked)

	duplication := utils.Reduce(numbersToBeChecked, false, func(duplication bool, n int, i int, numbersToBeChecked []int) bool {
		if i == 0 {
			return duplication
		}

		return duplication || n == numbersToBeChecked[i-1]
	})

	return duplication
}

func transpose(matrix [][]int) [][]int {
	return utils.Map(matrix[0], func(n int, i int, slice []int) []int {
		return utils.Map(matrix, func(row []int, j int, slice [][]int) int {
			return row[i]
		})
	})
}

// transform 9x9 board to 3x3 board represented by each row of the transformation result
func transform(board [][]int) [][]int {
	transformed := make([][]int, len(board))

	for i := 0; i < 3; i++ {
		sliced := board[3*i : 3+3*i]
		transposed := transpose(sliced)

		for j := 0; j < 3; j++ {
			sliced := transposed[3*j : 3+3*j]
			// reduced is the 3x3 board
			reduced := utils.Reduce(sliced, []int{}, func(reduced []int, numbers []int, i int, sliced [][]int) []int {
				return append(reduced, numbers...)
			})

			sort.Ints(reduced)

			transformed[i*3+j] = reduced
		}
	}

	return transformed
}

func Solution(board [][]int) (string, error) {
	err := isValid(board)
	if err != nil {
		return "", err
	}

	duplicationInRow := utils.Reduce(board, false, func(duplicationInRow bool, row []int, i int, board [][]int) bool {
		return duplicationInRow || checkDuplication(row)
	})

	duplicationInColumn := utils.Reduce(transpose(board), false, func(duplicationInColumn bool, column []int, i int, transposedBoard [][]int) bool {
		return duplicationInColumn || checkDuplication(column)
	})

	duplicationInSmallBoard := utils.Reduce(transform(board), false, func(duplicationInSmallBoard bool, smallBoard []int, i int, transformedBoard [][]int) bool {
		return duplicationInSmallBoard || checkDuplication(smallBoard)
	})

	if duplicationInRow || duplicationInColumn || duplicationInSmallBoard {
		return "not valid", nil
	}

	return "valid", nil
}
