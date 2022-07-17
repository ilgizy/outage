package main

import (
	"PreventiveWork/internal/handlers"
	"PreventiveWork/internal/models"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

// @title           preventive-works
// @version         1.0
// @description     API для отслеживания профилактических работ
// @host      localhost:8101
func main() {

	var ds models.DataSource
	ds.New()

	router = gin.Default()
	handler := handlers.NewHandler(ds)
	handler.Register(router)
	router.Run()

}
