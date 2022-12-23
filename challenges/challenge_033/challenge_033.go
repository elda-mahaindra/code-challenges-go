/*
   Offset Arrays

   To settle the debate of 0-based vs 1-based indexing I have created a language where you must explicitly state the range of indices an array should have.

   For example, given an array definition "A[-1..1] = 1 2 3", you would have:
   A[-1] = 1
   A[0] = 2
   A[1] = 3

   You are given a list of n array definitions and your job is to figure out what number is found in a given index i of an array arr.
   Note that the indexing operations may be nested (in the above example, A[A[-1]] would produce result 3).

   Input
       • n: an integer represents the number of array assignments.
       • assignments: an array of n strings where each string represents an assigment
         in the form of array_identifier[first_index..last_index] = last_index - first_index + 1.
       • element: a string that represents the element to print in the form of arr[i].

   Output
   a single integer.

   Constraints:
   • 1 ≤ n ≤ 100
   • Each array name consists of only uppercase letters (A to Z)
   • Array lengths are between 1 and 100 (no empty arrays)
   • Indexing operations have at most 50 levels of nesting
   • Indices are always within bounds in the test cases

   Example 1:
       Input:
         n = 3,
         assignments = [
           "A[-1..1] = 1 2 3",
           "B[3..7] = 3 4 5 6 7",
           "C[-2..1] = 1 2 3 4",
         ],
         element = "A[0]"
       Output: 2

   source: codingame
*/

package challenge_033

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"code-challenges-go/utils"
)

const (
	INVALID_ARRAY_ASSIGNMENTS = "the length of input 'assignments' should be equal to input 'n' and each assignment should be a valid assignment according to the constraints"
	INVALID_ELEMENT           = "the input 'element' should be a valid index operation and have at most 50 levels of nesting"
	OUT_OF_RANGE_N            = "the value of input 'n' should be between 1 and 100"
)

type TArrayAssignment struct {
	identifier            string
	firstIndex, lastIndex int
	values                []int
}

func isValid(n int, assignments []string, element string) error {
	switch {
	case n < 1 || n > 100:
		return errors.New(OUT_OF_RANGE_N)
	case len(assignments) != n ||
		!utils.Reduce(assignments, true, func(valid bool, assignment string, i int, assignments []string) bool {
			openSquareBracketIndex := strings.Index(assignment, "[")
			closeSquareBracketIndex := strings.Index(assignment, "]")
			connectorIndex := strings.Index(assignment, "..")

			firstIndex, err := strconv.Atoi(assignment[openSquareBracketIndex+1 : connectorIndex])
			if err != nil {
				panic(err)
			}

			lastIndex, err := strconv.Atoi(assignment[connectorIndex+2 : closeSquareBracketIndex])
			if err != nil {
				panic(err)
			}

			arrayLength := lastIndex - firstIndex + 1

			regex, err := regexp.Compile(`^[A-Z]+$`)
			if err != nil {
				panic(err)
			}

			validIdentifier := regex.MatchString(assignment[0:openSquareBracketIndex])

			validLength := arrayLength >= 1 && arrayLength <= 100

			return valid && validIdentifier && validLength
		}):
		return errors.New(INVALID_ARRAY_ASSIGNMENTS)
	case !func(element string) bool {
		openSquareBracketCount := 0
		closeSquareBracketCount := 0

		for i := 0; i < len(element); i++ {
			if string([]rune(element)[i]) == "[" {
				regex, err := regexp.Compile(`^[A-Z-\d]$`)
				if err != nil {
					panic(err)
				}

				if !regex.MatchString(string([]rune(element)[i+1])) {
					return false
				}

				openSquareBracketCount += 1
			}

			if string([]rune(element)[i]) == "]" {
				regex, err := regexp.Compile(`^[\]\d]$`)
				if err != nil {
					panic(err)
				}

				if !regex.MatchString(string([]rune(element)[i-1])) {
					return false
				}

				closeSquareBracketCount += 1
			}
		}

		return openSquareBracketCount == closeSquareBracketCount &&
			openSquareBracketCount >= 1 &&
			openSquareBracketCount <= 50
	}(element):
		return errors.New(INVALID_ELEMENT)
	default:
		return nil
	}
}

func getValue(arrayAssignment TArrayAssignment, index int) int {
	return arrayAssignment.values[index-arrayAssignment.firstIndex]
}

func Solution(n int, assignments []string, element string) (int, error) {
	err := isValid(n, assignments, element)
	if err != nil {
		return -1, err
	}

	arrayAssignments := utils.Map(assignments, func(assignment string, i int, assignments []string) TArrayAssignment {
		splitted := strings.Split(assignment, " = ")
		leftSide := splitted[0]
		rightSide := splitted[1]

		values := utils.Map(strings.Split(rightSide, " "), func(numberString string, i int, splittedRightSide []string) int {
			parsedInt, err := strconv.Atoi(numberString)
			if err != nil {
				panic(err)
			}

			return parsedInt
		})

		openSquareBracketIndex := strings.Index(leftSide, "[")
		closeSquareBracketIndex := strings.Index(leftSide, "]")
		connectorIndex := strings.Index(leftSide, "..")

		identifier := leftSide[0:openSquareBracketIndex]
		firstIndex, err := strconv.Atoi(leftSide[openSquareBracketIndex+1 : connectorIndex])
		if err != nil {
			panic(err)
		}
		lastIndex, err := strconv.Atoi(leftSide[connectorIndex+2 : closeSquareBracketIndex])
		if err != nil {
			panic(err)
		}

		return TArrayAssignment{
			identifier: identifier,
			firstIndex: firstIndex,
			lastIndex:  lastIndex,
			values:     values,
		}
	})

	stringValue := element
	for strings.Contains(stringValue, "[") {
		m := regexp.MustCompile(`[A-Z]+\[[-]?[\d]+\]`)
		loc := m.FindStringIndex(stringValue)
		identifierIndex := loc[0]

		m = regexp.MustCompile(`\[[-]?[\d]+\]`)
		loc = m.FindStringIndex(stringValue)
		openSquareBracketIndex := loc[0]

		closeSquareBracketIndex := strings.Index(stringValue, "]")

		identifier := stringValue[identifierIndex:openSquareBracketIndex]
		index, err := strconv.Atoi(stringValue[openSquareBracketIndex+1 : closeSquareBracketIndex])
		if err != nil {
			panic(err)
		}

		arrayAssignment := TArrayAssignment{}
		for i := 0; i < len(arrayAssignments); i++ {
			if arrayAssignments[i].identifier == identifier {
				arrayAssignment = arrayAssignments[i]
				break
			}
		}

		currentValue := getValue(arrayAssignment, index)

		stringValue = strings.Replace(stringValue, fmt.Sprintf("%s[%d]", identifier, index), strconv.Itoa(currentValue), -1)
	}

	parsedInt, err := strconv.Atoi(stringValue)
	if err != nil {
		return 0, err
	}

	return parsedInt, nil
}
