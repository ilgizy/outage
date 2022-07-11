package main

import (
	"PreventiveWork/internal/handlers"
	"PreventiveWork/internal/models"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// @title           Swagger Example API
	// @version         1.0
	// @host      localhost:8080

	var ds models.DataSource
	ds.New()

	router = gin.Default()
	handler := handlers.NewHandler(ds)
	handler.Register(router)
	router.Run()

}
