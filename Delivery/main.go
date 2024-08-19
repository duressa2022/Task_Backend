package main

import (
	"application/project/Delivery/controllers"
	"application/project/Delivery/routers"
	repository "application/project/Repository"
	usecases "application/project/Usecases"
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// main function to setup connection and routes
func main() {
	client, err := ConnectDB()
	if err != nil {
		panic("error while connecting with database")
	}
	taskRepository := repository.NewTaskRepository(client, "taskcluster", "tasks")
	userRepository := repository.NewUserCollection(client, "taskcluster", "users")
	router := Setup(taskRepository, userRepository)
	router.Run("localhost:8000")

}

// create a function to connect to the database
func ConnectDB() (*mongo.Client, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		return nil, err
	}
	url := os.Getenv("URL")
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// create function for loading the url from .env
func LoadUrl() (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	return os.Getenv("URL"), nil
}

// create a function to setup the routing path
func Setup(task *repository.TaskRepository, user *repository.UserRepository) *gin.Engine {
	router := gin.Default()
	task_UseCase := usecases.NewTaskUseCase(task)
	task_handler := controllers.NewTaskController(task_UseCase)
	routers.TaskRouter(router, task_handler)

	user_usecase := usecases.NewUserUseCase(user)
	user_handler := controllers.NewUserHandler(user_usecase)
	routers.UserRouter(router, user_handler)
	return router
}
