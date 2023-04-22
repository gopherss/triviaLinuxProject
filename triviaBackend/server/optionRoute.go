package server

import appOption "trivia-backend/core/apps"

func (es *EchoServer) optionRoutes() {

	option := appOption.NewOptionHTTPService(es.ctx, es.db)

	es.GET("/option", option.ListOptionHandler)
	es.POST("/option/create", option.CreateOptionHandler)

}

func (es *EchoServer) routesIndexOption() {
	es.optionRoutes()
}
