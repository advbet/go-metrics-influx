package influx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SplitWithEscaping(t *testing.T) {
	tests := map[string]struct {
		Value     string
		Separator byte
		Result    []string
	}{
		"No escapes": {
			Value:     "abc,test,123,123",
			Separator: ',',
			Result: []string{
				"abc", "test", "123", "123",
			},
		},
		"With one escape": {
			Value:     "abc,te/st,123/,123",
			Separator: ',',
			Result: []string{
				"abc", "te/st", "123,123",
			},
		},
		"With multiple escapes": {
			Value:     "abc=test=123/=123=/=",
			Separator: '=',
			Result: []string{
				"abc", "test", "123=123", "=",
			},
		},
	}

	for name, test := range tests {
		test := test

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, test.Result, SplitWithEscaping(test.Value, test.Separator))
		})
	}
}
