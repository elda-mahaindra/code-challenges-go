/*
   Disordered First Contact

   Finally, we have received the first messages from aliens! Unfortunately, we cannot understand them because they have a very unique way of speaking.

   Here is how aliens encode their messages:
   - At each step of the encoding, they remove a bigger part from the beginning of the original message, starting from 1 character only. First, they take the first character, then 2 characters, then 3, etc...
   - Starting from an empty result string, they add each part taken from the original message alternatively at the end and at the beginning of the result string. They add the first part at the end, the second part at the beginning, the third part at the end, etc...

   Example
   abcdefghi becomes ghibcadef

   1) Take "a" from abcdefghi, add it at the end of an empty string -> a
   2) Take "bc" from bcdefghi, add it at the beginning of a -> bca
   3) Take "def" from defghi, add it at the end of bca -> bcadef
   4) Take the remaining characters "ghi" and add it at the beginning of bcadef -> ghibcadef

   Your job here is to decode or encode the messages to discuss with aliens.

   Input
       • N: an integer N indicating the number of times the message was transformed. If N is positive you have to decode i.e. retrieve the original message. If N is negative you have to encode i.e. transform the message.
       • message: a string message to be decoded or encoded.

   Output
   The encoded or decoded message, depend on the N.

   Constraints:
   • -10 ≤ N ≤ 10
   • 0 < message length < 1024

   Example 1:
       Input: N = 1, message = "ghibcadef"
       Output: "abcdefghi"

   source: codingame
*/

package challenge_020

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

const (
	OUT_OF_RANGE_MESSAGE = "the length of input 'message' should be between 1 and 1023"
	OUT_OF_RANGE_N       = "the value of input 'N' should be between -10 and 10"
)

type TStep struct {
	count int
	end   bool
}

func isValid(N int, message string) error {
	switch {
	case len(message) < 1 || len(message) > 1023:
		return errors.New(OUT_OF_RANGE_MESSAGE)
	case N < -10 || N > 10:
		return errors.New(OUT_OF_RANGE_N)
	default:
		return nil
	}
}

func encode(message string) string {
	letters := strings.Split(message, "")
	index := 0
	encoded := ""

	for len(letters) > 0 {
		var spliced = []string{}
		if len(letters) > index+1 {
			spliced = letters[0 : index+1]
			letters = letters[index+1:]
		} else {
			spliced = letters[0:]
			letters = []string{}
		}

		if index%2 != 0 {
			encoded = fmt.Sprintf("%s%s", strings.Join(spliced, ""), encoded)
		} else {
			encoded = fmt.Sprintf("%s%s", encoded, strings.Join(spliced, ""))
		}

		index++
	}

	return encoded
}

func getDecodeSteps(message string) []TStep {
	letters := strings.Split(message, "")
	index := 0
	steps := []TStep{}

	for len(letters) > 0 {
		var spliced = []string{}
		if len(letters) > index+1 {
			spliced = letters[0 : index+1]
			letters = letters[index+1:]
		} else {
			spliced = letters[0:]
			letters = []string{}
		}

		if len(steps) > 0 {
			steps = append([]TStep{{count: len(spliced), end: !steps[0].end}}, steps...)
		} else {
			steps = append([]TStep{{count: len(spliced), end: true}}, steps...)
		}

		index++
	}

	return steps
}

func decode(message string) string {
	decodeSteps := getDecodeSteps(message)

	letters := strings.Split(message, "")
	decoded := ""

	for _, step := range decodeSteps {
		count := step.count
		end := step.end

		var spliced = []string{}
		if end {
			spliced = letters[len(letters)-count : len(letters)-count+count]
			letters = letters[0 : len(letters)-count]
		} else {
			spliced = letters[0:count]
			letters = letters[count:]
		}

		decoded = fmt.Sprintf("%s%s", strings.Join(spliced, ""), decoded)
	}

	return decoded
}

func Solution(N int, message string) (string, error) {
	err := isValid(N, message)
	if err != nil {
		return "", err
	}

	result := message
	for i := 0; i < int(math.Abs(float64(N))); i++ {
		if N < 0 {
			result = encode(result)
		} else {
			result = decode(result)
		}
	}

	return result, nil
}
