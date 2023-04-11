package cmd

import (
	"context"
	"trivia-backend/env"
	"trivia-backend/internal/database"
	"trivia-backend/server"
)

func Start() {

	ctx := context.Background()

	_env := env.GetEnv(".env.development")

	db := database.NewPostgreSQLDatabase(_env).ConnectDB()

	server.NewEchoServer(ctx, db, _env.SERVER_PORT).Run()

}
