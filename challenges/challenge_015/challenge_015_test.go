package challenge_015_test

import (
	"errors"
	"fmt"
	"math"
	"testing"

	challenge "code-challenges-go/challenges/challenge_015"

	"github.com/stretchr/testify/require"
)

func TestChallenge_015(t *testing.T) {
	t.Parallel()

	type input struct {
		n      int
		values []int
	}

	successCases := []struct {
		input  input
		output int
	}{
		{
			input:  input{n: 6, values: []int{3, 2, 4, 2, 1, 5}},
			output: -3,
		},
		{
			input:  input{n: 6, values: []int{5, 3, 4, 2, 3, 1}},
			output: -4,
		},
		{
			input:  input{n: 5, values: []int{1, 2, 4, 4, 5}},
			output: 0,
		},
		{
			input:  input{n: 5, values: []int{3, 4, 7, 9, 10}},
			output: 0,
		},
		{
			input:  input{n: 6, values: []int{3, 2, 10, 7, 15, 14}},
			output: -3,
		},
	}

	failureCases := []struct {
		input input
		err   error
	}{
		{
			input: input{n: 0, values: []int{3, 2, 4, 2, 1, 5}},
			err:   errors.New(challenge.OUT_OF_RANGE_N),
		},
		{
			input: input{n: 6, values: []int{int(math.Pow(2, 31)), 2, 4, 2, 1, 5}},
			err:   errors.New(challenge.OUT_OF_RANGE_VALUE),
		},
	}

	for i := range successCases {
		tc := successCases[i]

		t.Run(fmt.Sprintf("success case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.n, tc.input.values)
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, tc.output, result)
		})
	}

	for i := range failureCases {
		tc := failureCases[i]

		t.Run(fmt.Sprintf("failure case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.n, tc.input.values)
			require.Error(t, err)
			require.Equal(t, 1, result)
			require.Equal(t, tc.err, err)
		})
	}
}
