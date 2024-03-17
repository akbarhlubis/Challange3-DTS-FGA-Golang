package main

import (
	"challange3_17Mar/config"
	"challange3_17Mar/controller"
	"challange3_17Mar/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	inDB := &controller.InDB{DB: db}

	router := gin.Default()

	router.Use(middleware.AuthMiddleware)
	router.GET("person/:id", inDB.GetPerson)
	router.GET("persons", inDB.GetPersons)
	router.POST("person", inDB.CreatePerson)
	router.PUT("person", inDB.UpdatePerson)
	router.DELETE("person/:id", inDB.DeletePerson)
	router.Run(":8080")
}
