package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

//ShowPreventiveWork
//@Tags         PreventiveWork
//@Summary      отображение профилактической работы по id
//@Param        id   path      string  true  "PreventiveWork id"
//@Accept       json
//@Produce      json
//@Success      200  {object}  models.PreventiveWork
// @Failure      404  {object}  int
//@Router       /{id} [get]
func (h *handler) ShowPreventiveWork(c *gin.Context) {
	id := c.Param("id")
	data := h.ds.FindPreventiveWorkByID(id, context.TODO())
	if data == nil {
		c.String(http.StatusNotFound, "Профилактической работы с запрашиваемым id не существует")
		h.logger.Infof("Профилактической работы с id: %s не существует", id)
		return
	}
	c.Data(http.StatusOK, gin.MIMEJSON, data)
}
