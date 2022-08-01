package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// NewEvent
// @Tags         NewPreventiveWork
// @Summary      добавление новой профилактической работы
// @Param        id   path      string  true  "id профилактической работы"
// @Param        status    formData     string  true  "Статус события"
// @Param        create_at    formData     string  true  "Дата создания события"
// @Param        deadline    formData     string  true  "Дата окончания события"
// @Param        description    formData     string  true  "Описание события"
// @Success      200
// @Router       /{id}/new_event [put]
func (h *handler) NewEvent(c *gin.Context) {
	h.logger.Info("создание нового события в профилактической работе")
	id := c.Param("id")
	status := c.PostForm("status")
	createAtString := c.PostForm("create_at")
	deadlineSTring := c.PostForm("deadline")
	description := c.PostForm("description")

	createAt, err := time.Parse("2006-01-02 15:04:05", createAtString)
	if err != nil {
		c.Status(http.StatusBadRequest)
		h.logger.Debug("дата создания введена неверно")
	}

	deadline, err := time.Parse("2006-01-02 15:04:05", deadlineSTring)
	if err != nil {
		c.Status(http.StatusBadRequest)
		h.logger.Debug("дата окончания введена неверно")
	}

	err = h.ds.AddNewEvent(context.TODO(), id, createAt, deadline, description, status)
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.Status(http.StatusOK)
	}
}