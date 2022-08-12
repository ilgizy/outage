package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, host, port, username, password, database string) (db *mongo.Database, err error) {
	var mongoDBURL string
	mongoDBURL = fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)

	clientOptions := options.Client().ApplyURI(mongoDBURL)
	clientOptions.SetAuth(options.Credential{
		Username: username,
		Password: password,
	})

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongoDB due to error: %v", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping mongoDB due to error: %v", err)
	}

	return client.Database(database), nil
}
