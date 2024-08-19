package controllers

import (
	domain "application/project/Domain"
	usecases "application/project/Usecases"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// create a struct for handleing tasks
type TaskController struct {
	task_UseCase *usecases.TaskUseCase
}

// create the method for creating newtaskController
func NewTaskController(task_UseCase *usecases.TaskUseCase) *TaskController {
	return &TaskController{
		task_UseCase: task_UseCase,
	}
}

// create a handler to create a task on the dtatabase
func (t *TaskController) CreateTask(c *gin.Context) {
	user_id, ok := c.Get("user_id")
	if !ok {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error of"})
		return
	}
	userid, ok := user_id.(string)
	if !ok {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error of"})
		return
	}
	var task domain.Task
	if err := c.BindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error while creating"})
		return
	}
	id, err := primitive.ObjectIDFromHex(userid)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error of"})
		return
	}
	task.UserID = id
	err = t.task_UseCase.CreateTask(&task)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error while creating"})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "task created"})

}

// create a handler to get tasks from the database
func (t *TaskController) GetTasks(c *gin.Context) {
	user_id, ok := c.Get("user_id")
	if !ok {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error of"})
		return
	}
	userid, ok := user_id.(string)
	if !ok {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error of"})
		return
	}
	tasks, err := t.task_UseCase.GetAllTasks(userid)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error while getting"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": tasks})

}

// create a handler to get task based on the id
func (t *TaskController) GetTask(c *gin.Context) {
	task, err := t.task_UseCase.GetByID(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error while getting"})
		return
	}
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": task})
}

// create a handler to get task based on the title
func (t *TaskController) GetByTitle(c *gin.Context) {
	task, err := t.task_UseCase.GetByTitle(c.Param("title"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error while getting..."})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": task})
}

// create a handler for updatting the task
func (t *TaskController) UpdateTask(c *gin.Context) {
	var task domain.Task
	if err := c.BindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error while binding.."})
		return
	}
	err := t.task_UseCase.UpdateTask(&task)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"messsage": "error while updating"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "task updated"})
}

// create a handler for deleting the task by using id
func (t *TaskController) DeleteByID(c *gin.Context) {
	err := t.task_UseCase.DeleteByID(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error while deleting"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "task deleted"})
}

// create a handler for deleting task by using title
func (t *TaskController) DeleteByTitle(c *gin.Context) {
	err := t.task_UseCase.DeleteByTitle(c.Param("title"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error while deleting"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "task is deleted"})
}

// create a handler for getting task by using status
func (t *TaskController) GetByStatus(c *gin.Context) {
	condition := map[string]interface{}{}
	condition["status"] = c.Param("status")
	task, err := t.task_UseCase.GetByCondition(condition)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error while getting"})
		return
	}
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": task})
}
