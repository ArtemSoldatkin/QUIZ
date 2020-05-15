package quiz

import "testing"

func TestQuizInit(t *testing.T) {
	quiz := Quiz{}
	quiz.Init()
	if quiz.ID == "" {
		t.Errorf("ID must not ba an empty string")
	}
}

func TestQuizEdit(t *testing.T) {
	quiz := Quiz{Title: "test"}
	quiz.Edit("new title")
	if quiz.Title != "new title" {
		t.Errorf("Got: %s must: %s", quiz.Title, "new title")
	}
}

func TestQuizShuffle(t *testing.T) {

	q1 := QA{}
	q2 := QA{}
	q3 := QA{}
	q4 := QA{}

	qArr := []*QA{&q1, &q2, &q3, &q4}
	quiz := Quiz{QAS: qArr}

	quiz.Shuffle()
	c := 0
	for i := 0; i < len(qArr); i++ {
		if quiz.QAS[i] == qArr[i] {
			c++
		}
	}
	if c == len(qArr)-1 {
		t.Errorf("Question is not shuffled %d", c)
	}
}
