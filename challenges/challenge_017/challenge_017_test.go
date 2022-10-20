package challenge_017_test

import (
	"errors"
	"fmt"
	"testing"

	challenge "code-challenges-go/challenges/challenge_017"

	"github.com/stretchr/testify/require"
)

func TestChallenge_017(t *testing.T) {
	t.Parallel()

	type input struct {
		W, H        int
		t1, t2, t3  float64
		pictureRows []string
	}

	successCases := []struct {
		input  input
		output []string
	}{
		{
			input: input{
				W:  5,
				H:  5,
				t1: 1,
				t2: 2,
				t3: 3,
				pictureRows: []string{
					"A.... .A...",
					"..... .....",
					"..... .....",
					"..... .....",
					"..... .....",
				},
			},
			output: []string{
				"..A..",
				".....",
				".....",
				".....",
				".....",
			},
		},
		{
			input: input{
				W:  5,
				H:  5,
				t1: 1,
				t2: 2,
				t3: 3,
				pictureRows: []string{
					"A.... .....",
					"..... A....",
					"..... .....",
					"..... .....",
					"..... .....",
				},
			},
			output: []string{
				".....",
				".....",
				"A....",
				".....",
				".....",
			},
		},
		{
			input: input{
				W:  5,
				H:  5,
				t1: 1,
				t2: 2,
				t3: 3,
				pictureRows: []string{
					"A.... .....",
					"..... .A...",
					"..... .....",
					"..... .....",
					"..... .....",
				},
			},
			output: []string{
				".....",
				".....",
				"..A..",
				".....",
				".....",
			},
		},
		{
			input: input{
				W:  5,
				H:  5,
				t1: 1,
				t2: 2,
				t3: 3,
				pictureRows: []string{
					"..... .....",
					"..... .A...",
					"..A.. .....",
					"..... .....",
					"..... .....",
				},
			},
			output: []string{
				"A....",
				".....",
				".....",
				".....",
				".....",
			},
		},
		{
			input: input{
				W:  6,
				H:  6,
				t1: 1,
				t2: 5,
				t3: 6,
				pictureRows: []string{
					"A..... ....A.",
					"...... ......",
					"...... ......",
					"...... ......",
					"...... ......",
					"...... ......",
				},
			},
			output: []string{
				".....A",
				"......",
				"......",
				"......",
				"......",
				"......",
			},
		},
		{
			input: input{
				W:  6,
				H:  6,
				t1: 1,
				t2: 3,
				t3: 5,
				pictureRows: []string{
					"A..... .A....",
					"...... B.....",
					"B..... ......",
					"...... ......",
					"...... ......",
					"...... ......",
				},
			},
			output: []string{
				"B.A...",
				"......",
				"......",
				"......",
				"......",
				"......",
			},
		},
		{
			input: input{
				W:  6,
				H:  6,
				t1: 1,
				t2: 6,
				t3: 11,
				pictureRows: []string{
					"..H... ......",
					"...... ..H...",
					"E...G. .E.G..",
					"...... ..F...",
					"..F... ......",
					"...... ......",
				},
			},
			output: []string{
				"......",
				"......",
				"..E...",
				"......",
				"......",
				"......",
			},
		},
		{
			input: input{
				W:  5,
				H:  5,
				t1: 0,
				t2: 1255,
				t3: 9999,
				pictureRows: []string{
					"..... .....",
					".A... .A...",
					"..... .....",
					"...D. ...D.",
					"..... .....",
				},
			},
			output: []string{
				".....",
				".A...",
				".....",
				"...D.",
				".....",
			},
		},
		{
			input: input{
				W:  20,
				H:  20,
				t1: 25,
				t2: 75,
				t3: 100,
				pictureRows: []string{
					".................O.. G...................",
					".....N...........U.. ...............W....",
					".............L.R.... ...................C",
					".................... ...E................",
					"..........Z..V.H.... ..............K.....",
					"................X... ...........T........",
					".............P...... ............A.......",
					".............A...... .....P...FLI......N.",
					".Q.............T.... ....................",
					"..................F. ........D...........",
					".................... ......S..Y.........M",
					"......K............W .........B....Z.....",
					"...............Y.... ....................",
					"..............S..... ....V.............J.",
					"...........JE......D .........O..........",
					"...M................ ..X...........U.....",
					"......B..G...C....I. ....................",
					".................... ....................",
					".................... ..Q................R",
					".................... .......H............",
				},
			},
			output: []string{
				"..................K.",
				"....................",
				".......I............",
				".........T..........",
				"....................",
				"...........A........",
				"..D.F...............",
				".P..................",
				"..S.......B.........",
				"......Y.L...........",
				"....................",
				"....................",
				"....................",
				"....................",
				"................Z...",
				"....................",
				"....................",
				"....................",
				"....................",
				"....................",
			},
		},
		{
			input: input{
				W:  10,
				H:  10,
				t1: 100,
				t2: 200,
				t3: 300,
				pictureRows: []string{
					"A......... .A........",
					"B......... ..B.......",
					"C......... ...C......",
					"D......... ....D.....",
					"E......... .....E....",
					".........F ........F.",
					".........G .......G..",
					".........H ......H...",
					".........I .....I....",
					".........J ....J.....",
				},
			},
			output: []string{
				"..A.......",
				"....B.....",
				"......C...",
				"........D.",
				"..........",
				".......F..",
				".....G....",
				"...H......",
				".I........",
				"..........",
			},
		},
	}

	failureCases := []struct {
		input input
		err   error
	}{
		{
			input: input{
				W:  0,
				H:  5,
				t1: 1,
				t2: 2,
				t3: 3,
				pictureRows: []string{
					"A.... .A...",
					"..... .....",
					"..... .....",
					"..... .....",
					"..... .....",
				},
			},
			err: errors.New(challenge.OUT_OF_RANGE_W),
		},
		{
			input: input{
				W:  5,
				H:  21,
				t1: 1,
				t2: 2,
				t3: 3,
				pictureRows: []string{
					"A.... .A...",
					"..... .....",
					"..... .....",
					"..... .....",
					"..... .....",
				},
			},
			err: errors.New(challenge.OUT_OF_RANGE_H),
		},
		{
			input: input{
				W:  5,
				H:  5,
				t1: 3,
				t2: 2,
				t3: 3,
				pictureRows: []string{
					"A.... .A...",
					"..... .....",
					"..... .....",
					"..... .....",
					"..... .....",
				},
			},
			err: errors.New(challenge.OUT_OF_RANGE_T1),
		},
		{
			input: input{
				W:  5,
				H:  5,
				t1: 1,
				t2: 4,
				t3: 3,
				pictureRows: []string{
					"A.... .A...",
					"..... .....",
					"..... .....",
					"..... .....",
					"..... .....",
				},
			},
			err: errors.New(challenge.OUT_OF_RANGE_T2),
		},
		{
			input: input{
				W:  5,
				H:  5,
				t1: 1,
				t2: 2,
				t3: 10001,
				pictureRows: []string{
					"A.... .A...",
					"..... .....",
					"..... .....",
					"..... .....",
					"..... .....",
				},
			},
			err: errors.New(challenge.OUT_OF_RANGE_T3),
		},
		{
			input: input{
				W:  5,
				H:  5,
				t1: 1,
				t2: 2,
				t3: 3,
				pictureRows: []string{
					"A...... .A.....",
					"..... .....",
					"..... .....",
					"..... .....",
					"..... .....",
				},
			},
			err: errors.New(challenge.INVALID_PICTURE_ROWS),
		},
	}

	for i := range successCases {
		tc := successCases[i]

		t.Run(fmt.Sprintf("success case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.W, tc.input.H, tc.input.t1, tc.input.t2, tc.input.t3, tc.input.pictureRows)
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, tc.output, result)
		})
	}

	for i := range failureCases {
		tc := failureCases[i]

		t.Run(fmt.Sprintf("failure case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.W, tc.input.H, tc.input.t1, tc.input.t2, tc.input.t3, tc.input.pictureRows)
			require.Error(t, err)
			require.Nil(t, result)
			require.Equal(t, tc.err, err)
		})
	}
}
