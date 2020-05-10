package main

import (
	"fmt"
	qz "quiz/quiz"
)

func main() {
	quiz := qz.QA{Question: "Test?", Answers: []*qz.Answer{&qz.Answer{Text: "Nope"}, &qz.Answer{Text: "Yes", IsRight: true}}}
	quiz.Shuffle()
	fmt.Println(quiz)

}
