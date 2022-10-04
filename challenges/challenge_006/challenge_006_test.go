package challenge_006_test

import (
	"errors"
	"fmt"
	"testing"

	challenge "code-challenges-go/challenges/challenge_006"

	"github.com/stretchr/testify/require"
)

func TestChallenge_006(t *testing.T) {
	t.Parallel()

	type input struct {
		N      int
		inputs string
	}

	successCases := []struct {
		input  input
		output int
	}{
		{
			input:  input{N: 5, inputs: "1 -2 -8 4 5"},
			output: 1,
		},
		{
			input:  input{N: 3, inputs: "-12 -5 -137"},
			output: -5,
		},
		{
			input:  input{N: 6, inputs: "42 -5 12 21 5 24"},
			output: 5,
		},
		{
			input:  input{N: 6, inputs: "42 5 12 21 -5 24"},
			output: 5,
		},
		{
			input:  input{N: 10, inputs: "-5 -4 -2 12 -40 4 2 18 11 5"},
			output: 2,
		},
	}

	failureCases := []struct {
		input input
		err   error
	}{
		{
			input: input{N: 3, inputs: "1 -2 -8 4 5"},
			err:   errors.New(challenge.INVALID_INPUTS),
		},
		{
			input: input{N: -1, inputs: "1 -2 -8 4 5"},
			err:   errors.New(challenge.OUT_OF_RANGE_N),
		},
	}

	for i := range successCases {
		tc := successCases[i]

		t.Run(fmt.Sprintf("success case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.N, tc.input.inputs)
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, tc.output, result)
		})
	}

	for i := range failureCases {
		tc := failureCases[i]

		t.Run(fmt.Sprintf("failure case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.N, tc.input.inputs)
			require.Error(t, err)
			require.Equal(t, -274, result)
			require.Equal(t, tc.err, err)
		})
	}
}
