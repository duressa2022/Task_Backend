package usecases

import (
	domain "application/project/Domain"
	infrastructure "application/project/Infrastructure"
	repository "application/project/Repository"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// create the struct to repesent the usercase and user rep
type UserUseCase struct {
	UserRepository *repository.UserRepository
}

// create a method to create the new user from the repo
func NewUserUseCase(user_repository *repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		UserRepository: user_repository,
	}
}

// create a function to register the user into the database
func (u *UserUseCase) RegisterUser(user *domain.User) error {
	_, err := u.UserRepository.GetByUserName(user.Username)
	if err == nil {
		return errors.New("user by this username already exist")
	}
	hashed, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashed
	user.Role = "user"
	return u.UserRepository.CreatUser(user)

}

// create a method for deleting the user from the database
func (u *UserUseCase) DeleteUser(username string) error {
	user, err := u.UserRepository.GetByUserName(username)
	if err != nil {
		return err
	}
	if user.Role == "admin" {
		return errors.New("error of authority")
	}
	return u.UserRepository.DeleteUser(username)

}

// create a method for updating the suer information in database
func (u *UserUseCase) UpdateUser(user *domain.User) error {
	_, err := u.UserRepository.GetByUserName(user.Username)
	if err != nil {
		return err
	}
	hashed, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashed
	return u.UpdateUser(user)

}

// create the function to get user by using the id
func (u *UserUseCase) GetUser(id primitive.ObjectID) (*domain.User, error) {
	return u.UserRepository.GetByID(id.String())
}

// create the method for logging the user in to the system
func (u *UserUseCase) LoginUser(username string, password string) (string, error) {
	user, err := u.UserRepository.GetByUserName(username)
	if err != nil {
		return "", err
	}
	if err := infrastructure.ComparePassword(password, user.Password); err != nil {
		return "", err
	}
	token, err := infrastructure.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", err
	}
	return token, nil

}

// create the method for adding the admin into the system
func (u *UserUseCase) RegisterAdmin(user *domain.User) error {
	user, err := u.UserRepository.GetByUserName(user.Username)
	if err != nil {
		return err
	}
	hashed, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashed
	user.Role = "admin"
	return u.UserRepository.CreatUser(user)
}
