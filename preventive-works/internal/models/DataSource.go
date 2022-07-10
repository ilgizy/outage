package models

import "time"

type DataSource struct {
	Service        []Service        `json:"service"`
	Event          []Event          `json:"event"`
	PreventiveWork []PreventiveWork `json:"preventive_work"`
}

func (ds *DataSource) New() {
	ds.Service = []Service{
		{Name: "Go", Id: 0},
		{Name: "С#", Id: 1},
	}

	ds.PreventiveWork = []PreventiveWork{
		{Id: 0, CreateAt: time.Date(2022, time.March, 4, 15, 15, 15, 15, time.UTC), Deadline: time.Date(2022, time.March, 4, 18, 15, 15, 15, time.UTC), Title: "Задача1", Description: "Описание задачи 1", CountEvent: 3, IdService: 1},
		{Id: 1, CreateAt: time.Date(2022, time.June, 4, 15, 15, 15, 15, time.UTC), Deadline: time.Date(2022, time.March, 4, 18, 15, 15, 15, time.UTC), Title: "Задача2", Description: "Описание задачи 2", CountEvent: 3, IdService: 0},
		{Id: 2, CreateAt: time.Date(2022, time.February, 4, 15, 15, 15, 15, time.UTC), Deadline: time.Date(2022, time.March, 4, 18, 15, 15, 15, time.UTC), Title: "Задача3", Description: "Описание задачи 3", CountEvent: 3, IdService: 1},
	}

	ds.Event = []Event{
		{Id: 0, CreateAt: time.Date(2022, time.March, 4, 15, 15, 15, 15, time.UTC), Deadline: time.Date(2022, time.March, 4, 18, 15, 15, 15, time.UTC), Description: "Создана задача", Status: "Запланировано", IdPreventiveWork: 0},
		{Id: 1, CreateAt: time.Date(2022, time.March, 4, 16, 15, 15, 15, time.UTC), Deadline: time.Date(2022, time.March, 4, 16, 30, 15, 15, time.UTC), Description: "Сделано", Status: "Обновлено", IdPreventiveWork: 0},
		{Id: 2, CreateAt: time.Date(2022, time.March, 4, 17, 15, 15, 15, time.UTC), Deadline: time.Date(2022, time.March, 4, 17, 30, 15, 15, time.UTC), Description: "Закрыто", Status: "Завершено", IdPreventiveWork: 0},

		{Id: 3, CreateAt: time.Date(2022, time.March, 4, 15, 15, 15, 15, time.UTC), Deadline: time.Date(2022, time.March, 4, 18, 15, 15, 15, time.UTC), Description: "Создана задача", Status: "Запланировано", IdPreventiveWork: 1},
		{Id: 4, CreateAt: time.Date(2022, time.March, 4, 16, 15, 15, 15, time.UTC), Deadline: time.Date(2022, time.March, 4, 16, 30, 15, 15, time.UTC), Description: "Сделано", Status: "Обновлено", IdPreventiveWork: 1},
		{Id: 5, CreateAt: time.Date(2022, time.March, 4, 15, 15, 15, 15, time.UTC), Deadline: time.Date(2022, time.March, 4, 18, 15, 15, 15, time.UTC), Description: "Закрыто", Status: "Завершено", IdPreventiveWork: 1},

		{Id: 6, CreateAt: time.Date(2022, time.March, 4, 15, 15, 15, 15, time.UTC), Deadline: time.Date(2022, time.March, 4, 18, 15, 15, 15, time.UTC), Description: "Создана задача", Status: "Запланировано", IdPreventiveWork: 2},
		{Id: 7, CreateAt: time.Date(2022, time.March, 4, 16, 15, 15, 15, time.UTC), Deadline: time.Date(2022, time.March, 4, 16, 30, 15, 15, time.UTC), Description: "Сделано", Status: "Обновлено", IdPreventiveWork: 2},
		{Id: 8, CreateAt: time.Date(2022, time.March, 4, 15, 15, 15, 15, time.UTC), Deadline: time.Date(2022, time.March, 4, 18, 15, 15, 15, time.UTC), Description: "Закрыто", Status: "Завершено", IdPreventiveWork: 2},
	}
}

func (ds *DataSource) AddNewPreventiveWork(idService int, nameService string, idPreventiveWork int, createAt time.Time, deadline time.Time, title string, description string) {
	flag := true
	for _, service := range ds.Service {
		if idService == service.Id {
			flag = false
		}
	}
	if flag {
		service := Service{
			Name: nameService,
			Id:   idService,
		}
		ds.Service = append(ds.Service, service)
	}

	preventiveWork := PreventiveWork{
		Id:          idPreventiveWork,
		CreateAt:    createAt,
		Deadline:    deadline,
		Title:       title,
		Description: description,
		CountEvent:  1,
		IdService:   idService,
	}
	ds.PreventiveWork = append(ds.PreventiveWork, preventiveWork)

	event := Event{
		Id:               0,
		CreateAt:         createAt,
		Deadline:         deadline,
		Description:      description,
		Status:           "Запланированно",
		IdPreventiveWork: idPreventiveWork,
	}

	ds.Event = append(ds.Event, event)
}

func (ds *DataSource) addNewEvent(idEvent int, idPreventiveWork int, createAt time.Time, deadline time.Time, description string, status string) {
	event := Event{
		Id:               idEvent,
		CreateAt:         createAt,
		Deadline:         deadline,
		Description:      description,
		Status:           status,
		IdPreventiveWork: idPreventiveWork,
	}
	ds.Event = append(ds.Event, event)

}
