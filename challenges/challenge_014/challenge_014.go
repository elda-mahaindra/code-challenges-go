/*
   Scrabble

   When playing Scrabble©, each player draws 7 letters and must find a word that scores the most points using these letters.

   A player doesn't necessarily have to make a 7-letter word; the word can be shorter.
   The only constraint is that the word must be made using the 7 letters which the player has drawn.

   For example, with the letters  etaenhs, some possible words are: ethane, hates, sane, ant.

   Your objective is to find the word that scores the most points using the available letters (1 to 7 letters).

   Rules
   In Scrabble©, each letter is weighted with a score depending on how difficult it is to place that letter in a word.
   You will see below a list showing the letters corresponding to each point:

    1 : e, a, i, o, n, r, t, l, s, u
    2 : d, g
    3 : b, c, m, p
    4 : f, h, v, w, y
    5 : k
    8 : j, x
   10 : q, z

   The word banjo earns you 3 + 1 + 1 + 8 + 1 = 14 points.

   A dictionary of authorized words is provided as input for the program.
   The program must find the word in the dictionary which wins the most points for the seven given letters (a letter can only be used once).
   If two words win the same number of points, then the word which appears first in the order of the given dictionary should be chosen.

   All words will only be composed of alphabetical characters in lower case. There will always be at least one possible word.

   Input
       • N: the number of words in dictionary.
       • dictionary: an array of string which represents a dictionary, where each string represents a word in dictionary.
       • letters: the 7 letters available for the current game.

   Output
   The word that scores the most points using the available letters (1 to 7 letters). The word must belong to the dictionary.
   Each letter must be used at most once in the solution. There is always a solution.

   Constraints:
       • 0 < N < 100000
       • Words in the dictionary have a maximum length of 30 characters.

   Example 1:
     Input: N = 5, dictionary = ["because", "first", "these", "could", "which"], letters = "hicquwh"
     ]
     Output: "which"

   source: codingame
*/

package challenge_014

import (
	"errors"
	"strings"

	"code-challenges-go/utils"
)

const (
	INVALID_DICTIONARY   = "the length of input 'dictionary' should equal to the input 'N'"
	OUT_OF_RANGE_LETTERS = "the length of input 'letters' should equal to 7"
	OUT_OF_RANGE_N       = "the value of input 'N' should be between 1 and 99999"
	OUT_OF_RANGE_WORD    = "each string inside the input 'dictionary' should have maximum of 30 characters length"
)

type scrabblePoint struct {
	point   int
	letters string
}

var scrabblePoints = []scrabblePoint{
	{point: 1, letters: "eaionrtlsu"},
	{point: 2, letters: "dg"},
	{point: 3, letters: "bcmp"},
	{point: 4, letters: "fhvwy"},
	{point: 5, letters: "k"},
	{point: 8, letters: "jx"},
	{point: 10, letters: "qz"}}

func isValid(N int, dictionary []string, letters string) error {
	switch {
	case N < 1 || N > 99999:
		return errors.New(OUT_OF_RANGE_N)
	case len(letters) != 7:
		return errors.New(OUT_OF_RANGE_LETTERS)
	case len(dictionary) != N:
		return errors.New(INVALID_DICTIONARY)
	case !utils.Reduce(dictionary, true, func(valid bool, word string, i int, dictionary []string) bool {
		return valid && len(word) <= 30
	}):
		return errors.New(OUT_OF_RANGE_WORD)
	default:
		return nil
	}
}

func calculatePoint(letters, word string) int {
	if len(word) > len(letters) {
		return 0
	}

	remainingLetters := strings.Split(letters, "")
	matchingChars := ""
	totalPoint := 0

	for i := 0; i < len(word); i++ {
		indexFound := -1
		for index, letter := range remainingLetters {
			if letter == string([]rune(word)[i]) {
				indexFound = index
				break
			}
		}

		if indexFound >= 0 {
			point := -1
			for _, scrabblePoint := range scrabblePoints {
				splitted := strings.Split(scrabblePoint.letters, "")
				if utils.Includes(splitted, string([]rune(word)[i])) {
					point = scrabblePoint.point
				}
			}

			totalPoint += point

			// Remove the element at indexFound from remainingLetters.
			copy(remainingLetters[indexFound:], remainingLetters[indexFound+1:]) // Shift remainingLetters[indexFound+1:] left one index.
			remainingLetters[len(remainingLetters)-1] = ""                       // Erase last element (write zero value).
			remainingLetters = remainingLetters[:len(remainingLetters)-1]        // Truncate slice.

			matchingChars += string([]rune(word)[i])
		}
	}

	if matchingChars == word {
		return totalPoint
	}

	return 0
}

func Solution(N int, dictionary []string, letters string) (string, error) {
	err := isValid(N, dictionary, letters)
	if err != nil {
		return "", err
	}

	highestPointHolder := struct {
		word  string
		point int
	}{
		word:  "",
		point: 0,
	}

	for _, word := range dictionary {
		point := calculatePoint(letters, word)

		if point > highestPointHolder.point {
			highestPointHolder = struct {
				word  string
				point int
			}{
				word:  word,
				point: point,
			}
		}
	}

	return highestPointHolder.word, nil
}
