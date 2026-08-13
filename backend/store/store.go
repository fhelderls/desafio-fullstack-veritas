package store

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"sync"

	"desafio-fullstack-veritas/models"
)

const dataFile = "tasks.json"

// TaskStore mantém as tasks em memória, protegidas por um RWMutex
// para permitir leituras concorrentes sem bloquear escritas.
type TaskStore struct {
	mu     sync.RWMutex
	tasks  map[string]models.Task
	nextID int
}

// NewTaskStore cria um TaskStore e carrega o estado salvo em tasks.json, se existir.
func NewTaskStore() *TaskStore {
	s := &TaskStore{
		tasks: make(map[string]models.Task),
	}
	s.load()
	return s
}

// load lê tasks.json e popula o mapa em memória. Se o arquivo não existir
// ou estiver corrompido, o store simplesmente começa vazio.
func (s *TaskStore) load() {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		return
	}

	var tasks []models.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return
	}

	for _, t := range tasks {
		s.tasks[t.ID] = t
		if id, err := strconv.Atoi(t.ID); err == nil && id > s.nextID {
			s.nextID = id
		}
	}
}

// save grava o estado atual em tasks.json. Só pode ser chamado por métodos
// que já detêm o lock de escrita (s.mu.Lock) — chamar s.mu.Lock() de novo
// aqui causaria deadlock, já que sync.Mutex não é reentrante.
func (s *TaskStore) save() {
	tasks := make([]models.Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		tasks = append(tasks, t)
	}

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return
	}

	os.WriteFile(dataFile, data, 0644)
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
func (s *TaskStore) Create(title, description, status, dueDate string) models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.nextID++
	id := strconv.Itoa(s.nextID)
	task := models.Task{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      status,
		DueDate:     dueDate,
	}
	s.tasks[id] = task
	s.save()
	return task
}

// Update sobrescreve título, descrição, status e data de vencimento de uma
// task existente. É o mesmo método usado para mover uma task entre colunas
// do Kanban, já que mover é apenas uma mudança de Status.
func (s *TaskStore) Update(id, title, description, status, dueDate string) (models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[id]
	if !exists {
		return models.Task{}, errors.New("task not found")
	}

	task.Title = title
	task.Description = description
	task.Status = status
	task.DueDate = dueDate
	s.tasks[id] = task
	s.save()
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
	s.save()
	return nil
}
