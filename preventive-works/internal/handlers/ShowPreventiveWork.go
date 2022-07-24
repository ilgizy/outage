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
//@Router       /{id} [get]
func (h *handler) ShowPreventiveWork(c *gin.Context) {
	id := c.Param("id")
	data := h.ds.FindPreventiveWorkByID(id, context.TODO())
	if data == nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.Data(http.StatusOK, gin.MIMEJSON, data)
}
