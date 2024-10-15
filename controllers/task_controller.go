package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"task-api/models"
	"task-api/services"

	"github.com/gorilla/mux"
)

type TaskController struct {
	TaskService *services.TaskService
}

// get task by user
func (c *TaskController) GetTasks(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int) // get el user_id from context

	tasks, err := c.TaskService.GetTasksByUserID(userID)
	if err != nil {
		http.Error(w, "failed to get task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// create task by user
func (c *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int) // get el user_id from context

	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Bad data", http.StatusBadRequest)
		return
	}

	task.UserID = userID
	err = c.TaskService.CreateTask(&task)
	if err != nil {
		http.Error(w, "Error creating task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// update task by user
func (c *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int) // get el user_id from context
	taskID, _ := strconv.Atoi(mux.Vars(r)["id"])

	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "wrong data", http.StatusBadRequest)
		return
	}

	task.ID = taskID
	task.UserID = userID
	err = c.TaskService.UpdateTask(&task)
	if err != nil {
		http.Error(w, "failed to update the task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Eliminar tarea
func (c *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int) // get el user_id from context
	taskID, _ := strconv.Atoi(mux.Vars(r)["id"])

	err := c.TaskService.DeleteTask(taskID, userID)
	if err != nil {
		http.Error(w, "failed to remove the task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
