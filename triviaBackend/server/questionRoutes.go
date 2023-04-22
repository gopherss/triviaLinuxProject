package server

import (
	appQuestion "trivia-backend/core/apps"
)

func (es *EchoServer) questionRoutes() {

	question := appQuestion.NewQuestionHTTPService(es.ctx, es.db)

	es.GET("/question", question.ListQuestionHandler)
	es.POST("/question/create", question.CreateQuestionHandler)

}

func (es *EchoServer) routesIndexQuestion() {
	es.questionRoutes()
}
