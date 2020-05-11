package api

import (
	"encoding/json"
	"net/http"
	"quiz/quiz"

	"github.com/gorilla/mux"
)

// QUIZs - list of quiz
var QUIZs []*quiz.Quiz

// CreateQUIZ - create quiz
func CreateQUIZ(w http.ResponseWriter, r *http.Request) {
	var params map[string]string
	json.NewDecoder(r.Body).Decode(&params)
	title := params["title"]
	newQuiz := quiz.Quiz{Title: title}
	newQuiz.Init()
	QUIZs = append(QUIZs, &newQuiz)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// AddQuestion - add question to quiz by id
func AddQuestion(w http.ResponseWriter, r *http.Request) {
	var qa quiz.QA
	json.NewDecoder(r.Body).Decode(&qa)
	qParams := mux.Vars(r)
	id := qParams["id"]
	quiz := findQuiz(id)
	if quiz == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}
	quiz.AddQA(&qa)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// NewRouter - create router
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", CreateQUIZ).Methods("POST")
	r.HandleFunc("/add-question/{id}", AddQuestion).Methods("PUT")
	return r
}
