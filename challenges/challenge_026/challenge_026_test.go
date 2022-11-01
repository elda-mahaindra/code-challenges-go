package challenge_026_test

import (
	"errors"
	"fmt"
	"testing"

	challenge "code-challenges-go/challenges/challenge_026"

	"github.com/stretchr/testify/require"
)

func TestChallenge_026(t *testing.T) {
	t.Parallel()

	type input struct {
		s           string
		changeCount int
		rawChanges  []string
	}

	successCases := []struct {
		input  input
		output []string
	}{
		{
			input: input{
				s:           `Hello world`,
				changeCount: 4,
				rawChanges:  []string{"0|11|!", "0|5|,\\n", "0|7| w", "0|10|\\n"},
			},
			output: []string{"Hello,", " w worl", "d!"},
		},
		{
			input: input{
				s:           "He said that . To which I replied .",
				changeCount: 2,
				rawChanges:  []string{"0|13|I'm not good enough for the job", "0|34|\"Your lose!\""},
			},
			output: []string{"He said that I'm not good enough for the job. To which I replied \"Your lose!\"."},
		},
		{
			input: input{
				s:           "main\\nHello World}",
				changeCount: 4,
				rawChanges: []string{"0|0|void ",
					"1|0|  Console.WriteLine(\"",
					"0|4|()\\n{",
					"1|11|\");\\n"},
			},
			output: []string{"void main()", "{", "  Console.WriteLine(\"Hello World\");", "}"},
		},
		{
			input: input{
				s:           "\",,,\\n\"\\n-",
				changeCount: 5,
				rawChanges: []string{
					"0|1|You've gotta dance like there's nobody watching",
					"0|2|\\nLove like you'll\\nnever be hurt",
					"0|3|\\nSing like there's nobody listening",
					"1|0|And live like it's heaven on earth.",
					"2|1| William W. Purkey",
				},
			},
			output: []string{
				"\"You've gotta dance like there's nobody watching,",
				"Love like you'll",
				"never be hurt,",
				"Sing like there's nobody listening,",
				"And live like it's heaven on earth.\"",
				"- William W. Purkey",
			},
		},
	}

	failureCases := []struct {
		input input
		err   error
	}{
		{
			input: input{
				s:           "",
				changeCount: 4,
				rawChanges:  []string{"0|11|!", "0|5|,\\n", "0|7| w", "0|10|\\n"},
			},
			err: errors.New(challenge.OUT_OF_RANGE_S),
		},
		{
			input: input{
				s:           "Hello world",
				changeCount: 11,
				rawChanges:  []string{"0|11|!", "0|5|,\\n", "0|7| w", "0|10|\\n"},
			},
			err: errors.New(challenge.OUT_OF_RANGE_CHANGE_COUNT),
		},
		{
			input: input{
				s:           "Hello world",
				changeCount: 4,
				rawChanges:  []string{"0|11|!", "0|5|,\\n", "0|7| w"},
			},
			err: errors.New(challenge.INVALID_RAW_CHANGE),
		},
		{
			input: input{
				s:           "Hello world",
				changeCount: 4,
				rawChanges:  []string{"0|11|!", "0|5|,\\n", "0|7| w", "0|10"},
			},
			err: errors.New(challenge.INVALID_RAW_CHANGE),
		},
		{
			input: input{
				s:           "Hello world",
				changeCount: 4,
				rawChanges:  []string{"0|_|!", "0|5|,\\n", "0|7| w", "0|10|\\n"},
			},
			err: errors.New(challenge.INVALID_RAW_CHANGE),
		},
	}

	for i := range successCases {
		tc := successCases[i]

		t.Run(fmt.Sprintf("success case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.s, tc.input.changeCount, tc.input.rawChanges)
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, tc.output, result)
		})
	}

	for i := range failureCases {
		tc := failureCases[i]

		t.Run(fmt.Sprintf("failure case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.s, tc.input.changeCount, tc.input.rawChanges)
			require.Error(t, err)
			require.Nil(t, result)
			require.Equal(t, tc.err, err)
		})
	}
}
