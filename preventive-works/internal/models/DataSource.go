package models

import (
	"encoding/json"
	"time"
)

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

//добавление новой профилактической работы
func (ds *DataSource) AddNewPreventiveWork(nameService string, createAt time.Time, deadline time.Time, title string, description string) {
	flag := true
	var s = Service{}
	for _, service := range ds.Service {
		if nameService == service.Name {
			flag = false
			s = service
		}
	}
	if flag {
		service := Service{
			Name: nameService,
			Id:   len(ds.Service),
		}
		s = service
		ds.Service = append(ds.Service, service)
	}

	preventiveWork := PreventiveWork{
		Id:          len(ds.PreventiveWork),
		CreateAt:    createAt,
		Deadline:    deadline,
		Title:       title,
		Description: description,
		CountEvent:  1,
		IdService:   s.Id,
	}

	event := Event{
		Id:               0,
		CreateAt:         createAt,
		Deadline:         deadline,
		Description:      description,
		Status:           "Запланированно",
		IdPreventiveWork: len(ds.PreventiveWork),
	}
	ds.PreventiveWork = append(ds.PreventiveWork, preventiveWork)
	ds.Event = append(ds.Event, event)
}

// добавление нового события в профилактическую работу
func (ds *DataSource) AddNewEvent(idPreventiveWork int, createAt time.Time, deadline time.Time, description string, status string) {
	event := Event{
		Id:               len(ds.Event),
		CreateAt:         createAt,
		Deadline:         deadline,
		Description:      description,
		Status:           status,
		IdPreventiveWork: idPreventiveWork,
	}
	ds.Event = append(ds.Event, event)
}

//Возвращает профилактическую работу в формате json по ее id
func (ds *DataSource) FindPreventiveWorkByID(id int) []byte {
	for _, work := range ds.PreventiveWork {
		if work.Id == id {
			var events []Event
			for _, event := range ds.Event {
				if event.IdPreventiveWork == id {
					events = append(events, event)
				}

			}
			work.Events = events
			preventiveWork, _ := json.Marshal(&work)
			return preventiveWork
		}
	}
	return nil
}

//Возвращает список всех сервисов в формате json
func (ds DataSource) GetServiceJson() []byte {
	var services []byte
	for _, service := range ds.Service {
		serviceJSON, _ := json.Marshal(service)
		services = append(services, serviceJSON...)
	}
	return services
}

//Возвращает список всех профилактических работ в формате json
func (ds DataSource) GetPreventiveWorkJson() []byte {
	var preventiveWork []byte
	for _, work := range ds.PreventiveWork {
		workJSON, _ := json.Marshal(work)
		preventiveWork = append(preventiveWork, workJSON...)
	}
	return preventiveWork
}

//Возвращает список всех событий в формате json
func (ds DataSource) GetEventJson() []byte {
	var events []byte
	for _, event := range ds.Event {
		eventJSON, _ := json.Marshal(event)
		events = append(events, eventJSON...)
	}
	return events
}
