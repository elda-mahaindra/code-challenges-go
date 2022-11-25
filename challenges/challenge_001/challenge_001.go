/*
   Most Repeated Character

   Given a string 's', return the most frequent character (an alphabet letter) inside string 's'.

   Constraints:
     • 1 <= s.length <= 100
     • s contains lower-case and upper-case English letters and numbers

   Example 1:
     Input: s = "abcddefda1111133333333"
     Output: 'd'

   Example 2:
     Input: s= "AA0AB0BB0ccc0aa0aw00wo0BBBw123123"
     Output: 'B'

   source: turing
*/

package challenge_001

import (
	"errors"
	"regexp"
	"sort"
	"strings"

	"code-challenges-go/utils"
)

const (
	INVALID_S      string = "the string input 's' should contain at least one alphabet letter"
	OUT_OF_RANGE_S string = "the length of input 's' should be between 1 and 100"
)

type status struct {
	char      string
	occurence int32
}

func isValid(s string) error {
	splitted := strings.Split(s, "")

	switch {
	case len(s) < 1 || len(s) > 100:
		return errors.New(OUT_OF_RANGE_S)
	// ensure there is at least one alphabet letter in the input string
	case !utils.Reduce(splitted, false, func(valid bool, letter string, i int, splitted []string) bool {
		regex, err := regexp.Compile("[a-zA-Z]")
		if err != nil {
			valid = valid && false
		}

		valid = valid || regex.MatchString(letter)

		return valid
	}):
		return errors.New(INVALID_S)
	default:
		return nil
	}
}

func Solution(s string) (string, error) {
	err := isValid(s)
	if err != nil {
		return "", err
	}

	splitted := strings.Split(s, "")
	sort.Strings(splitted)

	regex, err := regexp.Compile("[a-zA-Z]")
	if err != nil {
		return "", err
	}

	statuses := utils.Reduce(splitted, []status{}, func(statuses []status, v string, i int, splitted []string) []status {
		if !regex.MatchString(splitted[i]) {
			return statuses
		}

		if len(statuses) == 0 {
			return append(statuses, status{
				char:      splitted[i],
				occurence: 1,
			})
		} else {
			latest := statuses[len(statuses)-1]
			if splitted[i] == latest.char {
				statuses = statuses[0 : len(statuses)-1]
				return append(statuses, status{
					char:      splitted[i],
					occurence: latest.occurence + 1,
				})
			} else {
				return append(statuses, status{
					char:      splitted[i],
					occurence: 1,
				})
			}
		}
	})

	sort.SliceStable(statuses, func(i, j int) bool {
		return statuses[i].occurence > statuses[j].occurence
	})

	return statuses[0].char, nil
}
