package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// NewEvent
// @Tags         NewEvent
// @Summary      добавление новой профилактической работы
// @Param        id   path      string  true  "id профилактической работы"
// @Param        status    formData     string  true  "Статус события"
// @Param        create_at    formData     string  true  "Дата создания события"
// @Param        deadline    formData     string  true  "Дата окончания события"
// @Param        description    formData     string  true  "Описание события"
// @Success      200
// @Failure      404  {object}  int
// @Failure      500  {object}  int
// @Router       /{id}/new_event [put]
func (h *handler) NewEvent(c *gin.Context) {
	h.logger.Info("создание нового события в профилактической работе")
	id := c.Param("id")
	status := c.PostForm("status")
	createAtString := c.PostForm("create_at")
	deadlineString := c.PostForm("deadline")
	description := c.PostForm("description")

	createAt, err := time.Parse("2006-01-02 15:04:05", createAtString)
	if err != nil {
		c.Status(http.StatusBadRequest)
		h.logger.Debug("дата создания введена неверно", err)
		return
	}

	deadline, err := time.Parse("2006-01-02 15:04:05", deadlineString)
	if err != nil {
		c.Status(http.StatusBadRequest)
		h.logger.Debug("дата окончания введена неверно", err)
		return
	}

	if deadline.Before(createAt) {
		c.Status(http.StatusInternalServerError)
		h.logger.Debug("дата окончания не может быть раньше даты создания")
		return
	}
	err = h.ds.AddNewEvent(context.TODO(), id, createAt, deadline, description, status)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.Status(http.StatusOK)
	}
}
