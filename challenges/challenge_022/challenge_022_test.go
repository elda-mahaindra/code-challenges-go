package challenge_022_test

import (
	"errors"
	"fmt"
	"testing"

	challenge "code-challenges-go/challenges/challenge_022"

	"github.com/stretchr/testify/require"
)

func TestChallenge_022(t *testing.T) {
	t.Parallel()

	type input struct {
		h    int
		rows []string
	}

	successCases := []struct {
		input  input
		output []string
	}{
		{
			input: input{
				h:    4,
				rows: []string{"1 3 2 1", "1 3 2 1", "1 3 2 1", "1 3 2 1"},
			},
			output: []string{".OOO..O", ".OOO..O", ".OOO..O", ".OOO..O"},
		},
		{
			input: input{
				h:    4,
				rows: []string{"0 1 1 1 1", "0 1 1 1 1", "0 1 1 1 1", "0 1 1 1 1"},
			},
			output: []string{"O.O.", "O.O.", "O.O.", "O.O."},
		},
		{
			input: input{
				h: 8,
				rows: []string{
					"0 1 1 1 1 1 1 1 1",
					"1 1 1 1 1 1 1 1",
					"0 1 1 1 1 1 1 1 1",
					"1 1 1 1 1 1 1 1",
					"0 1 1 1 1 1 1 1 1",
					"1 1 1 1 1 1 1 1",
					"0 1 1 1 1 1 1 1 1",
					"1 1 1 1 1 1 1 1",
				},
			},
			output: []string{
				"O.O.O.O.",
				".O.O.O.O",
				"O.O.O.O.",
				".O.O.O.O",
				"O.O.O.O.",
				".O.O.O.O",
				"O.O.O.O.",
				".O.O.O.O",
			},
		},
		{
			input: input{
				h:    4,
				rows: []string{"8", "0 8", "8", "0 8"},
			},
			output: []string{"........", "OOOOOOOO", "........", "OOOOOOOO"},
		},
		{
			input: input{
				h: 8,
				rows: []string{
					"45",
					"2 2 3 2 2 3 2 1 1 1 3 1 2 2 3 2 2 1 3 1 1 4 1",
					"1 1 2 1 1 1 2 1 1 1 2 1 1 1 1 2 2 1 1 1 2 1 1 1 2 1 1 2 1 2 1 1 4",
					"1 1 4 1 2 1 1 1 2 1 1 1 1 1 1 1 1 1 1 1 4 1 2 1 1 1 1 1 1 1 1 3 2",
					"1 1 4 1 2 1 1 1 2 1 1 1 1 1 2 2 1 1 1 2 1 4 1 1 3 1 1 1 4",
					"1 1 2 1 1 1 2 1 1 1 2 1 1 1 1 1 3 1 1 1 2 1 1 1 2 1 1 1 3 1 1 1 4",
					"2 2 3 2 2 3 2 1 1 1 3 1 2 2 2 1 2 1 1 1 3 1 1 4 1",
					"45",
				},
			},
			output: []string{
				".............................................",
				"..OO...OO..OOO..O.O...O..OO...OO..O...O.OOOO.",
				".O..O.O..O.O..O.O.OO..O.O..O.O..O.OO.OO.O....",
				".O....O..O.O..O.O.O.O.O.O....O..O.O.O.O.OOO..",
				".O....O..O.O..O.O.O..OO.O.OO.OOOO.O...O.O....",
				".O..O.O..O.O..O.O.O...O.O..O.O..O.O...O.O....",
				"..OO...OO..OOO..O.O...O..OO..O..O.O...O.OOOO.",
				".............................................",
			},
		},
		{
			input: input{
				h:    4,
				rows: []string{"0 1 1 2", "0 2 1 1", "0 1 1 1", "1 1 1 1"},
			},
			output: []string{"INVALID"},
		},
		{
			input: input{
				h: 5,
				rows: []string{
					"0 1 2 1 1 1 2 2 2 1 1 1 1 1 2 1 1 2 1 1 1 1 1 1 1 1 1 2 1 2 1 1 1 1 1 1 1 3 1 2",
					"2 3 1 1 6 1 1 1 1 1 2 2 2 1 1 1 4 1 1 1 3 2 1 2 1 1 1 1 3 1",
					"1 1 3 5 2 3 1 2 1 3 2 2 2 3 1 2 4 2 4 5 1",
					"0 1 1 1 1 3 3 1 1 2 2 2 2 1 1 3 1 1 2 1 1 2 2 1 3 1 1 1 2 4 1 1",
					"2 2 1 1 1 3 1 1 2 1 6 1 1 4 1 1 2 1 1 1 1 1 1 1 2 2 3 1 4",
				},
			},
			output: []string{
				"O..O.O..OO..O.O.O..O.OO.O.O.O.O.OO.OO.O.O.O.OOO.OO",
				"..OOO.O......O.O.O..OO..O.O....O.O...OO.OO.O.O...O",
				".O...OOOOO..OOO.OO.OOO..OO..OOO.OO....OO....OOOOO.",
				"O.O.OOO...O.OO..OO..O.OOO.O..O.OO..O...O.O..OOOO.O",
				"..OO.O.OOO.O..O......O.OOOO.O..O.O.O.O..OO...O....",
			},
		},
	}

	failureCases := []struct {
		input input
		err   error
	}{
		{
			input: input{
				h:    200,
				rows: []string{"1 3 2 1", "1 3 2 1", "1 3 2 1", "1 3 2 1"},
			},
			err: errors.New(challenge.OUT_OF_RANGE_H),
		},
		{
			input: input{
				h:    4,
				rows: []string{"1 3 2 1", "1 3 2 1", "1 3 2 1"},
			},
			err: errors.New(challenge.INVALID_ROWS),
		},
	}

	for i := range successCases {
		tc := successCases[i]

		t.Run(fmt.Sprintf("success case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.h, tc.input.rows)
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, tc.output, result)
		})
	}

	for i := range failureCases {
		tc := failureCases[i]

		t.Run(fmt.Sprintf("failure case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.h, tc.input.rows)
			require.Error(t, err)
			require.Nil(t, result)
			require.Equal(t, tc.err, err)
		})
	}
}
