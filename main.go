package main

import "fmt"

func main() {

	question := []string{"2 + 2", "3 + 3", "4 + 4"}
	answerArray := []int{4, 6, 8}
	var answer int
	var CorrectAnswer int = 0

	for i := 0; i < len(question); i++ {

		fmt.Print("\nQuestion ", i, ": ", question[i], " = ")
		fmt.Scan(&answer)

		if answer == answerArray[i] {
			CorrectAnswer++
		}
	}

	fmt.Printf("\nU have done %d/%d \n", CorrectAnswer, len(question))
}
