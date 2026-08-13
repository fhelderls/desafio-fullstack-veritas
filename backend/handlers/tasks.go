package handlers

import (
	store "desafio-fullstack-veritas/store"
	"encoding/json"
	"net/http"
	"time"
)

type TaskHandler struct {
	store *store.TaskStore
}

func NewTaskHandler(s *store.TaskStore) *TaskHandler {
	return &TaskHandler{store: s}
}

type taskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	DueDate     string `json:"due_date"`
}

var validStatuses = map[string]bool{
	"todo":        true,
	"in_progress": true,
	"done":        true,
}

func isValidStatus(status string) bool {
	return validStatuses[status]
}

// dueDateLayout segue a data de referencia do pacote time do Go
// (02/01/2006 = mes/dia/ano na notacao americana); aqui usamos so a
// parte de data, no formato ISO "AAAA-MM-DD".
const dueDateLayout = "2006-01-02"

// isValidDueDate aceita string vazia (data de vencimento e opcional) ou
// uma data valida no layout AAAA-MM-DD.
func isValidDueDate(dueDate string) bool {
	if dueDate == "" {
		return true
	}
	_, err := time.Parse(dueDateLayout, dueDate)
	return err == nil
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var input taskInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if input.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}
	if input.Status == "" {
		input.Status = "todo"

	}

	if !isValidStatus(input.Status) {
		http.Error(w, "Invalid status", http.StatusBadRequest)
		return
	}

	if !isValidDueDate(input.DueDate) {
		http.Error(w, "Invalid due date", http.StatusBadRequest)
		return
	}

	task := h.store.Create(input.Title, input.Description, input.Status, input.DueDate)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.store.GetAllTasks()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var input taskInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if input.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	if !isValidStatus(input.Status) {
		http.Error(w, "Invalid status", http.StatusBadRequest)
		return
	}

	if !isValidDueDate(input.DueDate) {
		http.Error(w, "Invalid due date", http.StatusBadRequest)
		return
	}

	task, err := h.store.Update(id, input.Title, input.Description, input.Status, input.DueDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.store.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
