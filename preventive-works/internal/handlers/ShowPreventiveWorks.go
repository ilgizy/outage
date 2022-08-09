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
// @Failure      404  {object}  int
// @Router       / [get]
func (h *handler) ShowPreventiveWorks(c *gin.Context) {
	data := h.ds.GetPreventiveWorkJson(context.TODO())
	if data == nil {
		c.String(http.StatusNotFound, "Профилактические работы отсутствуют")
		h.logger.Info("Профилактические работы отсутствуют")
		return
	}
	c.Data(http.StatusOK, gin.MIMEJSON, data)
}
