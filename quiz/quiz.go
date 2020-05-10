package quiz

import (
	"errors"
	"math/rand"
	"time"
)

// Quiz - type of quiz
type Quiz struct {
	Title  string
	QAS    []*QA
	Result *QuizResult
}

// Shuffle - shuffle quiz questions and answers
func (q *Quiz) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(q.QAS), func(i, j int) { q.QAS[i], q.QAS[j] = q.QAS[j], q.QAS[i] })
	for _, qa := range q.QAS {
		qa.Shuffle()
	}
}

// Find - find q&a by id
func (q Quiz) Find(id string) (qa *QA) {
	for _, q := range q.QAS {
		if q.ID == id {
			return q
		}
	}
	return nil
}

// EditQA - edit q&a by id
func (q Quiz) EditQA(id string, question string, answers []*Answer) {
	qa := q.Find(id)
	if qa == nil {
		return
	}
	qa.Edit(question, answers)
}

// SetAnswer - set user's answer to question
func (q Quiz) SetAnswer(id string, answers []int) (bool, error) {
	qa := q.Find(id)
	if qa == nil {
		return false, errors.New("Q&A is not found")
	}
	return qa.SetAnswer(answers), nil
}

// GetResults - return results of quiz if quiz complete
func (q Quiz) GetResults() *QuizResult {
	q.Result.CalculateResult(q.QAS)
	return q.Result
}
