package repositories

import (
	"database/sql"
	"task-api/models"
)

type TaskRepository struct {
	DB *sql.DB
}

// get task by user_id
func (r *TaskRepository) GetTasksByUserID(userID int) ([]models.Task, error) {
	query := "SELECT id, title, description, status FROM tasks WHERE user_id = ?"
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// create task
func (r *TaskRepository) CreateTask(task *models.Task) error {
	query := "INSERT INTO tasks (title, description, status, user_id) VALUES (?, ?, ?, ?)"
	_, err := r.DB.Exec(query, task.Title, task.Description, task.Status, task.UserID)
	return err
}

// update task
func (r *TaskRepository) UpdateTask(task *models.Task) error {
	query := "UPDATE tasks SET title = ?, description = ?, status = ? WHERE id = ? AND user_id = ?"
	_, err := r.DB.Exec(query, task.Title, task.Description, task.Status, task.ID, task.UserID)
	return err
}

// remove task
func (r *TaskRepository) DeleteTask(taskID int, userID int) error {
	query := "DELETE FROM tasks WHERE id = ? AND user_id = ?"
	_, err := r.DB.Exec(query, taskID, userID)
	return err
}
