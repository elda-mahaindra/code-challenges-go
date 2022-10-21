/*
   Encryption/Decryption of Enigma Machine

   During World War II, the Germans were using an encryption code called Enigma – which was basically an encryption machine that encrypted messages for transmission. The Enigma code went many years unbroken. Here's How the basic machine works:

   First Caesar shift is applied using an incrementing number:
   If String is AAA and starting number is 4 then output will be EFG.
   A + 4 = E
   A + 4 + 1 = F
   A + 4 + 1 + 1 = G

   Now map EFG to first ROTOR such as:
   ABCDEFGHIJKLMNOPQRSTUVWXYZ
   BDFHJLCPRTXVZNYEIWGAKMUSQO
   So EFG becomes JLC. Then it is passed through 2 more rotors to get the final value.

   If the second ROTOR is AJDKSIRUXBLHWTMCQGZNPYFVOE, we apply the substitution step again thus:
   ABCDEFGHIJKLMNOPQRSTUVWXYZ
   AJDKSIRUXBLHWTMCQGZNPYFVOE
   So JLC becomes BHD.

   If the third ROTOR is EKMFLGDQVZNTOWYHXUSPAIBRCJ, then the final substitution is:
   ABCDEFGHIJKLMNOPQRSTUVWXYZ
   EKMFLGDQVZNTOWYHXUSPAIBRCJ
   So BHD becomes KQF.

   Final output is sent via Radio Transmitter.

   Input
       • operation: a string represents the type of number which value is either "ENCODE" or "DECODE".
       • pseudoRandomNumber:  an integer represents the starting shift.
       • rotors: an array of three strings represents the three rotors that will be used in the process.
       • message: a string represents the message to be processed.

   Output
   Encoded or decoded string.

   Constraints:
       • 0 ≤ pseudoRandomNumber < 26
       • message consists only of uppercase letters (A-Z)
       • 1 ≤ message length < 50

   Example 1:
       Input: operation = "ENCODE", pseudoRandomNumber = 4, rotors = ["BDFHJLCPRTXVZNYEIWGAKMUSQO", "AJDKSIRUXBLHWTMCQGZNPYFVOE", "EKMFLGDQVZNTOWYHXUSPAIBRCJ"], message = "AAA"
       Output: "KQF"

   source: codingame
*/

package challenge_019

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"code-challenges-go/utils"
)

const (
	INVALID_MESSAGE                   = "the value of input 'message' should be between 1 and 49 letters and consists only of uppercase letters (A-Z)"
	INVALID_ROTORS                    = "each rotor inside the value input 'rotor' should contain exactly 26 A-Z letters in any order"
	OUT_OF_RANGE_PSEUDO_RANDOM_NUMBER = "the value of input 'pseudoRandomNumber' should be between 0 and 25"
	OUT_OF_RANGE_ROTORS               = "the length of input 'rotors' should be 3"
)

const uppercaseAlphabetLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func isValid(pseudoRandomNumber int, rotors []string, message string) error {
	regex, err := regexp.Compile("^[A-Z]{1,49}$")
	if err != nil {
		return err
	}

	switch {
	case pseudoRandomNumber < 0 || pseudoRandomNumber >= 26:
		return errors.New(OUT_OF_RANGE_PSEUDO_RANDOM_NUMBER)
	case len(rotors) != 3:
		return errors.New(OUT_OF_RANGE_ROTORS)
	case !regex.MatchString(message):
		return errors.New(INVALID_MESSAGE)
	case utils.Reduce(rotors, false, func(invalid bool, rotor string, i int, rotors []string) bool {
		sorted := strings.Split(rotor, "")
		sort.Strings(sorted)

		duplication := utils.Reduce(sorted, false, func(duplication bool, letter string, i int, letters []string) bool {
			if i == 0 {
				return duplication
			}

			return duplication || letter == letters[i-1]
		})

		return invalid || duplication || len(rotor) != len(uppercaseAlphabetLetters)
	}):
		return errors.New(INVALID_ROTORS)
	default:
		return nil
	}
}

func encode(pseudoRandomNumber int, rotors []string, message string) string {
	shifted := strings.Join(
		utils.Map(strings.Split(message, ""), func(letter string, i int, letters []string) string {
			temp1 := int([]rune(letter)[0]) - int([]rune("A")[0])
			temp2 := temp1 + pseudoRandomNumber + i
			temp3 := temp2 % len(uppercaseAlphabetLetters)
			charCode := temp3 + int([]rune("A")[0])

			return string(rune(charCode))
		}), "")

	mappedToRotors := utils.Reduce(rotors, shifted, func(mappedToRotors, rotor string, i int, rotors []string) string {
		return strings.Join(utils.Map(strings.Split(mappedToRotors, ""), func(letter string, i int, letters []string) string {
			return string([]rune(rotor)[int([]rune(letter)[0])-int([]rune("A")[0])])
		}), "")
	})

	return mappedToRotors
}

func decode(pseudoRandomNumber int, rotors []string, message string) string {
	mappedBackFromRotors := message
	for i := len(rotors) - 1; i >= 0; i-- {
		mappedBackFromRotors = utils.Reduce(strings.Split(mappedBackFromRotors, ""), "", func(mappedBackFromRotors, letter string, _ int, _ []string) string {
			indexFound := -1
			for index, l := range strings.Split(rotors[i], "") {
				if l == letter {
					indexFound = index
					break
				}
			}

			return fmt.Sprintf("%s%s", mappedBackFromRotors, string([]rune(uppercaseAlphabetLetters)[indexFound]))
		})
	}

	shiftedBack := strings.Join(
		utils.Map(strings.Split(mappedBackFromRotors, ""), func(letter string, i int, letters []string) string {
			temp1 := int([]rune(letter)[0]) - int([]rune("A")[0])
			temp2 := temp1 - pseudoRandomNumber - i
			temp3 := temp2 % len(uppercaseAlphabetLetters)
			charCode := temp3

			if temp3 >= 0 {
				charCode += int([]rune("A")[0])
			} else {
				charCode += int([]rune("Z")[0]) + 1
			}

			return string(rune(charCode))
		}), "")

	return shiftedBack
}

func Solution(operation string, pseudoRandomNumber int, rotors []string, message string) (string, error) {
	err := isValid(pseudoRandomNumber, rotors, message)
	if err != nil {
		return "", err
	}

	switch operation {
	case "ENCODE":
		{
			encoded := encode(pseudoRandomNumber, rotors, message)

			return encoded, nil
		}
	case "DECODE":
		{
			decoded := decode(pseudoRandomNumber, rotors, message)

			return decoded, nil
		}
	default:
		{
			return "", nil
		}
	}
}
