/*
   Table of Contents

   You are writing a book, and the table of contents is the only thing left to do.
   Sadly, the necessary packages are not working well, so you will have to implement one yourself.

   To generate the table of contents, your program will read N entries, describing a section with its level, title and page.
   • The level is given by the number of > at the start of the entry.
   • The title will not contain any space nor > characters.
   • The page is an integer, separated from the title by a space.

   Your program will then output the table of contents with the right format, N lines containing :
   • An indentation to reflect the level, 4 spaces per level.
   • The number of the section
   • Its title
   • A variable number of dots, for each line to be lengthofline long (including the page number)
   • The page number

   Input
       • lengthOfLine: an integer that represents the length of a line in the formatted result.
       • N: an integer that represents the number of entries.
       • entries: an array of string where each string represents an entry in bad format.

   Output
   An array of N strings where each string represents the entry in good format.

   Constraints:
   • 1 ≤ N ≤ 30
   • 30 ≤ lengthOfLine ≤ 50

   Example 1:
       Input:
           lengthOfLine = 40,
           N = 5,
           entries = [
               "Title1 4",
               ">Subtitle1 5",
               ">>Subsubtitle1 5",
               ">Subtitle2 6",
               "Title2 10",
           ]
       Output: [
           "1 Title1...............................4",
           "    1 Subtitle1........................5",
           "        1 Subsubtitle1.................5",
           "    2 Subtitle2........................6",
           "2 Title2..............................10",
       ]

   source: codingame
*/

package challenge_029

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"code-challenges-go/utils"
)

const (
	INVALID_ENTRIES             = "the length of input 'entries' should be equal to input 'N'"
	OUT_OF_RANGE_LENGTH_OF_LINE = "the value of input 'lengthOfLine' should be between 30 and 50"
	OUT_OF_RANGE_N              = "the value of input 'N' should be between 1 and 30"
)

type TContent struct {
	level             int
	pageNumber, title string
	sectionNumber     int
}

func isValid(lengthOfLine, N int, entries []string) error {
	switch {
	case lengthOfLine < 30 || lengthOfLine > 50:
		return errors.New(OUT_OF_RANGE_LENGTH_OF_LINE)
	case N < 1 || N > 30:
		return errors.New(OUT_OF_RANGE_N)
	case len(entries) != N:
		return errors.New(INVALID_ENTRIES)
	default:
		return nil
	}
}

func Solution(lengthOfLine, N int, entries []string) ([]string, error) {
	err := isValid(lengthOfLine, N, entries)
	if err != nil {
		return nil, err
	}

	contents := utils.Reduce(entries, []TContent{}, func(contents []TContent, entry string, i int, entries []string) []TContent {
		splitted := strings.Split(entry, " ")

		pageNumber := splitted[1]
		title := strings.ReplaceAll(splitted[0], ">", "")

		level := 0
		for i := 0; i < len(splitted[0]); i++ {
			if string([]rune(splitted[0])[i]) != ">" {
				break
			}

			level += 1
		}

		sectionNumber := 1

		if i != 0 {
			lastContent := contents[len(contents)-1]

			if lastContent.level == level {
				sectionNumber = lastContent.sectionNumber + 1
			} else if lastContent.level > level {
				filtered := utils.Filter(contents, func(content TContent) bool {
					return content.level == level
				})

				if len(filtered) > 0 {
					sectionNumber = filtered[len(filtered)-1].sectionNumber + 1
				}
			}
		}

		return append(contents, TContent{
			level:         level,
			pageNumber:    pageNumber,
			title:         title,
			sectionNumber: sectionNumber,
		})
	})

	goodFormattedContents := utils.Map(contents, func(content TContent, i int, contents []TContent) string {
		level := content.level
		pageNumber := content.pageNumber
		title := content.title
		sectionNumber := content.sectionNumber

		goodFormattedContent := ""
		for i := 0; i < level; i++ {
			goodFormattedContent = fmt.Sprintf("%s    ", goodFormattedContent)
		}

		goodFormattedContent = fmt.Sprintf("%s%s %s", goodFormattedContent, strconv.Itoa(sectionNumber), title)
		dotPaddingLength := lengthOfLine - len(goodFormattedContent) - len(pageNumber)

		for j := 0; j < dotPaddingLength; j++ {
			goodFormattedContent = fmt.Sprintf("%s.", goodFormattedContent)
		}

		return fmt.Sprintf("%s%s", goodFormattedContent, pageNumber)
	})

	return goodFormattedContents, nil
}
