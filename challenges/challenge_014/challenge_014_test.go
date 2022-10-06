package challenge_014_test

import (
	"errors"
	"fmt"
	"testing"

	challenge "code-challenges-go/challenges/challenge_014"

	"github.com/stretchr/testify/require"
)

func TestChallenge_014(t *testing.T) {
	t.Parallel()

	type input struct {
		N          int
		dictionary []string
		letters    string
	}

	successCases := []struct {
		input  input
		output string
	}{
		{
			input: input{
				N:          5,
				dictionary: []string{"because", "first", "these", "could", "which"},
				letters:    "hicquwh",
			},
			output: "which",
		},
		{
			input: input{
				N: 10,
				dictionary: []string{
					"some", "first", "potsie", "day", "could",
					"postie", "from", "have", "back", "this",
				},
				letters: "sopitez",
			},
			output: "potsie",
		},
		{
			input: input{
				N: 10,
				dictionary: []string{
					"after", "repots", "user", "powers", "these",
					"time", "know", "from", "could", "people",
				},
				letters: "tsropwe",
			},
			output: "powers",
		},
		{
			input: input{
				N: 10,
				dictionary: []string{
					"arrest", "rarest", "raster", "raters", "sartre",
					"starer", "waster", "waters", "wrest", "wrase",
				},
				letters: "arwtsre",
			},
			output: "waster",
		},
		{
			input: input{
				N:          5,
				dictionary: []string{"entire", "tween", "soft", "would", "test"},
				letters:    "etiewrn",
			},
			output: "tween",
		},
		{
			input: input{
				N:          5,
				dictionary: []string{"qzyoq", "azejuy", "kqjsdh", "aeiou", "qsjkdh"},
				letters:    "qzaeiou",
			},
			output: "aeiou",
		},
		{
			input: input{
				N: 10,
				dictionary: []string{
					"after", "repots", "poowers", "powers", "these",
					"time", "know", "from", "could", "people",
				},
				letters: "tsropwe",
			},
			output: "powers",
		},
	}

	failureCases := []struct {
		input input
		err   error
	}{
		{
			input: input{
				N:          0,
				dictionary: []string{"because", "first", "these", "could", "which"},
				letters:    "hicquwh",
			},
			err: errors.New(challenge.OUT_OF_RANGE_N),
		},
		{
			input: input{
				N:          5,
				dictionary: []string{"because", "first", "these", "could", "which"},
				letters:    "hic",
			},
			err: errors.New(challenge.OUT_OF_RANGE_LETTERS),
		},
		{
			input: input{
				N:          5,
				dictionary: []string{"because", "first", "these"},
				letters:    "hicquwh",
			},
			err: errors.New(challenge.INVALID_DICTIONARY),
		},
		{
			input: input{
				N:          5,
				dictionary: []string{"becausebecausebecausebecausebecause", "first", "these", "could", "which"},
				letters:    "hicquwh",
			},
			err: errors.New(challenge.OUT_OF_RANGE_WORD),
		},
	}

	for i := range successCases {
		tc := successCases[i]

		t.Run(fmt.Sprintf("success case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.N, tc.input.dictionary, tc.input.letters)
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, tc.output, result)
		})
	}

	for i := range failureCases {
		tc := failureCases[i]

		t.Run(fmt.Sprintf("failure case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.N, tc.input.dictionary, tc.input.letters)
			require.Error(t, err)
			require.Equal(t, "", result)
			require.Equal(t, tc.err, err)
		})
	}
}
