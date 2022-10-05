package challenge_010_test

import (
	"errors"
	"fmt"
	"testing"

	challenge "code-challenges-go/challenges/challenge_010"

	"github.com/stretchr/testify/require"
)

func TestChallenge_010(t *testing.T) {
	t.Parallel()

	type input struct {
		forest []string
	}

	successCases := []struct {
		input  input
		output string
	}{
		{
			input: input{
				forest: []string{
					"######",
					"######",
					"######",
					"######",
					"######",
					"******",
				},
			},
			output: "12",
		},
		{
			input: input{
				forest: []string{
					"######",
					"######",
					"######",
					"######",
					"======",
					"******",
				},
			},
			output: "6",
		},
		{
			input: input{
				forest: []string{
					"======",
					"======",
					"==*===",
					"======",
					"======",
					"======",
				},
			},
			output: "JUST RUN",
		},
		{
			input: input{
				forest: []string{
					"######",
					"######",
					"######",
					"##*###",
					"######",
					"######",
				},
			},
			output: "24",
		},
		{
			input: input{
				forest: []string{
					"******",
					"oooooo",
					"oooooo",
					"oooooo",
					"oooooo",
					"oooooo",
				},
			},
			output: "JUST RUN",
		},
		{
			input: input{
				forest: []string{
					"******",
					"#*****",
					"******",
					"******",
					"******",
					"*****#",
				},
			},
			output: "JUST RUN",
		},
		{
			input: input{
				forest: []string{
					"#o##o#",
					"#o*#o#",
					"*o##o*",
					"#o##o#",
					"#o**o#",
					"#o##o#",
				},
			},
			output: "JUST RUN",
		},
		{
			input: input{
				forest: []string{
					"======",
					"=###==",
					"======",
					"=###==",
					"==oo==",
					"======",
				},
			},
			output: "RELAX",
		},
	}

	failureCases := []struct {
		input input
		err   error
	}{
		{
			input: input{
				forest: []string{
					"#####",
					"#####",
					"#####",
					"#####",
					"*****",
				},
			},
			err: errors.New(challenge.INVALID_FOREST),
		},
	}

	for i := range successCases {
		tc := successCases[i]

		t.Run(fmt.Sprintf("success case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.forest)
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, tc.output, result)
		})
	}

	for i := range failureCases {
		tc := failureCases[i]

		t.Run(fmt.Sprintf("failure case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.forest)
			require.Error(t, err)
			require.Equal(t, "", result)
			require.Equal(t, tc.err, err)
		})
	}
}
