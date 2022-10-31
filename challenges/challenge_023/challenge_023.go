/*
   Bijective Numeration

   Decimal isn't the only base ten positional numbering system!

   Write a simple calculator that operates entirely in "decimary" notation to sum a list of positive integer values.
   Decimary is a fun name for what mathematics calls bijective numeration base ten, which uses A to represent ten and lacks a zero digit.

   The system is better illustrated than explained. Starting from one, these numerals are written:

   1, 2, 3, 4, 5, 6, 7, 8, 9, A = ten, 11, 12, 13, 14, 15, 16, 17, 18, 19, 1A = ten plus ten = twenty, 21, 22, 23, 24, 25, 26, 27, 28, 29, 2A = thirty, 31, ...

   88, 89, 8A = ninety, 91, 92, 93, 94, 95, 96, 97, 98, 99, 9A = ninety plus ten = one hundred, A1 = a hundred and one, A2, A3, A4, A5, A6, A7, A8, A9, AA = a hundred and ten, 111, ...

   199, 19A = two hundred, 1A1 = two hundred and one, ...

   Input
       • count: an integer count represents the number of decimaries.
       • decimaryInputs: a string that consist of 'count' decimary values separated by a space.

   Output
   A string in "decimary" notation representing the summation.

   Constraints:
   • 2 ≤ count < 10
   • 1 ≤ length of each decimary representation < 10

   Example 1:
       Input: count = 3, message = "A A 12",
       Output: "32"

   source: codingame
*/

package challenge_023

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"code-challenges-go/utils"
)

const (
	INVALID_DECIMARY   = "the number of decimaries inside input 'decimaries' should equal to the input 'count' and each decimary should be a valid decimary according to the constraints"
	OUT_OF_RANGE_COUNT = "the value of input 'count' should be between 2 and 9"
)

func isValid(count int, decimaryInputs string) error {
	decimaries := strings.Split(decimaryInputs, " ")

	switch {
	case count < 2 || count > 9:
		return errors.New(OUT_OF_RANGE_COUNT)
	case len(decimaries) != count || !utils.Reduce(decimaries, true, func(valid bool, decimary string, i int, decimaries []string) bool {
		return valid && len(decimary) >= 1 && len(decimary) < 10
	}):
		return errors.New(INVALID_DECIMARY)
	default:
		return nil
	}
}

/*
  encode and decode flows
  10  = A = (10 * 10^0) = 10
    => 1 0, proceed from right to left
      =>  step 1:
          digit is 0
          take 1 from next digit on the left, it become 10
          return 10 as A
      =>  step 2:
          digit is 1
          taken 1 by previous digit on the right, it become 1 - 1 = 0
          no other digits on the left so return 0
      =>  result:
          0A, the leading zero should be ignored
          A

  20  = 1A  = ( 1 * 10^1) + (10 * 10^0)  = 10 + 10  = 20
    => 2 0, proceed from right to left
      =>  step 1:
          digit is 0
          take 1 from next digit on the left, it become 10
          return 10 as A
      =>  step 2:
          digit is 2
          taken 1 by previous digit on the right, it become 2 - 1 = 1
          return 1
      =>  result:
          1A

  100 = 9A  = ( 9 * 10^1) + (10 * 10^0)  = 90 + 10  = 100
    => 1 0 0, proceed from right to left
      =>  step 1:
          digit is 0
          take 1 from next digit on the left, it become 10
          return 10 as A
      =>  step 2:
          digit is 0
          take 1 from next digit on the left, it become 10
          taken by previous digit on the right, it become 10 - 1 = 9
          return 9
      =>  result:
          9A

  101 = A1  = (10 * 10^1) + ( 1 * 10^0)  = 100 + 1  = 101
    => 1 0 1, proceed from right to left
      =>  step 1:
          digit is 1
          return 1
      =>  step 2:
          digit is 0
          take 1 from next digit on the left, it become 10
          return 10 as A
      =>  digit is 1
          taken 1 by previous digit on the right, it become 1 - 1 = 0
          no other digits on the left so return 0
      =>  result:
          0A1, the leading zero should be ignored
          A1

  110 = AA  = (10 * 10^1) + (10 * 10^0)  = 100 + 10 = 110

  120 = 11A = (1 * 10^2) + (1 * 10^1) + (10 * 10^0)  = 100 + 10 + 10 = 120

  130 = 12A = (1 * 10^2) + (2 * 10^1) + (10 * 10^0)  = 100 + 20 + 10 = 130

  200 = 19A = (1 * 10^2) + (9 * 10^1) + (10 * 10^0)  = 100 + 90 + 10 = 200

  201 = 1A1 = (1 * 10^2) + (10 * 10^1) + (1 * 10^0) = 100 + 100 + 1 = 201

  205 = 1A5 = (1 * 10^2) + (10 * 10^1) + (5 * 10^0) = 100 + 100 + 5 = 205

  210 = 1AA = (1 * 10^2) + (10 * 10^1) + (10 * 10^0) = 100 + 100 + 10 = 210

  860 = 85A - (8 * 10^2) + (5 * 10^1) + (10 * 10^0) = 800 + 50 + 10 = 860
    => 8 6 0, proceed from right to left
      =>  step 1:
          digit is 0
          take 1 from next digit on the left, it become 10
          return 10 as A
      =>  step 2:
          digit is 6
          taken by previous digit on the right, it become 6 - 1 = 5
          return 5
      =>  step 3:
          digit is 8
          return 8
      =>  result:
          85A

  8060 = 7A5A = (7 * 10^3) + (10 * 10^2) + (5 * 10^1) + (10 * 10^0) = 7000 +1000 + 50 + 10 = 8060

  8010 = 79AA = (7 * 10^3) + (9 * 10^2) + (10 * 10^1) + (10 * 10^0) = 7000 + 900 + 100 + 10 = 8010
    => 8 0 1 0, proceed from right to left
      =>  step 1:
          digit is 0
          take 1 from next digit on the left, it become 10
          return 10 as A
      =>  step 2:
          digit is 1
          taken by previous digit on the right, it become 1 - 1 = 0
          take 1 from next digit on the left, it become 10
          return 10 as A
      =>  step 3:
          digit is 0
          take 1 from next digit on the left, it become 10
          taken by previous digit on the right, it become 10 - 1 = 9
          return 9
      =>  step 4:
          digit is 8
          taken by previous digit on the right, it become 8 - 1 = 7
          return 7
      =>  result:
          79AA
*/

