package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {

	var answer int
	var CorrectAnswer, WrongAnswers int = 0, 0

	f, err := os.Open("problems.csv")
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
