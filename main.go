package main

import (
	"log"
	"net/http"
	"os"
	"task-api/controllers"
	"task-api/dal/database"
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
	authService := &services.AuthService{
		UserRepo:  userRepo,
		JwtSecret: os.Getenv("JWT_SECRET"),
	}
	authController := &controllers.AuthController{AuthService: authService}

	router := mux.NewRouter()

	// init routes
	router.HandleFunc("/login", authController.Login).Methods("POST")
	router.HandleFunc("/register", authController.Register).Methods("POST")

	// task routes
	taskController := &controllers.TaskController{}
	router.HandleFunc("/tasks", taskController.GetTasks).Methods("GET")
	router.HandleFunc("/tasks", taskController.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", taskController.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", taskController.DeleteTask).Methods("DELETE")

	// init server
	log.Println("Servidor escuchando en :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
