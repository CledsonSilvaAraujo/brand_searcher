package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func connectToMongoDB(uri string) error {
	clientOptions := options.Client().ApplyURI(uri)
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}
	return client.Ping(context.TODO(), nil)
}

// func storeResults(results []GoogleResult) error {
// 	collection := client.Database("your_db").Collection("results")
// 	_, err := collection.InsertMany(context.TODO(), results)
// 	return err
// }
