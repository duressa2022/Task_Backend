package repository

import (
	domain "application/project/Domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// create a structure for holding user information
type UserRepository struct {
	UserCollection *mongo.Collection
}

// create a function for providing user collection
func NewUserCollection(client *mongo.Client, database string, collection string) *UserRepository {
	user_collection := client.Database(database).Collection(collection)
	return &UserRepository{
		UserCollection: user_collection,
	}
}

// create a method for adding new user into the datbase
func (r *UserRepository) CreatUser(user *domain.User) error {
	user.ID = primitive.NewObjectID()
	_, err := r.UserCollection.InsertOne(context.TODO(), user)
	return err
}

// create a method for getting user based on user id
func (r *UserRepository) GetByID(user_id string) (*domain.User, error) {
	var user *domain.User
	id, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return nil, err
	}
	err = r.UserCollection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// create a method for getting user by using username
func (r *UserRepository) GetByUserName(username string) (*domain.User, error) {
	var user *domain.User
	err := r.UserCollection.FindOne(context.TODO(), bson.D{{Key: "username", Value: username}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// create a method for updating user information
func (r *UserRepository) UpdateUser(user *domain.User) error {
	updatedArea := bson.M{
		"username":  user.Username,
		".password": user.Password,
	}
	_, err := r.UserCollection.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: user.ID}}, bson.M{"$set": updatedArea})
	return err
}

// create a method for deleting user by using username
func (r *UserRepository) DeleteUser(username string) error {
	_, err := r.UserCollection.DeleteOne(context.TODO(), bson.D{{Key: "username", Value: username}})
	return err

}

// create a method for deleting user by user id
func (r *UserRepository) DeleteByID(user_id string) error {
	id, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return err
	}
	_, err = r.UserCollection.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: id}})
	return err

}

// create a method for getting all user from the databse
func (r *UserRepository) GetAllUser() ([]*domain.User, error) {
	var users []*domain.User
	cur, err := r.UserCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var user domain.User
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

// create for getting user from the database on conditions
func (r *UserRepository) GetBYCondition(condition map[string]interface{}) ([]*domain.User, error) {
	var users []*domain.User
	filter := bson.M{}
	for key, value := range condition {
		filter[key] = value
	}
	cur, err := r.UserCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var user domain.User
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil

}

// create a method for counting the number od user is collection
func (r *UserRepository) CountUser() int {
	number, err := r.UserCollection.CountDocuments(context.TODO(), bson.D{{}})
	if err != nil {
		return 0
	}
	return int(number)
}
