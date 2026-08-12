package store

import (
	"os"
	"testing"
)

// setupTestStore cria um TaskStore isolado: muda o diretório de trabalho
// para um diretório temporário (apagado automaticamente no fim do teste),
// assim o tasks.json de teste nunca toca o arquivo real do projeto.
func setupTestStore(t *testing.T) *TaskStore {
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

	return NewTaskStore()
}

func TestCreate(t *testing.T) {
	s := setupTestStore(t)

	task := s.Create("Titulo", "Descricao", "todo")

	if task.ID == "" {
		t.Error("expected a non-empty ID")
	}
	if task.Title != "Titulo" {
		t.Errorf("expected title 'Titulo', got %q", task.Title)
	}

	all := s.GetAllTasks()
	if len(all) != 1 {
		t.Fatalf("expected 1 task, got %d", len(all))
	}
}

func TestUpdate(t *testing.T) {
	s := setupTestStore(t)
	task := s.Create("Original", "desc", "todo")

	updated, err := s.Update(task.ID, "Atualizado", "nova desc", "in_progress")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if updated.Title != "Atualizado" || updated.Status != "in_progress" {
		t.Errorf("update did not apply correctly: %+v", updated)
	}
}

func TestUpdateNotFound(t *testing.T) {
	s := setupTestStore(t)

	_, err := s.Update("id-inexistente", "x", "y", "todo")
	if err == nil {
		t.Error("expected an error for a non-existent id, got nil")
	}
}

func TestDelete(t *testing.T) {
	s := setupTestStore(t)
	task := s.Create("Para deletar", "desc", "todo")

	if err := s.Delete(task.ID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	all := s.GetAllTasks()
	if len(all) != 0 {
		t.Errorf("expected 0 tasks after delete, got %d", len(all))
	}
}

func TestDeleteNotFound(t *testing.T) {
	s := setupTestStore(t)

	if err := s.Delete("id-inexistente"); err == nil {
		t.Error("expected an error for a non-existent id, got nil")
	}
}

func TestPersistsAcrossRestart(t *testing.T) {
	s := setupTestStore(t)
	s.Create("Sobrevive", "desc", "todo")

	// simula um restart: cria um novo TaskStore no mesmo diretorio de teste
	reloaded := NewTaskStore()

	all := reloaded.GetAllTasks()
	if len(all) != 1 {
		t.Fatalf("expected 1 task after reload, got %d", len(all))
	}
	if all[0].Title != "Sobrevive" {
		t.Errorf("expected title 'Sobrevive', got %q", all[0].Title)
	}
}
