/*
   Temperatures

   Analyze records of temperature to find the closest to zero.

   Rules
   Write a program that prints the temperature closest to 0 among input data.
   If two numbers are equally close to zero, positive integer has to be considered closest to zero (for instance, if the temperatures are -5 and 5, then display 5).

   Input
       • N:  the number of temperatures to analyze.
       • inputs: A string with the N temperatures expressed as integers ranging from -273 to 5526 separated by space

   Output
   Display 0 (zero) if no temperatures are provided. Otherwise, display the temperature closest to 0.

   Constraints:
       • 0 ≤ N < 10000

   Example 1:
       Input: N = 5, inputs = "1 -2 -8 4 5"
       Output: 1

   Example 2:
       Input: N = 0, inputs = ""
       Output: 0

   source: codingame
*/

package challenge_006

import (
	"errors"
	"math"
	"sort"
	"strconv"
	"strings"

	"code-challenges-go/utils"
)

const (
	INVALID_INPUTS = "the numbers of temperatures contained inside the input 'inputs' should be equal to the input 'N'"
	OUT_OF_RANGE_N = "the value of input 'N' should be between 0 and 9999"
)

func isValid(N int, inputs string) error {
	inputLength := len(strings.Split(inputs, " "))

	switch {
	case N < 0 || N >= 10000:
		return errors.New(OUT_OF_RANGE_N)
	case inputLength != N:
		return errors.New(INVALID_INPUTS)
	default:
		return nil
	}
}

func difference(a, b int) int {
	return int(math.Abs(float64(a - b)))
}

func sliceOfStringToSliceOfInt(strs []string) ([]int, error) {
	result := []int{}

	for _, str := range strs {
		number, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}

		result = append(result, number)
	}

	return result, nil
}

func Solution(N int, inputs string) (int, error) {
	err := isValid(N, inputs)
	if err != nil {
		return -274, err
	}

	splitted := strings.Split(inputs, " ")
	temps, err := sliceOfStringToSliceOfInt(splitted)
	if err != nil {
		return -274, nil
	}

	lowestTemps := utils.Reduce(temps, []int{}, func(lowestTemps []int, temp int, i int, temps []int) []int {
		if i == 0 {
			return append(lowestTemps, temp)
		}

		diff1 := difference(0, temp)
		diff2 := difference(0, lowestTemps[0])

		if diff1 > diff2 {
			return lowestTemps
		}

		if diff1 < diff2 {
			return []int{temp}
		}

		return append(lowestTemps, temp)
	})

	sort.SliceStable(lowestTemps, func(i, j int) bool {
		return lowestTemps[i] > lowestTemps[j]
	})

	return lowestTemps[0], nil
}
