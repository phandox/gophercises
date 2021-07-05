package main

import (
	"flag"
	"fmt"
	"github.com/phandox/gophercises/quiz"
	"os"
	"time"
)

func main() {
	qFile := flag.String("qFile", "", "file path to csv file with questions")
	duration := flag.Duration("duration", 10*time.Second, "how much time before quiz ends (default 10s)")
	flag.Parse()

	if len(*qFile) == 0 {
		panic("path can't be empty")
	}
	if *duration <= 0 {
		panic("time for quiz must be positive non-zero")
	}

	f, err := os.Open(*qFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	inout := quiz.Terminal{
		Reader: os.Stdin,
		Writer: os.Stdout,
	}
	fmt.Println(quiz.Game(f, inout, *duration))
}
