package handlers

import (
	"PreventiveWork/internal/models"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (h *handler) ShowPreventiveWork(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	data := h.ds.FindPreventiveWorkByID(id)
	if data == nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.Data(http.StatusOK, gin.MIMEJSON, data)
}

func (h *handler) ShowPreventiveWorks(c *gin.Context) {
	data := h.ds.GetPreventiveWorkJson()
	if len(data) == 0 {
		c.String(http.StatusOK, "Профилактические работы отсутствуют")
		return
	}
	c.Data(http.StatusOK, gin.MIMEJSON, h.ds.GetPreventiveWorkJson())

	//c.AsciiJSON(http.StatusOK, string(h.ds.GetPreventiveWorkJson()))
}
