package quiz

import (
	"encoding/csv"
	"fmt"
	"io"
)

func Start(src io.Reader, ans io.Reader, out io.Writer) (ok int, wrong int) {
	qt := loadQuestions(src)
	for _, qa := range qt {
		_, err := fmt.Fprintf(out, "%s : ", qa[0])
		if err != nil {
			panic(err)
		}
		a := fetchAnswer(ans)
		switch a == qa[1] {
		case true:
			ok++
		case false:
			wrong++
		}
	}
	return ok, wrong
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
