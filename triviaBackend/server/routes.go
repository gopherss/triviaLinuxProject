package server

import appQuestion "trivia-backend/core/apps"

func (es *EchoServer) questionRoutes() {

	question := appQuestion.NewQuestionHTTPService(es.ctx, es.db)

	es.GET("/question", question.ListHandler)
	es.POST("/question/create", question.CreateHandler)

}

func (es *EchoServer) routes() {
	es.questionRoutes()
}
