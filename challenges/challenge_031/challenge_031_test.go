package challenge_031_test

import (
	"errors"
	"fmt"
	"testing"

	challenge "code-challenges-go/challenges/challenge_031"

	"github.com/stretchr/testify/require"
)

func TestChallenge_031(t *testing.T) {
	t.Parallel()

	type input struct {
		recipe string
	}

	successCases := []struct {
		input  input
		output []string
	}{
		{
			input: input{
				recipe: "1sp 1/ 1bS 1_ 1/ 1bS nl 1( 1sp 1o 1. 1o 1sp 1) nl 1sp 1> 1sp 1^ 1sp 1< nl 2sp 3|",
			},
			output: []string{" /\\_/\\", "( o.o )", " > ^ <", "  |||"},
		},
		{
			input: input{
				recipe: "18sp 1: nl 18sp 1: nl 7sp 2, 9sp 1: 9sp 2, nl 7sp 2: 9sp 1: 9sp 2: nl 2, 5sp 2: 9sp 1: 9sp 2: 5sp 2, nl 2: 5sp 2: 9sp 1: 9sp 2: 5sp 2: nl 1sp 1sQ 2: 1. 3sp 1sQ 2: 1. 6sp 1: 6sp 1. 2: 1sQ 3sp 1. 2: 1sQ nl 4sp 1sQ 2: 1. 2sp 1sQ 2: 1. 2sp 1_ 1/ 1~ 1bS 1_ 2sp 1. 2: 1sQ 2sp 1. 2: 1sQ nl 6sp 1sQ 2: 1. 2sp 3: 1/ 5sp 1bS 3: 2sp 1. 2: 1sQ nl 8sp 1sQ 5: 1( 7sp 1) 5: 1sQ nl 15sp 1bS 1sp 3_ 1sp 1/ nl 9sp 1. 5: 1/ 1` 3sp 1` 1bS 5: 1. nl 7sp 1. 2: 1sQ 3sp 1. 1: 1bS 1o 1sp 1o 1/ 1: 1. 3sp 1sQ 2: 1. nl 5sp 1. 2: 1sQ 3sp 1. 2: 2sp 1: 1sQ 1: 2sp 2: 1. 3sp 1sQ 2: 1. nl 3sp 1. 2: 1sQ 4sp 2: 1sQ 3sp 1sQ 1sp 1sQ 3sp 1sQ 2: 4sp 1sQ 2: 1. nl 2sp 2: 6sp 2: 13sp 2: 6sp 2: nl 2sp 2^ 6sp 2: 13sp 2: 6sp 2^ nl 10sp 2: 13sp 2: nl 10sp 2^ 13sp 2^",
			},
			output: []string{
				"                  :",
				"                  :",
				"       ,,         :         ,,",
				"       ::         :         ::",
				",,     ::         :         ::     ,,",
				"::     ::         :         ::     ::",
				" '::.   '::.      :      .::'   .::'",
				"    '::.  '::.  _/~\\_  .::'  .::'",
				"      '::.  :::/     \\:::  .::'",
				"        ':::::(       ):::::'",
				"               \\ ___ /",
				"         .:::::/`   `\\:::::.",
				"       .::'   .:\\o o/:.   '::.",
				"     .::'   .::  :':  ::.   '::.",
				"   .::'    ::'   ' '   '::    '::.",
				"  ::      ::             ::      ::",
				"  ^^      ::             ::      ^^",
				"          ::             ::",
				"          ^^             ^^",
			},
		},
		{
			input: input{
				recipe: "13sp 2x nl 12sp 1/ 1sp 1. 1| 1_ nl 11sp 1/ 1( 1_ 1) 1_ 1< nl 10sp 1/ 2sp 1( nl 1sp 2( 4_ 1. 1- 1sQ 4sp 1) nl 2sp 2bS 9sp 1/ nl 3sp 1bS 1sQ 1- 1. 1- 1. 1- 1sQ 1` 1/ nl 1sp 1_ 2sp 1bS 6_ 1/ nl 1( 1_ 1) 3sp 1_ 1| 1_ 1bS 1_ nl 16sQ",
			},
			output: []string{
				"             xx",
				"            / .|_",
				"           /(_)_<",
				"          /  (",
				" ((____.-'    )",
				"  \\\\         /",
				"   \\'-.-.-'`/",
				" _  \\______/",
				"(_)   _|_\\_",
				"''''''''''''''''",
			},
		},
		{
			input: input{
				recipe: "nl 1sp 2_ 1( 1sp 1) 1_ nl 1( 6sp 1( 1o 4_ nl 1sp 1| 10sp 1| nl 1sp 1| 6sp 1( 2_ 1/ nl 3sp 1bS 5sp 1/ 3sp 3_ nl 3sp 1/ 5sp 1bS 2sp 1bS 3_ 1/ nl 1sp 1/ 4sp 1^ 4sp 1/ 5sp 1bS nl 1| 3sp 1| 2sp 1| 2_ 1| 1_ 1H 1U 2N 1Y 1sp 1| nl 1| 4sp 1bS 6_ 1) 4_ 1/ nl 1sp 1bS 9sp 1/ nl 3sp 1bS 5sp 1/ 1_ nl 4sp 1| 2sp 1( 1sp 2_ 1) nl 4sp 1( 4_ 1)",
			},
			output: []string{
				"",
				" __( )_",
				"(      (o____",
				" |          |",
				" |      (__/",
				"   \\     /   ___",
				"   /     \\  \\___/",
				" /    ^    /     \\",
				"|   |  |__|_HUNNY |",
				"|    \\______)____/",
				" \\         /",
				"   \\     /_",
				"    |  ( __)",
				"    (____)",
			},
		},
		{
			input: input{
				recipe: "1sp 1[ 1] 1[ 1] 1[ 1] 1sp 1/ 2sQ 1bS 1sp 1[ 1] 1[ 1] 1[ 1] nl 2sp 1| 2: 1| 1sp 1/ 4_ 1bS 1sp 1| 2: 1| nl 2sp 1| 1[ 1] 1| 1_ 1| 4: 1| 1_ 1| 1[ 1] 1| nl 2sp 1| 6: 2_ 6: 1| nl 2sp 1| 5: 1/ 2| 1bS 5: 1| nl 2sp 1| 1: 1# 3: 4| 2: 1# 2: 1| nl 2sp 168 nl 1sp 88 2sp 88 nl 1sp 78 4sp 78",
			},
			output: []string{
				" [][][] /''\\ [][][]",
				"  |::| /____\\ |::|",
				"  |[]|_|::::|_|[]|",
				"  |::::::__::::::|",
				"  |:::::/||\\:::::|",
				"  |:#:::||||::#::|",
				"  8888888888888888",
				" 88888888  88888888",
				" 8888888    8888888",
			},
		},
		{
			input: input{
				recipe: "1# nl",
			},
			output: []string{"#", ""},
		},
		{
			input: input{
				recipe: "10sp 5_ nl 9sp 1/ 4sp 1/ 1bS nl 8sp 1/ 4sp 1/ 2sp 1bS nl 7sp 1/ 4sp 1/ 4sp 1bS nl 6sp 1/ 4sp 1/ 2sp 1/ 1bS 2sp 1bS nl 5sp 1/ 4sp 1/ 2sp 1/ 2sp 1bS 2sp 1bS nl 4sp 1/ 4sp 1/ 2sp 1/ 1bS 3sp 1bS 2sp 1bS nl 3sp 1/ 4sp 1/ 2sp 1/ 2sp 1bS 3sp 1bS 2sp 1bS nl 2sp 1/ 4sp 1/ 2_ 1/ 4_ 1bS 3sp 1bS 2sp 1bS nl 1sp 1/ 14sp 1bS 3sp 1bS 2sp 1bS nl 1/ 16_ 1bS 3sp 1bS 2sp 1bS nl 1bS 21_ 1bS 1sp 1/",
			},
			output: []string{
				"          _____",
				"         /    /\\",
				"        /    /  \\",
				"       /    /    \\",
				"      /    /  /\\  \\",
				"     /    /  /  \\  \\",
				"    /    /  /\\   \\  \\",
				"   /    /  /  \\   \\  \\",
				"  /    /__/____\\   \\  \\",
				" /              \\   \\  \\",
				"/________________\\   \\  \\",
				"\\_____________________\\ /",
			},
		},
		{
			input: input{
				recipe: "6sp 1_ 7sp 1_ nl 5sp 1( 1_ 1bS 5sp 1/ 1_ 1) nl 7sp 2) 3sp 2( nl 5sp 1. 1- 7sQ 1- 1. 2sp nl 1sp 1/ 1^ 1bS 1/ 2sp 1_ 1. 3sp 1_ 1. 2sp 1bS 1/ 1^ 1bS nl 1sp 1bS 1( 3sp 1/ 2_ 1bS 1sp 1/ 2_ 1bS 3sp 1) 1/ nl 2sp 1bS 1, 2sp 1bS 1o 1_ 1/ 1_ 1bS 1o 1_ 1/ 2sp 1, 1/ nl 4sp 1bS 4sp 1( 1_ 1) 4sp 1/ nl 5sp 1` 1- 1. 1sQ 3= 1sQ 1. 1- 1sQ nl 6sp 2_ 1) 1sp 1- 1sp 1( 2_ 3sp nl 5sp 1/ 2sp 1` 3~ 1` 2sp 1bS nl 4sp 1/ 2sp 1/ 5sp 1bS 2sp 1bS nl 4sp 1bS 1sp 1: 7sp 1; 1sp 1/ nl 5sp 1bS 1| 2= 1( 1* 1) 2= 1| 1/ nl 6sp 1: 7sp 1: nl 7sp 1bS 2sp 1| 2sp 1/ nl 5sp 3_ 1) 1= 1| 1= 1( 3_ nl 4sp 1{ 4_ 1/ 1sp 1bS 4_ 1}",
			},
			output: []string{
				"      _       _",
				"     (_\\     /_)",
				"       ))   ((",
				"     .-'''''''-.  ",
				" /^\\/  _.   _.  \\/^\\",
				" \\(   /__\\ /__\\   )/",
				"  \\,  \\o_/_\\o_/  ,/",
				"    \\    (_)    /",
				"     `-.'==='.-'",
				"      __) - (__   ",
				"     /  `~~~`  \\",
				"    /  /     \\  \\",
				"    \\ :       ; /",
				"     \\|==(*)==|/",
				"      :       :",
				"       \\  |  /",
				"     ___)=|=(___",
				"    {____/ \\____}",
			},
		},
		{
			input: input{
				recipe: "15sp 1O 1o 1* 1o 1O 1o 1O 1o nl 10sp 1o 1O 1o 1O 1* 1O 1o 1* 1o 1O 1o 1O 1o 1O 1o 1O 1@ 1o nl 8sp 1o 1O 1o 1O 1o 1O 18 1O 1o 1O 1@ 1o 1@ 1o 1O 1o 1O 1o 1O 1o nl 8sp 1o 1O 1o 1O 1@ 1O 1o 18 1o 1O 2o 1@ 1o 1O 1o 18 1@ 1O 1o 1O nl 7sp 2o 1O 1o 18 1o 1O 1o 1O 1o 1O 1@ 20 1O 1o 1O 1* 1O 1o 1O 1o 1O 1o nl 4sp 2o 1O 1o 1O 1o 1O 1@ 1O 1o 1O 1o 1O 1o 10 1o 1O 1@ 1O 1o 1O 1o 1O 1o 1O 1@ 1O 1o nl 2sp 2o 1O 1o 18 1o 1O 1o 1O 1o 1O 1o 1O 1@ 1O 1* 2o 1O 1o 18 1o 1O 1@ 1O 1o 1O 1o 1O 1o 1O 1o nl 2sp 1o 1O 1o 1O 1o 1O 1o 1* 18 1O 1o 1O 1o 1O 1o 1O 1o 1O 18 1O 1o 1O 1o 1O 1o 1O 1o 18 1o 1O 1o 1O nl 2sp 1* 2o 1O 2o 1O 2o 1O 2o 1O 2o 1@ 1* 2o 1O 2o 1O 18 1o 1O 2o 1O 2o 1* nl 6sp 1O 2sp 1* 2o 1O 1sp 1x 1sp 4X 1* 1O 1o 1* 2sp 1@ 2sp 1* 1o nl 11sp 1O 1o 2sp 1x 4X 1* 2o nl 16sp 2X 1@ 1X 1x nl 16sp 4X nl 16sp 4X nl 15sp 1x 5X 1x nl 13sp 1x 8X 1x nl 9sp 16X 1x nl 35=",
			},
			output: []string{
				"               Oo*oOoOo",
				"          oOoO*Oo*oOoOoOoO@o",
				"        oOoOoO8OoO@o@oOoOoOo",
				"        oOoO@Oo8oOoo@oOo8@OoO",
				"       ooOo8oOoOoO@00OoO*OoOoOo",
				"    ooOoOoO@OoOoOo0oO@OoOoOoO@Oo",
				"  ooOo8oOoOoOoO@O*ooOo8oO@OoOoOoOo",
				"  oOoOoOo*8OoOoOoOoO8OoOoOoOo8oOoO",
				"  *ooOooOooOooOoo@*ooOooO8oOooOoo*",
				"      O  *ooO x XXXX*Oo*  @  *o",
				"           Oo  xXXXX*oo",
				"                XX@Xx",
				"                XXXX",
				"                XXXX",
				"               xXXXXXx",
				"             xXXXXXXXXx",
				"         XXXXXXXXXXXXXXXXx",
				"===================================",
			},
		},
		{
			input: input{
				recipe: "4sp 2_ nl 2sp 1. 1^ 1o 1sp 1~ 1bS nl 1sp 1Y 1sp 1/ 1sQ 1~ 1) 1sp 1} 6sp 5_ nl 1sp 1l 1/ 2sp 1/ 1sp 1/ 4sp 1, 1- 1~ 5sp 2~ 2- 1. 1, 1_ nl 4sp 1( 1sp 1( 4sp 1/ 2sp 1~ 1- 1. 1_ 9sp 1^ 1. nl 5sp 1bS 1sp 1sQ 2- 1sQ 2- 1. 4sp 1sQ 1- 1. 1_ 7sp 1bS nl 6sp 1sQ 1- 1. 8_ 5sp 1~ 2- 1. 1, 2_ 1sp 1^ 1. nl 16sp 1sQ 1~ 1r 1- 1. 1, 3_ 1. 1- 1sQ 1- 1. 1sp 1^ 1. nl 17sp 1Y 1I 4sp 1bS 6sp 1~ 1- 1. 1bS nl 17sp 2| 5sp 1bS 8sp 1` 1bS nl 17sp 2| 5sp 2/ nl 17sp 2| 4sp 2/ nl 17sp 1( 1) 3sp 2/ nl 17sp 2| 2sp 2/ 4sp 1~ 1F 1r 1a 1n nl 17sp 2| 1sp 1( 1sp 1c nl 4sp 3_ 1. 1_ 1sp 2_ 2sp 3_ 1I 1| 2_ 1` 2- 2_ 1. 1_ 1sp 2_ 2sp 1_ nl 2sp 1sQ 1~ 5sp 1~ 2sp 1sQ 1~ 3sp 2: 2sp 2~ 2sQ 4sp 1~ 2sp 2~ nl 17sp 2: nl 17sp 1. 1: nl 18sp 1. nl 2sp 32~",
			},
			output: []string{
				"    __",
				"  .^o ~\\",
				" Y /'~) }      _____",
				" l/  / /    ,-~     ~~--.,_",
				"    ( (    /  ~-._         ^.",
				"     \\ '--'--.    '-._       \\",
				"      '-.________     ~--.,__ ^.",
				"                '~r-.,___.-'-. ^.",
				"                 YI    \\      ~-.\\",
				"                 ||     \\        `\\",
				"                 ||     //",
				"                 ||    //",
				"                 ()   //",
				"                 ||  //    ~Fran",
				"                 || ( c",
				"    ___._ __  ___I|__`--__._ __  _",
				"  '~     ~  '~   ::  ~~''    ~  ~~",
				"                 ::",
				"                 .:",
				"                  .",
				"  ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~",
			},
		},
	}

	failureCases := []struct {
		input input
		err   error
	}{
		{
			input: input{
				recipe: "1/nl",
			},
			err: errors.New(challenge.OUT_OF_RANGE_RECIPE),
		},
		{
			input: input{
				recipe: "1sp \" 1/ 1bS 1_ 1/ 1bS nl 1( 1sp 1o 1. 1o 1sp 1) nl 1sp 1> 1sp 1^ 1sp 1< nl 2sp 3|",
			},
			err: errors.New(challenge.INVALID_RECIPE),
		},
		{
			input: input{
				recipe: "1sp 1/ 1bS 1_ 1/ 1bS 1( 1sp 1o 1. 1o 1sp 1) 1sp 1> 1sp 1^ 1sp 1< 2sp 3|",
			},
			err: errors.New(challenge.INVALID_RECIPE),
		},
	}

	for i := range successCases {
		tc := successCases[i]

		t.Run(fmt.Sprintf("success case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.recipe)
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, tc.output, result)
		})
	}

	for i := range failureCases {
		tc := failureCases[i]

		t.Run(fmt.Sprintf("failure case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.recipe)
			require.Error(t, err)
			require.Nil(t, result)
			require.Equal(t, tc.err, err)
		})
	}
}
