package challenge_005_test

import (
	"errors"
	"fmt"
	"testing"

	challenge "code-challenges-go/challenges/challenge_005"

	"github.com/stretchr/testify/require"
)

func TestChallenge_005(t *testing.T) {
	t.Parallel()

	type input struct {
		T string
	}

	successCases := []struct {
		input  input
		output []string
	}{
		{
			input: input{T: "E"},
			output: []string{
				"### ",
				"#   ",
				"##  ",
				"#   ",
				"### ",
			},
		},
		{
			input: input{T: "MANHATTAN"},
			output: []string{
				"# #  #  ### # #  #  ### ###  #  ### ",
				"### # # # # # # # #  #   #  # # # # ",
				"### ### # # ### ###  #   #  ### # # ",
				"# # # # # # # # # #  #   #  # # # # ",
				"# # # # # # # # # #  #   #  # # # # ",
			},
		},
		{
			input: input{T: "ManhAtTan"},
			output: []string{
				"# #  #  ### # #  #  ### ###  #  ### ",
				"### # # # # # # # #  #   #  # # # # ",
				"### ### # # ### ###  #   #  ### # # ",
				"# # # # # # # # # #  #   #  # # # # ",
				"# # # # # # # # # #  #   #  # # # # ",
			},
		},
		{
			input: input{T: "M@NH@TT@N"},
			output: []string{
				"# # ### ### # # ### ### ### ### ### ",
				"###   # # # # #   #  #   #    # # # ",
				"###  ## # # ###  ##  #   #   ## # # ",
				"# #     # # # #      #   #      # # ",
				"# #  #  # # # #  #   #   #   #  # # ",
			},
		},
	}

	failureCases := []struct {
		input input
		err   error
	}{
		{
			input: input{
				T: "vFSnqyFMDeLxOGvcuKaFAVwNIEAABjtaAoZeOscdZqIIDluAAqSpoZrOsxZmHiJotsthzgzcQeKxmtTdjPkRGfXPTYwerPVhBWHkbJPIZBWtEviVrrpDtYhxSdmaOCsHPDDucToiANkONOdnIKDclAaOAlOEJXvwSzXchbbgVXiIgKomizGnLEPWEJyhAKKbkhgHawps",
			},
			err: errors.New(challenge.OUT_OF_RANGE_N),
		},
	}

	for i := range successCases {
		tc := successCases[i]

		t.Run(fmt.Sprintf("success case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.T)
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, tc.output, result)
		})
	}

	for i := range failureCases {
		tc := failureCases[i]

		t.Run(fmt.Sprintf("failure case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.T)
			require.Error(t, err)
			require.Nil(t, result)
			require.Equal(t, tc.err, err)
		})
	}
}
