package challenge_025_test

import (
	"errors"
	"fmt"
	"testing"

	challenge "code-challenges-go/challenges/challenge_025"

	"github.com/stretchr/testify/require"
)

func TestChallenge_025(t *testing.T) {
	t.Parallel()

	type input struct {
		side, diameter float64
	}

	successCases := []struct {
		input  input
		output int
	}{
		{
			input:  input{side: 3, diameter: 1},
			output: 2,
		},
		{
			input:  input{side: 12, diameter: 3},
			output: 4,
		},
		{
			input:  input{side: 12, diameter: 6},
			output: 0,
		},
		{
			input:  input{side: 12, diameter: 5},
			output: 3,
		},
		{
			input:  input{side: 30, diameter: 3.25},
			output: 27,
		},
		{
			input:  input{side: 6, diameter: 3.1},
			output: 3,
		},
		{
			input:  input{side: 14, diameter: 4},
			output: 6,
		},
		{
			input:  input{side: 34.8, diameter: 2.5},
			output: 77,
		},
		{
			input:  input{side: 99.99, diameter: 5.001},
			output: 147,
		},
		{
			input:  input{side: 89.89, diameter: 45.5},
			output: 3,
		},
		{
			input:  input{side: 11.99, diameter: 4},
			output: 7,
		},
		{
			input:  input{side: 96.5, diameter: 2.2},
			output: 600,
		},
		{
			input:  input{side: 0.95, diameter: 0.215},
			output: 8,
		},
		{
			input:  input{side: 98.95, diameter: 0.215},
			output: 58089,
		},
	}

	failureCases := []struct {
		input input
		err   error
	}{
		{
			input: input{side: 3, diameter: 0},
			err:   errors.New(challenge.OUT_OF_RANGE_DIAMETER),
		},
		{
			input: input{side: 100, diameter: 1},
			err:   errors.New(challenge.OUT_OF_RANGE_SIDE),
		},
	}

	for i := range successCases {
		tc := successCases[i]

		t.Run(fmt.Sprintf("success case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.side, tc.input.diameter)
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, tc.output, result)
		})
	}

	for i := range failureCases {
		tc := failureCases[i]

		t.Run(fmt.Sprintf("failure case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.side, tc.input.diameter)
			require.Error(t, err)
			require.Equal(t, -1, result)
			require.Equal(t, tc.err, err)
		})
	}
}
