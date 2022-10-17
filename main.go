package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of questions,answer")
	timeLimit := flag.Int("time", 30, "a time limit for the quiz in seconds")
	flag.Parse()

	var CorrectAnswer int = 0

	f, err := os.Open(*csvFileName)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

problemloop:
	for i, p := range problems {

		fmt.Print("Question", i+1, ": ", p.q, " = ")
		answerCh := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerCh:
			if answer == p.a {
				CorrectAnswer++
			}
		}
	}

	fmt.Printf("\nU have done %d/%d \n", CorrectAnswer, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
