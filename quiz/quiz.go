package quiz

import (
	"encoding/csv"
	"fmt"
	"io"
	"time"
)

// Terminal wraps Reader and Writer as anonymous members of struct.
// Pass your Reader and Writer instances to interact with user
type Terminal struct {
	io.Reader
	io.Writer
}

//Score stores result of game. InTime tells if all questions were answered in time limit.
type Score struct {
	Ok     int
	Wrong  int
	InTime bool
}

func (s Score) String() string {
	if s.InTime {
		return fmt.Sprintf("\nFinish! Your score(ok/wrong) is: %d/%d", s.Ok, s.Wrong)
	}
	return fmt.Sprintf("\nTimes up! Your score(ok/wrong) is: %d/%d", s.Ok, s.Wrong)
}

//Game starts quiz with questions as 'question,answer' in CSV format and
//counts the time limit for answering.
func Game(questions io.Reader, input Terminal, timeout time.Duration) Score {
	qt := loadQuestions(questions)
	s := Score{Wrong: len(qt), InTime: true}
	done := make(chan struct{})
	go func() {
		start(qt, input.Reader, input.Writer, &s)
		close(done)
	}()
	select {
	case <-done:
		return s
	case <-time.After(timeout):
		s.InTime = false
		return s
	}
}

func start(questions [][]string, ans io.Reader, out io.Writer, s *Score) {
	for _, qa := range questions {
		if !askQuestion(qa, ans, out) {
			continue
		}
		s.Ok++
		s.Wrong--
	}
}

func askQuestion(qa []string, in io.Reader, out io.Writer) bool {
	_, err := fmt.Fprintf(out, "%s : ", qa[0])
	if err != nil {
		panic(err)
	}
	a := fetchAnswer(in)
	return a == qa[1]
}

func fetchAnswer(ans io.Reader) string {
	var a string
	_, err := fmt.Fscanln(ans, &a)
	// newlines and EOF are empty answer
	if err != nil {
		return ""
	}
	return a
}

func loadQuestions(src io.Reader) [][]string {
	r := csv.NewReader(src)
	var qt [][]string
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		qt = append(qt, rec)
	}
	return qt
}
