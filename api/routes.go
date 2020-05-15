package api

import (
	"encoding/json"
	"net/http"
	qz "quiz/quiz"

	"github.com/gorilla/mux"
)

// QUIZs - list of quiz
var QUIZs []*qz.Quiz

// GetQuizList - get quiz list
func GetQuizList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(QUIZs)
}

// CreateQUIZ - create quiz
func CreateQUIZ(w http.ResponseWriter, r *http.Request) {
	var params map[string]string
	json.NewDecoder(r.Body).Decode(&params)
	title := params["title"]
	newQuiz := qz.Quiz{Title: title}
	newQuiz.Init()
	QUIZs = append(QUIZs, &newQuiz)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// GetQuiz - get quiz by id
func GetQuiz(w http.ResponseWriter, r *http.Request) {
	qParam := mux.Vars(r)
	id := qParam["id"]
	quiz, _ := findQuiz(id)
	if quiz == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(quiz)
}

// EditQuiz - edit quiz (title)
func EditQuiz(w http.ResponseWriter, r *http.Request) {
	var params map[string]string
	json.NewDecoder(r.Body).Decode(&params)
	qParam := mux.Vars(r)
	id := qParam["id"]
	quiz, _ := findQuiz(id)
	if quiz == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}
	quiz.Edit(params["title"])
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// RemoveQuiz - remove quiz from list
func RemoveQuiz(w http.ResponseWriter, r *http.Request) {
	qParam := mux.Vars(r)
	id := qParam["id"]
	err := removeQuiz(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Quiz is not found"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// ShuffleQuiz - shuffle questions and answers in quiz
func ShuffleQuiz(w http.ResponseWriter, r *http.Request) {
	qParam := mux.Vars(r)
	id := qParam["id"]
	quiz, _ := findQuiz(id)
	if quiz == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}
	quiz.Shuffle()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// GetQuestion - get question from quiz
func GetQuestion(w http.ResponseWriter, r *http.Request) {
	qParams := mux.Vars(r)
	quizID := qParams["id"]
	questionID := qParams["qa_id"]
	quiz, _ := findQuiz(quizID)
	if quiz == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}
	qa, _ := quiz.Find(questionID)
	if qa == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(qa)
}

// AddQuestion - add question by id
func AddQuestion(w http.ResponseWriter, r *http.Request) {
	var qa qz.QA
	json.NewDecoder(r.Body).Decode(&qa)
	qParams := mux.Vars(r)
	id := qParams["id"]
	quiz, _ := findQuiz(id)
	if quiz == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}
	quiz.AddQA(&qa)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// EditQuestion - edit question by id
func EditQuestion(w http.ResponseWriter, r *http.Request) {
	var qa qz.QA
	json.NewDecoder(r.Body).Decode(&qa)
	qParams := mux.Vars(r)
	id := qParams["id"]
	qaID := qParams["qa_id"]
	quiz, _ := findQuiz(id)
	if quiz == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}
	quiz.EditQA(qaID, qa.Question, qa.Answers)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

//RemoveQuestion - remove question from list by id
func RemoveQuestion(w http.ResponseWriter, r *http.Request) {
	qParams := mux.Vars(r)
	quizID := qParams["id"]
	qaID := qParams["qa_id"]
	quiz, _ := findQuiz(quizID)
	if quiz == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Quiz is not found"))
		return
	}
	err := quiz.RemoveQA(qaID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Q&A not found"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// ShuffleAnswers - shuffle answers in q&a
func ShuffleAnswers(w http.ResponseWriter, r *http.Request) {
	qParams := mux.Vars(r)
	quizID := qParams["id"]
	questionID := qParams["qa_id"]
	quiz, _ := findQuiz(quizID)
	if quiz == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}
	qa, _ := quiz.Find(questionID)
	if qa == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}
	qa.Shuffle()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// SetAnswer - set user's answer to question by id
func SetAnswer(w http.ResponseWriter, r *http.Request) {
	var answers []int
	json.NewDecoder(r.Body).Decode(&answers)
	qParams := mux.Vars(r)
	quizID := qParams["id"]
	questionID := qParams["qa_id"]
	quiz, _ := findQuiz(quizID)
	if quiz == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}
	res, err := quiz.SetAnswer(questionID, answers)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}
	result := make(map[string]bool)
	result["result"] = res
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

// GetResults - get quiz results by id
func GetResults(w http.ResponseWriter, r *http.Request) {
	qParams := mux.Vars(r)
	quizID := qParams["id"]
	quiz, _ := findQuiz(quizID)
	if quiz == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Quiz is not found"))
		return
	}
	result := quiz.GetResults()
	json.NewEncoder(w).Encode(result)
}

// NewRouter - create router
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", GetQuizList).Methods("GET")
	r.HandleFunc("/create-quiz", CreateQUIZ).Methods("POST")
	r.HandleFunc("/get-quiz/{id}", GetQuiz).Methods("GET")
	r.HandleFunc("/edit-quiz/{id}", EditQuiz).Methods("PUT")
	r.HandleFunc("/remove-quiz/{id}", RemoveQuiz).Methods("DELETE")
	r.HandleFunc("/shuffle-quiz/{id}", ShuffleQuiz).Methods("PUT")
	r.HandleFunc("/get-question/{id}/{qa_id}", GetQuestion).Methods("GET")
	r.HandleFunc("/add-question/{id}", AddQuestion).Methods("POST")
	r.HandleFunc("/edit-question/{id}/{qa_id}", EditQuestion).Methods("PUT")
	r.HandleFunc("/remove-question/{id}/{qa_id}", RemoveQuestion).Methods("DELETE")
	r.HandleFunc("/shuffle-answers/{id}/{qa_id}", ShuffleAnswers).Methods("PUT")
	r.HandleFunc("/set-answer/{id}/{qa_id}", SetAnswer).Methods("POST")
	r.HandleFunc("/get-results/{id}", GetResults).Methods("GET")
	return r
}
