package quiz

import "testing"

func TestCalculateResult(t *testing.T) {

	qa1 := QA{IsAnswered: true, Result: true, IsChanged: true}
	qa2 := QA{Result: true}
	qa3 := QA{IsAnswered: true}

	qRes := QuizResult{}
	err := qRes.CalculateResult([]*QA{&qa1, &qa2, &qa3})
	if err == nil {
		t.Errorf("Should return error: Quiz is not completed")
	}

	err = qRes.CalculateResult([]*QA{&qa1, &qa3})
	if err != nil {
		t.Errorf("Shoul return nil")
	} else {
		if qRes.Total != 2 && qRes.Right != 1 && qRes.Changed != 1 {
			t.Errorf("Total got: %d must: %d\nRight got: %d must: %d\nChanged: got: %d must: %d", qRes.Total, 2, qRes.Right, 1, qRes.Changed, 1)
		}
	}

}
