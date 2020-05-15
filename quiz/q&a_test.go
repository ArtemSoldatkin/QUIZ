package quiz

import (
	"testing"
)

func TestQAInit(t *testing.T) {
	qaTest := QA{}
	qaTest.Init()
	if qaTest.ID == "" {
		t.Errorf("Quiz id is empty")
	}
}

func TestQAEdit(t *testing.T) {
	var answers []*Answer
	ans1 := Answer{Text: "ans1", IsRight: true}
	answers = append(answers, &ans1)
	qaTest := QA{Question: "test question", Answers: answers}

	var newAnswers []*Answer
	qaTest.Edit("new question", newAnswers)
	if qaTest.Question != "new question" || len(qaTest.Answers) != 0 {
		t.Errorf("QA is not edited")
	}

	qaTest.Edit("", newAnswers)
	if qaTest.Question == "" {
		t.Errorf("QA question can't be empty")
	}
}

func TestQAShuffle(t *testing.T) {
	var answers []*Answer
	a1 := Answer{Text: "a1", IsRight: true}
	a2 := Answer{Text: "a2", IsRight: true}
	a3 := Answer{Text: "a3", IsRight: true}
	a4 := Answer{Text: "a4", IsRight: true}
	answers = append(answers, &a1, &a2, &a3, &a4)
	qa := QA{Question: "test question", Answers: answers}

	qa.Shuffle()
	c := 0
	for i := 0; i < len(answers); i++ {
		if qa.Answers[i] == answers[i] {
			c++
		}
	}
	if c == len(answers)-1 {
		t.Errorf("Answers is not shuffled %d", c)
	}
}

func TestQACheckResult(t *testing.T) {
	var answers []*Answer
	a1 := Answer{Text: "a1", IsRight: true}
	a2 := Answer{Text: "a2"}
	a3 := Answer{Text: "a3", IsRight: true}
	a4 := Answer{Text: "a4"}
	answers = append(answers, &a1, &a2, &a3, &a4)
	qa := QA{Question: "test question", Answers: answers}

	result := qa.checkResult([]int{0, 2})
	if !result {
		t.Errorf("Got %t must %t ", result, true)
	}

	result = qa.checkResult([]int{0})
	if result {
		t.Errorf("Got %t must %t ", result, false)
	}

	result = qa.checkResult([]int{0, 3})
	if result {
		t.Errorf("Got %t must %t ", result, false)
	}
}

func TestQASetAnswer(t *testing.T) {
	var answers []*Answer
	a1 := Answer{Text: "a1", IsRight: true}
	a2 := Answer{Text: "a2"}
	a3 := Answer{Text: "a3", IsRight: true}
	a4 := Answer{Text: "a4"}
	answers = append(answers, &a1, &a2, &a3, &a4)
	qa := QA{Question: "test question", Answers: answers}

	result := qa.SetAnswer([]int{0, 2})

	if !result {
		t.Errorf("Got %t must %t ", result, true)
	}
}
