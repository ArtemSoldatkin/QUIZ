# QUIZ Api (GoLang / Gorilla Mux)

_To start:_<br/> docker-compose up -d --build

## API Documentation

_Get list of quiz_<br/> Return list of saved quiz.

-   URL<br> /
-   Method<br/> GET
-   URL Params<br/> None
-   Data Params<br/> None
-   Success Response:<br/> Code: 200 OK<br/> Content: [{"ID":"5d1d18a4-e630-4bd9-a9a3-1cefbef41403","Title":"quiz
    1","QAS":null,"Result":{"Total":0,"Right":0,"Changed":0}}]
-   Error Response:<br/> None

_Create quiz_<br/> Create quiz and add to list.

-   URL<br> /create-quiz
-   Method<br/> POST
-   URL Params<br/> None
-   Data Params<br/> title=[string]
-   Success Response:<br/> Code: 200 OK<br/> Content: "OK"
-   Error Response:<br/> None

_Get quiz_<br/> Get quiz by id.

-   URL<br> /get-quiz/:id
-   Method<br/> GET
-   URL Params<br/> id=[string]
-   Data Params<br/> None
-   Success Response:<br/> Code: 200 OK<br/> Content: {"ID":"5d1d18a4-e630-4bd9-a9a3-1cefbef41403","Title":"quiz
    1","QAS":null,"Result":{"Total":0,"Right":0,"Changed":0}}
-   Error Response:<br/> Code: 404 Not Found<br/> Content: "Quiz is not found"

_Edit quiz_<br/> Edit quiz by id.

-   URL<br> /edit-quiz/:id
-   Method<br/> PUT
-   URL Params<br/> id=[string]
-   Data Params<br/> title=[string]
-   Success Response:<br/> Code: 200 OK<br/> Content: "OK"
-   Error Response:<br/> Code: 404 Not Found<br/> Content: "Quiz is not found"

_Remove quiz_<br/> Remove quiz frm list by id.

-   URL<br> /remove-quiz/:id
-   Method<br/> DELETE
-   URL Params<br/> id=[string]
-   Data Params<br/> None
-   Success Response:<br/> Code: 200 OK<br/> Content: "OK"
-   Error Response:<br/> Code: 404 Not Found<br/> Content: "Quiz is not found"

_Shuffle quiz_<br/> Shuffle quiz by id.

-   URL<br> /shuffle-quiz/:id
-   Method<br/> PUT
-   URL Params<br/> id=[string]
-   Data Params<br/> None
-   Success Response:<br/> Code: 200 OK<br/> Content: "OK"
-   Error Response:<br/> Code: 404 Not Found<br/> Content: "Quiz is not found"

_Get question_<br/> Get question from quiz by id.

-   URL<br> /get-question/:id/:qa_id
-   Method<br/> GET
-   URL Params<br/> id=[string]<br/> qa_id=[string]
-   Data Params<br/> None
-   Success Response:<br/> Code: 200 OK<br/> Content:
    {"ID":"a395eef5-a364-4153-8cf8-a8bbd47a3803","question":"test","answers":null,"IsAnswered":false,"Result":false,"IsChanged":false}
-   Error Response:<br/> Code: 404 Not Found<br/> Content: "Quiz is not found" OR "Question is not found"

_Add question_<br/> Add question to quiz.

-   URL<br> /add-question/:id
-   Method<br/> POST
-   URL Params<br/> id=[string]
-   Data Params<br/> question=[string]
-   Success Response:<br/> Code: 200 OK<br/> Content: "OK"
-   Error Response:<br/> Code: 404 Not Found<br/> Content: "Quiz is not found"

_Edit question_<br/> Edit question from quiz by id.

-   URL<br> /edit-question/:id/:qa_id
-   Method<br/> PUT
-   URL Params<br/> id=[string]<br/> qa_id=[string]
-   Data Params<br/> question=[string]<br/> answers=[{text:[string], is_rigth: [boolean]}]
-   Success Response:<br/> Code: 200 OK<br/> Content: "OK"
-   Error Response:<br/> Code: 404 Not Found<br/> Content: "Quiz is not found"

_Remove question_<br/> Remove question from quiz by id.

-   URL<br> /remove-question/:id/:qa_id
-   Method<br/> DELETE
-   URL Params<br/> id=[string]<br/> qa_id=[string]
-   Data Params<br/> None
-   Success Response:<br/> Code: 200 OK<br/> Content: "OK"
-   Error Response:<br/> Code: 404 Not Found<br/> Content: "Quiz is not found" OR "Question is not found"

_Shuffle question_<br/> Shuffle question's answers from quiz by id.

-   URL<br> /shuffle-answers/:id/:qa_id
-   Method<br/> PUT
-   URL Params<br/> id=[string]<br/> qa_id=[string]
-   Data Params<br/> None
-   Success Response:<br/> Code: 200 OK<br/> Content: "OK"
-   Error Response:<br/> Code: 404 Not Found<br/> Content: "Quiz is not found" OR "Question is not found"

_(WIP)Set answer_<br/> Set answer to question from quiz by id.

-   URL<br> /set-answer/:id/:qa_id
-   Method<br/> POST
-   URL Params<br/> id=[string]<br/> qa_id=[string]
-   Data Params<br/> [integer]
-   Success Response:<br/> Code: 200 OK<br/> Content: {"result":false}
-   Error Response:<br/> None

_Get results_<br/> Get quiz results by id.

-   URL<br> /get-results/:id
-   Method<br/> GET
-   URL Params<br/> id=[string]
-   Data Params<br/> None
-   Success Response:<br/> Code: 200 OK<br/> Content: {"Total":1,"Right":1,"Changed":1}
-   Error Response:<br/> Code: 404 Not Found<br/> Content: "Quiz is not found"
