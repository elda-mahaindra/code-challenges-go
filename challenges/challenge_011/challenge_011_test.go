package challenge_011_test

import (
	"errors"
	"fmt"
	"testing"

	challenge "code-challenges-go/challenges/challenge_011"

	"github.com/stretchr/testify/require"
)

func TestChallenge_011(t *testing.T) {
	t.Parallel()

	type input struct {
		mimeTypes, fileNames []string
	}

	successCases := []struct {
		input  input
		output []string
	}{
		{
			input: input{
				mimeTypes: []string{"html text/html", "png image/png", "gif image/gif"},
				fileNames: []string{"animated.gif", "portrait.png", "index.html"},
			},
			output: []string{"image/gif", "image/png", "text/html"},
		},
		{
			input: input{
				mimeTypes: []string{"txt text/plain", "xml text/xml", "flv video/x-flv"},
				fileNames: []string{"image.png", "animated.gif", "script.js", "source.cpp"},
			},
			output: []string{"UNKNOWN", "UNKNOWN", "UNKNOWN", "UNKNOWN"},
		},
		{
			input: input{
				mimeTypes: []string{"wav audio/x-wav", "mp3 audio/mpeg", "pdf application/pdf"},
				fileNames: []string{"a", "a.wav", "b.wav.tmp", "test.vmp3", "pdf", ".pdf", "mp3", "report..pdf", "defaultwav", ".mp3.", "final."},
			},
			output: []string{"UNKNOWN", "audio/x-wav", "UNKNOWN", "UNKNOWN", "UNKNOWN", "application/pdf", "UNKNOWN", "application/pdf", "UNKNOWN", "UNKNOWN", "UNKNOWN"},
		},
		{
			input: input{
				mimeTypes: []string{"png image/png", "TIFF image/TIFF", "css text/css", "TXT text/plain"},
				fileNames: []string{"example.TXT", "referecnce.txt", "strangename.tiff", "resolv.CSS", "matrix.TiFF", "lanDsCape.Png", "extract.cSs"},
			},
			output: []string{"text/plain", "text/plain", "image/TIFF", "text/css", "image/TIFF", "image/png", "text/css"},
		},
	}

	failureCases := []struct {
		input input
		err   error
	}{
		{
			input: input{
				mimeTypes: []string{},
				fileNames: []string{"animated.gif", "portrait.png", "index.html"},
			},
			err: errors.New(challenge.OUT_OF_RANGE_MIME_TYPES),
		},
		{
			input: input{
				mimeTypes: []string{"html text/html", "png image/png", "gif image/gif"},
				fileNames: []string{},
			},
			err: errors.New(challenge.OUT_OF_RANGE_FILE_NAMES),
		},
		{
			input: input{
				mimeTypes: []string{"h-t-m-l text/html", "png image/png", "gif image/gif"},
				fileNames: []string{"animated.gif", "portrait.png", "index.html"},
			},
			err: errors.New(challenge.INVALID_FILE_EXTENSION),
		},
		{
			input: input{
				mimeTypes: []string{"html text/html", "png image/png", "gif image/gifimage/gifimage/gifimage/gifimage/gifimage/gif"},
				fileNames: []string{"animated.gif", "portrait.png", "index.html"},
			},
			err: errors.New(challenge.INVALID_MIME_TYPE),
		},
		{
			input: input{
				mimeTypes: []string{"html text/html", "png image/png", "gif image/gif"},
				fileNames: []string{"ani mated.gif", "portrait.png", "index.html"},
			},
			err: errors.New(challenge.INVALID_FILE_NAME),
		},
	}

	for i := range successCases {
		tc := successCases[i]

		t.Run(fmt.Sprintf("success case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.mimeTypes, tc.input.fileNames)
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, tc.output, result)
		})
	}

	for i := range failureCases {
		tc := failureCases[i]

		t.Run(fmt.Sprintf("failure case %d", i), func(t *testing.T) {
			t.Parallel()

			result, err := challenge.Solution(tc.input.mimeTypes, tc.input.fileNames)
			require.Error(t, err)
			require.Nil(t, result)
			require.Equal(t, tc.err, err)
		})
	}
}
