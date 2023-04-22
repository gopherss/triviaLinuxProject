package server

import "trivia-backend/core/apps"

func (es *EchoServer) levelRoutes() {

	level := apps.NewLevelHTTPService(es.ctx, es.db)

	es.GET("/level", level.ListLevelHandler)
	es.POST("/level/create", level.CreateQuestionHandler)
	es.PUT("/level/edit/:id", level.EditLevelHandler)
	es.DELETE("/level/delete/:id", level.RemoveLevelHandler)

}

func (es *EchoServer) routesIndexLevel() {
	es.levelRoutes()
}
