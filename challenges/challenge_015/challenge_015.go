/*
   Stock Exchange Losses

   A finance company is carrying out a study on the worst stock investments and would like to acquire a program to do so.
   The program must be able to analyze a chronological series of stock values in order to show the largest loss that it is possible to
   make by buying a share at a given time t0 and by selling it at a later date t1.
   The loss will be expressed as the difference in value between t0 and t1. If there is no loss, the loss will be worth 0.

   Input
       • n: the number n of stock values available.
       • values: an array of numbers represents the stock values arranged in order, from the date of their introduction on the stock market v1 until the last known value vn.

   Output
   The maximal loss p, expressed negatively if there is a loss, otherwise 0.

   Constraints:
       • 0 < n < 100000
       • 0 < v < 2^31

   Example 1:
       Input: n = 6, values = "3 2 4 2 1 5"
       Output: -3, because the negative movement from v2 (4) to v4 (1) is the maximum loss for this case.

   Example 2:
       Input: n = 6, values = "5 3 4 2 3 1"
       Output: -4, because the negative movement from v0 (5) to v5 (1) is the maximum loss for this case.

   Example 3:
       Input: n = 5, values = "1 2 4 4 5"
       Output: 0, because there are no negative movement for this case.

   source: codingame
*/

package challenge_015

import (
	"errors"
	"math"

	"code-challenges-go/utils"
)

const (
	OUT_OF_RANGE_N     = "the value of input 'n' should be between 1 and 99999"
	OUT_OF_RANGE_VALUE = "each value 'vi' inside the value of input 'values' should be between 1 and 2^31"
)

func isValid(n int, values []int) error {
	switch {
	case n <= 0 || n >= 100000:
		return errors.New(OUT_OF_RANGE_N)
	// ensure each vi inside values following the constraint
	case utils.Reduce(values, false, func(invalid bool, value int, i int, values []int) bool {
		return invalid || value <= 0 || value >= int(math.Pow(2, 31))
	}):
		return errors.New(OUT_OF_RANGE_VALUE)
	default:
		return nil
	}
}

func Solution(n int, values []int) (int, error) {
	err := isValid(n, values)
	if err != nil {
		return 1, err
	}

	currentHighestValue := values[0]
	largestLoss := 0

	for i := 1; i < len(values); i++ {
		if values[i] == values[i-1] {
			continue
		}

		if values[i] > currentHighestValue {
			currentHighestValue = values[i]
			continue
		}

		if values[i] < values[i-1] {
			loss := values[i] - currentHighestValue

			if loss < largestLoss {
				largestLoss = loss
			}
		}
	}

	return largestLoss, nil
}
