package models

import (
	"PreventiveWork/pkg/client/mongodb"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type DataSource struct {
	db *mongo.Database
}

func (ds *DataSource) New() {
	client, err := mongodb.NewClient(context.TODO(), "mongo", "27017", "root", "root", "PreventiveWork")
	if err != nil {
		return
	}
	ds.db = client
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
func (ds DataSource) GetServiceJson(ctx context.Context) []byte {
	var services []byte
	collection := ds.db.Collection("Service")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result Service
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		serviceJSON, _ := json.Marshal(result)
		services = append(services, serviceJSON...)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return services
}

//Возвращает список всех профилактических работ в формате json
func (ds DataSource) GetPreventiveWorkJson(ctx context.Context) []byte {
	var preventiveWork []byte
	collection := ds.db.Collection("PreventiveWork")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result PreventiveWork
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		preventiveWorkJSON, _ := json.Marshal(result)
		preventiveWork = append(preventiveWork, preventiveWorkJSON...)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return preventiveWork
}

//Возвращает список всех событий в формате json
func (ds DataSource) GetEventJson(ctx context.Context) []byte {
	var events []byte
	collection := ds.db.Collection("Events")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result Event
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		eventsJSON, _ := json.Marshal(result)
		events = append(events, eventsJSON...)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return events
}
