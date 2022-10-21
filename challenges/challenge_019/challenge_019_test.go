package challenge_019_test

import (
	"errors"
	"fmt"
	"testing"

	challenge "code-challenges-go/challenges/challenge_019"

	"github.com/stretchr/testify/require"
)

func TestChallenge_019(t *testing.T) {
	t.Parallel()

	type input struct {
		operation          string
		pseudoRandomNumber int
		rotors             []string
		message            string
	}

	successCases := []struct {
		input  input
		output string
	}{
		{
			input: input{
				operation:          "ENCODE",
				pseudoRandomNumber: 4,
				rotors: []string{
					"BDFHJLCPRTXVZNYEIWGAKMUSQO",
					"AJDKSIRUXBLHWTMCQGZNPYFVOE",
					"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
				},
				message: "AAA",
			},
			output: "KQF",
		},
		{
			input: input{
				operation:          "DECODE",
				pseudoRandomNumber: 4,
				rotors: []string{
					"BDFHJLCPRTXVZNYEIWGAKMUSQO",
					"AJDKSIRUXBLHWTMCQGZNPYFVOE",
					"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
				},
				message: "KQF",
			},
			output: "AAA",
		},
		{
			input: input{
				operation:          "ENCODE",
				pseudoRandomNumber: 25,
				rotors: []string{
					"BDFHJLCPRTXVZNYEIWGAKMUSQO",
					"AJDKSIRUXBLHWTMCQGZNPYFVOE",
					"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
				},
				message: "AAA",
			},
			output: "OZN",
		},
		{
			input: input{
				operation:          "DECODE",
				pseudoRandomNumber: 25,
				rotors: []string{
					"BDFHJLCPRTXVZNYEIWGAKMUSQO",
					"AJDKSIRUXBLHWTMCQGZNPYFVOE",
					"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
				},
				message: "OZN",
			},
			output: "AAA",
		},
		{
			input: input{
				operation:          "ENCODE",
				pseudoRandomNumber: 4,
				rotors: []string{
					"BDFHJLCPRTXVZNYEIWGAKMUSQO",
					"AJDKSIRUXBLHWTMCQGZNPYFVOE",
					"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
				},
				message: "XXX",
			},
			output: "NVA",
		},
		{
			input: input{
				operation:          "DECODE",
				pseudoRandomNumber: 4,
				rotors: []string{
					"BDFHJLCPRTXVZNYEIWGAKMUSQO",
					"AJDKSIRUXBLHWTMCQGZNPYFVOE",
					"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
				},
				message: "NVA",
			},
			output: "XXX",
		},
		{
			input: input{
				operation:          "ENCODE",
				pseudoRandomNumber: 4,
				rotors: []string{
					"BDFHJLCPRTXVZNYEIWGAKMUSQO",
					"AJDKSIRUXBLHWTMCQGZNPYFVOE",
					"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
				},
				message: "YYY",
			},
			output: "VAK",
		},
		{
			input: input{
				operation:          "DECODE",
				pseudoRandomNumber: 4,
				rotors: []string{
					"BDFHJLCPRTXVZNYEIWGAKMUSQO",
					"AJDKSIRUXBLHWTMCQGZNPYFVOE",
					"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
				},
				message: "VAK",
			},
			output: "YYY",
		},
		{
			input: input{
				operation:          "ENCODE",
				pseudoRandomNumber: 7,
				rotors: []string{
					"BDFHJLCPRTXVZNYEIWGAKMUSQO",
					"AJDKSIRUXBLHWTMCQGZNPYFVOE",
					"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
				},
				message: "WEATHERREPORTWINDYTODAY",
			},
			output: "ALWAURKQEQQWLRAWZHUYKVN",
		},
		{
			input: input{
				operation:          "DECODE",
				pseudoRandomNumber: 9,
				rotors: []string{
					"BDFHJLCPRTXVZNYEIWGAKMUSQO",
					"AJDKSIRUXBLHWTMCQGZNPYFVOE",
					"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
				},
				message: "PQSACVVTOISXFXCIAMQEM",
			},
			output: "EVERYONEISWELCOMEHERE",
		},
		{
			input: input{
				operation:          "ENCODE",
				pseudoRandomNumber: 9,
				rotors: []string{
					"BDFHJLCPRTXVZNYEIWGAKMUSQO",
					"AJDKSIRUXBLHWTMCQGZNPYFVOE",
					"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
				},
				message: "EVERYONEISWELCOMEHERE",
			},
			output: "PQSACVVTOISXFXCIAMQEM",
		},
		{
			input: input{
				operation:          "ENCODE",
				pseudoRandomNumber: 9,
				rotors: []string{
					"BDFHJLCPRTXVZNYEIWGAKMUSQO",
					"AJDKSIRUXBLHWTMCQGZNPYFVOE",
					"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
				},
				message: "EVERYONEISWELCOMEHEREEVERYONEISWELCOMEHERE",
			},
			output: "PQSACVVTOISXFXCIAMQEMDZIXFJJSTQIENEFQXVZYV",
		},
		{
			input: input{
				operation:          "DECODE",
				pseudoRandomNumber: 5,
				rotors: []string{
					"BDFHJLCPRTXVZNYEIWGAKMUSQO",
					"AJDKSIRUXBLHWTMCQGZNPYFVOE",
					"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
				},
				message: "XPCXAUPHYQALKJMGKRWPGYHFTKRFFFNOUTZCABUAEHQLGXREZ",
			},
			output: "THEQUICKBROWNFOXJUMPSOVERALAZYSPHINXOFBLACKQUARTZ",
		},
	}

	failureCases := []struct {
		input input
		err   error
	}{
		{
			input: input{
				operation:          "ENCODE",
				pseudoRandomNumber: 26,
				rotors: []string{
					"BDFHJLCPRTXVZNYEIWGAKMUSQO",
					"AJDKSIRUXBLHWTMCQGZNPYFVOE",
					"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
				},
				message: "AAA",
			},
			err: errors.New(challenge.OUT_OF_RANGE_PSEUDO_RANDOM_NUMBER),
		},
		{
			input: input{
				operation:          "ENCODE",
				pseudoRandomNumber: 4,
				rotors: []string{
					"BDFHJLCPRTXVZNYEIWGAKMUSQO",
					"AJDKSIRUXBLHWTMCQGZNPYFVOE",
				},
				message: "AAA",
			},
			err: errors.New(challenge.OUT_OF_RANGE_ROTORS),
		},
		{
			input: input{
				operation:          "ENCODE",
				pseudoRandomNumber: 4,
				rotors: []string{
					"BDFHJLCPRTXVZNYEIWGAKMUSQO",
					"AJDKSIRUXBLHWTMCQGZNPYFVOE",
					"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
				},
				message: "aaa",
			},
			err: errors.New(challenge.INVALID_MESSAGE),
		},
		{
			input: input{
				operation:          "ENCODE",
				pseudoRandomNumber: 4,
				rotors: []string{
					"BBBBBBCPRTXVZNYEIWGAKMUSQO",
					"AJDKSIRUXBLHWTMCQGZNPYFVOE",
					"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
				},
				message: "AAA",
			},
			err: errors.New(challenge.INVALID_ROTORS),
		},
	}

	for i := range successCases {
		tc := successCases[i]

		t.Run(fmt.Sprintf("success case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.operation, tc.input.pseudoRandomNumber, tc.input.rotors, tc.input.message)
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, tc.output, result)
		})
	}

	for i := range failureCases {
		tc := failureCases[i]

		t.Run(fmt.Sprintf("failure case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.operation, tc.input.pseudoRandomNumber, tc.input.rotors, tc.input.message)
			require.Error(t, err)
			require.Equal(t, "", result)
			require.Equal(t, tc.err, err)
		})
	}
}
