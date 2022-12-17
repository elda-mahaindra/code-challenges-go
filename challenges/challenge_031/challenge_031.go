/*
   Retro Typewriter Art

   Back in the day, people had fun turning "recipes" into surprise images using typewriters.

   Use the provided recipe to create a recognizable image.

   Chunks in the recipe are separated by a space.
   Each chunk will tell you either
   nl meaning NewLine (aka Carriage Return)
   ~or~
   how many of the character and what character

   For example:
       4z means zzzz
       1{ means {
       10= means ==========
       5bS means \\\\\ (see Abbreviations list below)
       27 means 77
       123 means 333333333333
       (If a chunk is composed only of numbers, the character is the last digit.)

       So if part of the recipe is
       2* 15sp 1x 4sQ nl
       ...that tells you to show
       **               x''''
       and then go to a new line.

   Abbreviations used:
       sp = space
       bS = backSlash \
       sQ = singleQuote '
       and
       nl = NewLine

   Sources/references:
       https://asciiart.cc
       https://loriemerson.net/2013/01/18/d-i-y-typewriter-art/
       https://www.youtube.com/watch?v=kyK5WvpFxqo

   Input
       • recipe: a string represents the recipe.

   Output
   An array of strings showing the image created by the recipe.

   Constraints:
   • 5 ≤ length of recipe ≤ 1000
   • There won't be any double quotes (") in the recipe
   • recipe will contain at least 1 nl

   Example 1:
       Input: recipe = "1sp 1/ 1bS 1_ 1/ 1bS nl 1( 1sp 1o 1. 1o 1sp 1) nl 1sp 1> 1sp 1^ 1sp 1< nl 2sp 3|"
       Output: [
            " /\_/\",
            "( o.o )",
            " > ^ <",
            "  |||",
       ]

   source: codingame
*/

package challenge_031

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"code-challenges-go/utils"
)

const (
	INVALID_RECIPE      = "the input 'recipe' should not have any double quotes but have at least one 'nl'"
	OUT_OF_RANGE_RECIPE = "the length of input 'recipe' should be between 5 and 1000"
)

func isValid(recipe string) error {
	switch {
	case len(recipe) < 5 || len(recipe) > 1000:
		return errors.New(OUT_OF_RANGE_RECIPE)
	case strings.Contains(recipe, "\"") || !strings.Contains(recipe, "nl"):
		return errors.New(INVALID_RECIPE)
	default:
		return nil
	}
}

func translateChunk(chunk string) string {
	if chunk == "nl" {
		return "nl"
	}

	regex, err := regexp.Compile(`^[\d]+$`)
	if err != nil {
		panic(err)
	}

	if regex.MatchString(chunk) {
		char := string([]rune(chunk)[len(chunk)-1])
		times, err := strconv.Atoi(chunk[:len(chunk)-1])
		if err != nil {
			panic(err)
		}

		translated := ""
		for i := 0; i < times; i++ {
			translated = fmt.Sprintf("%s%s", translated, char)
		}

		return translated
	}

	index := -1
	index = strings.Index(chunk, "sp")
	if index >= 0 {
		times, err := strconv.Atoi(chunk[:index])
		if err != nil {
			panic(err)
		}

		translated := ""
		for i := 0; i < times; i++ {
			translated = fmt.Sprintf("%s%s", translated, " ")
		}

		return translated
	}

	index = strings.Index(chunk, "bS")
	if index >= 0 {
		times, err := strconv.Atoi(chunk[:index])
		if err != nil {
			panic(err)
		}

		translated := ""
		for i := 0; i < times; i++ {
			translated = fmt.Sprintf("%s%s", translated, "\\")
		}

		return translated
	}

	index = strings.Index(chunk, "sQ")
	if index >= 0 {
		times, err := strconv.Atoi(chunk[:index])
		if err != nil {
			panic(err)
		}

		translated := ""
		for i := 0; i < times; i++ {
			translated = fmt.Sprintf("%s%s", translated, "'")
		}

		return translated
	}

	char := string([]rune(chunk)[len(chunk)-1])
	times, err := strconv.Atoi(chunk[:len(chunk)-1])
	if err != nil {
		panic(err)
	}

	translated := ""
	for i := 0; i < times; i++ {
		translated = fmt.Sprintf("%s%s", translated, char)
	}

	return translated
}

func Solution(recipe string) ([]string, error) {
	err := isValid(recipe)
	if err != nil {
		return nil, err
	}

	chunks := strings.Split(recipe, " ")

	result := utils.Reduce(chunks, []string{""}, func(result []string, chunk string, i int, chunks []string) []string {
		translated := translateChunk(chunk)

		if translated != "nl" {
			mapped := utils.Map(result, func(line string, i int, result []string) string {
				if i == len(result)-1 {
					return fmt.Sprintf("%s%s", line, translated)
				}
				return line
			})

			return mapped
		}

		return append(result, "")
	})

	return result, nil
}
