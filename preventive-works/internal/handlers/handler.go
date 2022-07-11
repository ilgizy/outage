package handlers

import (
	"PreventiveWork/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
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
	router.GET("/prevntive_works", func(c *gin.Context) {
		c.Data(http.StatusOK, gin.MIMEJSON, h.ds.GetPreventiveWorkJson())
		//c.AsciiJSON(http.StatusOK, string(h.ds.GetPreventiveWorkJson()))
	})

	//router.GET("/prevntive_works/:id")
}
