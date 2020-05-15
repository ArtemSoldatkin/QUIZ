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

	quiz, _ := findQuiz(q1.ID)
	if quiz != &q1 {
		t.Errorf("Found wrong quiz")
	}

	emptyString, _ := findQuiz("")
	if emptyString != nil {
		t.Errorf("Got: %v must: \"\"", emptyString)
	}

	wrongID, _ := findQuiz("wrong id")
	if wrongID != nil {
		t.Errorf("Got: %v must: \"\"", emptyString)
	}

}

func TestRemoveQuiz(t *testing.T) {
	q1 := quiz.Quiz{Title: "Q1"}
	q2 := quiz.Quiz{Title: "Q2"}
	q3 := quiz.Quiz{Title: "Q3"}
	q1.Init()
	q2.Init()
	q3.Init()

	QUIZs = append(QUIZs, &q1, &q2, &q3)

	removeQuiz(q2.ID)

	for _, quiz := range QUIZs {
		if quiz == &q2 {
			t.Errorf("QUIZ is not been deleted")
			break
		}
	}
}
