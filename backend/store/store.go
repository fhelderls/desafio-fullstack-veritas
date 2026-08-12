package store

import (
	"errors"
	"strconv"
	"sync"

	"desafio-fullstack-veritas/models"
)

// TaskStore mantém as tasks em memória, protegidas por um RWMutex
// para permitir leituras concorrentes sem bloquear escritas.
type TaskStore struct {
	mu     sync.RWMutex
	tasks  map[string]models.Task
	nextID int
}

// NewTaskStore cria um TaskStore vazio, pronto para uso.
func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks: make(map[string]models.Task),
	}
}

// GetAllTasks retorna todas as tasks. Usa RLock porque é uma operação
// de leitura, permitindo múltiplas chamadas simultâneas.
func (s *TaskStore) GetAllTasks() []models.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]models.Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		result = append(result, t)
	}
	return result
}

// Create adiciona uma nova task com ID sequencial gerado internamente.
func (s *TaskStore) Create(title, description, status string) models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.nextID++
	id := strconv.Itoa(s.nextID)
	task := models.Task{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      status,
	}
	s.tasks[id] = task
	return task
}

// Update sobrescreve título, descrição e status de uma task existente.
// É o mesmo método usado para mover uma task entre colunas do Kanban,
// já que mover é apenas uma mudança de Status.
func (s *TaskStore) Update(id, title, description, status string) (models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[id]
	if !exists {
		return models.Task{}, errors.New("task not found")
	}

	task.Title = title
	task.Description = description
	task.Status = status
	s.tasks[id] = task
	return task, nil
}

// Delete remove uma task pelo ID. Retorna erro se a task não existir.
func (s *TaskStore) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; !exists {
		return errors.New("task not found")
	}

	delete(s.tasks, id)
	return nil
}
