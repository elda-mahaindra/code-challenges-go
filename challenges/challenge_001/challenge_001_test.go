package challenge_001_test

import (
	"errors"
	"fmt"
	"testing"

	challenge "code-challenges-go/challenges/challenge_001"

	"github.com/stretchr/testify/require"
)

func TestChallenge_001(t *testing.T) {
	t.Parallel()

	type input struct {
		s string
	}

	successCases := []struct {
		input  input
		output string
	}{
		{
			input:  input{s: "abcddefda1111133333333"},
			output: "d",
		},
		{
			input:  input{s: "AA0AB0BB0ccc0aa0aw00wo0BBBw123123"},
			output: "B",
		},
		{
			input:  input{s: "aaaaaaaaaaaBCGDhsjwnq"},
			output: "a",
		},
	}

	failureCases := []struct {
		input input
		err   error
	}{
		{
			input: input{s: "0000000012312313123"},
			err:   errors.New(challenge.INVALID_S),
		},
		{
			input: input{s: "zvchW39Kt5mQQLXXE8IekkBAVOjwRVaKaHGwSI44uuJ005n1luXnfwW8pBW7S1ujXdrintWaZ7nK6dBsjch4BRPYazB2kdcAgfZak"},
			err:   errors.New(challenge.OUT_OF_RANGE_S),
		},
	}

	for i := range successCases {
		tc := successCases[i]

		t.Run(fmt.Sprintf("success case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.s)
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, tc.output, result)
		})
	}

	for i := range failureCases {
		tc := failureCases[i]

		t.Run(fmt.Sprintf("failure case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.s)
			require.Error(t, err)
			require.Equal(t, "", result)
			require.Equal(t, tc.err, err)
		})
	}
}
