package challenge_002_test

import (
	"errors"
	"fmt"
	"testing"

	challenge "code-challenges-go/challenges/challenge_002"

	"github.com/stretchr/testify/require"
)

func TestChallenge_002(t *testing.T) {
	t.Parallel()

	type input struct {
		board [][]int
	}

	successCases := []struct {
		input  input
		output string
	}{
		{
			input: input{board: [][]int{
				{7, 9, 2, 1, 5, 4, 3, 8, 6},
				{6, 4, 3, 8, 2, 7, 1, 5, 9},
				{8, 5, 1, 3, 9, 6, 7, 2, 4},
				{2, 6, 5, 9, 7, 3, 8, 4, 1},
				{4, 8, 9, 5, 6, 1, 2, 7, 3},
				{3, 1, 7, 4, 8, 2, 9, 6, 5},
				{1, 3, 6, 7, 4, 8, 5, 9, 2},
				{9, 7, 4, 2, 1, 5, 6, 3, 8},
				{5, 2, 8, 6, 3, 9, 4, 1, 7},
			}},
			output: "valid",
		},
		{
			input: input{board: [][]int{
				{5, 5, 5, 5, 5, 5, 5, 5, 5},
				{5, 5, 5, 5, 5, 5, 5, 5, 5},
				{5, 5, 5, 5, 5, 5, 5, 5, 5},
				{5, 5, 5, 5, 5, 5, 5, 5, 5},
				{5, 5, 5, 5, 5, 5, 5, 5, 5},
				{5, 5, 5, 5, 5, 5, 5, 5, 5},
				{5, 5, 5, 5, 5, 5, 5, 5, 5},
				{5, 5, 5, 5, 5, 5, 5, 5, 5},
				{5, 5, 5, 5, 5, 5, 5, 5, 5},
			}},
			output: "not valid",
		},
	}

	failureCases := []struct {
		input input
		err   error
	}{
		{
			input: input{board: [][]int{
				{7, 9, 2, 1, 5, 4, 3},
				{6, 4, 3, 8, 2, 7, 1},
				{8, 5, 1, 3, 9, 6, 7},
				{2, 6, 5, 9, 7, 3, 8},
				{4, 8, 9, 5, 6, 1, 2},
				{3, 1, 7, 4, 8, 2, 9},
				{1, 3, 6, 7, 4, 8, 5},
			}},
			err: errors.New(challenge.INVALID_BOARD_DIMENSION),
		},
		{
			input: input{board: [][]int{
				{7, 0, 2, 1, 5, 4, 3, 8, 6},
				{6, 4, 3, 8, 2, 7, 1, 5, 0},
				{8, 5, 1, 3, 0, 6, 7, 2, 4},
				{2, 6, 5, 0, 7, 3, 8, 4, 1},
				{4, 8, 0, 5, 6, 1, 2, 7, 3},
				{3, 1, 7, 4, 8, 2, 0, 6, 5},
				{1, 3, 6, 7, 4, 8, 5, 0, 2},
				{0, 7, 4, 2, 1, 5, 6, 3, 8},
				{5, 2, 8, 6, 3, 0, 4, 1, 7},
			}},
			err: errors.New(challenge.OUT_OF_RANGE_BOARD_VALUE),
		},
	}

	for i := range successCases {
		tc := successCases[i]

		t.Run(fmt.Sprintf("success case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.board)
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, tc.output, result)
		})
	}

	for i := range failureCases {
		tc := failureCases[i]

		t.Run(fmt.Sprintf("failure case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.board)
			require.Error(t, err)
			require.Equal(t, "", result)
			require.Equal(t, tc.err, err)
		})
	}
}
