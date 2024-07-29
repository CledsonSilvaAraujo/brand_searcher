package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SearchResult struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Term  string             `bson:"term"`
	Title string             `bson:"title,omitempty"`
	Link  string             `bson:"link,omitempty"`
}
