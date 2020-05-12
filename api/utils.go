package api

import "quiz/quiz"

func findQuiz(id string) *quiz.Quiz {
	if id == "" {
		return nil
	}
	for _, q := range QUIZs {
		if q.ID == id {
			return q
		}
	}
	return nil
}
