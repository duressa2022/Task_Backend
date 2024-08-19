package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// create a type for handling the collection
type MongoCollection struct {
	Collection *mongo.Collection
}

// create a interface for working with the collection
type Collection interface {
	FindOne(cxt context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.SingleResult, error)
	InsertOne(cxt context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	InsertMany(cxt context.Context, documents interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
	Find(cxt context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.SingleResult, error)
	DeleteOne(cxt context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	DeleteMany(cxt context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	UpdateOne(cxt context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	UpdateMany(cxt context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	CountDocuments(cxt context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error)
}

// create a method for implementing the interface here
func (c *MongoCollection) FindOne(cxt context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.SingleResult, error) {
	singleResult := c.Collection.FindOne(cxt, filter)
	return singleResult, nil

}
