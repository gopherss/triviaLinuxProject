package database

import (
	"fmt"
	"log"
	"trivia-backend/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreSQL struct {
	config env.EnvApp
}

func (p *PostgreSQL) ConnectDB() *gorm.DB {

	chainConnection := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s",
		p.config.DB_USERNAME,
		p.config.DB_PASSWORD,
		p.config.DB_HOST,
		p.config.DB_PORT,
		p.config.DB_DATABASE,
	)

	db, err := gorm.Open(postgres.Open(chainConnection), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database")
	} else {
		fmt.Println("Connection success.")
	}

	/* db.AutoMigrate(&models.Level{})
	db.AutoMigrate(&models.Question{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Option{})
	*/
	return db
}

func NewPostgreSQLDatabase(ec env.EnvApp) Database {
	return &PostgreSQL{config: ec}
}
