package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"task-api/models"
	"task-api/services"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type TaskController struct {
	TaskService *services.TaskService
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// get task by user
func (controller *TaskController) GetTasks(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int) // get user_id from context

	tasks, err := controller.TaskService.GetTasksByUserID(userID)
	if err != nil {
		http.Error(w, "failed to get task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// create task by user
func (controller *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int) // get el user_id from context

	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	task.UserID = userID

	validationErr := validate.Struct(task)
	if validationErr != nil {
		http.Error(w, validationErr.Error(), http.StatusBadRequest)
		return
	}
	err = controller.TaskService.CreateTask(&task)
	if err != nil {
		http.Error(w, "Error creating task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// update task by user
func (controller *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
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
	validationErr := validate.Struct(task)
	if validationErr != nil {
		http.Error(w, validationErr.Error(), http.StatusBadRequest)
		return
	}
	err = controller.TaskService.UpdateTask(&task)
	if err != nil {
		http.Error(w, "failed to update the task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// remove task
func (controller *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int) // get el user_id from context
	taskID, _ := strconv.Atoi(mux.Vars(r)["id"])

	err := controller.TaskService.DeleteTask(taskID, userID)
	if err != nil {
		http.Error(w, "failed to remove the task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
