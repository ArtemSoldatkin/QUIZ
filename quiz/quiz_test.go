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

func TestQuizFind(t *testing.T) {

	q1 := QA{}
	q2 := QA{}
	q3 := QA{}
	q4 := QA{}
	q1.Init()
	q2.Init()
	q3.Init()
	q4.Init()

	qArr := []*QA{&q1, &q2, &q3, &q4}
	quiz := Quiz{QAS: qArr}

	qa, index := quiz.Find(q2.ID)
	if qa != &q2 || index != 1 {
		t.Errorf("Got qa - %v idx - %d must: qa - %v idx - %d", qa, index, &q2, 1)
	}

	qa, index = quiz.Find("")
	if qa != nil || index != -1 {
		t.Errorf("Got: qa - %v idx - %d must: qa - nil, idx - -1", qa, index)
	}

	qa, index = quiz.Find("wrong_id")
	if qa != nil || index != -1 {
		t.Errorf("Got: qa - %v idx - %d must: qa - nil, idx - -1", qa, index)
	}

}

func TestQuizAddQA(t *testing.T) {
	quiz := Quiz{}
	q1 := QA{}
	quiz.AddQA(&q1)
	if len(quiz.QAS) == 0 {
		t.Errorf("QAS is empty")
	}
	for _, q := range quiz.QAS {
		if q != &q1 {
			t.Errorf("Got: %v must: %v", q, q1)
		}
	}
}

func TestQuizEditQA(t *testing.T) {

	a1 := Answer{}
	q1 := QA{Question: "test"}
	q1.Init()
	quiz := Quiz{QAS: []*QA{&q1}}

	quiz.EditQA(q1.ID, "new", []*Answer{&a1})
	if q1.Question != "new" {
		t.Errorf("Got: %s must: \"new\"", q1.Question)
	}
	if len(q1.Answers) == 0 {
		t.Errorf("Answers is empty")
	}
	for _, a := range q1.Answers {
		if a != &a1 {
			t.Errorf("Got: %v must: %v", a, a1)
		}
	}
}

func TestQuizRemoveQA(t *testing.T) {
	q1 := QA{}
	q2 := QA{}
	q3 := QA{}
	q4 := QA{}
	q1.Init()
	q2.Init()
	q3.Init()
	q4.Init()

	qArr := []*QA{&q1, &q2, &q3, &q4}
	quiz := Quiz{QAS: qArr}

	err := quiz.RemoveQA(q2.ID)
	if err != nil {
		t.Errorf("Can't find q&a")
	}

	if len(quiz.QAS) != 3 {
		t.Errorf("Error len, got: %d, must: 3", len(quiz.QAS))
	}

	for _, q := range quiz.QAS {
		if q == &q2 {
			t.Errorf("Element is not removed")
		}
	}

	err = quiz.RemoveQA("")
	if err == nil {
		t.Errorf("Don't shown error with empty id")
	}

	err = quiz.RemoveQA("wrong_id")
	if err == nil {
		t.Errorf("Don't shownn error with wrong id")
	}

}

func TestQuizSetAnswer(t *testing.T) {

	a1 := Answer{IsRight: true}
	a2 := Answer{}

	qa := QA{Answers: []*Answer{&a1, &a2}}
	qa.Init()

	quiz := Quiz{QAS: []*QA{&qa}}

	res, err := quiz.SetAnswer("", []int{})
	if err == nil {
		t.Errorf("Don't shown error with empty id")
	}

	res, err = quiz.SetAnswer(qa.ID, []int{0})
	if !res {
		t.Errorf("Wrong result")
	}

	res, err = quiz.SetAnswer(qa.ID, []int{1})
	if res {
		t.Errorf("Wrong result")
	}
}

func TestQuizGetResults(t *testing.T) {
	a1 := Answer{IsRight: true}
	a2 := Answer{}

	qa := QA{Answers: []*Answer{&a1, &a2}}
	qa.Init()

	quiz := Quiz{QAS: []*QA{&qa}}
	quiz.Init()
	quiz.SetAnswer(qa.ID, []int{0})
	res := quiz.GetResults()
	if res == nil {
		t.Errorf("Empty return")
	}
	if res.Right != 1 && res.Total != 1 {
		t.Errorf("Get right: %d total: %d, must right: 1 total: 1 ", res.Right, res.Total)
	}
}
