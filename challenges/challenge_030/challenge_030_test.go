package challenge_030_test

import (
	"errors"
	"fmt"
	"testing"

	challenge "code-challenges-go/challenges/challenge_030"

	"github.com/stretchr/testify/require"
)

func TestChallenge_030(t *testing.T) {
	t.Parallel()

	type input struct {
		W, H    int
		diagram []string
	}

	successCases := []struct {
		input  input
		output []string
	}{
		{
			input: input{
				W: 7,
				H: 7,
				diagram: []string{"A  B  C",
					"|  |  |",
					"|--|  |",
					"|  |--|",
					"|  |--|",
					"|  |  |",
					"1  2  3"},
			},
			output: []string{"A2", "B1", "C3"},
		},
		{
			input: input{
				W: 13,
				H: 8,
				diagram: []string{"A  B  C  D  E",
					"|  |  |  |  |",
					"|  |--|  |  |",
					"|--|  |  |  |",
					"|  |  |--|  |",
					"|  |--|  |--|",
					"|  |  |  |  |",
					"1  2  3  4  5"},
			},
			output: []string{"A3", "B5", "C1", "D2", "E4"},
		},
		{
			input: input{
				W: 16,
				H: 14,
				diagram: []string{"F  E  D  C  B  A",
					"|  |--|  |  |  |",
					"|--|  |--|  |--|",
					"|  |--|  |--|  |",
					"|  |  |  |  |--|",
					"|  |--|  |--|  |",
					"|  |  |--|  |  |",
					"|  |  |--|  |--|",
					"|--|  |  |--|  |",
					"|  |  |--|  |  |",
					"|--|  |  |  |--|",
					"|  |--|  |  |  |",
					"|  |  |--|  |  |",
					"0  1  2  3  4  5"},
			},
			output: []string{"F3", "E1", "D0", "C2", "B5", "A4"},
		},
		{
			input: input{
				W: 22,
				H: 18,
				diagram: []string{"P  Q  R  S  T  U  V  W",
					"|  |  |  |  |--|  |  |",
					"|  |  |--|  |  |  |--|",
					"|  |--|  |--|  |  |  |",
					"|--|  |--|  |  |  |--|",
					"|--|  |  |  |  |--|  |",
					"|  |--|  |  |--|  |--|",
					"|  |  |  |--|  |--|  |",
					"|--|  |  |  |--|  |  |",
					"|  |  |--|  |  |  |  |",
					"|  |  |  |--|  |  |--|",
					"|  |  |  |  |--|  |  |",
					"|--|  |  |  |  |  |  |",
					"|--|  |--|  |  |  |--|",
					"|  |--|  |  |--|  |  |",
					"|  |  |--|  |  |  |--|",
					"|--|  |--|  |  |--|  |",
					"1  2  3  4  5  6  7  8"},
			},
			output: []string{"P3", "Q7", "R8", "S5", "T6", "U2", "V4", "W1"},
		},
		{
			input: input{
				W: 28,
				H: 20,
				diagram: []string{"A  B  C  D  E  F  G  H  I  J",
					"|--|  |--|  |--|  |--|  |--|",
					"|  |--|  |--|  |--|  |--|  |",
					"|--|  |--|  |--|  |--|  |--|",
					"|--|  |--|  |--|  |--|  |--|",
					"|  |--|  |--|  |--|  |--|  |",
					"|  |--|  |--|  |--|  |--|  |",
					"|--|  |--|  |--|  |--|  |--|",
					"|--|  |--|  |--|  |--|  |--|",
					"|  |--|  |--|  |--|  |--|  |",
					"|--|  |--|  |--|  |--|  |--|",
					"|  |--|  |--|  |--|  |--|  |",
					"|--|  |--|  |--|  |--|  |--|",
					"|--|  |--|  |--|  |--|  |--|",
					"|--|  |--|  |--|  |--|  |--|",
					"|  |--|  |--|  |--|  |--|  |",
					"|--|  |--|  |--|  |--|  |--|",
					"|--|  |--|  |--|  |--|  |--|",
					"|  |--|  |--|  |--|  |--|  |",
					"0  1  2  3  4  5  6  7  8  9"},
			},
			output: []string{"A1", "B3", "C0", "D5", "E2", "F7", "G4", "H9", "I6", "J8"},
		},
		{
			input: input{
				W: 76,
				H: 23,
				diagram: []string{"~  !  @  #  $  %  ^  &  *  (  )  +  `  1  2  3  4  5  6  7  8  9  0  =  \\  /",
					"|  |--|  |  |--|  |  |--|  |--|  |  |--|  |  |  |--|  |--|  |  |--|  |  |--|",
					"|--|  |--|  |  |  |--|  |--|  |--|  |  |  |--|  |  |--|  |--|  |  |  |--|  |",
					"|  |--|  |--|  |  |  |  |  |--|  |--|  |  |  |  |--|  |--|  |--|  |--|  |--|",
					"|--|  |--|  |  |  |--|  |--|  |--|  |  |  |--|  |--|  |--|  |  |  |--|  |--|",
					"|--|  |  |  |  |--|  |  |--|  |  |  |  |--|  |--|  |--|  |--|  |--|  |--|  |",
					"|  |--|  |  |--|  |--|  |  |--|  |  |--|  |--|  |  |  |--|  |  |--|  |--|  |",
					"|  |  |  |--|  |--|  |--|  |  |  |--|  |--|  |  |--|  |--|  |--|  |--|  |--|",
					"|--|  |  |  |--|  |--|  |--|  |  |  |--|  |--|  |--|  |  |--|  |  |--|  |--|",
					"|  |  |--|  |  |  |  |--|  |  |--|  |  |  |  |  |  |--|  |  |  |--|  |--|  |",
					"|  |  |  |--|  |  |--|  |  |  |  |--|  |  |--|  |--|  |--|  |--|  |--|  |--|",
					"|  |--|  |--|  |  |  |  |  |--|  |--|  |  |  |  |--|  |--|  |--|  |--|  |--|",
					"|--|  |--|  |  |  |--|  |--|  |--|  |  |  |--|  |--|  |--|  |  |  |--|  |--|",
					"|--|  |  |  |  |--|  |  |--|  |  |  |  |--|  |--|  |--|  |--|  |--|  |--|  |",
					"|--|  |--|  |  |  |--|  |--|  |--|  |  |  |--|  |  |--|  |  |  |--|  |  |--|",
					"|  |--|  |  |--|  |--|  |  |--|  |  |--|  |--|  |  |  |--|  |  |--|  |--|  |",
					"|  |--|  |  |--|  |  |  |  |--|  |  |--|  |  |--|  |--|  |--|  |--|  |--|  |",
					"|--|  |  |--|  |  |  |  |--|  |  |--|  |--|  |  |--|  |--|  |--|  |--|  |--|",
					"|--|  |--|  |  |  |--|  |--|  |--|  |  |  |--|  |  |--|  |  |  |--|  |  |--|",
					"|  |--|  |  |--|  |  |--|  |--|  |  |  |--|  |--|  |  |--|  |--|  |--|  |--|",
					"|  |  |  |--|  |  |--|  |  |  |  |--|  |  |--|  |  |--|  |--|  |--|  |--|  |",
					"|--|  |--|  |--|  |--|  |--|  |--|  |--|  |--|  |--|  |--|  |  |  |  |  |--|",
					"a  A  b  B  c  C  d  D  e  E  f  F  g  G  h  H  i  I  j  J  k  K  l  L  m  M"},
			},
			output: []string{"~E",
				"!F",
				"@C",
				"#c",
				"$G",
				"%B",
				"^A",
				"&h",
				"*a",
				"(g",
				")b",
				"+f",
				"`I",
				"1d",
				"2D",
				"3i",
				"4J",
				"5e",
				"6M",
				"7k",
				"8L",
				"9l",
				"0H",
				"=K",
				"\\j",
				"/m"},
		},
	}

	failureCases := []struct {
		input input
		err   error
	}{
		{
			input: input{
				W: 3,
				H: 7,
				diagram: []string{"A  B  C",
					"|  |  |",
					"|--|  |",
					"|  |--|",
					"|  |--|",
					"|  |  |",
					"1  2  3"},
			},
			err: errors.New(challenge.OUT_OF_RANGE_W),
		},
		{
			input: input{
				W: 7,
				H: 3,
				diagram: []string{"A  B  C",
					"|  |  |",
					"|--|  |",
					"|  |--|",
					"|  |--|",
					"|  |  |",
					"1  2  3"},
			},
			err: errors.New(challenge.OUT_OF_RANGE_H),
		},
		{
			input: input{
				W: 7,
				H: 7,
				diagram: []string{"A  B  C",
					"|  |  |",
					"|--|  |",
					"|  |--|",
					"|  |--|",
					"1  2  3"},
			},
			err: errors.New(challenge.INVALID_DIAGRAM),
		},
		{
			input: input{
				W: 7,
				H: 7,
				diagram: []string{"A  B  C",
					"| |  |",
					"|-|  |",
					"| |--|",
					"| |--|",
					"| |  |",
					"1 2  3"},
			},
			err: errors.New(challenge.INVALID_DIAGRAM),
		},
	}

	for i := range successCases {
		tc := successCases[i]

		t.Run(fmt.Sprintf("success case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.W, tc.input.H, tc.input.diagram)
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, tc.output, result)
		})
	}

	for i := range failureCases {
		tc := failureCases[i]

		t.Run(fmt.Sprintf("failure case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.W, tc.input.H, tc.input.diagram)
			require.Error(t, err)
			require.Nil(t, result)
			require.Equal(t, tc.err, err)
		})
	}
}
