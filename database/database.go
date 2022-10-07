package database

import (
	"fmt"
	"github.com/diaspangestu/FGA-Golang-Practice-Rest-Api/structs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	host     = "localhost"
	user     = "dias"
	password = "testing123"
	dbPort   = "5432"
	dbname   = "practice"
)

func DBInit() *gorm.DB {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)

	DB, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	DB.AutoMigrate(structs.Person{})
	return DB
}
