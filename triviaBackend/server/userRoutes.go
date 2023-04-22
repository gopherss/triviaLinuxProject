package server

import appUser "trivia-backend/core/apps"

func (es *EchoServer) userRoutes() {
	users := appUser.NewUserHTTPService(es.ctx, es.db)

	es.GET("/user", users.ListUserHandler)
	es.POST("/user/create", users.CreateUserHandler)
}

func (es *EchoServer) routesIndexUser() {
	es.userRoutes()
}
