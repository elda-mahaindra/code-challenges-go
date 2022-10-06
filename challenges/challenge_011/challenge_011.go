/*
   MIME Type

   MIME types are used in numerous internet protocols to associate a media type (html, image, video ...) with the content sent.
   The MIME type is generally inferred from the extension of the file to be sent.

   You have to write a program that makes it possible to detect the MIME type of a file based on its name.

   Rule
   You are provided with a table which associates MIME types to file extensions.
   You are also given a list of names of files to be transferred and for each one of these files, you must find the MIME type to be used.

   The extension of a file is defined as the substring which follows the last occurrence, if any, of the dot character within the file name.
   If the extension for a given file can be found in the association table (case insensitive, e.g. TXT is treated the same way as txt), then print the corresponding MIME type.
   If it is not possible to find the MIME type corresponding to a file, or if the file doesn’t have an extension, print UNKNOWN.

   Input
       • mimeTypes: an array of strings which each string represents a file extension per line and the corresponding MIME type (separated by a blank space).
       • fileNames: an array of strings which each string represents a file name that it's corresponding MIME type need to be found in association table.

   Output
   An array of strings which each string represents a corresponding MIME type of the input 'fileNames' or UNKNOWN.

   Constraints:
       • 0 < mimeTypes length < 10000
       • 0 < fileNames length < 10000
       • File extensions are composed of a maximum of 10 alphanumerical ASCII characters.
       • MIME types are composed of a maximum 50 alphanumerical and punctuation ASCII characters.
       • File names are composed of a maximum of 256 alphanumerical ASCII characters and dots (full stops).
       • There are no spaces in the file names, extensions or MIME types.

   Example 1:
       Input: mimeTypes = ["html text/html", "png image/png", "gif image/gif"], fileNames = ["animated.gif", "portrait.png", "index.html"]
       Output: ["gif image/gif", "png image/png", "html text/html"]

   source: codingame
*/

package challenge_011

import (
	"errors"
	"regexp"
	"strings"

	"code-challenges-go/utils"
)

const (
	INVALID_FILE_EXTENSION  = "file extensions are composed of a maximum of 10 alphanumerical ascii characters without any space"
	INVALID_FILE_NAME       = "file names are composed of a maximum of 256 alphanumerical ascii characters and dots (full stops) without any space"
	INVALID_MIME_TYPE       = "mime types are composed of a maximum 50 alphanumerical and punctuation ascii characters without any space"
	OUT_OF_RANGE_FILE_NAMES = "the length of input 'fileNames' should be between 1 and 9999"
	OUT_OF_RANGE_MIME_TYPES = "the length of input 'mimeTypes' should be between 1 and 9999"
)

type mimeInfo struct {
	extension, mimeType string
}

func isValid(mimeTypes, fileNames []string) error {
	switch {
	case len(mimeTypes) <= 0 || len(mimeTypes) >= 10000:
		return errors.New(OUT_OF_RANGE_MIME_TYPES)
	case len(fileNames) <= 0 || len(fileNames) >= 10000:
		return errors.New(OUT_OF_RANGE_FILE_NAMES)
	case !utils.Reduce(mimeTypes, true, func(valid bool, mimeType string, i int, mimeTypes []string) bool {
		splitted := strings.Split(mimeType, " ")

		regex, err := regexp.Compile("^[0-9A-Za-z]{1,10}$")
		if err != nil {
			return valid && false
		}

		return valid && regex.MatchString(splitted[0])
	}):
		return errors.New(INVALID_FILE_EXTENSION)
	case !utils.Reduce(mimeTypes, true, func(valid bool, mimeType string, i int, mimeTypes []string) bool {
		splitted := strings.Split(mimeType, " ")

		regex, err := regexp.Compile(`^[\+\-\.\/0-9A-Za-z]{1,49}$`)
		if err != nil {
			return valid && false
		}

		return valid && regex.MatchString(splitted[1])
	}):
		return errors.New(INVALID_MIME_TYPE)
	case !utils.Reduce(fileNames, true, func(valid bool, fileName string, i int, fileNames []string) bool {
		regex, err := regexp.Compile(`^[\.0-9A-Za-z]{1,256}$`)
		if err != nil {
			return valid && false
		}

		return valid && regex.MatchString(fileName)
	}):
		return errors.New(INVALID_FILE_NAME)
	default:
		return nil
	}
}

func Solution(mimeTypes, fileNames []string) ([]string, error) {
	err := isValid(mimeTypes, fileNames)
	if err != nil {
		return nil, err
	}

	mimeTable := utils.Map(mimeTypes, func(mimeType string, i int, mimeTypes []string) mimeInfo {
		splitted := strings.Split(mimeType, " ")

		return mimeInfo{extension: splitted[0], mimeType: splitted[1]}
	})

	result := utils.Map(fileNames, func(fileName string, i int, filenames []string) string {
		splitted := strings.Split(fileName, "")
		dotFound := utils.Includes(splitted, ".")

		if !dotFound {
			return "UNKNOWN"
		} else {
			splitted := strings.Split(fileName, ".")
			extension := splitted[len(splitted)-1]

			mimeInfoFound, found := func(mimeTable []mimeInfo) (mimeInfo, bool) {
				for _, mimeInfo := range mimeTable {
					if strings.EqualFold(strings.ToLower(mimeInfo.extension), strings.ToLower(extension)) {
						return mimeInfo, true
					}
				}

				return mimeInfo{}, false
			}(mimeTable)

			if found {
				return mimeInfoFound.mimeType
			}

			return "UNKNOWN"
		}
	})

	return result, nil
}
