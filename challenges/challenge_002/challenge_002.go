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
)

const (
	INVALID_INPUT_DIMENSION = "the input 'board' should be a 9 x 9 sudoku board"
	INVALID_INPUT_RANGE     = "the input 'board' should contain only numbers between 1 to 9"
)

const boardSideLength = 9

func isValid(board [][]int) error {
	checkValidDimension := func(board [][]int, expectedXLength, expectedYLength int) bool {
		validXLength := true
		for i := 0; i < len(board); i++ {
			validXLength = validXLength && len(board[i]) == expectedXLength
		}

		validYLength := len(board) == expectedYLength

		return validXLength && validYLength
	}

	checkValidRange := func(board [][]int, start, end int) bool {
		if end <= start {
			return false
		}

		valid := true
		for j := 0; j < len(board); j++ {
			currentRow := board[j]

			for i := 0; i < len(currentRow); i++ {
				n := board[j][i]
				valid = valid && n >= start && n <= end
			}
		}

		return valid
	}

	switch {
	case !checkValidDimension(board, boardSideLength, boardSideLength):
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

	duplication := false
	for i := 0; i < len(numbersToBeChecked); i++ {
		if i == 0 {
			continue
		}

		duplication = duplication || numbersToBeChecked[i] == numbersToBeChecked[i-1]
	}

	return duplication
}

func transpose(matrix [][]int) [][]int {
	lengthX := len(matrix[0])
	lengthY := len(matrix)

	transposed := make([][]int, lengthX)
	for i := range transposed {
		transposed[i] = make([]int, lengthY)
	}

	for i := 0; i < lengthX; i++ {
		for j := 0; j < lengthY; j++ {
			transposed[i][j] = matrix[j][i]
		}
	}
	return transposed
}

// transform 9x9 board to 3x3 board represented by each row of the transformation result
func transform(board [][]int, boardSideLength, sideLength int) [][]int {
	transformed := make([][]int, len(board))

	for i := 0; i < boardSideLength/sideLength; i++ {
		sliced := board[3*i : 3+3*i]
		transposed := transpose(sliced)

		for j := 0; j < boardSideLength/sideLength; j++ {
			sliced := transposed[3*j : 3+3*j]
			// reduced is the 3x3 board
			reduced := func(sliced [][]int) []int {
				reduced := []int{}

				for k := 0; k < len(sliced); k++ {
					reduced = append(reduced, sliced[k]...)
				}

				return reduced
			}(sliced)

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

	duplicationInRow := func(board [][]int) bool {
		duplicationInRow := false

		for i := 0; i < len(board); i++ {
			duplicationInRow = duplicationInRow || checkDuplication(board[i])
		}

		return duplicationInRow
	}(board)

	duplicationInColumn := func(board [][]int) bool {
		duplicationInColumn := false

		for i := 0; i < len(board); i++ {
			duplicationInRow = duplicationInRow || checkDuplication(board[i])
		}

		return duplicationInColumn
	}(transpose(board))

	duplicationInSmallBoard := func(board [][]int) bool {
		duplicationInSmallBoard := false

		for i := 0; i < len(board); i++ {
			duplicationInRow = duplicationInRow || checkDuplication(board[i])
		}

		return duplicationInSmallBoard
	}(transform(board, boardSideLength, 3))

	if duplicationInRow || duplicationInColumn || duplicationInSmallBoard {
		return "not valid", nil
	}

	return "valid", nil
}
