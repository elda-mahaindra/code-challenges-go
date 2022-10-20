package challenge_018_test

import (
	"errors"
	"fmt"
	"testing"

	challenge "code-challenges-go/challenges/challenge_018"

	"github.com/stretchr/testify/require"
)

func TestChallenge_018(t *testing.T) {
	t.Parallel()

	type input struct {
		encoded string
	}

	successCases := []struct {
		input  input
		output string
	}{
		{
			input:  input{encoded: "0 0 00 00000 0 0"},
			output: "A",
		},
		{
			input:  input{encoded: "0 0 00 00000 0 00 00 00000 0 0"},
			output: "AA",
		},
		{
			input: input{
				encoded: "0 0 00 0000 0 0 00 0 0 000 00 00 0 0 00 0 0 000 00 0 0 0 00 0 0 000 00 000 0 0000 00 00 0 0 00 0 0 0 00 0 0 0 00 00000 0 0 00 00 0 00 00 00 0 00 00 00 0 0 00 0 0 000 00 00 0 0 00 0 0 0",
			},
			output: "Bruce Lee",
		},
		{
			input: input{
				encoded: "0 0 00 0000 0 0 00 0 0 000 00 00 0 0 00 0 0 000 00 0 0 0 00 0 0 000 00 000 0 0000 00 00 0 0 00 0 0 0 00 0 0 0 00 00000 0 0 00 00 0 00 00 00 0 00 00 00 0 0 00 0 0 000 00 00 0 0 00 0 0 0 00 0 0 0 00 00000 0 000 00 0 0 00 00 0 0 000 00 00 0 00 00 0 0 0 00 00000 0 0 00 0000 0 0000 00 0 0 0 00 000 0 000 00 0 0 0 00 0 0 000 00 000 0 0000 00 0 0 0 00 0 0 00 00 0 0 0 00 00000 0 0 00 00 0 000 00 0 0 00 00 0 0 0000000 00 00 0 0 00 0 0 000 00 00 0 0 00 0 0 00 00 0 0 0 00 00 0 0000 00 00 0 00",
			},
			output: "Bruce Lee vs Chuck Norris",
		},
		{
			input:  input{encoded: "Bruce Lee"},
			output: "INVALID",
		},
		{
			input: input{
				encoded: "0 0 00 00 0 0 00 00 0 0000 00 0 0 0 00 000 0 0 00 00000 0 00 00 0 0 0 00 00 0 0000 00 00 0 00 00 0 0 0 00 00000 0 00 00 000 0 0000 00 0 0 000000 00 0 0 000 00 0 0 000 00 00 0 0000 00 0 0 0 00 00 0 000 00 00 0 0 00 00 0 00 00 00 0 0 00 0 0 0000 00 00 0 0 00 0 0 00 00 00 0 0 00 0 0 000 00 00 0 0 00 000 0 0 00 00000 0 00 00 0000 0 0 00 0 0 0 00 00000 0 00 00 00 0 000000 00 00 0 0 00 0 0 00 00 00 0 0 00 0 0 000 00 0000 0 0000 00 0 0 0 00 000 0 0 00 00000 0 00 00 0000 0 000 00 000 0 0000 00 000 0 0000 00 0 0 000000 00 0 0 00 00 0 0 0000 00 0000 0 00 00 0 0 00 00 00 0 00 00 0 0 0 00 00 0 0000 00 00 0 0000 00 0 0 0 00 000 0 00 00 0 0 00 00 0 0 000 00 00 0 0 00 0 0 000 00 0 0 000 00 0 0 000 00 0 0 0 00 000 0 0 00 00000 0 000 00 0 0 0 00 00 0 00 00 0 0 0000 00 0 0 0 00 00000 0 00 00 00 0 00000 00 0 0 0000 00 0 0 0 00 00000 0 00 00 00 0 0 00 00 0 00 00 0 0 0000000 00 0 0 00000 00 0 0 000 00 00 0 0 00 00000 0 0 00 00 0 000 00 0 0 00 00 0 0 0 00 00 0 000 00 0000 0 000 00 00 0 00000 00 0000 0 0000 00 00 0 0 00 0 0 00 00 0000 0 0 00 0 0 0 00 00000 0 0 00 000 0 00 00 0 0 00 00 0000 0 000 00 0 0 00 00 00 0 00 00 0 0 00 00 00 0 000 00 00 0 00 00 0 0 0 00 00000 0 00 00 0 0 0 00 00 0 000 00 0 0 000 00 00 0 0 00 00000 0 00 00 0000 0 0 00 0 0 0 00 00000 0 000 00 0 0 00000 00 0 0 000000 00 0 0 000000 00 00 0 0 00 00 0 00 00 00 0 0 00 0 0 000 00 0 0 000 00 00 0 0 00 00000 0 00 00 000 0 0 00 0 0 00 00 0000 0 0000 00 00 0 0 00 0 0 000 00 00 0 0 00 0 0 00 00 00 0 0 00 0 0 000 00 0 0 00 00 000 0 0 00 0 0 000 00 00 0 0 00 00000 0 0 00 0000 0 0000 00 0 0 0 00 000 0 000 00 0 0 0 00 0 0 000 00 000 0 0000 00 0 0 0 00 0 0 00 00 0 0 0 00 00000 0 0 00 00 0 000 00 0 0 00 00 0 0 0000000 00 00 0 0 00 0 0 000 00 00 0 0 00 0 0 00 00 0 0 0 00 00 0 0000 00 00 0 00 00 0 0 0 00 00000 0 00 00 000 0 0000 00 0000 0 000 00 0 0 000 00 00 0 0 00 00000 0 00 00 00 0 00000 00 0 0 0000 00 0 0 0 00 00000 0 000 00 0 0 0 00 0 0 0000 00 00000 0 0 00 00000 0 0 00 00 0 000 00 0 0 00 00 0 0 0 00 00 0 000 00 0000 0 000 00 00 0 00000 00 0000 0 0000 00 00 0 0 00 0 0 00 00 0000 0 0 00 0 0 0 00 00000 0 0 00 000 0 00 00 0 0 00 00 0000 0 000 00 0 0 00 00 00 0 00 00 0 0 00 00 00 0 000 00 00 0 00 00 0 0 0 00 00000 0 00 00 0 0 0 00 00 0 000 00 0 0 000 00 00 0 0 00 00000 0 00 00 0000 0 0 00 0 0 0 00 00000 0 00 00 000 0 0000 00 0000 0 0000 00 00 0 0 00 0 0 00 00 00 0 0 00 00 0 00 00 000 0 0 00 0 0 00 00 0 0 000000 00 0000 0 0000 00 00 0 0 00 0 0 00 00 00 0 0 00 000 0 0 00 00000 0 00 00 000 0 0 00 0 0 00 00 0 0 00000000 00 0000 0 0 00 0 0 000 00 0",
			},
			output: "It is considered a great accomplishment to go down Niagara Falls in a wooden barrel. Chuck Norris can go up Niagara Falls in a cardboard box.",
		},
		{
			input: input{
				encoded: "00 0 0 0 00 00 0 0 00 0 0 0 00 0 0 0 00 0 0 0 00 0 0 0 00 00 0 0 00 000 0 00 00 0 0 00 00 0000",
			},
			output: "%*#0",
		},
		{
			input:  input{encoded: "0 0 00 00 0 0 00 00 0 0000 00 0 and it's not valid..."},
			output: "INVALID",
		},
		{
			input:  input{encoded: "0 0 00 00000 000 0"},
			output: "INVALID",
		},
		{
			input:  input{encoded: "0 0 00 00000 00"},
			output: "INVALID",
		},
		{
			input:  input{encoded: "0 0 00 0000 0 0"},
			output: "INVALID",
		},
	}

	failureCases := []struct {
		input input
		err   error
	}{
		{
			input: input{encoded: ""},
			err:   errors.New(challenge.OUT_OF_RANGE_ENCODED),
		},
	}

	for i := range successCases {
		tc := successCases[i]

		t.Run(fmt.Sprintf("success case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.encoded)
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, tc.output, result)
		})
	}

	for i := range failureCases {
		tc := failureCases[i]

		t.Run(fmt.Sprintf("failure case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.encoded)
			require.Error(t, err)
			require.Equal(t, "", result)
			require.Equal(t, tc.err, err)
		})
	}
}
