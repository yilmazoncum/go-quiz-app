package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of questions,answer")
	flag.Parse()

	var answer int
	var CorrectAnswer, WrongAnswers int = 0, 0

	f, err := os.Open(*csvFileName)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)

	for i := 0; ; i++ {

		questions, err := r.Read()

		if err == io.EOF {
			fmt.Println("Questions Done :)")
			break
		} else if err != nil {
			log.Fatal(err)
		}

		fmt.Print("Question", i, ": ", questions[0], " = ")
		fmt.Scan(&answer)

		x, err := strconv.Atoi(questions[1])
		if answer == x {
			CorrectAnswer++
		} else {
			WrongAnswers++
		}

	}

	fmt.Printf("\nU have done %d/%d \n", CorrectAnswer, WrongAnswers+CorrectAnswer)
}
