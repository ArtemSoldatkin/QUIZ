package api

import (
	"errors"
	"quiz/quiz"
)

func findQuiz(id string) (*quiz.Quiz, int) {
	if id == "" {
		return nil, -1
	}
	for i, q := range QUIZs {
		if q.ID == id {
			return q, i
		}
	}
	return nil, -1
}

func removeQuiz(id string) error {
	_, i := findQuiz(id)
	if i == -1 {
		return errors.New("Quiz is not found")
	}
	copy(QUIZs[i:], QUIZs[i+1:])
	QUIZs[len(QUIZs)-1] = nil
	QUIZs = QUIZs[:len(QUIZs)-1]
	return nil
}
