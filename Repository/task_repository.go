package repository

import (
	domain "application/project/Domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// create for holding information /collection of tasks
type TaskRepository struct {
	TaskCollection *mongo.Collection
}

// create a methosd for creating new task repository
func NewTaskRepository(client *mongo.Client, database string, collection string) *TaskRepository {
	task_collection := client.Database(database).Collection(collection)
	return &TaskRepository{TaskCollection: task_collection}
}

// create a method to insert/post one task into the database
func (r *TaskRepository) CreateTask(task *domain.Task) error {
	task.ID = primitive.NewObjectID()
	_, err := r.TaskCollection.InsertOne(context.TODO(), task)
	if err != nil {
		return err
	}
	return nil
}

// create a method to get all tasks from the database
func (r *TaskRepository) GetAllTasks(id string) ([]*domain.Task, error) {
	userid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return []*domain.Task{}, err
	}
	var tasks []*domain.Task
	cur, err := r.TaskCollection.Find(context.TODO(), bson.D{{Key: "_user_id", Value: userid}})
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var task domain.Task
		err := cur.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}

// create a method for getting task by using id
func (r *TaskRepository) GetByID(task_id string) (*domain.Task, error) {
	id, err := primitive.ObjectIDFromHex(task_id)
	if err != nil {
		return nil, err
	}
	var task *domain.Task
	err = r.TaskCollection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}}).Decode(&task)
	if err != nil {
		return nil, err
	}
	return task, nil

}

// ceate a method for updating the task by using id
func (r *TaskRepository) UpdateTask(task *domain.Task) error {
	updatedArea := bson.M{
		"title":       task.Title,
		"description": task.Description,
		"duedate":     task.DueDate,
		"status":      task.Status,
	}
	_, err := r.TaskCollection.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: task.ID}}, bson.M{"$set": updatedArea})
	return err
}

// create a method for deleting task by using id of the task
func (r *TaskRepository) DeleteTask(task_id string) error {
	id, err := primitive.ObjectIDFromHex(task_id)
	if err != nil {
		return err
	}
	_, err = r.TaskCollection.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: id}})
	return err

}

// create a method for deleting task by using title of the task
func (r *TaskRepository) DeleteByTitle(title string) error {
	_, err := r.TaskCollection.DeleteMany(context.TODO(), bson.D{{Key: "title", Value: title}})
	return err
}

// create a task by using creatain condition give in the map
func (r *TaskRepository) GetByCondition(condition map[string]interface{}) ([]*domain.Task, error) {
	var tasks []*domain.Task
	filter := bson.M{}
	for key, value := range condition {
		filter[key] = value
	}
	cur, err := r.TaskCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var task domain.Task
		err := cur.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil

}

// create a method for checking the existance of the title of the task in dtatbase
func (r *TaskRepository) CheckTitle(title string) error {
	_, err := r.TaskCollection.Find(context.TODO(), bson.D{{Key: "title", Value: title}})
	return err

}
