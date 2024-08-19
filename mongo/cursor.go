package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// create a type for working with the cursor
type MongoCursor struct {
	cursor *mongo.Cursor
}

// create an interface for working with the cursor
type Cursor interface {
	Close(context.Context) error
	Next(context.Context) bool
	Decode(interface{}) error
	All(context.Context, interface{}) error
}

// create a method for closing the cursor
func (c *MongoCursor) Close(cxt context.Context) error {
	return c.cursor.Close(cxt)
}

// create a method for moveing to the next
func (c *MongoCursor) Next(cxt context.Context) bool {
	return c.cursor.Next(cxt)
}

// create a method for decoding the data
func (c *MongoCursor) Decode(v interface{}) error {
	return c.cursor.Decode(v)
}

// create a method for decoding all to the the val
func (c *MongoCursor) All(cxt context.Context, v interface{}) error {
	return c.cursor.All(cxt, v)
}
