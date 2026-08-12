package main

import (
	"desafio-fullstack-veritas/handlers"
	"desafio-fullstack-veritas/store"
	"log"
	"net/http"
)

func main() {
	taskStore := store.NewTaskStore()
	handler := handlers.NewTaskHandler(taskStore)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /tasks", handler.ListTasks)
	mux.HandleFunc("POST /tasks", handler.CreateTask)
	mux.HandleFunc("PUT /tasks/{id}", handler.UpdateTask)
	mux.HandleFunc("DELETE /tasks/{id}", handler.DeleteTask)

	log.Println("Server is running on http://localhost:3001")
	log.Fatal(http.ListenAndServe(":3001", withCors(mux)))

}

func withCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
