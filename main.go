package main

import (
	"To-Do_List/cli"
	"To-Do_List/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter() // Создаем новый роутер

	// Определяем маршруты
	router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")           // Получить все задачи
	router.HandleFunc("/tasks/{id}", handlers.GetTaskByID).Methods("GET")   // Получить задачу по ID
	router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")        // Создать новую задачу
	router.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")    // Обновить задачу
	router.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE") // Удалить задачу

	// Запускаем сервер на порту 8080
	go func() {
		log.Println("Server is running on port 8080...")
		log.Fatal(http.ListenAndServe(":8080", router))
	}()
	// Запускаем CLI
	cli.RunCLI()
}
