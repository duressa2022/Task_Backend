package mongo

import "go.mongodb.org/mongo-driver/mongo"

//create the struct for handling database
type MongoDatabase struct {
	db *mongo.Database
}

//create am interface for handling methods
type Database interface {
	Collection(name string) *Collection
	Client() *MongoClient
}

//create a function to establish the collection
func (d *MongoDatabase) Collection(name string) *MongoCollection {
	collection := d.db.Collection(name)
	return &MongoCollection{
		Collection: collection,
	}
}

//create a function to create a client
func (d *MongoDatabase) Client() *MongoClient {
	client := d.db.Client()
	return &MongoClient{
		client: client,
	}
}
