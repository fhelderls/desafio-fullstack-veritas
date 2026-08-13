package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"desafio-fullstack-veritas/store"
)

// newTestHandler cria um TaskHandler com um TaskStore isolado num diretorio
// temporario, para nao deixar tasks.json de teste espalhado pelo projeto.
func newTestHandler(t *testing.T) *TaskHandler {
	t.Helper()

	originalWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get working directory: %v", err)
	}
	tmpDir := t.TempDir()
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("failed to chdir to temp dir: %v", err)
	}
	t.Cleanup(func() {
		os.Chdir(originalWd)
	})

	return NewTaskHandler(store.NewTaskStore())
}

func TestCreateTask(t *testing.T) {
	h := newTestHandler(t)

	body := bytes.NewBufferString(`{"title":"Nova tarefa","description":"desc","status":"todo"}`)
	req := httptest.NewRequest(http.MethodPost, "/tasks", body)
	rec := httptest.NewRecorder()

	h.CreateTask(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", rec.Code)
	}

	var task map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &task); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if task["title"] != "Nova tarefa" {
		t.Errorf("expected title 'Nova tarefa', got %v", task["title"])
	}
}

func TestCreateTaskMissingTitle(t *testing.T) {
	h := newTestHandler(t)

	body := bytes.NewBufferString(`{"title":"","description":"desc","status":"todo"}`)
	req := httptest.NewRequest(http.MethodPost, "/tasks", body)
	rec := httptest.NewRecorder()

	h.CreateTask(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status 400 for missing title, got %d", rec.Code)
	}
}

func TestCreateTaskInvalidStatus(t *testing.T) {
	h := newTestHandler(t)

	body := bytes.NewBufferString(`{"title":"Tarefa","description":"desc","status":"invalido"}`)
	req := httptest.NewRequest(http.MethodPost, "/tasks", body)
	rec := httptest.NewRecorder()

	h.CreateTask(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status 400 for invalid status, got %d", rec.Code)
	}
}

func TestCreateTaskInvalidDueDate(t *testing.T) {
	h := newTestHandler(t)

	body := bytes.NewBufferString(`{"title":"Tarefa","description":"desc","status":"todo","due_date":"20-08-2026"}`)
	req := httptest.NewRequest(http.MethodPost, "/tasks", body)
	rec := httptest.NewRecorder()

	h.CreateTask(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status 400 for invalid due date, got %d", rec.Code)
	}
}

func TestListTasks(t *testing.T) {
	h := newTestHandler(t)

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	rec := httptest.NewRecorder()

	h.ListTasks(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	var tasks []map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &tasks); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if len(tasks) != 0 {
		t.Errorf("expected an empty list, got %d tasks", len(tasks))
	}
}

func TestDeleteTaskNotFound(t *testing.T) {
	h := newTestHandler(t)

	req := httptest.NewRequest(http.MethodDelete, "/tasks/id-inexistente", nil)
	req.SetPathValue("id", "id-inexistente")
	rec := httptest.NewRecorder()

	h.DeleteTask(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", rec.Code)
	}
}
