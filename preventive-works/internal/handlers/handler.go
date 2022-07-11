package handlers

import (
	"PreventiveWork/docs"
	"PreventiveWork/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	"strconv"
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

	docs.SwaggerInfo.Title = "preventive-works"
	docs.SwaggerInfo.Description = "API для отслеживания профилактических работ"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/preventive_works"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// ShowPreventiveWork
// @Tags         PreventiveWork
// @Summary      отображение профилактической работы по id
// @Param        id   path      int  true  "PreventiveWork id"
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.PreventiveWork
// @Router       /{id} [get]
func (h *handler) ShowPreventiveWork(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.Data(http.StatusOK, gin.MIMEJSON, h.ds.FindPreventiveWorkByID(id))
	return
}

// ShowPreventiveWorks
// @Tags         PreventiveWorks
// @Summary      отображение всех профилактических работ
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.PreventiveWork
// @Router       / [get]
func (h *handler) ShowPreventiveWorks(c *gin.Context) {
	c.Data(http.StatusOK, gin.MIMEJSON, h.ds.GetPreventiveWorkJson())
	//c.AsciiJSON(http.StatusOK, string(h.ds.GetPreventiveWorkJson()))
}
