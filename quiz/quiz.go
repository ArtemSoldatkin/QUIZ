package quiz

import (
	"errors"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// Quiz - type of quiz
type Quiz struct {
	ID     string
	Title  string
	QAS    []*QA
	Result *QuizResult
}

// Init - initialize quiz
func (q *Quiz) Init() {
	q.ID = uuid.New().String()
}

// Edit - edit quiz
func (q *Quiz) Edit(title string) {
	q.Title = title
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
func (q Quiz) Find(id string) (*QA, int) {
	if id == "" {
		return nil, -1
	}
	for i, qa := range q.QAS {
		if qa.ID == id {
			return qa, i
		}
	}
	return nil, -1
}

// AddQA - add new q&a to quiz
func (q Quiz) AddQA(qa *QA) {
	q.QAS = append(q.QAS, qa)
}

// EditQA - edit q&a by id
func (q Quiz) EditQA(id string, question string, answers []*Answer) {
	qa, _ := q.Find(id)
	if qa == nil {
		return
	}
	qa.Edit(question, answers)
}

// RemoveQA - remove q&a from q&a list
func (q *Quiz) RemoveQA(id string) error {
	qa, i := q.Find(id)
	if qa == nil {
		return errors.New("Q&A is not found")
	}
	copy(q.QAS[i:], q.QAS[i+1:])
	q.QAS[len(q.QAS)-1] = nil
	q.QAS = q.QAS[:len(q.QAS)-1]
	return nil
}

// SetAnswer - set user's answer to question by id
func (q Quiz) SetAnswer(id string, answers []int) (bool, error) {
	qa, _ := q.Find(id)
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
