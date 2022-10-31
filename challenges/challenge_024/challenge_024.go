/*
   1D Spreadsheet

   You are given a 1-dimensional spreadsheet. You are to resolve the formulae and give the value of all its cells.

   Each input cell's content is provided as an operation with two operands arg1 and arg2.

   There are 4 types of operations:
   • VALUE arg1 arg2: The cell's value is arg1, (arg2 is not used and will be "_" to aid parsing).
   • ADD arg1 arg2: The cell's value is arg1 + arg2.
   • SUB arg1 arg2: The cell's value is arg1 - arg2.
   • MULT arg1 arg2: The cell's value is arg1 × arg2.

   Arguments can be of two types:
   • Reference $ref:
       If an argument starts with a dollar sign, it is a interpreted as a reference and its value is equal to the value of the cell by that number ref, 0-indexed.
       For example, "$0" will have the value of the result of the first cell.
       Note that a cell can reference a cell after itself!
   • Value val:
       If an argument is a pure number, its value is val.
       For example: "3" will have the value 3.

   There won't be any cyclic references: a cell that reference itself or a cell that references it, directly or indirectly.

   Input
       • N: an integer N represents the number of cells.
       • operations: an array of strings where each string represents the operation type, arg1, and arg2 separated by a space.

   Output
   An array of N numbers where each string represents the value of each cell.

   Constraints:
   • 1 ≤ N ≤ 100
   • -10000 ≤ val ≤ 10000
   • $0 ≤ $ref ≤ $(N - 1)
   • val ∈ Z
   • ref ∈ N
   • there are no cyclic references

   Example 1:
       Input: N = 2, operations = ["VALUE 3 _", "ADD $0 4"]
       Output: ["3", "7"]

   source: codingame
*/

package challenge_024

import (
	"errors"
	"strconv"
	"strings"

	"code-challenges-go/utils"
)

const (
	INVALID_OPERATION = "the length of input 'operations' should equal to the input 'N'"
	OUT_OF_RANGE_N    = "the value of input 'N' should be between 1 and 100"
)

type TCell struct {
	op, arg1, arg2 string
	counted        bool
	value          int
}

func isValid(N int, operations []string) error {
	switch {
	case N < 1 || N > 100:
		return errors.New(OUT_OF_RANGE_N)
	case len(operations) != N:
		return errors.New(INVALID_OPERATION)
	default:
		return nil
	}
}

func count(operationType string, arg1, arg2 int) int {
	switch operationType {
	case "VALUE":
		{
			return arg1
		}
	case "ADD":
		{
			return arg1 + arg2
		}
	case "SUB":
		{
			return arg1 - arg2
		}
		// case "MULT"
	default:
		{
			return arg1 * arg2
		}
	}
}

func checkCounted(cells []TCell) bool {
	return utils.Reduce(cells, true, func(allCounted bool, cell TCell, i int, cells []TCell) bool {
		return allCounted && cell.counted
	})
}

func Solution(N int, operations []string) ([]int, error) {
	err := isValid(N, operations)
	if err != nil {
		return nil, err
	}

	cells := utils.Map(operations, func(operation string, i int, operations []string) TCell {
		splitted := strings.Split(operation, " ")
		op := splitted[0]
		arg1 := splitted[1]
		arg2 := splitted[2]

		counted := false
		value := 0

		if !strings.Contains(arg1, "$") && !strings.Contains(arg2, "$") {
			intParsedArg1, err := strconv.Atoi(arg1)
			if err != nil {
				panic(err)
			}

			intParsedArg2 := 0
			if op != "VALUE" {
				parsed, err := strconv.Atoi(arg2)
				if err != nil {
					panic(err)
				}

				intParsedArg2 = parsed
			}

			value = count(op, intParsedArg1, intParsedArg2)
			counted = true
		}

		return TCell{
			op:      op,
			arg1:    arg1,
			arg2:    arg2,
			counted: counted,
			value:   value,
		}
	})

	allCounted := checkCounted(cells)

	for !allCounted {
		cells = utils.Map(cells, func(cell TCell, i int, cells []TCell) TCell {
			op := cell.op
			arg1 := cell.arg1
			arg2 := cell.arg2
			counted := cell.counted

			if !counted {
				intParsedArg1 := 0
				intParsedArg2 := 0

				if strings.Contains(arg1, "$") {
					intParsed, err := strconv.Atoi(arg1[1:])
					if err != nil {
						panic(err)
					}
					cellRef1 := cells[intParsed]

					if !cellRef1.counted {
						return cell
					}

					intParsedArg1 = cellRef1.value
				} else {
					intParsed, err := strconv.Atoi(arg1)
					if err != nil {
						panic(err)
					}

					intParsedArg1 = intParsed
				}

				if strings.Contains(arg2, "$") {
					intParsed := 0
					if op != "VALUE" {
						parsed, err := strconv.Atoi(arg2[1:])
						if err != nil {
							panic(err)
						}

						intParsed = parsed
					}
					cellRef2 := cells[intParsed]

					if !cellRef2.counted {
						return cell
					}

					intParsedArg2 = cellRef2.value
				} else {
					intParsed := 0
					if op != "VALUE" {
						parsed, err := strconv.Atoi(arg2)
						if err != nil {
							panic(err)
						}

						intParsed = parsed
					}

					intParsedArg2 = intParsed
				}

				return TCell{
					op:      op,
					arg1:    arg1,
					arg2:    arg2,
					counted: true,
					value:   count(op, intParsedArg1, intParsedArg2),
				}
			}

			return cell
		})

		allCounted = checkCounted(cells)
	}

	mapped := utils.Map(cells, func(cell TCell, i int, cells []TCell) int {
		return cell.value
	})

	return mapped, nil
}
