package models

import (
	"PreventiveWork/pkg/client/mongodb"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
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
func (ds *DataSource) AddNewPreventiveWork(ctx context.Context, nameService string, createAt time.Time, deadline time.Time, title string, description string) {
	//проверяет есть ли сервис с таким именем, если есть, то запоминаем его id, если нет, то добавляем новый
	services := ds.getServices(ctx)
	flag := true
	var idService primitive.ObjectID
	for _, service := range services {
		if nameService == service.Name {
			flag = false
			idService = service.Id
		}
	}
	if flag {
		ds.addService(ctx, nameService)
	}

	//создание первого события в профилактической работе
	event := Event{
		CreateAt:    createAt,
		Deadline:    deadline,
		Description: description,
		Status:      "Запланированно",
	}
	var events []Event
	events = append(events, event)
	preventiveWork := PreventiveWork{
		CreateAt:    createAt,
		Deadline:    deadline,
		Title:       title,
		Description: description,
		IdService:   idService,
		Events:      events,
	}

	//добавление профилактической работы в базу данных
	collection := ds.db.Collection("PreventiveWork")
	_, err := collection.InsertOne(ctx, preventiveWork)
	if err != nil {
		log.Fatal(err)
	}
}

// добавление нового события в профилактическую работу
func (ds *DataSource) AddNewEvent(ctx context.Context, idPreventiveWork string, createAt time.Time, deadline time.Time, description string, status string) {
	event := Event{
		CreateAt:    createAt,
		Deadline:    deadline,
		Description: description,
		Status:      status,
	}

	idObject, _ := primitive.ObjectIDFromHex(idPreventiveWork)
	collection := ds.db.Collection("PreventiveWork")
	filter := bson.M{"_id": idObject}
	update := bson.M{"$push": bson.M{"events": event}}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
}

//Возвращает профилактическую работу в формате json по ее id
func (ds *DataSource) FindPreventiveWorkByID(id string, ctx context.Context) []byte {
	var result PreventiveWork
	collection := ds.db.Collection("PreventiveWork")
	idObject, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", idObject}}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil
	} else if err != nil {
		log.Fatal(err)
	}
	preventiveWork, _ := json.Marshal(&result)
	return preventiveWork
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

// возвращает список всех сервисов
func (ds DataSource) getServices(ctx context.Context) []Service {
	var services []Service
	collection := ds.db.Collection("Service")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result Service
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		services = append(services, result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return nil
	}
	return services
}

//добавление нового сервиса
func (ds DataSource) addService(ctx context.Context, nameService string) primitive.ObjectID {
	idService := primitive.NewObjectID()
	s := Service{
		Name: nameService,
		Id:   idService,
	}
	collection := ds.db.Collection("Service")
	_, err := collection.InsertOne(ctx, s)
	if err != nil {
		log.Fatal(err)
	}
	return idService
}
