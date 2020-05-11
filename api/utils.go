package api

import "quiz/quiz"

func findQuiz(id string) *quiz.Quiz {
	for _, q := range QUIZs {
		if q.ID == id {
			return q
		}
	}
	return nil
}
