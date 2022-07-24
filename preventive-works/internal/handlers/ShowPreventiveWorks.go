package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ShowPreventiveWorks
// @Tags         PreventiveWorks
// @Summary      отображение всех профилактических работ
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.PreventiveWork
// @Router       / [get]
func (h *handler) ShowPreventiveWorks(c *gin.Context) {
	data := h.ds.GetPreventiveWorkJson(context.TODO())
	if len(data) == 0 {
		c.String(http.StatusOK, "Профилактические работы отсутствуют")
		return
	}
	c.Data(http.StatusOK, gin.MIMEJSON, data)

	//c.AsciiJSON(http.StatusOK, string(h.ds.GetPreventiveWorkJson()))
}
