package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// NewPreventiveWork
// @Tags         NewPreventiveWork
// @Summary      добавление новой профилактической работы
// @Param        name_service    formData     string  true  "Имя сервиса"
// @Param        create_at    formData     string  true  "Дата создания профил. работы"
// @Param        deadline    formData     string  true  "Дата окончания профил. работы"
// @Param        title    formData     string  true  "Название профил. работы"
// @Param        description    formData     string  true  "Описание профил. работы"
// @Success      200
// @Failure      404  {object}  int
// @Failure      500  {object}  int
// @Router       /new_work [post]
func (h *handler) NewPreventiveWork(c *gin.Context) {
	h.logger.Info("создание новой профилактической работы")
	nameService := c.PostForm("name_service")
	createAtString := c.PostForm("create_at")
	deadlineSTring := c.PostForm("deadline")
	title := c.PostForm("title")
	description := c.PostForm("description")

	createAt, err := time.Parse("2006-01-02 15:04:05", createAtString)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		h.logger.Debug("дата окончания введена неверно")
	}
	deadline, err := time.Parse("2006-01-02 15:04:05", deadlineSTring)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		h.logger.Debug("дата окончания введена неверно")
	}

	if deadline.Before(createAt) {
		c.Status(http.StatusInternalServerError)
		h.logger.Debug("дата окончания не может быть раньше даты создания")
	}

	err = h.ds.AddNewPreventiveWork(context.TODO(), nameService, createAt, deadline, title, description)
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.Status(http.StatusOK)
	}
}
