package database

import (
	"context"
	"log"
	"strings"
	"time"

	"backend/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = Client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")
}

func GetCollection(databaseName, collectionName string) *mongo.Collection {
	return Client.Database(databaseName).Collection(collectionName)
}

func SaveResultsToMongo(term string, results string) {
	collection := GetCollection("brand_monitor", "search_results")

	links := strings.Split(results, "\n")

	for _, link := range links {

		searchResult := models.SearchResult{
			ID:   primitive.NewObjectID(),
			Term: term,
			Link: link,
		}
		if link == "" {
			log.Printf("Link not found: %v", link)
			continue
		}
		_, err := collection.InsertOne(context.TODO(), searchResult)
		if err != nil {
			log.Printf("Could not insert result: %v", err)
		} else {
			log.Println("Result inserted successfully")
		}
	}
}
