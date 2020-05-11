package quiz

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// Answer - type of answer
type Answer struct {
	Text    string `json:"text"`
	IsRight bool   `json:"is_right"`
}

// QA - type of Question & Answer
type QA struct {
	ID                            string
	Question                      string    `json:"question"`
	Answers                       []*Answer `json:"answers"`
	IsAnswered, Result, IsChanged bool
}

// Init - initialize question & answers
func (qa *QA) Init() {
	qa.ID = uuid.New().String()
}

// Edit - edit answers
func (qa *QA) Edit(question string, answers []*Answer) {
	if question != "" {
		qa.Question = question
	}
	qa.Answers = answers
}

// Shuffle - shuffle answers in Q&A
func (qa *QA) Shuffle() {
	a := qa.Answers
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
}

// checkResult - compare user answers with the right answers
func (qa QA) checkResult(result []int) (isRight bool) {
	for _, r := range result {
		if r < len(qa.Answers) && qa.Answers[r].IsRight {
			continue
		}
		return false
	}
	return true
}

// SetAnswer - set answer to question
func (qa *QA) SetAnswer(result []int) bool {
	if !qa.IsAnswered {
		qa.IsAnswered = true
	} else if !qa.IsChanged {
		qa.IsChanged = true
	}
	qa.Result = qa.checkResult(result)
	return qa.Result
}
