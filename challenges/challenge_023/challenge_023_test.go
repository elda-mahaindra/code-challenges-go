package challenge_023_test

import (
	"errors"
	"fmt"
	"testing"

	challenge "code-challenges-go/challenges/challenge_023"

	"github.com/stretchr/testify/require"
)

func TestChallenge_023(t *testing.T) {
	t.Parallel()

	type input struct {
		count          int
		decimaryInputs string
	}

	successCases := []struct {
		input  input
		output string
	}{
		{
			input: input{
				count:          3,
				decimaryInputs: "A A 12",
			},
			output: "32",
		},
		{
			input: input{
				count:          4,
				decimaryInputs: "9A A2 1A 12",
			},
			output: "234",
		},
		{
			input: input{
				count:          3,
				decimaryInputs: "1AA A2A AA5",
			},
			output: "2345",
		},
		{
			input: input{
				count:          4,
				decimaryInputs: "1 2 3 4",
			},
			output: "A",
		},
		{
			input: input{
				count:          8,
				decimaryInputs: "512 256 128 64 32 16 8 8",
			},
			output: "A24",
		},
		{
			input: input{
				count:          2,
				decimaryInputs: "19 91",
			},
			output: "AA",
		},
		{
			input: input{
				count:          3,
				decimaryInputs: "99A 9A9 A1",
			},
			output: "1AAA",
		},
		{
			input: input{
				count:          4,
				decimaryInputs: "499A 2A1A AA9 911",
			},
			output: "9A3A",
		},
		{
			input: input{
				count:          6,
				decimaryInputs: "123456789 AAA21 AA3A A54A A67A A8A9AAA",
			},
			output: "1344AA18A",
		},
	}

	failureCases := []struct {
		input input
		err   error
	}{
		{
			input: input{
				count:          10,
				decimaryInputs: "A A 12",
			},
			err: errors.New(challenge.OUT_OF_RANGE_COUNT),
		},
		{
			input: input{
				count:          3,
				decimaryInputs: "123456789A A 12",
			},
			err: errors.New(challenge.INVALID_DECIMARY),
		},
	}

	for i := range successCases {
		tc := successCases[i]

		t.Run(fmt.Sprintf("success case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.count, tc.input.decimaryInputs)
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, tc.output, result)
		})
	}

	for i := range failureCases {
		tc := failureCases[i]

		t.Run(fmt.Sprintf("failure case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.count, tc.input.decimaryInputs)
			require.Error(t, err)
			require.Equal(t, "", result)
			require.Equal(t, tc.err, err)
		})
	}
}
