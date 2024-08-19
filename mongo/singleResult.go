package mongo

import "go.mongodb.org/mongo-driver/mongo"

//create for handling the single result type
type MongoSingleResult struct {
	singleResult *mongo.SingleResult
}

//create a interface for singleResult
type SingleResult interface {
	Decode(v interface{}) error
}

//create an implemantation for decode method

func (s *MongoSingleResult) Decode(v interface{}) error {
	return s.singleResult.Decode(v)
}
