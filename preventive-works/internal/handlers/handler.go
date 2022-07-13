package handlers

import (
	"PreventiveWork/docs"
	"PreventiveWork/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	"strconv"
	"time"
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
	router.POST("/preventive_works/new_work", h.NewPreventiveWork)
	router.PUT("/preventive_works/:id/new_event", h.NewEvent)

	docs.SwaggerInfo.Title = "preventive-works"
	docs.SwaggerInfo.Description = "API для отслеживания профилактических работ"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/preventive_works"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// ShowPreventiveWork
// @Tags         PreventiveWork
// @Summary      отображение профилактической работы по id
// @Param        id   path      int  true  "PreventiveWork id"
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.PreventiveWork
// @Router       /{id} [get]
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

// ShowPreventiveWorks
// @Tags         PreventiveWorks
// @Summary      отображение всех профилактических работ
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.PreventiveWork
// @Router       / [get]
func (h *handler) ShowPreventiveWorks(c *gin.Context) {
	data := h.ds.GetPreventiveWorkJson()
	if len(data) == 0 {
		c.String(http.StatusOK, "Профилактические работы отсутствуют")
		return
	}
	c.Data(http.StatusOK, gin.MIMEJSON, h.ds.GetPreventiveWorkJson())

	//c.AsciiJSON(http.StatusOK, string(h.ds.GetPreventiveWorkJson()))
}

// NewPreventiveWork
// @Tags         NewPreventiveWork
// @Summary      добавление новой профилактической работы
// @Param        name_service    formData     string  true  "Имя сервиса"
// @Param        create_at    formData     string  true  "Дата создания профил. работы"
// @Param        deadline    formData     string  true  "Дата окончания профил. работы"
// @Param        title    formData     string  true  "Название профил. работы"
// @Param        description    formData     string  true  "Описание профил. работы"
// @Success      200
// @Router       /new_work [post]
func (h *handler) NewPreventiveWork(c *gin.Context) {
	nameService := c.PostForm("name_service")
	createAtString := c.PostForm("create_at")
	deadlineSTring := c.PostForm("deadline")
	title := c.PostForm("title")
	description := c.PostForm("description")

	createAt, err := time.Parse("2006-01-02 15:04:05", createAtString)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	deadline, err := time.Parse("2006-01-02 15:04:05", deadlineSTring)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	h.ds.AddNewPreventiveWork(nameService, createAt, deadline, title, description)
	c.Status(http.StatusOK)
}

// NewEvent
// @Tags         NewPreventiveWork
// @Summary      добавление новой профилактической работы
// @Param        id   path      int  true  "id профилактической работы"
// @Param        status    formData     string  true  "Статус события"
// @Param        create_at    formData     string  true  "Дата создания события"
// @Param        deadline    formData     string  true  "Дата окончания события"
// @Param        description    formData     string  true  "Описание события"
// @Success      200
// @Router       /{id}/new_event [put]
func (h *handler) NewEvent(c *gin.Context) {
	//idEvent int, idPreventiveWork int, createAt time.Time, deadline time.Time, description string, status string
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	status := c.PostForm("status")
	createAtString := c.PostForm("create_at")
	deadlineSTring := c.PostForm("deadline")
	description := c.PostForm("description")

	createAt, err := time.Parse("2006-01-02 15:04:05", createAtString)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	deadline, err := time.Parse("2006-01-02 15:04:05", deadlineSTring)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	h.ds.AddNewEvent(id, createAt, deadline, description, status)
	c.Status(http.StatusOK)
}
