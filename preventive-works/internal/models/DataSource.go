package models

import (
	"PreventiveWork/pkg/client/mongodb"
	"PreventiveWork/pkg/logging"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type DataSource struct {
	db     *mongo.Database
	logger logging.Logger
}

func (ds *DataSource) New(logger logging.Logger) {
	ds.logger = logger
	client, err := mongodb.NewClient(context.TODO(), "mongo", "27017", "root", "root", "PreventiveWork")
	if err != nil {
		logger.Fatal("нет подключения к базе данных")
	}
	logger.Info("успешное подключение к базе данных")
	ds.db = client
}

//добавление новой профилактической работы
func (ds *DataSource) AddNewPreventiveWork(ctx context.Context, nameService string, createAt time.Time, deadline time.Time, title string, description string) (id string, err error) {
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
		idService, err = ds.addService(ctx, nameService)
		if err != nil {
			return "", err
		}
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
	idWork, err := collection.InsertOne(ctx, preventiveWork)
	if err != nil {
		ds.logger.Fatal(err)
		return "", err
	}
	return idWork.InsertedID.(primitive.ObjectID).Hex(), nil
}

// добавление нового события в профилактическую работу
func (ds *DataSource) AddNewEvent(ctx context.Context, idPreventiveWork string, createAt time.Time, deadline time.Time, description string, status string) (err error) {
	event := Event{
		CreateAt:    createAt,
		Deadline:    deadline,
		Description: description,
		Status:      status,
	}

	idObject, err := primitive.ObjectIDFromHex(idPreventiveWork)
	if err != nil {
		ds.logger.Fatal(err)
		return err
	}
	collection := ds.db.Collection("PreventiveWork")
	filter := bson.M{"_id": idObject}
	update := bson.M{"$push": bson.M{"events": event}}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		ds.logger.Fatal(err)
		return err
	}
	return nil
}

//Возвращает профилактическую работу в формате json по ее id
func (ds *DataSource) FindPreventiveWorkByID(id string, ctx context.Context) []byte {
	var result PreventiveWork
	collection := ds.db.Collection("PreventiveWork")
	idObject, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", idObject}}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		ds.logger.Error(err)
		return nil
	} else if err != nil {
		ds.logger.Error(err)
		return nil
	}
	preventiveWork, _ := json.Marshal(&result)
	return preventiveWork
}

//Возвращает список всех профилактических работ в формате json
func (ds DataSource) GetPreventiveWorkJson(ctx context.Context) []byte {
	var preventiveWork []byte
	collection := ds.db.Collection("PreventiveWork")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		ds.logger.Error(err)
		return nil
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result PreventiveWork
		err := cur.Decode(&result)
		if err != nil {
			ds.logger.Error(err)
			return nil
		}
		preventiveWorkJSON, _ := json.Marshal(result)
		preventiveWork = append(preventiveWork, preventiveWorkJSON...)
	}
	if err := cur.Err(); err != nil {
		ds.logger.Error(err)
		return nil
	}
	return preventiveWork
}

// возвращает список всех сервисов
func (ds DataSource) getServices(ctx context.Context) []Service {
	var services []Service
	collection := ds.db.Collection("Service")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		ds.logger.Error(err)
		return nil
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result Service
		err := cur.Decode(&result)
		if err != nil {
			ds.logger.Error(err)
			return nil
		}
		services = append(services, result)
	}
	if err := cur.Err(); err != nil {
		ds.logger.Error(err)
		return nil
	}
	return services
}

//добавление нового сервиса
func (ds DataSource) addService(ctx context.Context, nameService string) (primitive.ObjectID, error) {
	idService := primitive.NewObjectID()
	s := Service{
		Name: nameService,
		Id:   idService,
	}
	collection := ds.db.Collection("Service")
	_, err := collection.InsertOne(ctx, s)
	if err != nil {
		ds.logger.Fatal(err)
		return idService, err
	}
	return idService, nil
}