func decimaryToDecimal(decimary string) int {
	splitted := strings.Split(decimary, "")

	decimal := utils.Reduce(splitted, 0, func(decimal int, letter string, i int, splitted []string) int {
		maxPower := len(decimary) - 1
		power := math.Abs(float64(i - maxPower))

		if letter == "A" {
			return decimal + int(10*math.Pow(10, power))
		} else {
			parsedInt, err := strconv.Atoi(letter)
			if err != nil {
				panic(err)
			}

			return decimal + int(float64(parsedInt)*math.Pow(10, power))
		}
	})

	return decimal
}

func decimalToDecimary(decimal int) string {
	splitted := strings.Split(strconv.Itoa(decimal), "")
	hasZero := utils.Includes(splitted, "0")

	if !hasZero {
		return strconv.Itoa(decimal)
	}

	decimary := ""

	remainder := decimal
	for remainder > 0 {
		splitted := strings.Split(strconv.Itoa(remainder), "")
		lastIndex := len(splitted) - 1
		lastDigit := splitted[lastIndex]

		if lastDigit != "0" {
			decimary = fmt.Sprintf("%s%s", lastDigit, decimary)
			splitted = splitted[0:lastIndex]

			if len(splitted) > 0 {
				parsedInt, err := strconv.Atoi(strings.Join(splitted, ""))
				if err != nil {
					panic(err)
				}

				remainder = parsedInt
			} else {
				remainder = 0
			}
		} else {
			decimary = fmt.Sprintf("%s%s", "A", decimary)
			splitted = splitted[0:lastIndex]

			if len(splitted) > 0 {
				parsedInt, err := strconv.Atoi(strings.Join(splitted, ""))
				if err != nil {
					panic(err)
				}

				remainder = parsedInt - 1
			} else {
				remainder = 0
			}
		}
	}

	return decimary
}

func Solution(count int, decimaryInputs string) (string, error) {
	err := isValid(count, decimaryInputs)
	if err != nil {
		return "", err
	}

	decimaries := strings.Split(decimaryInputs, " ")

	decimals := utils.Map(decimaries, func(decimary string, i int, decimaries []string) int {
		return decimaryToDecimal(decimary)
	})

	sum := utils.Reduce(decimals, 0, func(sum, decimal, i int, decimals []int) int {
		return sum + decimal
	})

	decimary := decimalToDecimary(sum)

	return decimary, nil
}
