package challenge_029_test

import (
	"errors"
	"fmt"
	"testing"

	challenge "code-challenges-go/challenges/challenge_029"

	"github.com/stretchr/testify/require"
)

func TestChallenge_029(t *testing.T) {
	t.Parallel()

	type input struct {
		lengthOfLine, N int
		entries         []string
	}

	successCases := []struct {
		input  input
		output []string
	}{
		{
			input: input{
				lengthOfLine: 40,
				N:            5,
				entries: []string{
					"Title1 4",
					">Subtitle1 5",
					">>Subsubtitle1 5",
					">Subtitle2 6",
					"Title2 10",
				},
			},
			output: []string{
				"1 Title1...............................4",
				"    1 Subtitle1........................5",
				"        1 Subsubtitle1.................5",
				"    2 Subtitle2........................6",
				"2 Title2..............................10",
			},
		},
		{
			input: input{
				lengthOfLine: 30,
				N:            3,
				entries:      []string{"One 5", "Two 50", "AppendixA 100"},
			},
			output: []string{
				"1 One........................5",
				"2 Two.......................50",
				"3 AppendixA................100",
			},
		},
		{
			input: input{
				lengthOfLine: 45,
				N:            18,
				entries: []string{
					"A 1",
					">AA 5",
					">>AAA 8",
					">>>AAAA 8",
					">>>>AAAAA 9",
					">>AAB 10",
					">>>AABA 12",
					">>>>AABAA 12",
					">>>>AABAB 13",
					">>>>>AABABA 14",
					">AB 15",
					">>ABA 20",
					">>ABB 25",
					">>>ABBA 26",
					">>>>ABBAA 27",
					">>>>>ABBAAA 28",
					">AC 29",
					"B 5005",
				},
			},
			output: []string{
				"1 A.........................................1",
				"    1 AA....................................5",
				"        1 AAA...............................8",
				"            1 AAAA..........................8",
				"                1 AAAAA.....................9",
				"        2 AAB..............................10",
				"            1 AABA.........................12",
				"                1 AABAA....................12",
				"                2 AABAB....................13",
				"                    1 AABABA...............14",
				"    2 AB...................................15",
				"        1 ABA..............................20",
				"        2 ABB..............................25",
				"            1 ABBA.........................26",
				"                1 ABBAA....................27",
				"                    1 ABBAAA...............28",
				"    3 AC...................................29",
				"2 B......................................5005",
			},
		},
		{
			input: input{
				lengthOfLine: 50,
				N:            13,
				entries: []string{
					"Sudamerica 1",
					">Argentina 5",
					">>BuenosAires 8",
					">>Cordoba 10",
					">Brasil 15",
					">>SaoPaulo 20",
					">>Fortaleza 25",
					"Asia 30",
					">Japan 32",
					">>Yokohama 35",
					">>Tokio 40",
					">Iran 42",
					">>Teheran 45",
				},
			},
			output: []string{
				"1 Sudamerica.....................................1",
				"    1 Argentina..................................5",
				"        1 BuenosAires............................8",
				"        2 Cordoba...............................10",
				"    2 Brasil....................................15",
				"        1 SaoPaulo..............................20",
				"        2 Fortaleza.............................25",
				"2 Asia..........................................30",
				"    1 Japan.....................................32",
				"        1 Yokohama..............................35",
				"        2 Tokio.................................40",
				"    2 Iran......................................42",
				"        1 Teheran...............................45",
			},
		},
	}

	failureCases := []struct {
		input input
		err   error
	}{
		{
			input: input{
				lengthOfLine: 29,
				N:            5,
				entries: []string{
					"Title1 4",
					">Subtitle1 5",
					">>Subsubtitle1 5",
					">Subtitle2 6",
					"Title2 10",
				},
			},
			err: errors.New(challenge.OUT_OF_RANGE_LENGTH_OF_LINE),
		},
		{
			input: input{
				lengthOfLine: 40,
				N:            31,
				entries: []string{
					"Title1 4",
					">Subtitle1 5",
					">>Subsubtitle1 5",
					">Subtitle2 6",
					"Title2 10",
				},
			},
			err: errors.New(challenge.OUT_OF_RANGE_N),
		},
		{
			input: input{
				lengthOfLine: 40,
				N:            5,
				entries: []string{
					"Title1 4",
					">Subtitle1 5",
					">>Subsubtitle1 5",
					">Subtitle2 6",
				},
			},
			err: errors.New(challenge.INVALID_ENTRIES),
		},
	}

	for i := range successCases {
		tc := successCases[i]

		t.Run(fmt.Sprintf("success case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.lengthOfLine, tc.input.N, tc.input.entries)
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, tc.output, result)
		})
	}

	for i := range failureCases {
		tc := failureCases[i]

		t.Run(fmt.Sprintf("failure case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.lengthOfLine, tc.input.N, tc.input.entries)
			require.Error(t, err)
			require.Nil(t, result)
			require.Equal(t, tc.err, err)
		})
	}
}
