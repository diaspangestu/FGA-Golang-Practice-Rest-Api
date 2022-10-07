package main

import (
	"github.com/diaspangestu/FGA-Golang-Practice-Rest-Api/controllers"
	"github.com/diaspangestu/FGA-Golang-Practice-Rest-Api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.DBInit()
	inDB := &controllers.Controller{Masterdb: db}

	router := gin.Default()

	router.GET("/person/:id", inDB.GetPerson)
	router.GET("/persons", inDB.GetAllPerson)
	router.POST("/person", inDB.CreatePerson)
	router.PUT("/person/:id", inDB.UpdatePerson)
	router.DELETE("/person/:id", inDB.DeletePerson)

	router.Run(":8080")
}
