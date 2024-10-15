package main

import (
	"log"
	"net/http"
	"os"
	"task-api/controllers"
	"task-api/dal/database"
	"task-api/middleware"
	"task-api/repositories"
	"task-api/services"

	"github.com/gorilla/mux"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// init repo
	userRepo := &repositories.UserRepository{DB: db}

	// auth
	jwtSecret := os.Getenv("JWT_SECRET")
	authService := &services.AuthService{
		UserRepo:  userRepo,
		JwtSecret: jwtSecret,
	}
	authController := &controllers.AuthController{AuthService: authService}

	taskRepo := &repositories.TaskRepository{DB: db}
	taskService := &services.TaskService{
		TaskRepo: taskRepo,
	}
	taskController := &controllers.TaskController{TaskService: taskService}

	router := mux.NewRouter()

	// init routes
	router.HandleFunc("/login", authController.Login).Methods("POST")
	router.HandleFunc("/register", authController.Register).Methods("POST")

	// task routes
	router.HandleFunc("/tasks", middleware.JwtMiddleware(taskController.GetTasks, jwtSecret)).Methods("GET")
	router.HandleFunc("/tasks", middleware.JwtMiddleware(taskController.CreateTask, jwtSecret)).Methods("POST")
	router.HandleFunc("/tasks/{id}", middleware.JwtMiddleware(taskController.UpdateTask, jwtSecret)).Methods("PUT")
	router.HandleFunc("/tasks/{id}", middleware.JwtMiddleware(taskController.DeleteTask, jwtSecret)).Methods("DELETE")

	corsRouter := middleware.CORSMiddleware(router)
	// init server
	log.Println("Listening :8080")
	if err := http.ListenAndServe(":8080", corsRouter); err != nil {
		log.Fatal(err)
	}
}
