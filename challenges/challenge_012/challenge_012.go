/*
   Unary

   Binary with 0 and 1 is good, but binary with only 0, or almost, is even better!

   Write a program that takes an incoming message as input and displays as output the message encoded using this method.

   Rules
   Here is the encoding principle:
   • The input message consists of ASCII characters (7-bit)
   • The encoded output message consists of blocks of 0
   • A block is separated from another block by a space
   • Two consecutive blocks are used to produce a series of same value bits (only 1 or 0 values):
       - First block: it is always 0 or 00. If it is 0, then the series contains 1, if not, it contains 0
       - Second block: the number of 0 in this block is the number of bits in the series

   Let’s take a simple example with a message which consists of only one character: Capital C. C in binary is represented as 1000011, so with this method, this gives:
   • 0 0 (the first series consists of only a single 1)
   • 00 0000 (the second series consists of four 0)
   • 0 00 (the third consists of two 1)
   So C is coded as: 0 0 00 0000 0 00


   Second example, we want to encode the message CC (i.e. the 14 bits 10000111000011) :
   • 0 0 (one single 1)
   • 00 0000 (four 0)
   • 0 000 (three 1)
   • 00 0000 (four 0)
   • 0 00 (two 1)
   So CC is coded as: 0 0 00 0000 0 000 00 0000 0 00

   Input
       • message: a string message to be encoded consisting of N ASCII characters (without carriage return).

   Output
   The encoded message.

   Constraints:
       • 0 < N < 100

   Example 1:
       Input: message = "C"
       Output: "0 0 00 0000 0 00"

   Example 2:
       Input: message = "CC"
       Output: "0 0 00 0000 0 000 00 0000 0 00"

   source: codingame
*/

package challenge_012

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"code-challenges-go/utils"
)

const (
	OUT_OF_RANGE_MESSAGE = "the length of input 'message' should be between 1 and 99"
	INVALID_BINARY       = "expecting a binary string which each value represented by '0' or '1'"
)

type TStatus struct {
	char      string
	occurence int
}

func isValid(message string) error {
	switch {
	case len(message) <= 0 || len(message) >= 100:
		return errors.New(OUT_OF_RANGE_MESSAGE)
	default:
		return nil
	}
}

func textToBinary(text string) string {
	splitted := strings.Split(text, "")

	mapped := utils.Map(splitted, func(s string, i int, splitted []string) string {
		char := []rune(s)[0]
		ascii := int(char)
		binary := strconv.FormatInt(int64(ascii), 2)
		padded := fmt.Sprintf("%07s", binary)

		return padded
	})

	return strings.Join(mapped, "")
}

func binaryToUnary(binary string) (string, error) {
	regex, err := regexp.Compile(`^[0-1]{1,}`)
	if err != nil {
		return "", err
	}

	if !regex.MatchString(binary) {
		return "", errors.New(INVALID_BINARY)
	}

	splitted := strings.Split(binary, "")

	statuses := utils.Reduce(splitted, []TStatus{}, func(reduced []TStatus, char string, i int, splitted []string) []TStatus {
		if i == 0 {
			return []TStatus{{char: char, occurence: 1}}
		} else if char != reduced[len(reduced)-1].char {
			return append(reduced, TStatus{char: char, occurence: 1})
		} else {
			return utils.Map(reduced, func(stat TStatus, i int, statuses []TStatus) TStatus {
				if i == len(reduced)-1 {
					return TStatus{char: stat.char, occurence: stat.occurence + 1}
				}

				return stat
			})
		}
	})

	unary := utils.Reduce(statuses, "", func(reduced string, stat TStatus, i int, statuses []TStatus) string {
		result := reduced

		if i != 0 {
			result += " "
		}

		if stat.char == "1" {
			result += "0 "
		} else {
			result += "00 "
		}

		for i := 0; i < stat.occurence; i++ {
			result += "0"
		}

		return result
	})

	return unary, nil
}

func Solution(message string) (string, error) {
	err := isValid(message)
	if err != nil {
		return "", err
	}

	binary := textToBinary(message)

	unary, err := binaryToUnary(binary)
	if err != nil {
		return "", nil
	}

	return unary, nil
}
