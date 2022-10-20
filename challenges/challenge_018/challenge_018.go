/*
   Bruce Lee

   Your program must decode the encoded message from the Chuck Norris encoding project.

   It is strongly recommended to have done the Chuck Norris project.
   Link -> https://www.codingame.com/training/easy/chuck-norris

   Here are some reminders about the Chuck Norris encoding method:
   - The encoded message is unary, containing only sequences of zeroes separated by spaces.
   - These sequences of zeroes always come in pairs.
   - The first sequence of a pair can be either 0 or 00, which represent the binary bits 1 and 0 respectively.
   - The second sequence of a pair is made of k zeroes, where k is the number of time the previous bit has to be printed in order to decode the message.

   For instance, if we want to encode the character A, we first start to write down the 7-bit ASCII code for A which is 1000001 in binary. (We only use 7 bits because the first bit is always zero so it's ignored).
   Then we turn the binary into unary as follows:
   1000001 -> 0 0 (bit 1, one time)
   1000001 -> 00 00000 (bit 0, five times)
   1000001 -> 0 0 (bit 1, one time)
   Therefore, the encoded message is 0 0 00 00000 0 0.

   You are asked to do the reverse process, and thus print A when given the message 0 0 00 00000 0 0.
   If the input is invalid, just print INVALID.

   Input
       • encoded: an encoded string message of N characters.

   Output
   The decoded message, or the word INVALID when the input is not valid.

   Constraints:
       • 0 < N < 4096

   Example 1:
       Input: encoded = "0 0 00 00000 0 0"
       Output: "A"

   source: codingame
*/

package challenge_018

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"code-challenges-go/utils"
)

const (
	OUT_OF_RANGE_ENCODED = "the length of input 'encoded' should be between 1 and 4095"
	INVALID_UNARY        = "expecting a valid unary string"
)

func isValid(encoded string) error {
	switch {
	case len(encoded) < 1 || len(encoded) > 4095:
		return errors.New(OUT_OF_RANGE_ENCODED)
	default:
		return nil
	}
}

func unaryToBinaryStrings(unary string) (string, error) {
	splitted := strings.Split(unary, " ")

	if len(splitted)%2 != 0 {
		return "", errors.New(INVALID_UNARY)
	}

	binaryStrings := ""
	for i := 0; i < len(splitted); i += 2 {
		firstBlock := splitted[i]
		secondBlock := splitted[i+1]

		if firstBlock != "0" && firstBlock != "00" {
			return "", errors.New(INVALID_UNARY)
		}

		for j := 0; j < len(secondBlock); j++ {
			if firstBlock == "0" {
				binaryStrings = fmt.Sprintf("%s%s", binaryStrings, "1")
			} else {
				binaryStrings = fmt.Sprintf("%s%s", binaryStrings, "0")
			}
		}
	}

	return binaryStrings, nil
}

func binaryToChar(binary string) (string, error) {
	ascii, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		return "", err
	}

	char := rune(ascii)

	return string(char), nil
}

func Solution(encoded string) (string, error) {
	err := isValid(encoded)
	if err != nil {
		return "", err
	}

	binaryStrings, err := unaryToBinaryStrings(encoded)
	if err != nil {
		return "INVALID", nil
	}

	if len(binaryStrings)%7 != 0 {
		return "INVALID", nil
	}

	binaries := []string{}
	for i := 0; i < len(binaryStrings); i += 7 {
		binaries = append(binaries, string([]rune(binaryStrings)[i:i+7]))
	}

	decoded := utils.Reduce(binaries, "", func(decoded, binary string, i int, binaries []string) string {
		char, err := binaryToChar(binary)
		if err != nil {
			return decoded
		}

		return fmt.Sprintf("%s%s", decoded, char)
	})

	return decoded, nil
}
