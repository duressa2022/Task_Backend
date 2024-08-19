package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// create struct for representing the task...........models for a task
type Task struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	DueDate     time.Time          `json:"duedate" bson:"duedate"`
	Status      string             `json:"status" bson:"status"`
	UserID      primitive.ObjectID `json:"user_id" bson:"_user_id"`
}

// create an interface to for handling the taskRep
type TaskRepository interface {
	CreateTask(task *Task) error
	GetAllTasks(id string) ([]*Task, error)
	GetByID(id string) (*Task, error)
	UpdateTask(task *Task) error
	DeleteTask(id string) error
	DeleteByTitle(title string) error
	GetByCondition(condition map[string]interface{}) ([]*Task, error)
	CheckTitle(title string) error
}
