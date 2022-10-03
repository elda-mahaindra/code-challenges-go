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
)

const (
	ALPHABET_NOT_FOUND string = "the string input 's' should contain at least one alphabet letter"
	OUT_OF_RANGE_S     string = "the length of input 's' should be between 1 and 100"
)

type status struct {
	char      string
	occurence int32
}

func isValid(s string) error {
	switch {
	case len(s) < 1 || len(s) > 100:
		return errors.New(OUT_OF_RANGE_S)
	case !func(s string) bool {
		splitted := strings.Split(s, "")

		containAlphabet := false
		for _, letter := range splitted {
			regex, err := regexp.Compile("[a-zA-Z]")
			if err != nil {
				containAlphabet = containAlphabet && false
			}

			containAlphabet = containAlphabet || regex.MatchString(letter)
		}

		return containAlphabet
	}(s):
		return errors.New(ALPHABET_NOT_FOUND)
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

	statuses := []status{}
	for i := 0; i < len(splitted); i++ {
		regex, err := regexp.Compile("[a-zA-Z]")
		if err != nil {
			continue
		}

		if !regex.MatchString(splitted[i]) {
			continue
		}

		if len(statuses) == 0 {
			statuses = append(statuses, status{
				char:      splitted[i],
				occurence: 1,
			})
		} else {
			latest := statuses[len(statuses)-1]
			if splitted[i] == latest.char {
				statuses = statuses[0 : len(statuses)-1]
				statuses = append(statuses, status{
					char:      splitted[i],
					occurence: latest.occurence + 1,
				})
			} else {
				statuses = append(statuses, status{
					char:      splitted[i],
					occurence: 1,
				})
			}
		}
	}

	sort.SliceStable(statuses, func(i, j int) bool {
		return statuses[i].occurence > statuses[j].occurence
	})

	return statuses[0].char, nil
}
