/*
   Happy Numbers

   A happy number is defined by the following process:
   Starting with any positive integer, replace the number by the sum of the squares of its digits in base-ten.
   Repeat the process until the number either equals 1 (where it will stay), or it loops endlessly in a cycle that does not include 1.
   Those numbers for which this process ends in 1 are happy numbers, while those that do not end in 1 are unhappy numbers.

   Given a list of numbers, classify each of them as happy or unhappy.

   Input
       • N: an integer that represents the number of numbers to test.
       • strNums: an array of string where each string represents the positive integer to test.

   Output
   An array of N strings where each string consists of two parts separated by a space which are the number which has been tested and an ascii art of :) or :( to indicates wheter the number is a happy or unhappy number.

   Constraints:
   • 1 ≤ N ≤ 100
   • 0 < strNums ≤ 10^26

   Example 1:
       Input: N = 2, strNums = ["23", "24"]
       Output: ["23 :)", "24 :("]

   source: codingame
*/

package challenge_027

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"code-challenges-go/utils"
)

const (
	INVALID_NUM    = "the length of input 'nums' should be equal to input 'N' and each element inside it should represent an integer value between 1 and 10^26"
	OUT_OF_RANGE_N = "the value of input 'N' should be between 1 and 100"
)

func isValid(N int, strNums []string) error {
	switch {
	case N < 1 || N > 100:
		return errors.New(OUT_OF_RANGE_N)
	case len(strNums) != N || !utils.Reduce(strNums, true, func(valid bool, strNum string, i int, strNums []string) bool {
		splitted := strings.Split(strNum, "")

		if len(strNum) == 1 && strNum == "0" {
			return valid && false
		}

		noNegativeSign := !utils.Includes(splitted, "-")
		if !noNegativeSign {
			return valid && false
		}

		parsable := utils.Reduce(splitted, true, func(parsable bool, digit string, i int, splitted []string) bool {
			_, err := strconv.Atoi(digit)
			if err != nil {
				return parsable && false
			}

			return parsable && true
		})
		if !parsable {
			return valid && false
		}

		return valid && len(strNum) <= 27
	}):
		return errors.New(INVALID_NUM)
	default:
		return nil
	}
}

func checkBaseTenDuplication(strNums []string, a string) bool {
	if !(len(strNums) > 0) {
		return false
	}

	duplication := utils.Reduce(strNums, false, func(duplication bool, b string, i int, strNums []string) bool {
		aStrings := strings.Split(a, "")
		sort.Strings(aStrings)

		bStrings := strings.Split(b, "")
		sort.Strings(bStrings)

		if len(aStrings) != len(bStrings) {
			return duplication || false
		}

		return duplication || utils.Reduce(aStrings, true, func(duplication bool, _ string, i int, aStrings []string) bool {
			return duplication && aStrings[i] == bStrings[i]
		})
	})

	return duplication
}

func transform(strNums string) string {
	splitted := strings.Split(strNums, "")

	reduced := utils.Reduce(splitted, 0, func(reduced int, baseTen string, i int, splitted []string) int {
		parsedInt, err := strconv.Atoi(baseTen)
		if err != nil {
			panic(err)
		}

		return reduced + int(math.Pow(float64(parsedInt), 2))
	})

	return strconv.Itoa(reduced)
}

func Solution(N int, strNums []string) ([]string, error) {
	err := isValid(N, strNums)
	if err != nil {
		return nil, err
	}

	mapped := utils.Map(strNums, func(strNum string, i int, strNums []string) string {
		transformResults := []string{strNum}

		foundHappy := false
		foundDuplication := false

		for !foundHappy && !foundDuplication {
			transformed := transform(transformResults[len(transformResults)-1])

			foundHappy = transformed == "1"
			foundDuplication = checkBaseTenDuplication(transformResults, transformed)

			transformResults = append(transformResults, transformed)
		}

		if foundHappy {
			return fmt.Sprintf("%s :)", strNum)
		} else {
			return fmt.Sprintf("%s :(", strNum)
		}
	})

	return mapped, nil
}
