package controllers

import (
	domain "application/project/Domain"
	usecases "application/project/Usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

// create struct for representing user handler
type UserHandler struct {
	userUseCase *usecases.UserUseCase
}

// create method for creating new instance of user handler
func NewUserHandler(user_usecase *usecases.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: user_usecase,
	}
}

// create a method/handler  for creating new user
func (u *UserHandler) CreatUser(c *gin.Context) {
	var user domain.User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error binding the user"})
		return
	}
	err := u.userUseCase.RegisterUser(&user)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error while adding user"})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "user created"})

}

// create a method/handler for login th user into the system
func (u *UserHandler) LoginUser(c *gin.Context) {
	var user domain.User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error while binding.."})
		return
	}
	token, err := u.userUseCase.LoginUser(user.Username, user.Password)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error while logging"})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24, "", "", false, true)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "logged in to the system"})
}

// create a method/handler for updating the user
func (u *UserHandler) UpdateUser(c *gin.Context) {
	var user domain.User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error while binding"})
		return
	}
	err := u.userUseCase.UpdateUser(&user)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error while updating"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "user updated"})
}

// create a method/handler for updating the user
func (u *UserHandler) DeleteUserByID(c *gin.Context) {
	err := u.userUseCase.DeleteUser(c.Param("username"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error while deleting"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "user updated"})

}

// create a method/handler for registring thr user
func (u *UserHandler) RegisterAdmin(c *gin.Context) {
	var user domain.User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error while creating"})
		return
	}
	err := u.userUseCase.RegisterAdmin(&user)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error while creaing"})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "admin created"})
}
