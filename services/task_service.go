package services

import (
	"task-api/models"
	"task-api/repositories"
)

type TaskService struct {
	TaskRepo *repositories.TaskRepository
}

// Obtener tareas por user_id
func (s *TaskService) GetTasksByUserID(userID int) ([]models.Task, error) {
	return s.TaskRepo.GetTasksByUserID(userID)
}

// Crear tarea
func (s *TaskService) CreateTask(task *models.Task) error {
	return s.TaskRepo.CreateTask(task)
}

// Actualizar tarea
func (s *TaskService) UpdateTask(task *models.Task) error {
	return s.TaskRepo.UpdateTask(task)
}

// Eliminar tarea
func (s *TaskService) DeleteTask(taskID int, userID int) error {
	return s.TaskRepo.DeleteTask(taskID, userID)
}
