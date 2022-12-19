/*
   Prefix Code

   Given a fixed set of characters, a code is a table that gives the encoding to use for each character.

   A prefix code is a code with the prefix property, which is that there is no character with an encoding that is a prefix (initial segment) of the encoding of another character.


   Your goal is to decode an encoded string using the given prefix code, or say that is not possible.

   Example of encoding.
   Given the string "abracadabra" and the prefix code:
   a -> 1
   b -> 001
   c -> 011
   d -> 010
   r -> 000
   The resulting encoding is: 10010001011101010010001

   Thus, if your are given the code above and the input 10010001011101010010001, you should output the string "abracadabra".

   With the same prefix code, if the input is 0000, then you should tell that there is an error at index 3.
   Indeed, the first three characters of this input can be decoded to give an 'r', but that leaves 0, which cannot be decoded.

   Input
       • n: an integer represents the number of association in the prefix-code table.
       • associations: an array of n strings where each string represents the association of a binary code (Bi) and an integer (Ci) separated by a space,
           which tells that the character with ASCII code Ci will be encoded by Bi.
       • s: a string that represents the binary code of an encoded string.

   Output
   • If it is not possible to decode the encoded string, print DECODE FAIL AT INDEX i
       with i is the first index in the encoded string where the decoding fails (index starts from 0).
   • Otherwise print the decoded string.

   Constraints:
   • 0 ≤ n ≤ 127
   • 0 ≤ Ci ≤ 127
   • length of S <= 5000
   • length of Bi <= 5000

   Example 1:
       Input:
           n = 5,
           associations = [
               "1 97",
               "001 98",
               "000 114",
               "011 99",
               "010 100",
           ],
           s = "10010001011101010010001"
       Output: "abracadabra"

   source: codingame
*/

package challenge_032

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"code-challenges-go/utils"
)

const (
	INVALID_ASSOCIATIONS = "the number of associations inside input 'associations' should equal to the input 'N' and each associations should be a valid associations according to the constraints"
	OUT_OF_RANGE_N       = "the value of input 'N' should be between 0 and 127"
	OUT_OF_RANGE_S       = "the length of input 's' should less than or equal to 5000"
)

type TAssociation struct {
	b string // binary code
	c int    // ASCII code
}

func isValid(n int, associations []string, s string) error {
	switch {
	case n < 0 || n > 127:
		return errors.New(OUT_OF_RANGE_N)
	case len(s) > 5000:
		return errors.New(OUT_OF_RANGE_S)
	case len(associations) != n || !utils.Reduce(associations, true, func(valid bool, association string, i int, associations []string) bool {
		splitted := strings.Split(association, " ")

		b := splitted[0]
		c := splitted[1]

		parsedC, err := strconv.Atoi(c)
		if err != nil {
			panic(err)
		}

		validB := len(b) <= 5000
		validC := parsedC >= 0 && parsedC <= 127

		return valid && validB && validC

	}):
		return errors.New(INVALID_ASSOCIATIONS)
	default:
		return nil
	}
}

func Solution(n int, associations []string, s string) (string, error) {
	err := isValid(n, associations, s)
	if err != nil {
		return "", err
	}

	associationsTable := utils.Map(associations, func(association string, i int, associations []string) TAssociation {
		splitted := strings.Split(association, " ")

		parsedC, err := strconv.Atoi(splitted[1])
		if err != nil {
			panic(err)
		}

		return TAssociation{b: splitted[0], c: parsedC}
	})

	decoded := ""
	remainingS := s
	decodedCount := 0
	decodingFail := false

	for len(remainingS) > 0 {
		if decodingFail {
			break
		}

		for i := 0; i < len(remainingS); i++ {
			currentB := remainingS[:i+1]

			found := false
			associationFound := TAssociation{}
			for _, association := range associationsTable {
				if association.b == currentB {
					associationFound = association
					found = true
				}
			}

			if found {
				decoded = fmt.Sprintf("%s%s", decoded, string(rune(associationFound.c)))
				decodedCount += len(associationFound.b)
				remainingS = remainingS[len(associationFound.b):]
				break
			} else if i == len(remainingS)-1 {
				decodingFail = true
				break
			}
		}
	}

	if decodingFail {
		return fmt.Sprintf("DECODE FAIL AT INDEX %d", decodedCount), nil
	}

	return decoded, nil
}
