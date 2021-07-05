package quiz_test

import (
	. "github.com/phandox/gophercises/quiz"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
	"time"
)

const questionsOk = `"what is 1+1?",2
5+5,10`

func TestAnswersCounters(t *testing.T) {
	tests := []struct {
		name      string
		questions string
		ans       string
		want      Score
	}{
		{
			"correct and incorrect",
			questionsOk,
			`2
11`,
			Score{Ok: 1, Wrong: 1, InTime: true},
		},
		{
			"no answers",
			questionsOk,
			"",
			Score{Wrong: 2, InTime: true},
		},
		{
			"no questions",
			"",
			"",
			Score{InTime: true},
		},
		{
			"newline answers",
			questionsOk,
			`

`,
			Score{Wrong: 2, InTime: true},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			term := Terminal{
				Reader: strings.NewReader(test.ans),
				Writer: io.Discard,
			}
			got := Game(strings.NewReader(test.questions), term, time.Duration(10)*time.Second)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestGameTimeout(t *testing.T) {
	tests := []struct {
		name string
		in   string
		ans  string
		want Score
	}{
		{
			"no answers in time limit",
			questionsOk,
			``,
			Score{
				Ok:     0,
				Wrong:  2,
				InTime: false,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var timeout time.Duration
			term := Terminal{
				Reader: strings.NewReader(test.ans),
				Writer: io.Discard,
			}
			switch test.want.InTime {
			case true:
				timeout = time.Duration(10) * time.Second
			case false:
				timeout = time.Duration(1) * time.Nanosecond
			}
			got := Game(strings.NewReader(test.in), term, timeout)
			assert.Equal(t, test.want, got)

		})
	}
}
