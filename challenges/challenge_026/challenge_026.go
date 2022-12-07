/*
   Insert to String

   You're given a string and a list of changes to do on that string.
   Each change includes a line number, a column number, and a string to add at that location.
   Commit all of the changes to the given string so they wouldn't interfere with each other.

   ("\n" should be replaced with a newline.)

   Input
       • s: a string to manipulate.
       • changeCount: an integer that represents the number of changes.
       • rawChanges: an array of strings where each string represents the change.

   Output
   an array of strings where each string represents the manipulated string in one row.

   Constraints:
   • the length of s ≥ 1
   • 1 ≤ changeCount ≤ 10
   • each string inside 'rawChanges' should be in the format of <Line number>|<Column number>|<String to be added>

   Example 1:
        Input:
          s = "Hello world",
          changeCount = 4,
          rawChanges = ["0|11|!", "0|5|,\\n", "0|7| w", "0|10|\\n"],
        Output: ["Hello,", " w worl", "d!"]

   source: codingame
*/

package challenge_026

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"code-challenges-go/utils"
)

const (
	INVALID_RAW_CHANGE        = "the length of input 'rawChanges' should equal to the input 'changeCount' and each change should be a valid change"
	OUT_OF_RANGE_CHANGE_COUNT = "the length of input 'changeCount' should be between 1 and 10"
	OUT_OF_RANGE_S            = "the length of input 's' should be more than 0"
)

type TChange struct {
	lineNumber, columnNumber int
	pasted                   string
}

func isValid(s string, changeCount int, rawChanges []string) error {
	switch {
	case len(s) < 1:
		return errors.New(OUT_OF_RANGE_S)
	case changeCount < 1 || changeCount > 10:
		return errors.New(OUT_OF_RANGE_CHANGE_COUNT)
	case len(rawChanges) != changeCount || !utils.Reduce(rawChanges, true, func(valid bool, rawChange string, i int, rawChanges []string) bool {
		splitted := strings.Split(rawChange, "|")

		if len(splitted) != 3 {
			return false
		}

		_, err := strconv.Atoi(splitted[0])
		if err != nil {
			return false
		}

		_, err = strconv.Atoi(splitted[1])
		if err != nil {
			return false
		}

		return valid
	}):
		return errors.New(INVALID_RAW_CHANGE)
	default:
		return nil
	}
}

func Solution(s string, changeCount int, rawChanges []string) ([]string, error) {
	err := isValid(s, changeCount, rawChanges)
	if err != nil {
		return nil, err
	}

	targets := strings.Split(s, "\\n")

	changes := utils.Map(rawChanges, func(rawChange string, i int, rawChanges []string) TChange {
		splitted := strings.Split(rawChange, "|")

		lineNumber, err := strconv.Atoi(splitted[0])
		if err != nil {
			panic(err)
		}

		columnNumber, err := strconv.Atoi(splitted[1])
		if err != nil {
			panic(err)
		}

		return TChange{
			lineNumber:   lineNumber,
			columnNumber: columnNumber,
			pasted:       splitted[2],
		}
	})

	sortedChanges := changes
	sort.SliceStable(sortedChanges, func(i, j int) bool {
		if sortedChanges[i].lineNumber == sortedChanges[j].lineNumber {
			return sortedChanges[i].columnNumber > sortedChanges[j].columnNumber
		} else {
			return sortedChanges[i].lineNumber > sortedChanges[j].lineNumber
		}
	})

	for _, change := range sortedChanges {
		lineNumber := change.lineNumber
		columnNumber := change.columnNumber
		pasted := change.pasted

		target := targets[lineNumber]

		targets[lineNumber] = fmt.Sprintf("%s%s%s", target[:columnNumber], pasted, target[columnNumber:])
	}

	reduced := utils.Reduce(targets, []string{}, func(reduced []string, target string, i int, targets []string) []string {
		splitted := strings.Split(target, "\\n")

		return append(reduced, splitted...)
	})

	return reduced, nil
}
