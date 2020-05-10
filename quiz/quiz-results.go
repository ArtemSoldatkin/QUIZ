package quiz

import "errors"

// QuizResult - quiz result
type QuizResult struct {
	Total, Right, Changed int
}

// CalculateResult - calculate results of quiz
func (q *QuizResult) CalculateResult(qas []*QA) error {
	for _, qa := range qas {
		if !qa.IsAnswered {
			q.Right = 0
			q.Changed = 0
			return errors.New("Quiz is not completed")
		}
		if qa.Result {
			q.Right++
		}
		if qa.IsChanged {
			q.Changed++
		}
	}
	q.Total = len(qas)
	return nil
}
