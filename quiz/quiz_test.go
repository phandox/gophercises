package quiz_test

import (
	. "github.com/phandox/gophercises/quiz"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

const questionsOk = `"what is 1+1?",2
5+5,10`

func TestAnswersCounters(t *testing.T) {
	tests := []struct {
		name      string
		questions string
		ans       string
		ansOk     int
		ansWrong  int
	}{
		{
			"correct and incorrect",
			questionsOk,
			`2
11`,
			1,
			1,
		},
		{
			"no answers",
			questionsOk,
			"",
			0,
			2,
		},
		{
			"no questions",
			"",
			"",
			0,
			0,
		},
		{
			"newline answers",
			questionsOk,
			`

`,
			0,
			2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotOk, gotWrong := Start(strings.NewReader(test.questions), strings.NewReader(test.ans), io.Discard)
			assert.Equal(t, test.ansOk, gotOk)
			assert.Equal(t, test.ansWrong, gotWrong)
		})
	}
}
