package api

import (
	"quiz/quiz"
	"testing"
)

func TestFindQuiz(t *testing.T) {
	q1 := quiz.Quiz{Title: "Q1"}
	q2 := quiz.Quiz{Title: "Q2"}
	q3 := quiz.Quiz{Title: "Q3"}
	q1.Init()
	q2.Init()
	q3.Init()

	QUIZs = append(QUIZs, &q1, &q2, &q3)

	quiz := findQuiz(q1.ID)
	if quiz != &q1 {
		t.Errorf("Found wrong quiz")
	}

	emptyString := findQuiz("")
	if emptyString != nil {
		t.Errorf("Got: %v must: \"\"", emptyString)
	}

	wrongID := findQuiz("wrong id")
	if wrongID != nil {
		t.Errorf("Got: %v must: \"\"", emptyString)
	}

}
