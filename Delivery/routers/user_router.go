package routers

import (
	"application/project/Delivery/controllers"
	infrastructure "application/project/Infrastructure"

	"github.com/gin-gonic/gin"
)

// create a method or function for setting up a route
func UserRouter(router *gin.Engine, userController *controllers.UserHandler) {
	freeRouter := router.Group("/")
	freeRouter.POST("/register", userController.CreatUser)
	freeRouter.POST("/login", userController.LoginUser)

	userRouter := router.Group("/", infrastructure.AuthMiddleWare())
	userRouter.PUT("/user/update", userController.UpdateUser)

	freeadminRouter := router.Group("/admin")
	freeadminRouter.POST("/admin/register", userController.RegisterAdmin)
	freeadminRouter.POST("/admin/login", userController.LoginUser)

	adminRouter := router.Group("/admin", infrastructure.AdminMiddleWare(), infrastructure.AdminMiddleWare())
	adminRouter.PUT("/admin/update", userController.UpdateUser)
	adminRouter.DELETE("/admin/delete/:id", userController.DeleteUserByID)
	adminRouter.POST("/admin/register/admin", userController.RegisterAdmin)
}
