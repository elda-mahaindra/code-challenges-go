package challenge_021_test

import (
	"errors"
	"fmt"
	"testing"

	challenge "code-challenges-go/challenges/challenge_021"

	"github.com/stretchr/testify/require"
)

func TestChallenge_021(t *testing.T) {
	t.Parallel()

	type input struct {
		n       int
		rows    []string
		message string
	}

	successCases := []struct {
		input  input
		output string
	}{
		{
			input: input{
				n: 2,
				rows: []string{
					"A B",
					"C D",
				},
				message: "CBA",
			},
			output: "100100",
		},
		{
			input: input{
				n:       3,
				rows:    []string{"A B C D E F G H I", "J K L M N O P Q R", "S T U V W X Y Z"},
				message: "HELLOWORLD",
			},
			output: "07041212152415181203",
		},
		{
			input: input{
				n: 7,
				rows: []string{
					"Z Y X W",
					"V U T S",
					"R Q P O",
					"N M L K",
					"J I H G",
					"F E D C",
					"B A",
				},
				message: "HELLOWORLD",
			},
			output: "42513232230323203252",
		},
		{
			input: input{
				n: 6,
				rows: []string{
					"H U P N I",
					"J W C F T",
					"B Z A Q Y",
					"L O S X R",
					"M E V D G",
					"K # @ ! ?",
				},
				message: "SXD@OE?AR",
			},
			output: "323343523141542234",
		},
	}

	failureCases := []struct {
		input input
		err   error
	}{
		{
			input: input{
				n: 11,
				rows: []string{
					"A B",
					"C D",
				},
				message: "CBA",
			},
			err: errors.New(challenge.OUT_OF_RANGE_N),
		},
		{
			input: input{
				n: 3,
				rows: []string{
					"A B",
					"C D",
				},
				message: "CBA",
			},
			err: errors.New(challenge.INVALID_ROWS),
		},
		{
			input: input{
				n: 2,
				rows: []string{
					"AB",
					"C D",
				},
				message: "CBA",
			},
			err: errors.New(challenge.INVALID_ROWS),
		},
	}

	for i := range successCases {
		tc := successCases[i]

		t.Run(fmt.Sprintf("success case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.n, tc.input.rows, tc.input.message)
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, tc.output, result)
		})
	}

	for i := range failureCases {
		tc := failureCases[i]

		t.Run(fmt.Sprintf("failure case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.n, tc.input.rows, tc.input.message)
			require.Error(t, err)
			require.Equal(t, "", result)
			require.Equal(t, tc.err, err)
		})
	}
}
