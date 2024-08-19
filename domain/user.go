package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

//create struct for repesenting user information....user model
type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id, omitempty"`
	Username string             `json:"username"`
	Password string             `json:"password"`
	Role     string             `json:"role"`
}

//create an interface that is implemented by using user repo
type UserRepository interface {
	CreatUser(user *User) error
	GetByID(id string) (*User, error)
	GetByUserName(username string) (*User, error)
	UpdateUser(User *User) error
	DeleteUser(username string) error
	DeleteByID(id string) error
	GetAllUser() ([]*User, error)
	GetBYCondition(condition map[string]interface{}) ([]*User, error)
	CountUser() (int64, error)
}
