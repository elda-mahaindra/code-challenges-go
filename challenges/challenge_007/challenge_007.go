/*
   Horse-racing Duals

   Casablanca’s hippodrome is organizing a new type of horse racing: duals.
   During a dual, only two horses will participate in the race. In order for the race to be interesting, it is necessary to try to select two horses with similar strength.

   Write a program which, using a given number of strengths, identifies the two closest strengths and shows their difference with an integer (≥ 0).

   Input
       • N:  Number N of horses.
       • powers: an array of numbers represents the strength Pi of each horse.

   Output
   The difference D between the two closest strengths. D is an integer greater than or equal to 0.

   Constraints:
       • 0 < N < 10000
       • 0 < Pi ≤ 10000000

   Example 1:
       Input: N = 3, powers = [5, 8, 9]
       Output: 1

   source: codingame
*/

package challenge_007

import (
	"errors"
	"math"
	"sort"
)

const (
	INVALID_POWERS = "the numbers of powers contained inside the input 'powers' should be equal to the input 'N'"
	OUT_OF_RANGE_N = "the value of input 'N' should be between 1 and 9999"
	OUT_OF_RANGE_P = "each value 'Pi' inside the value of input 'powers' should be between 1 and 9999999"
)

func isValid(N int, powers []int) error {
	switch {
	case N < 0 || N >= 10000:
		return errors.New(OUT_OF_RANGE_N)
	case len(powers) != N:
		return errors.New(INVALID_POWERS)
	case func(powers []int) bool {
		invalid := false

		for _, power := range powers {
			invalid = invalid || power <= 0 || power > 10000000
		}

		return invalid
	}(powers):
		return errors.New(OUT_OF_RANGE_P)
	default:
		return nil
	}
}

func difference(a, b int) int {
	return int(math.Abs(float64(a - b)))
}

func Solution(N int, powers []int) (int, error) {
	err := isValid(N, powers)
	if err != nil {
		return -1, err
	}

	sort.Ints(powers)

	closestDiff := func(powers []int) int {
		closestDiff := 10000000

		for i, power := range powers {
			if i == 0 {
				continue
			}

			diff := difference(power, powers[i-1])

			if diff < closestDiff {
				closestDiff = diff
			}
		}

		return closestDiff
	}(powers)

	return closestDiff, nil
}
