package challenge_009_test

import (
	"errors"
	"fmt"
	"testing"

	challenge "code-challenges-go/challenges/challenge_009"

	"github.com/stretchr/testify/require"
)

func TestChallenge_009(t *testing.T) {
	t.Parallel()

	type input struct {
		w, h  int
		lines []string
	}

	successCases := []struct {
		input  input
		output []string
	}{
		{
			input: input{
				w: 16,
				h: 9,
				lines: []string{
					"................",
					"................",
					"................",
					"................",
					"................",
					"....x...........",
					"................",
					"................",
					"................",
				},
			},
			output: []string{
				"................",
				"................",
				"................",
				"................",
				"...111..........",
				"...1.1..........",
				"...111..........",
				"................",
				"................",
			},
		},
		{
			input: input{
				w: 16,
				h: 11,
				lines: []string{
					"..xxxxxx..x.x...",
					".xx...xxx....xxx",
					"x.xxxx.xxx...xxx",
					"xxxxxxxxxx..xxxx",
					"...xx..x..xxxx..",
					"xx.xx.xxxx..x...",
					"xxxxxx.....x..xx",
					"xx......xxx..xxx",
					"xxxxxxxxxxxxxxxx",
					"xxx.xxx......xx.",
					"........xxxxxxxx",
				},
			},
			output: []string{
				"13......32.2.332",
				"2..766...4223...",
				".7....7...214...",
				"..........44....",
				"456..66.75....42",
				"..6..5....45.432",
				"......34554.34..",
				"..766544...55...",
				"................",
				"...5...556667..5",
				"23222322........",
			},
		},
		{
			input: input{
				w: 26,
				h: 12,
				lines: []string{
					"..........................",
					"..........................",
					"..........................",
					"..........................",
					"..........................",
					"..........................",
					"..........................",
					"..........................",
					"..........................",
					"..........................",
					"..........................",
					"..........................",
				},
			},
			output: []string{
				"..........................",
				"..........................",
				"..........................",
				"..........................",
				"..........................",
				"..........................",
				"..........................",
				"..........................",
				"..........................",
				"..........................",
				"..........................",
				"..........................",
			},
		},
	}

	failureCases := []struct {
		input input
		err   error
	}{
		{
			input: input{
				w: 0,
				h: 9,
				lines: []string{
					"................",
					"................",
					"................",
					"................",
					"................",
					"....x...........",
					"................",
					"................",
					"................",
				},
			},
			err: errors.New(challenge.OUT_OF_RANGE_W),
		},
		{
			input: input{
				w: 16,
				h: 0,
				lines: []string{
					"................",
					"................",
					"................",
					"................",
					"................",
					"....x...........",
					"................",
					"................",
					"................",
				},
			},
			err: errors.New(challenge.OUT_OF_RANGE_H),
		},
		{
			input: input{
				w: 16,
				h: 9,
				lines: []string{
					"................",
					"................",
					"................",
					"................",
					"................",
					"....x...........",
					"................",
					"................",
				},
			},
			err: errors.New(challenge.INVALID_LINES),
		},
		{
			input: input{
				w: 16,
				h: 9,
				lines: []string{
					"................",
					"................",
					"................",
					"................",
					"................",
					"....x...........",
					"................",
					"................",
					"........",
				},
			},
			err: errors.New(challenge.INVALID_LINE),
		},
	}

	for i := range successCases {
		tc := successCases[i]

		t.Run(fmt.Sprintf("success case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.w, tc.input.h, tc.input.lines)
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, tc.output, result)
		})
	}

	for i := range failureCases {
		tc := failureCases[i]

		t.Run(fmt.Sprintf("failure case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.w, tc.input.h, tc.input.lines)
			require.Error(t, err)
			require.Nil(t, result)
			require.Equal(t, tc.err, err)
		})
	}
}
