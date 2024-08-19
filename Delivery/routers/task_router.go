package routers

import (
	"application/project/Delivery/controllers"
	infrastructure "application/project/Infrastructure"

	"github.com/gin-gonic/gin"
)

// create a function/method for settingup the route for tasks
func TaskRouter(router *gin.Engine, taskController *controllers.TaskController) {

	userRoute := router.Group("/", infrastructure.AuthMiddleWare())
	userRoute.GET("/tasks", taskController.GetTasks)
	userRoute.POST("/user/create", taskController.CreateTask)
	userRoute.GET("/tasks/:title", taskController.GetByTitle)
	userRoute.GET("/tasks/status/:status", taskController.GetByStatus)
	userRoute.POST("/tasks/create", taskController.CreateTask)
	userRoute.PUT("/tasks/update", taskController.UpdateTask)
	userRoute.DELETE("tasks/delete/:title", taskController.DeleteByTitle)

	adminRoute := router.Group("/admin", infrastructure.AuthMiddleWare(), infrastructure.AdminMiddleWare())
	adminRoute.POST("/admin/tasks", taskController.CreateTask)
	adminRoute.PUT("/admin/tasks/update", taskController.UpdateTask)
	adminRoute.DELETE("/admin/tasks/delete/:id", taskController.DeleteByID)
	adminRoute.DELETE("/admin/tasks/delete/title/:title", taskController.DeleteByTitle)
	adminRoute.GET("/admin/tasks", taskController.GetTasks)
	adminRoute.GET("/admin/tasks/:title", taskController.GetByTitle)
	adminRoute.GET("/admin/status/:status", taskController.GetByStatus)
}
