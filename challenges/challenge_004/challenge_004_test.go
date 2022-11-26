package challenge_004_test

import (
	"errors"
	"fmt"
	"testing"

	challenge "code-challenges-go/challenges/challenge_004"

	"github.com/stretchr/testify/require"
)

func TestChallenge_004(t *testing.T) {
	t.Parallel()

	type input struct {
		lightX, lightY, initialTx, initialTy int
	}

	successCases := []struct {
		input  input
		output string
	}{
		{
			input:  input{lightX: 3, lightY: 8, initialTx: 3, initialTy: 6},
			output: "S S",
		},
		{
			input:  input{lightX: 3, lightY: 6, initialTx: 3, initialTy: 8},
			output: "N N",
		},
		{
			input:  input{lightX: 31, lightY: 4, initialTx: 5, initialTy: 4},
			output: "E E E E E E E E E E E E E E E E E E E E E E E E E E",
		},
		{
			input:  input{lightX: 31, lightY: 4, initialTx: 31, initialTy: 17},
			output: "N N N N N N N N N N N N N",
		},
		{
			input:  input{lightX: 0, lightY: 17, initialTx: 31, initialTy: 4},
			output: "SW SW SW SW SW SW SW SW SW SW SW SW SW W W W W W W W W W W W W W W W W W W",
		},
		{
			input:  input{lightX: 36, lightY: 17, initialTx: 0, initialTy: 0},
			output: "SE SE SE SE SE SE SE SE SE SE SE SE SE SE SE SE SE E E E E E E E E E E E E E E E E E E E",
		},
	}

	failureCases := []struct {
		input input
		err   error
	}{
		{
			input: input{lightX: -1, lightY: 8, initialTx: 3, initialTy: 6},
			err:   errors.New(challenge.OUT_OF_RANGE_LIGHT_X),
		},
		{
			input: input{lightX: 3, lightY: -1, initialTx: 3, initialTy: 6},
			err:   errors.New(challenge.OUT_OF_RANGE_LIGHT_Y),
		},
		{
			input: input{lightX: 3, lightY: 8, initialTx: -1, initialTy: 6},
			err:   errors.New(challenge.OUT_OF_RANGE_INITIAL_TX),
		},
		{
			input: input{lightX: 3, lightY: 8, initialTx: 3, initialTy: -1},
			err:   errors.New(challenge.OUT_OF_RANGE_INITIAL_TY),
		},
	}

	for i := range successCases {
		tc := successCases[i]

		t.Run(fmt.Sprintf("success case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.lightX, tc.input.lightY, tc.input.initialTx, tc.input.initialTy)
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, tc.output, result)
		})
	}

	for i := range failureCases {
		tc := failureCases[i]

		t.Run(fmt.Sprintf("failure case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.lightX, tc.input.lightY, tc.input.initialTx, tc.input.initialTy)
			require.Error(t, err)
			require.Equal(t, "", result)
			require.Equal(t, tc.err, err)
		})
	}
}
