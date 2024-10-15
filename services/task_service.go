package services

import (
	"task-api/models"
	"task-api/repositories"
)

type TaskService struct {
	TaskRepo *repositories.TaskRepository
}

// get task by user id
func (s *TaskService) GetTasksByUserID(userID int) ([]models.Task, error) {
	return s.TaskRepo.GetTasksByUserID(userID)
}

// create task
func (s *TaskService) CreateTask(task *models.Task) error {
	return s.TaskRepo.CreateTask(task)
}

// update task
func (s *TaskService) UpdateTask(task *models.Task) error {
	return s.TaskRepo.UpdateTask(task)
}

// remove task
func (s *TaskService) DeleteTask(taskID int, userID int) error {
	return s.TaskRepo.DeleteTask(taskID, userID)
}
