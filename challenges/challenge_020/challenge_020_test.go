package challenge_020_test

import (
	"errors"
	"fmt"
	"testing"

	challenge "code-challenges-go/challenges/challenge_020"

	"github.com/stretchr/testify/require"
)

func TestChallenge_020(t *testing.T) {
	t.Parallel()

	type input struct {
		N       int
		message string
	}

	successCases := []struct {
		input  input
		output string
	}{
		{
			input: input{
				N:       1,
				message: "ghibcadef",
			},
			output: "abcdefghi",
		},
		{
			input: input{
				N:       -1,
				message: "hello world",
			},
			output: "worlelhlo d",
		},
		{
			input: input{
				N:       5,
				message: "hitoeplmu eneicldts aide  tsxt ",
			},
			output: "this is a mutliple encoded text",
		},
		{
			input: input{
				N:       -6,
				message: "hello worlds",
			},
			output: "hrlellwo ods",
		},
		{
			input: input{
				N:       3,
				message: " rius lorem. Duis risus nunc, condimentum at metun lacinia id. Pellentebortis. Suspendttis sed , maxis ornare nipulvinar. In v aliquam erat maximus bibenetus neque, tempus lovarius ipsnare vel. Donec , vitae sx enim. Sed vitaes sed nei ipFusces t. e at sum. Alt nibhgittidisse eu eteger id cursumque vel dui et libs.Maecenash. Suspendisse tristiqueeu condcondimentum atec orDui sitipsuorLem m dolteger quismus eget i ssim lacuss. Suspum feron arcu idvinar id eula elit in effiuspenlor. in blandem solm ne i psuc lorlicitudit ut acSIn luctus vcitur vae pulat arcu ferment maximus. Integerendisse hendrim. Inmentum nibh non dum.  amet, tur adlit. Fusceci pretium iacsi ut felibm neque, quis dignis orligsx nec sagi aliquam do maximuaodo nulla. isi quis, iquam esdu, npretium comMauris as. Ins elitque a mattittis. Morbi volutpat eroegestas irit vel ante ac dignisss nes scing elitconsecteoripi. Quisque msagiel puruuli mollis n enim est, ac bibendumissmentum. Ut dictum mi vel luctus rhoncus.tempor id.",
			},
			output: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque metus neque, sagittis sed condimentum at, maximus eget elit. Fusce condimentum nibh non erat maximus bibendum. Duis ornare nisi ut felis aliquam pulvinar. In vel purus nec orci pretium iaculis. Suspendisse hendrerit vel ante ac dignissim. Integer quis mollis nibh. Suspendisse tristique enim est, ac bibendum metus ornare vel. Donec egestas non arcu id maximus. Integer varius ipsum neque, quis dignissim lacus lacinia id. Pellentesque vel dui et libero tempus lobortis. Suspendisse pulvinar id ex nec sagittis. Morbi volutpat ligula at arcu fermentum fermentum. Ut maximus sed neque a mattis.Maecenas dictum mi vel luctus rhoncus. Suspendisse eu ex enim. Sed vitae aliquam dolor. In luctus velit in efficitur varius. Integer id cursus elit, vitae sagittis lorem. Duis risus nunc, condimentum at nisi quis, pretium commodo nulla. Mauris a ipsum nec lorem sollicitudin blandit ut ac est. Fusce at dui ipsum. Aliquam est nibh, tempor id.",
		},
	}

	failureCases := []struct {
		input input
		err   error
	}{
		{
			input: input{
				N:       1,
				message: "",
			},
			err: errors.New(challenge.OUT_OF_RANGE_MESSAGE),
		},
		{
			input: input{
				N:       11,
				message: "ghibcadef",
			},
			err: errors.New(challenge.OUT_OF_RANGE_N),
		},
	}

	for i := range successCases {
		tc := successCases[i]

		t.Run(fmt.Sprintf("success case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.N, tc.input.message)
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, tc.output, result)
		})
	}

	for i := range failureCases {
		tc := failureCases[i]

		t.Run(fmt.Sprintf("failure case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.N, tc.input.message)
			require.Error(t, err)
			require.Equal(t, "", result)
			require.Equal(t, tc.err, err)
		})
	}
}
