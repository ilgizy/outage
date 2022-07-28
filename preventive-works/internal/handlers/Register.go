package handlers

import (
	"PreventiveWork/docs"
	"PreventiveWork/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type Handler interface {
	Register(router *gin.Engine)
}

type handler struct {
	ds models.DataSource
}

func NewHandler(ds models.DataSource) Handler {
	return &handler{ds: ds}
}

func (h *handler) Register(router *gin.Engine) {

	router.GET("/preventive_works", h.ShowPreventiveWorks)
	router.GET("/preventive_works/:id", h.ShowPreventiveWork)
	router.POST("/preventive_works/new_work", h.NewPreventiveWork)
	router.PUT("/preventive_works/:id/new_event", h.NewEvent)

	docs.SwaggerInfo.Title = "preventive-works"
	docs.SwaggerInfo.Description = "API для отслеживания профилактических работ"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/preventive_works"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
