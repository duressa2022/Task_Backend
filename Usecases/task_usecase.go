package usecases

import (
	domain "application/project/Domain"
	repository "application/project/Repository"
)

// create the structure for handling the tasks usecase
type TaskUseCase struct {
	TaskRepository *repository.TaskRepository
}

// create the method for creating new usecase for the task
func NewTaskUseCase(task_repository *repository.TaskRepository) *TaskUseCase {
	return &TaskUseCase{
		TaskRepository: task_repository,
	}
}

// create the method for creating new task for the user
func (u *TaskUseCase) CreateTask(task *domain.Task) error {
	return u.TaskRepository.CreateTask(task)
}

// create the method for updaeting the task for the user
func (u *TaskUseCase) UpdateTask(task *domain.Task) error {
	return u.TaskRepository.UpdateTask(task)
}

// create the method for getting  all tasks
func (u *TaskUseCase) GetAllTasks(id string) ([]*domain.Task, error) {
	return u.TaskRepository.GetAllTasks(id)
}

// create the method for deleting the tasks by using id
func (u *TaskUseCase) DeleteByID(id string) error {
	return u.TaskRepository.DeleteTask(id)
}

// create the method for deleting the task by using title
func (u *TaskUseCase) DeleteByTitle(title string) error {
	return u.TaskRepository.DeleteByTitle(title)
}

// create the method getting the tasks by using conditions
func (u *TaskUseCase) GetByCondition(condition map[string]interface{}) ([]*domain.Task, error) {
	return u.TaskRepository.GetByCondition(condition)
}

// create the method for getting tasks by using id
func (u *TaskUseCase) GetByID(task_id string) (*domain.Task, error) {
	return u.TaskRepository.GetByID(task_id)

}

// create the method for getting tasks by using title
func (u *TaskUseCase) GetByTitle(title string) ([]*domain.Task, error) {
	condition := map[string]interface{}{}
	condition["title"] = title
	return u.TaskRepository.GetByCondition(condition)
}
