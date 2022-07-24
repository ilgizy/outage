package handlers

//// NewPreventiveWork
//// @Tags         NewPreventiveWork
//// @Summary      добавление новой профилактической работы
//// @Param        name_service    formData     string  true  "Имя сервиса"
//// @Param        create_at    formData     string  true  "Дата создания профил. работы"
//// @Param        deadline    formData     string  true  "Дата окончания профил. работы"
//// @Param        title    formData     string  true  "Название профил. работы"
//// @Param        description    formData     string  true  "Описание профил. работы"
//// @Success      200
//// @Router       /new_work [post]
//func (h *handler) NewPreventiveWork(c *gin.Context) {
//	nameService := c.PostForm("name_service")
//	createAtString := c.PostForm("create_at")
//	deadlineSTring := c.PostForm("deadline")
//	title := c.PostForm("title")
//	description := c.PostForm("description")
//
//	createAt, err := time.Parse("2006-01-02 15:04:05", createAtString)
//	if err != nil {
//		c.Status(http.StatusBadRequest)
//	}
//	deadline, err := time.Parse("2006-01-02 15:04:05", deadlineSTring)
//	if err != nil {
//		c.Status(http.StatusBadRequest)
//	}
//	h.ds.AddNewPreventiveWork(nameService, createAt, deadline, title, description)
//	c.Status(http.StatusOK)
//}
