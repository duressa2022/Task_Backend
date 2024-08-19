package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// create a type for handling the client struct
type MongoClient struct {
	client *mongo.Client
}

// create an interface for working with cliant
type Client interface {
	Database(string) MongoDatabase
	Connect(context.Context) error
	Disconnect(context.Context) error
	StartSession() (*mongo.Session, error)
	UseSession(cxt context.Context, f func(mongo.Session) error) error
	Ping(context.Context) error
}

// create a method for creating databse
func (c *MongoClient) Database(name string) *MongoDatabase {
	database := c.client.Database(name)
	return &MongoDatabase{
		db: database,
	}
}

// create a method connecting to the database
func (c *MongoClient) Connect(cxt context.Context) error {
	return c.client.Connect(cxt)
}

// create a method for starting session
func (c *MongoClient) Disconnect(cxt context.Context) error {
	return c.client.Disconnect(cxt)
}

