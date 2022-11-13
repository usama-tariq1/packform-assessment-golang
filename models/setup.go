package models

import (
	"fmt"

	deepgorm "github.com/survivorbat/gorm-deep-filtering"
	"github.com/usama-tariq1/leet-gin/config"
	"github.com/usama-tariq1/leet-gin/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {

	console := helper.Console{}

	env := config.Env{}
	host := env.Get("DB_HOST")
	user := env.Get("DB_USER")
	pass := env.Get("DB_PASS")
	db_name := env.Get("DB_NAME")
	port := env.Get("DB_PORT")
	// sslmode := env.Get("DB_SSL")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, pass, db_name, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.Use(deepgorm.New())
	if err != nil {
		console.Log("Error", "** Failed to connect to database! **")

		panic("Failed to connect to database!")
	}
	console.Log("Debug", "** Connected to Database **")

	DB = db
	return db
}
