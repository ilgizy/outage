package handlers

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Register(router *gin.Engine)
}

type handler struct {
}

func NewHandler() Handler {
	return &handler{}
}

func (h *handler) Register(router *gin.Engine) {

}
