package main

import (
	"flag"
	"fmt"
	"github.com/phandox/gophercises/quiz"
	"os"
)

func main() {
	qFile := flag.String("qFile", "", "file path to csv file with questions")
	flag.Parse()

	if len(*qFile) == 0 {
		panic("path can't be empty")
	}
	f, err := os.Open(*qFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	ok, wrong := quiz.Start(f, os.Stdin, os.Stdout)
	fmt.Printf("\nQuiz results: Correct/Wrong: %d/%d", ok, wrong)
}
