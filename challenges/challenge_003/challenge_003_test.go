package challenge_003_test

import (
	"errors"
	"fmt"
	"testing"

	challenge "code-challenges-go/challenges/challenge_003"

	"github.com/stretchr/testify/require"
)

func TestChallenge_003(t *testing.T) {
	t.Parallel()

	type input struct {
		h int
		n int
	}

	successCases := []struct {
		input  input
		output []string
	}{
		{
			input: input{h: 5, n: 1},
			output: []string{
				"..../\\....",
				".../..\\...",
				"../....\\..",
				"./......\\.",
				"/........\\",
			},
		},
		{
			input: input{h: 5, n: 2},
			output: []string{
				"..../\\......../\\....",
				".../..\\....../..\\...",
				"../....\\..../....\\..",
				"./......\\../......\\.",
				"/........\\/........\\",
			},
		},
		{
			input: input{h: 10, n: 1},
			output: []string{
				"........./\\.........",
				"......../..\\........",
				"......./....\\.......",
				"....../......\\......",
				"...../........\\.....",
				"..../..........\\....",
				".../............\\...",
				"../..............\\..",
				"./................\\.",
				"/..................\\",
			},
		},
		{
			input: input{h: 3, n: 6},
			output: []string{
				"../\\..../\\..../\\..../\\..../\\..../\\..",
				"./..\\../..\\../..\\../..\\../..\\../..\\.",
				"/....\\/....\\/....\\/....\\/....\\/....\\",
			},
		},
	}

	failureCases := []struct {
		input input
		err   error
	}{
		{
			input: input{h: 1, n: 1},
			err:   errors.New(challenge.OUT_OF_RANGE_H),
		},
		{
			input: input{h: 51, n: 1},
			err:   errors.New(challenge.OUT_OF_RANGE_H),
		},
		{
			input: input{h: 5, n: 0},
			err:   errors.New(challenge.OUT_OF_RANGE_N),
		},
		{
			input: input{h: 5, n: 11},
			err:   errors.New(challenge.OUT_OF_RANGE_N),
		},
	}

	for i := range successCases {
		tc := successCases[i]

		t.Run(fmt.Sprintf("success case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.h, tc.input.n)
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, tc.output, result)
		})
	}

	for i := range failureCases {
		tc := failureCases[i]

		t.Run(fmt.Sprintf("failure case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.h, tc.input.n)
			require.Error(t, err)
			require.Nil(t, result)
			require.Equal(t, tc.err, err)
		})
	}
}
