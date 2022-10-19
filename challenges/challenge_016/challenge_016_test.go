package challenge_016_test

import (
	"errors"
	"fmt"
	"testing"

	challenge "code-challenges-go/challenges/challenge_016"

	"github.com/stretchr/testify/require"
)

func TestChallenge_016(t *testing.T) {
	t.Parallel()

	type input struct {
		n, m                     int
		inputSignals, operations []string
	}

	successCases := []struct {
		input  input
		output []string
	}{
		{
			input: input{
				n: 2,
				m: 3,
				inputSignals: []string{
					"A __---___---___---___---___",
					"B ____---___---___---___---_",
				},
				operations: []string{
					"C AND A B",
					"D OR A B",
					"E XOR A B",
				},
			},
			output: []string{
				"C ____-_____-_____-_____-___",
				"D __-----_-----_-----_-----_",
				"E __--_--_--_--_--_--_--_--_",
			},
		},
		{
			input: input{
				n: 1,
				m: 1,
				inputSignals: []string{
					"A __---___---___---___---___",
				},
				operations: []string{
					"B NAND A A",
				},
			},
			output: []string{
				"B --___---___---___---___---",
			},
		},
		{
			input: input{
				n: 3,
				m: 3,
				inputSignals: []string{
					"CLK _-_-_-_-_-_-_-_-_-_-_-_-_-",
					"IN1 ___---___---___---___---__",
					"IN2 --__--__--__--__--__--__--",
				},
				operations: []string{
					"OUT1 AND CLK IN1",
					"OUT2 AND CLK IN2",
					"OUT3 AND IN1 IN2",
				},
			},
			output: []string{
				"OUT1 ___-_-___-_-___-_-___-_-__",
				"OUT2 _-___-___-___-___-___-___-",
				"OUT3 ____--___-______--___-____",
			},
		},
		{
			input: input{
				n: 3,
				m: 3,
				inputSignals: []string{
					"CLK _-_-_-_-_-_-_-_-_-_-_-_-_-",
					"IN1 ----____----____----____--",
					"IN2 --__--__--__--__--__--__--",
				},
				operations: []string{
					"OUT1 OR CLK IN1",
					"OUT2 OR CLK IN2",
					"OUT3 OR IN1 IN2",
				},
			},
			output: []string{
				"OUT1 ----_-_-----_-_-----_-_---",
				"OUT2 --_---_---_---_---_---_---",
				"OUT3 ------__------__------__--",
			},
		},
		{
			input: input{
				n: 3,
				m: 3,
				inputSignals: []string{
					"CLK _-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_",
					"IN1 __--__--__--__--__--__--__--__--_",
					"IN2 ___---___---___---___---___---___",
				},
				operations: []string{
					"OUT1 XOR IN1 CLK",
					"OUT2 XOR IN2 CLK",
					"OUT3 XOR IN2 IN1",
				},
			},
			output: []string{
				"OUT1 _--__--__--__--__--__--__--__--__",
				"OUT2 _-__-__-__-__-__-__-__-__-__-__-_",
				"OUT3 __-_----_-____-_----_-____-_----_",
			},
		},
		{
			input: input{
				n: 1,
				m: 1,
				inputSignals: []string{
					"IN0 -_--__---___----____-_--__---___",
				},
				operations: []string{
					"OUT OR IN0 IN0",
				},
			},
			output: []string{
				"OUT -_--__---___----____-_--__---___",
			},
		},
		{
			input: input{
				n: 3,
				m: 3,
				inputSignals: []string{
					"CLK _-_-_-_-_-_-_-_-_-_-_-_-_-",
					"IN1 ___---___---___---___---__",
					"IN2 --__--__--__--__--__--__--",
				},
				operations: []string{
					"OUT1 NAND CLK IN1",
					"OUT2 NAND CLK IN2",
					"OUT3 NAND IN1 IN2",
				},
			},
			output: []string{
				"OUT1 ---_-_---_-_---_-_---_-_--",
				"OUT2 -_---_---_---_---_---_---_",
				"OUT3 ----__---_------__---_----",
			},
		},
		{
			input: input{
				n: 3,
				m: 2,
				inputSignals: []string{
					"IN1 --__--__--__--__--__--__--__--__--__",
					"IN2 ____----____----____----____----____",
					"IN3 --------________--------________----",
				},
				operations: []string{
					"OUT1 NOR IN2 IN1",
					"OUT2 NOR IN2 IN3",
				},
			},
			output: []string{
				"OUT1 __--______--______--______--______--",
				"OUT2 ________----____________----________",
			},
		},
		{
			input: input{
				n: 4,
				m: 3,
				inputSignals: []string{
					"A -_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_",
					"B --__--__--__--____----____----______------",
					"C -_-__--__--__---___---___---___----____---",
					"D -----_____-----_____-----_____-----_____--",
				},
				operations: []string{
					"X XNOR A B",
					"Y XNOR B C",
					"Z XNOR C D",
				},
			},
			output: []string{"X -__--__--__--__-_--_-__-_--_-__-_-_--_-_-_",
				"Y -__-_-_-_-_-_-__--_------_--__-____-___---",
				"Z -_-____--_-__--_---_--______--_--------_--"},
		},
		{
			input: input{
				n: 4,
				m: 16,
				inputSignals: []string{
					"ZORGLUB ----____----____----____----____----____--",
					"MEGAMAN --____----____----____----____----____----",
					"ZOLTRON ---___---___------______------______-_-_-_",
					"PEW_PEW -_-_-_-_------_____----____---___--__--__-",
				},
				operations: []string{
					"OUTPUT1 AND ZORGLUB MEGAMAN",
					"OUTPUT2 OR ZORGLUB ZOLTRON",
					"OUTPUT3 XOR ZORGLUB PEW_PEW",
					"OUTPUT4 AND ZORGLUB ZORGLUB",
					"ROGUE_1 OR MEGAMAN MEGAMAN",
					"ROGUE_2 NAND MEGAMAN MEGAMAN",
					"ROGUE_3 NOR PEW_PEW PEW_PEW",
					"ROGUE_4 XNOR PEW_PEW MEGAMAN",
					"SQUAD_1 NAND PEW_PEW MEGAMAN",
					"SQUAD_2 OR ZOLTRON PEW_PEW",
					"SQUAD_3 NOR ZOLTRON PEW_PEW",
					"SQUAD_4 AND ZOLTRON PEW_PEW",
					"MIKADO1 AND MEGAMAN PEW_PEW",
					"MIKADO2 OR MEGAMAN PEW_PEW",
					"MIKADO3 XOR MEGAMAN MEGAMAN",
					"MIKADO4 XNOR ZOLTRON ZOLTRON",
				},
			},
			output: []string{
				"OUTPUT1 --______--______--______--______--______--",
				"OUTPUT2 ----__--------------____------__-----_-_--",
				"OUTPUT3 _-_--_-_____--__---_---_---_--__-__-_--_-_",
				"OUTPUT4 ----____----____----____----____----____--",
				"ROGUE_1 --____----____----____----____----____----",
				"ROGUE_2 __----____----____----____----____----____",
				"ROGUE_3 _-_-_-_-______-----____----___---__--__--_",
				"ROGUE_4 -__-_--_--________-___-___-______-_--_-__-",
				"SQUAD_1 _-----_-__------------_----------_----_--_",
				"SQUAD_2 ---_-_------------_----_------___--_---_--",
				"SQUAD_3 ___-_-____________-____-______---__-___-__",
				"SQUAD_4 -_-___-_-___--_____________---________-___",
				"MIKADO1 -_____-_--____________-__________-____-__-",
				"MIKADO2 ---_-_------------_-------_--------__-----",
				"MIKADO3 __________________________________________",
				"MIKADO4 ------------------------------------------",
			},
		},
	}

	failureCases := []struct {
		input input
		err   error
	}{
		{
			input: input{
				n: 5,
				m: 3,
				inputSignals: []string{
					"A __---___---___---___---___",
					"B ____---___---___---___---_",
				},
				operations: []string{
					"C AND A B",
					"D OR A B",
					"E XOR A B",
				},
			},
			err: errors.New(challenge.OUT_OF_RANGE_N),
		},
		{
			input: input{
				n: 2,
				m: 17,
				inputSignals: []string{
					"A __---___---___---___---___",
					"B ____---___---___---___---_",
				},
				operations: []string{
					"C AND A B",
					"D OR A B",
					"E XOR A B",
				},
			},
			err: errors.New(challenge.OUT_OF_RANGE_M),
		},
		{
			input: input{
				n: 2,
				m: 3,
				inputSignals: []string{
					"A __---___---___---___---___",
					"B ____---___---___---___---_",
					"Z ____---___---___---___---_",
				},
				operations: []string{
					"C AND A B",
					"D OR A B",
					"E XOR A B",
				},
			},
			err: errors.New(challenge.INVALID_INPUT_SIGNALS),
		},
		{
			input: input{
				n: 2,
				m: 3,
				inputSignals: []string{
					"A ..---...---...---...---...",
					"B ____---___---___---___---_",
				},
				operations: []string{
					"C AND A B",
					"D OR A B",
					"E XOR A B",
				},
			},
			err: errors.New(challenge.INVALID_INPUT_SIGNALS),
		},
		{
			input: input{
				n: 2,
				m: 3,
				inputSignals: []string{
					"A __---___---___---___---___",
					"B ____---___---___---___---_",
				},
				operations: []string{
					"C NANO A B",
					"D OR A B",
					"E XOR A B",
				},
			},
			err: errors.New(challenge.INVALID_OPERATIONS),
		},
		{
			input: input{
				n: 2,
				m: 3,
				inputSignals: []string{
					"A __---___---___---___---___",
					"B ____---___---___---___---_",
				},
				operations: []string{
					"C AND A Z",
					"D OR A B",
					"E XOR A B",
				},
			},
			err: errors.New(challenge.INVALID_OPERATIONS),
		},
	}

	for i := range successCases {
		tc := successCases[i]

		t.Run(fmt.Sprintf("success case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.n, tc.input.m, tc.input.inputSignals, tc.input.operations)
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, tc.output, result)
		})
	}

	for i := range failureCases {
		tc := failureCases[i]

		t.Run(fmt.Sprintf("failure case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.n, tc.input.m, tc.input.inputSignals, tc.input.operations)
			require.Error(t, err)
			require.Nil(t, result)
			require.Equal(t, tc.err, err)
		})
	}
}
