package mg

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New(dsnUri, db string) (*mongo.Client, *mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(dsnUri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, nil, err
	}
	return client, client.Database(db), nil
}
