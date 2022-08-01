package main

import (
	"PreventiveWork/internal/handlers"
	"PreventiveWork/internal/models"
	"PreventiveWork/pkg/logging"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

// @title           preventive-works
// @version         1.0
// @description     API для отслеживания профилактических работ
// @host      localhost:8101
func main() {
	logger := logging.GetLogger()

	var ds models.DataSource
	ds.New(logger)

	router = gin.Default()
	logger.Info("создан новый роутер")

	handler := handlers.NewHandler(ds, logger)
	handler.Register(router)

	router.Run()
}
