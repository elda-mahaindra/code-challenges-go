package challenge_013_test

import (
	"errors"
	"fmt"
	"testing"

	challenge "code-challenges-go/challenges/challenge_013"

	"github.com/stretchr/testify/require"
)

func TestChallenge_013(t *testing.T) {
	t.Parallel()

	type input struct {
		width, height int
		lines         []string
	}

	successCases := []struct {
		input  input
		output []string
	}{
		{
			input:  input{width: 2, height: 2, lines: []string{"00", "0."}},
			output: []string{"0 0 1 0 0 1", "1 0 -1 -1 -1 -1", "0 1 -1 -1 -1 -1"},
		},
		{
			input:  input{width: 5, height: 1, lines: []string{"0.0.0"}},
			output: []string{"0 0 2 0 -1 -1", "2 0 4 0 -1 -1", "4 0 -1 -1 -1 -1"},
		},
		{
			input: input{width: 1, height: 4, lines: []string{"0", "0", "0", "0"}},
			output: []string{
				"0 0 -1 -1 0 1",
				"0 1 -1 -1 0 2",
				"0 2 -1 -1 0 3",
				"0 3 -1 -1 -1 -1",
			},
		},
		{
			input: input{width: 3, height: 3, lines: []string{"0.0", "...", "0.0"}},
			output: []string{
				"0 0 2 0 0 2",
				"2 0 -1 -1 2 2",
				"0 2 2 2 -1 -1",
				"2 2 -1 -1 -1 -1",
			},
		},
		{
			input: input{width: 3, height: 3, lines: []string{"000", ".0.", ".0."}},
			output: []string{
				"0 0 1 0 -1 -1",
				"1 0 2 0 1 1",
				"2 0 -1 -1 -1 -1",
				"1 1 -1 -1 1 2",
				"1 2 -1 -1 -1 -1",
			},
		},
		{
			input: input{width: 4, height: 4, lines: []string{"0...", ".0..", "..0.", "...0"}},
			output: []string{
				"0 0 -1 -1 -1 -1",
				"1 1 -1 -1 -1 -1",
				"2 2 -1 -1 -1 -1",
				"3 3 -1 -1 -1 -1",
			},
		},
		{
			input: input{width: 4, height: 4, lines: []string{"00.0", "0.00", ".0.0", "000."}},
			output: []string{
				"0 0 1 0 0 1",
				"1 0 3 0 1 2",
				"3 0 -1 -1 3 1",
				"0 1 2 1 0 3",
				"2 1 3 1 2 3",
				"3 1 -1 -1 3 2",
				"1 2 3 2 1 3",
				"3 2 -1 -1 -1 -1",
				"0 3 1 3 -1 -1",
				"1 3 2 3 -1 -1",
				"2 3 -1 -1 -1 -1",
			},
		},
		{
			input: input{width: 7, height: 7, lines: []string{
				"..0....",
				".......",
				"..0.0.0",
				".......",
				"0.0.0..",
				".......",
				"....0..",
			}},
			output: []string{
				"2 0 -1 -1 2 2",
				"2 2 4 2 2 4",
				"4 2 6 2 4 4",
				"6 2 -1 -1 -1 -1",
				"0 4 2 4 -1 -1",
				"2 4 4 4 -1 -1",
				"4 4 -1 -1 4 6",
				"4 6 -1 -1 -1 -1",
			},
		},
	}

	failureCases := []struct {
		input input
		err   error
	}{
		{
			input: input{width: 0, height: 2, lines: []string{"00", "0."}},
			err:   errors.New(challenge.OUT_OF_RANGE_WIDTH),
		},
		{
			input: input{width: 2, height: 0, lines: []string{"00", "0."}},
			err:   errors.New(challenge.OUT_OF_RANGE_HEIGHT),
		},
		{
			input: input{width: 2, height: 2, lines: []string{"0", "0."}},
			err:   errors.New(challenge.INVALID_GRID_WIDTH),
		},
		{
			input: input{width: 2, height: 2, lines: []string{"00"}},
			err:   errors.New(challenge.INVALID_GRID_HEIGHT),
		},
	}

	for i := range successCases {
		tc := successCases[i]

		t.Run(fmt.Sprintf("success case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.width, tc.input.height, tc.input.lines)
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, tc.output, result)
		})
	}

	for i := range failureCases {
		tc := failureCases[i]

		t.Run(fmt.Sprintf("failure case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.width, tc.input.height, tc.input.lines)
			require.Error(t, err)
			require.Nil(t, result)
			require.Equal(t, tc.err, err)
		})
	}
}
